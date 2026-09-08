import { useEffect, useState } from 'react'
import './App.css'

interface HealthResponse {
  status: string
  message: string
}

function App() {
  const [apiData, setApiData] = useState<HealthResponse | null>(null)
  const [error, setError] = useState<string>('')

  useEffect(() => {
    async function checkBackend() {
      try {
        const response = await fetch('/api/v1/health')

        if (!response.ok) {
          throw new Error('Backend tidak memberikan respons yang valid')
        }

        const data: HealthResponse = await response.json()
        setApiData(data)
      } catch (error: unknown) {
        if (error instanceof Error) {
          setError(error.message)
          return
        }

        setError('Terjadi kesalahan yang tidak diketahui')
      }
    }

    checkBackend()
  }, [])

  return (
    <main>
      <h1>Ebook Store</h1>
      <p>Frontend React TypeScript dan backend Go</p>

      {!apiData && !error && <p>Menghubungkan ke backend...</p>}

      {apiData && (
        <section>
          <h2>Backend terhubung</h2>
          <p>Status: {apiData.status}</p>
          <p>Pesan: {apiData.message}</p>
        </section>
      )}

      {error && (
        <section>
          <h2>Koneksi gagal</h2>
          <p>{error}</p>
        </section>
      )}
    </main>
  )
}

export default App