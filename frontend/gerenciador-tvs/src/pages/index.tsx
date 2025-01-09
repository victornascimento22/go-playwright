import { TvCard } from "@/components/tv-card"

const TVS = [
  { id: "1", title: "TV Operação 1", ip: "192.168.1.101" },
  { id: "2", title: "TV Operação 2", ip: "192.168.1.102" },
  { id: "3", title: "TV Operação 3", ip: "192.168.1.103" },
  { id: "4", title: "TV Operação 4", ip: "192.168.1.104" },
  { id: "5", title: "TV Operação 5", ip: "192.168.1.105" },
  { id: "6", title: "TV Operação 6", ip: "192.168.1.106" },
]

export default function Home() {
  return (
    <main className="container py-8">
      <h1 className="text-3xl font-bold mb-8 text-white text-center drop-shadow-lg">
        Gerenciador de TVs
      </h1>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {TVS.map((tv) => (
          <TvCard
            key={tv.id}
            id={tv.id}
            title={tv.title}
            defaultIp={tv.ip}
          />
        ))}
      </div>
    </main>
  )
}
