"use client"

import { useState } from "react"
import { Monitor, RotateCw } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useToast } from "@/components/ui/use-toast"
import { displayService } from "@/services/api"
import { UrlList } from "./url-list"
import { StatusIndicator } from "./status-indicator"
import { useUrlCycle } from "@/hooks/useUrlCycle"
import type { TvCardProps, Url } from "@/types/tv"

export function TvCard({ id, title, defaultIp }: TvCardProps) {
  const { toast } = useToast()
  const [isLoading, setIsLoading] = useState(false)
  const [transitionTime, setTransitionTime] = useState(15)
  const [raspberryIp, setRaspberryIp] = useState(defaultIp)
  const [urls, setUrls] = useState<Url[]>([{ url: "", source: "generic" }])

  const { isRunning, currentUrlIndex, lastUpdate, startCycle, stopCycle, updateDisplay } = useUrlCycle({
    id,
    urls: urls.filter((url) => url.url.trim() !== ""),
    transitionTime,
    raspberryIp,
    onError: (error) => {
      toast({
        title: "Erro no ciclo de URLs",
        description: error,
        variant: "destructive",
      })
    },
  })

  const isValidIp = (ip: string) => {
    const ipRegex =
      /^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^localhost$/
    return ipRegex.test(ip)
  }

  const handleUpdateUrls = async () => {
    if (!isValidIp(raspberryIp)) {
      toast({
        title: "IP Inválido",
        description: "Por favor, insira um endereço IP válido",
        variant: "destructive",
      })
      return
    }

    if (urls.some((url) => !url.url)) {
      toast({
        title: "URL Inválida",
        description: "Por favor, preencha todas as URLs",
        variant: "destructive",
      })
      return
    }

    setIsLoading(true)
    try {
      await updateDisplay()
      toast({
        title: "URLs Atualizadas",
        description: "As configurações foram atualizadas com sucesso",
      })
    } catch (error) {
      toast({
        title: "Erro",
        description: error instanceof Error ? error.message : "Erro ao atualizar URLs",
        variant: "destructive",
      })
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <Card className="border-0 shadow-xl text-white h-full">
      <CardHeader>
        <div className="flex items-center justify-between">
          <CardTitle className="flex items-center gap-2">
            <Monitor className="h-5 w-5" />
            {title}
          </CardTitle>
          <StatusIndicator isOnline={isRunning} isChecking={false} />
        </div>
      </CardHeader>
      <CardContent className="space-y-4 bg-slate-900/60 backdrop-blur-sm rounded-b-lg">
        <div className="space-y-2">
          <UrlList
            urls={urls}
            onUrlChange={setUrls}
            onAniversarianteAdd={(url) => setUrls((current) => [...current, { url, source: "generic" }])}
          />
        </div>

        <div className="pt-4 border-t border-white/10 space-y-4">
          <div className="space-y-2">
            <Label htmlFor={`transition-${id}`}>Tempo de Transição (segundos)</Label>
            <Input
              id={`transition-${id}`}
              type="number"
              min={1}
              value={transitionTime}
              onChange={(e) => setTransitionTime(Number(e.target.value))}
              className="bg-white/90 text-black"
            />
          </div>

          <div className="space-y-2">
            <Label htmlFor={`ip-${id}`}>IP do Raspberry</Label>
            <Input
              id={`ip-${id}`}
              type="text"
              value={raspberryIp}
              onChange={(e) => setRaspberryIp(e.target.value)}
              className={`bg-white/90 text-black ${!isValidIp(raspberryIp) ? "border-red-500" : ""}`}
            />
          </div>

          <div className="grid grid-cols-2 gap-2">
            <Button className="bg-green-600 hover:bg-green-700 text-white" onClick={isRunning ? stopCycle : startCycle}>
              {isRunning ? "Parar Ciclo" : "Iniciar Ciclo"}
            </Button>
            <Button className="bg-slate-600 hover:bg-slate-700" onClick={handleUpdateUrls} disabled={isLoading}>
              {isLoading ? (
                <span className="flex items-center gap-2">
                  <RotateCw className="h-4 w-4 animate-spin" />
                  Atualizando...
                </span>
              ) : (
                "Atualizar URLs"
              )}
            </Button>
          </div>
        </div>

        {currentUrlIndex !== undefined && urls[currentUrlIndex] && (
          <div className="pt-2 text-sm text-white/60">URL atual: {urls[currentUrlIndex].url}</div>
        )}
      </CardContent>
    </Card>
  )
}

