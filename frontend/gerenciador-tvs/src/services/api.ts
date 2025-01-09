interface DisplayConfig {
  urls: Array<{ url: string; source: 'generic' | 'pbi' }>;
  transition_time: number;
  raspberry_ip: string;
}

export const displayService = {
  async updateDisplay(config: DisplayConfig) {
    const response = await fetch('http://localhost:8080/screenshots/update', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(config)
    });

    if (!response.ok) {
      const data = await response.json();
      throw new Error(data.error || 'Erro ao atualizar display');
    }

    return response.json();
  }
};

