interface Aniversariante {
    nome_completo: string
    nome_cracha: string
    aniversario_vida: string
    aniversario_empresa: string
    email: string
    url_aniversario_vida_tv: string
    url_aniversario_empresa_tv: string
  }
  
  interface AniversarianteResponse {
    url?: string
    error?: string
  }
  
  export const aniversariosService = {
    async getAniversariantesVida(): Promise<AniversarianteResponse> {
      try {
        const response = await fetch('http://localhost:8080/aniversario/getAniversariosVida')
        if (!response.ok) {
          throw new Error('Erro ao buscar aniversariantes')
        }
        const data: Aniversariante[] = await response.json()
        
        // Se não houver aniversariantes ou se o array estiver vazio
        if (!data || data.length === 0) {
          return { error: 'Não há aniversariantes hoje' }
        }
  
        // Pega a URL do primeiro aniversariante (assumindo que todos terão a mesma URL)
        const url = data[0].url_aniversario_vida_tv
  
        // Verifica se a URL é válida
        if (!url || url === 'NaN' || url === '') {
          return { error: 'URL do dashboard não disponível' }
        }
  
        return { url }
      } catch (error) {
        return { error: 'Erro ao buscar aniversariantes' }
      }
    },
  
    async getAniversariantesEmpresa(): Promise<AniversarianteResponse> {
      try {
        const response = await fetch('http://localhost:8080/aniversario/getAniversariosEmpresa')
        if (!response.ok) {
          throw new Error('Erro ao buscar aniversariantes')
        }
        const data: Aniversariante[] = await response.json()
        
        // Se não houver aniversariantes ou se o array estiver vazio
        if (!data || data.length === 0) {
          return { error: 'Não há aniversariantes hoje' }
        }
  
        // Pega a URL do primeiro aniversariante (assumindo que todos terão a mesma URL)
        const url = data[0].url_aniversario_empresa_tv
  
        // Verifica se a URL é válida
        if (!url || url === 'NaN' || url === '') {
          return { error: 'URL do dashboard não disponível' }
        }
  
        return { url }
      } catch (error) {
        return { error: 'Erro ao buscar aniversariantes' }
      }
    }
  }