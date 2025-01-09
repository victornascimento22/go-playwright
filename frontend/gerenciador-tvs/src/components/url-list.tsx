"use client"

import { Plus } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import type { Url, UrlSource } from "@/types/tv"

interface UrlListProps {
  urls: Url[]
  onUrlChange: (urls: Url[]) => void
}

export function UrlList({ urls, onUrlChange }: UrlListProps) {
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

  return (
    <div className="space-y-2">
      {urls.map((item, index) => (
        <div key={index} className="flex gap-2">
          <Input
            type="text"
            value={item.url}
            onChange={(e) => handleUrlChange(index, e.target.value)}
            placeholder="https://seu-dashboard.com"
            className="flex-1 bg-white/90"
          />
          <Select
            value={item.source}
            onValueChange={(value) => handleSourceChange(index, value as UrlSource)}
          >
            <SelectTrigger className="w-[140px] bg-white/90">
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
