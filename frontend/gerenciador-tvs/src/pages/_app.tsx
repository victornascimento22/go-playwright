import type { AppProps } from 'next/app'
import { Inter } from 'next/font/google'
import Head from 'next/head'
import '@/styles/globals.css'

const inter = Inter({ subsets: ['latin'] })

export default function App({ Component, pageProps }: AppProps) {
  return (
    <>
      <Head>
        <title>Gerenciador de TVs - Capital Trade</title>
        <link 
          rel="icon" 
          href="https://www.capitaltrade.srv.br/wp-content/uploads/2025/01/logo_CTIS@72x.png"
        />
      </Head>
      <div className={inter.className}>
        <Component {...pageProps} />
      </div>
    </>
  )
}