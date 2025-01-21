import { useState, useEffect } from 'react'
import { aniversariosService } from '@/services/aniversarios'

const STORAGE_KEY = 'last_aniversariantes_check'
const CHECK_HOUR = 7 // Hora do dia para verificar (7:00 AM)

interface UseAutoAniversariantesProps {
  onAniversarianteFound: (url: string) => void
}

export function useAutoAniversariantes({ onAniversarianteFound }: UseAutoAniversariantesProps) {
  const [isChecking, setIsChecking] = useState(false)
  const [lastCheck, setLastCheck] = useState<string>('')

  // Função para verificar se devemos fazer a checagem hoje
  const shouldCheckToday = () => {
    const lastCheckDate = localStorage.getItem(STORAGE_KEY)
    const today = new Date().toDateString()
    
    // Se nunca verificou ou se a última verificação não foi hoje
    return !lastCheckDate || lastCheckDate !== today
  }

  // Função para verificar se está na hora certa do dia
  const isCheckTime = () => {
    const now = new Date()
    return now.getHours() >= CHECK_HOUR
  }

  // Função principal de verificação
  const checkAniversariantes = async () => {
    if (isChecking) return
    
    setIsChecking(true)
    try {
      // Verifica aniversariantes de vida
      const vidaResponse = await aniversariosService.getAniversariantesVida()
      if (vidaResponse.url) {
        onAniversarianteFound(vidaResponse.url)
      }

      // Verifica aniversariantes de empresa
      const empresaResponse = await aniversariosService.getAniversariantesEmpresa()
      if (empresaResponse.url) {
        onAniversarianteFound(empresaResponse.url)
      }

      // Atualiza a data da última verificação
      const today = new Date().toDateString()
      localStorage.setItem(STORAGE_KEY, today)
      setLastCheck(today)
      
    } catch (error) {
      console.error('Erro ao verificar aniversariantes:', error)
    } finally {
      setIsChecking(false)
    }
  }

  // Efeito para verificação automática
  useEffect(() => {
    const checkIfNeeded = () => {
      if (shouldCheckToday() && isCheckTime()) {
        checkAniversariantes()
      }
    }

    // Verifica imediatamente ao montar o componente
    checkIfNeeded()

    // Configura um intervalo para verificar a cada hora
    const interval = setInterval(checkIfNeeded, 1000 * 60 * 60) // A cada hora

    return () => clearInterval(interval)
  }, [])

  // Função para forçar uma verificação (útil para testes)
  const forceCheck = () => {
    localStorage.removeItem(STORAGE_KEY)
    checkAniversariantes()
  }

  return {
    isChecking,
    lastCheck,
    forceCheck
  }
}

