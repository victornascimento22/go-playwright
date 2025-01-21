"use client"

import { useState } from 'react'
import { Plus, Gift, Briefcase, RotateCw } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { useToast } from "@/components/ui/use-toast"
import { aniversariosService } from "@/services/aniversarios"
import { useAutoAniversariantes } from "@/hooks/useAutoAniversariantes"
import type { Url, UrlSource } from "@/types/tv"

interface UrlListProps {
  urls: Url[]
  onUrlChange: (urls: Url[]) => void
  onAniversarianteAdd: (url: string) => void
}

export function UrlList({ urls, onUrlChange, onAniversarianteAdd }: UrlListProps) {
  const { toast } = useToast()
  const [isLoadingVida, setIsLoadingVida] = useState(false)
  const [isLoadingEmpresa, setIsLoadingEmpresa] = useState(false)

  // Usando o hook de verificação automática
  const { isChecking, lastCheck, forceCheck } = useAutoAniversariantes({
    onAniversarianteFound: (url) => {
      onAniversarianteAdd(url)
      toast({
        title: "Aniversariantes Encontrados",
        description: "URL dos aniversariantes foi adicionada automaticamente!",
      })
    }
  })

  const handleAddUrl = () => {
    onUrlChange([...urls, { url: "", source: "generic" }])
  }

  const handleRemoveUrl = (index: number) => {
    if (urls.length > 1) {
      onUrlChange(urls.filter((_, i) => i !== index))
    }
  }

  const handleUrlChange = (index: number, newUrl: string) => {
    onUrlChange(urls.map((item, i) => (i === index ? { ...item, url: newUrl } : item)))
  }

  const handleSourceChange = (index: number, source: UrlSource) => {
    onUrlChange(urls.map((item, i) => (i === index ? { ...item, source } : item)))
  }

  const handleBuscarAniversariantesVida = async () => {
    setIsLoadingVida(true)
    try {
      const response = await aniversariosService.getAniversariantesVida()
      if (response.url) {
        onAniversarianteAdd(response.url)
        toast({
          title: "Aniversariantes Encontrados",
          description: "URL dos aniversariantes foi adicionada com sucesso!",
        })
      } else {
        toast({
          title: "Nenhum Aniversariante",
          description: response.error || "Não há aniversariantes de vida hoje.",
          variant: "destructive",
        })
      }
    } catch (error) {
      toast({
        title: "Erro",
        description: "Erro ao buscar aniversariantes de vida.",
        variant: "destructive",
      })
    } finally {
      setIsLoadingVida(false)
    }
  }

  const handleBuscarAniversariantesEmpresa = async () => {
    setIsLoadingEmpresa(true)
    try {
      const response = await aniversariosService.getAniversariantesEmpresa()
      if (response.url) {
        onAniversarianteAdd(response.url)
        toast({
          title: "Aniversariantes Encontrados",
          description: "URL dos aniversariantes foi adicionada com sucesso!",
        })
      } else {
        toast({
          title: "Nenhum Aniversariante",
          description: response.error || "Não há aniversariantes de empresa hoje.",
          variant: "destructive",
        })
      }
    } catch (error) {
      toast({
        title: "Erro",
        description: "Erro ao buscar aniversariantes de empresa.",
        variant: "destructive",
      })
    } finally {
      setIsLoadingEmpresa(false)
    }
  }

  return (
    <div className="space-y-4">
      <div className="grid grid-cols-2 gap-2">
        <Button
          variant="outline"
          className="bg-white/10 text-white hover:bg-white/20 hover:text-white"
          onClick={handleBuscarAniversariantesVida}
          disabled={isLoadingVida}
        >
          <Gift className="mr-2 h-4 w-4" />
          {isLoadingVida ? "Buscando..." : "Aniversariantes Vida"}
        </Button>
        <Button
          variant="outline"
          className="bg-white/10 text-white hover:bg-white/20 hover:text-white"
          onClick={handleBuscarAniversariantesEmpresa}
          disabled={isLoadingEmpresa}
        >
          <Briefcase className="mr-2 h-4 w-4" />
          {isLoadingEmpresa ? "Buscando..." : "Aniversariantes Empresa"}
        </Button>
      </div>

      {/* Adiciona informações sobre a verificação automática */}
      <div className="flex items-center justify-between text-sm text-white/70">
        <span>
          {lastCheck 
            ? `Última verificação: ${new Date(lastCheck).toLocaleDateString()}`
            : 'Aguardando primeira verificação...'}
        </span>
        <Button
          variant="ghost"
          size="sm"
          className="text-white/70 hover:text-white"
          onClick={forceCheck}
          disabled={isChecking}
        >
          <RotateCw className={`h-4 w-4 mr-2 ${isChecking ? 'animate-spin' : ''}`} />
          Forçar verificação
        </Button>
      </div>

      <div className="space-y-2">
        {urls.map((item, index) => (
          <div key={index} className="flex gap-2">
            <Input
              type="text"
              value={item.url}
              onChange={(e) => handleUrlChange(index, e.target.value)}
              placeholder="https://seu-dashboard.com"
              className="flex-1 bg-white/90 text-black"
            />
            <Select
              value={item.source}
              onValueChange={(value) => handleSourceChange(index, value as UrlSource)}
            >
              <SelectTrigger className="w-[140px] bg-white/90 text-black">
                <SelectValue />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="generic">Generic</SelectItem>
                <SelectItem value="pbi">Power BI</SelectItem>
              </SelectContent>
            </Select>
            <Button
              variant="destructive"
              onClick={() => handleRemoveUrl(index)}
              disabled={urls.length <= 1}
              className="px-4"
            >
              ×
            </Button>
          </div>
        ))}
      </div>

      <Button
        variant="outline"
        className="w-full bg-white/10 text-white hover:bg-white/20 hover:text-white"
        onClick={handleAddUrl}
      >
        <Plus className="mr-2 h-4 w-4" />
        Adicionar URL
      </Button>
    </div>
  )
}

