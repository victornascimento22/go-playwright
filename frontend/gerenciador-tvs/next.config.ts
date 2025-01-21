/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    domains: ['www.capitaltrade.srv.br', 'app.powerbi.com'],
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'www.capitaltrade.srv.br',
        pathname: '/wp-content/uploads/**',
      },
    ],
  },
}

module.exports = nextConfig