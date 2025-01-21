import { Html, Head, Main, NextScript } from 'next/document'

export default function Document() {
  return (
    <Html lang="pt-BR">
      <Head>
        <link 
          rel="icon" 
          href="https://www.capitaltrade.srv.br/wp-content/uploads/2025/01/logo_CTIS@72x.png"
        />
        <title>Gerenciador de TVs - Capital Trade</title>
      </Head>
      <body className="min-h-screen bg-[url('https://www.capitaltrade.srv.br/wp-content/uploads/2025/01/background_CT_2024-1.png')] bg-cover bg-center bg-fixed">
        <Main />
        <NextScript />
      </body>
    </Html>
  )
}