import { PlayCircle, StopCircle, RefreshCcw } from "lucide-react"
import { Button } from "@/components/ui/button"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"

interface UrlControlsProps {
  isRunning: boolean
  onStart: () => void
  onStop: () => void
  onUpdate: () => void
  lastUpdate: Date | null
  disabled?: boolean
}

export function UrlControls({ isRunning, onStart, onStop, onUpdate, lastUpdate, disabled }: UrlControlsProps) {
  return (
    <div className="space-y-4">
      <div className="grid grid-cols-2 gap-2">
        <Button
          onClick={isRunning ? onStop : onStart}
          disabled={disabled}
          className={`w-full flex items-center justify-center ${
            isRunning ? "bg-red-600 hover:bg-red-700 text-white" : "bg-green-600 hover:bg-green-700 text-white"
          }`}
        >
          {isRunning ? (
            <>
              <StopCircle className="mr-2 h-4 w-4" />
              Parar Ciclo
            </>
          ) : (
            <>
              <PlayCircle className="mr-2 h-4 w-4" />
              Iniciar Ciclo
            </>
          )}
        </Button>
        <Button
          onClick={onUpdate}
          disabled={disabled || isRunning}
          variant="secondary"
          className="w-full flex items-center justify-center bg-slate-700 hover:bg-slate-600 text-white"
        >
          <RefreshCcw className="mr-2 h-4 w-4" />
          Atualizar URLs
        </Button>
      </div>
      {lastUpdate && (
        <p className="text-sm text-white/70 text-center">
          Última atualização: {format(lastUpdate, "dd/MM/yyyy HH:mm:ss", { locale: ptBR })}
        </p>
      )}
    </div>
  )
}

