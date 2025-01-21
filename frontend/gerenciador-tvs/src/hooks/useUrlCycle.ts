import { useState, useEffect, useRef } from "react"
import { displayService } from "@/services/api"
import type { Url } from "@/types/tv"

interface UseUrlCycleProps {
  id: string
  urls: Url[]
  transitionTime: number
  raspberryIp: string
  onError?: (error: string) => void
}

export function useUrlCycle({ id, urls, transitionTime, raspberryIp, onError }: UseUrlCycleProps) {
  const [isRunning, setIsRunning] = useState<boolean>(false)
  const [currentUrlIndex, setCurrentUrlIndex] = useState<number>(0)
  const [lastUpdate, setLastUpdate] = useState<Date | null>(null)
  const timeoutRef = useRef<NodeJS.Timeout | undefined>(undefined)
  const validUrls = urls.filter((url) => url.url.trim() !== "")

  const updateDisplay = async () => {
    try {
      console.log(`[TV ${id}] Atualizando display`)
      await displayService.updateDisplay({
        urls: validUrls,
        transition_time: transitionTime,
        raspberry_ip: raspberryIp,
      })
      setLastUpdate(new Date())
      return true
    } catch (error) {
      console.error(`[TV ${id}] Erro ao atualizar display:`, error)
      onError?.(error instanceof Error ? error.message : "Erro ao atualizar display")
      return false
    }
  }

  const stopCycle = () => {
    console.log(`[TV ${id}] Parando ciclo`)
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current)
      timeoutRef.current = undefined
    }
    setIsRunning(false)
    setCurrentUrlIndex(0)
  }

  const startCycle = async () => {
    console.log(`[TV ${id}] Iniciando ciclo com URLs:`, validUrls)
    if (validUrls.length === 0) {
      onError?.("Nenhuma URL configurada")
      return
    }

    setIsRunning(true)
    setCurrentUrlIndex(0)

    const success = await updateDisplay()
    if (!success) {
      stopCycle()
    }
  }

  useEffect(() => {
    if (!isRunning || validUrls.length === 0) {
      return
    }

    const cycleThroughUrls = async () => {
      const nextIndex = (currentUrlIndex + 1) % validUrls.length
      console.log(`[TV ${id}] Avançando para próxima URL:`, {
        currentIndex: currentUrlIndex,
        nextIndex,
        url: validUrls[nextIndex].url,
      })

      timeoutRef.current = setTimeout(async () => {
        setCurrentUrlIndex(nextIndex)
        const success = await updateDisplay()
        if (!success) {
          stopCycle()
        }
      }, transitionTime * 1000)
    }

    cycleThroughUrls()

    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current)
        timeoutRef.current = undefined
      }
    }
  }, [isRunning, currentUrlIndex, validUrls, transitionTime, raspberryIp])

  useEffect(() => {
    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current)
        timeoutRef.current = undefined
      }
    }
  }, [])

  return {
    isRunning,
    currentUrlIndex,
    lastUpdate,
    startCycle,
    stopCycle,
    updateDisplay,
  }
}

