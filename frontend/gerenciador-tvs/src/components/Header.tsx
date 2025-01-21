import Image from 'next/image'

export default function Header() {
  return (
    <header className="absolute top-0 left -0 p-4">
      <Image
        src="https://www.capitaltrade.srv.br/wp-content/uploads/2025/01/logo_sistemas_02.png"
        alt="Capital Trade Logo"
        width={150}
        height={50}
        objectFit="contain"
      />
    </header>
  )
}

