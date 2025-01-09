import { Trash2 } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import type { Url, UrlSource } from "@/types/tv"

interface UrlFormProps {
  url: Url
  index: number
  onUrlChange: (index: number, newUrl: string) => void
  onSourceChange: (index: number, source: UrlSource) => void
  onRemove: (index: number) => void
  canRemove: boolean
}

export function UrlForm({
  url,
  index,
  onUrlChange,
  onSourceChange,
  onRemove,
  canRemove
}: UrlFormProps) {
  return (
    <div className="flex gap-2">
      <Input
        type="url"
        value={url.url}
        onChange={(e) => onUrlChange(index, e.target.value)}
        placeholder="https://seu-dashboard.com"
        className="flex-1 text-black"
      />
      <Select
        value={url.source}
        onValueChange={(value) => onSourceChange(index, value as UrlSource)}
      >
        <SelectTrigger className="w-[140px] text-black">
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="generic">Generic</SelectItem>
          <SelectItem value="pbi">Power BI</SelectItem>
        </SelectContent>
      </Select>
      <Button
        variant="destructive"
        size="icon"
        onClick={() => onRemove(index)}
        disabled={!canRemove}
      >
        <Trash2 className="h-4 w-4" />
        <span className="sr-only">Remover URL</span>
      </Button>
    </div>
  )
}

