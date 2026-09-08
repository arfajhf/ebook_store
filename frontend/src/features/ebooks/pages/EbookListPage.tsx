import { useEffect, useState } from 'react'
import { getEbooks } from '../api/getEbooks'
import { EbookCard } from '../components/EbookCard'
import type { Ebook } from '../types/ebook'

export function EbookListPage() {
  const [ebooks, setEbooks] = useState<Ebook[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    const controller = new AbortController()

    async function loadEbooks() {
      try {
        const data = await getEbooks(controller.signal)
        setEbooks(data)
      } catch (error: unknown) {
        if (error instanceof Error && error.name === 'AbortError') {
          return
        }

        setError(
          error instanceof Error
            ? error.message
            : 'Terjadi kesalahan yang tidak diketahui',
        )
      } finally {
        setIsLoading(false)
      }
    }

    loadEbooks()

    return () => {
      controller.abort()
    }
  }, [])

  return (
    <main className="app-shell">
      <header className="hero">
        <div className="container">
          <span className="hero__label">Ebook Store</span>
          <h1>Temukan bahan ajar yang siap digunakan</h1>
          <p>
            Pilih ebook, modul ajar, panduan, dan lembar kerja
            untuk mendukung kegiatan pembelajaran.
          </p>
        </div>
      </header>

      <section className="catalog container">
        <div className="catalog__heading">
          <div>
            <span>Koleksi terbaru</span>
            <h2>Katalog ebook</h2>
          </div>

          {!isLoading && !error && (
            <p>{ebooks.length} ebook tersedia</p>
          )}
        </div>

        {isLoading && (
          <div className="state-message">Memuat katalog...</div>
        )}

        {error && (
          <div className="state-message state-message--error">
            {error}
          </div>
        )}

        {!isLoading && !error && ebooks.length === 0 && (
          <div className="state-message">
            Belum ada ebook yang tersedia.
          </div>
        )}

        {!isLoading && !error && ebooks.length > 0 && (
          <div className="ebook-grid">
            {ebooks.map((ebook) => (
              <EbookCard key={ebook.id} ebook={ebook} />
            ))}
          </div>
        )}
      </section>
    </main>
  )
}