import { useState, useEffect } from 'react'

interface StatusCheckProps {
  ip: string
}

export function useStatusCheck({ ip }: StatusCheckProps) {
  const [isOnline, setIsOnline] = useState<boolean>(false)
  const [isChecking, setIsChecking] = useState<boolean>(false)

  useEffect(() => {
    let mounted = true
    let timeoutId: NodeJS.Timeout

    const checkStatus = async () => {
      if (!mounted) return
      
      try {
        setIsChecking(true)
        const response = await fetch(`http://${ip}:8081/status`, {
          method: 'GET',
          // Adiciona um timeout de 5 segundos
          signal: AbortSignal.timeout(5000)
        })
        
        if (mounted) {
          setIsOnline(response.ok)
        }
      } catch (error) {
        if (mounted) {
          setIsOnline(false)
        }
      } finally {
        if (mounted) {
          setIsChecking(false)
          // Agenda próxima verificação
          timeoutId = setTimeout(checkStatus, 3000)
        }
      }
    }

    // Inicia verificação
    checkStatus()

    // Cleanup
    return () => {
      mounted = false
      clearTimeout(timeoutId)
    }
  }, [ip])

  return { isOnline, isChecking }
}

