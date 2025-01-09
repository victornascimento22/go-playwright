export type UrlSource = 'generic' | 'pbi'

export interface Url {
  url: string
  source: UrlSource
}

export interface TvConfig {
  urls: Url[]
  transition_time: number
  raspberry_ip: string
}

export interface TvCardProps {
  id: string
  title: string
  defaultIp: string
}

