import { ChevronLeft, ChevronRight } from 'lucide-react'
import { TvCard } from "@/components/tv-card"
import { Button } from "@/components/ui/button"
import { useState } from 'react'

const TVS = [
  { id: "1", title: "TV Sistemas", ip: "172.16.14.165" },
  { id: "2", title: "TV Operação 2", ip: "192.168.1.102" },
  { id: "3", title: "TV Operação 3", ip: "192.168.1.103" },
  { id: "4", title: "TV Operação 4", ip: "192.168.1.104" },
  { id: "5", title: "TV Operação 5", ip: "192.168.1.105" },
  { id: "6", title: "TV Operação 6", ip: "192.168.1.106" },
]

export default function Home() {
  const [currentIndex, setCurrentIndex] = useState(0)

  const handlePrevious = () => {
    setCurrentIndex((current) => (current > 0 ? current - 1 : TVS.length - 1))
  }

  const handleNext = () => {
    setCurrentIndex((current) => (current < TVS.length - 1 ? current + 1 : 0))
  }

  return (
    <main className="min-h-screen py-8">
      <h1 className="text-3xl font-bold mb-8 text-white text-center drop-shadow-lg">
        Gerenciador de TVs
      </h1>
      
      <div className="relative max-w-2xl mx-auto px-12">
        <div className="overflow-hidden">
          <div 
            className="flex transition-transform duration-300 ease-in-out"
            style={{ transform: `translateX(-${currentIndex * 100}%)` }}
          >
            {TVS.map((tv) => (
              <div key={tv.id} className="w-full flex-shrink-0 px-4">
                <TvCard
                  id={tv.id}
                  title={tv.title}
                  defaultIp={tv.ip}
                />
              </div>
            ))}
          </div>
        </div>
        
        <Button
          variant="ghost"
          size="icon"
          className="absolute left-0 top-1/2 -translate-y-1/2 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 z-10"
          onClick={handlePrevious}
        >
          <ChevronLeft className="h-6 w-6" />
          <span className="sr-only">TV anterior</span>
        </Button>

        <Button
          variant="ghost"
          size="icon"
          className="absolute right-0 top-1/2 -translate-y-1/2 bg-white/10 hover:bg-white/20 text-white rounded-full p-2 z-10"
          onClick={handleNext}
        >
          <ChevronRight className="h-6 w-6" />
          <span className="sr-only">Próxima TV</span>
        </Button>

        <div className="flex justify-center gap-2 mt-4">
          {TVS.map((_, index) => (
            <button
              key={index}
              className={`w-2 h-2 rounded-full transition-colors ${
                index === currentIndex ? 'bg-white' : 'bg-white/30'
              }`}
              onClick={() => setCurrentIndex(index)}
              aria-label={`Ir para TV ${index + 1}`}
            />
          ))}
        </div>
      </div>
    </main>
  )
}

