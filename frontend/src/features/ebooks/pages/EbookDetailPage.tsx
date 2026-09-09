import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router'
import { formatRupiah } from '../../../shared/utils/currency'
import { getEbookBySlug } from '../api/getEbookBySlug'
import { getEbooks } from '../api/getEbooks'
import { EbookCard } from '../components/EbookCard'
import type { Ebook } from '../types/ebook'

export function EbookDetailPage() {
    const { slug } = useParams<{ slug: string }>()

    const [ebook, setEbook] = useState<Ebook | null>(null)
    const [recommendations, setRecommendations] = useState<Ebook[]>([])
    const [isLoading, setIsLoading] = useState(true)
    const [error, setError] = useState('')

    useEffect(() => {
        const controller = new AbortController()

        async function loadPage() {
            if (!slug) {
                setError('Slug ebook tidak tersedia')
                setIsLoading(false)
                return
            }

            setIsLoading(true)
            setError('')
            window.scrollTo(0, 0)

            try {
                const [ebookData, allEbooks] = await Promise.all([
                    getEbookBySlug(slug, controller.signal),
                    getEbooks(controller.signal),
                ])

                const otherEbooks = allEbooks.filter(
                    (item) => item.slug !== ebookData.slug,
                )

                const sameCategory = otherEbooks.filter(
                    (item) => item.category === ebookData.category,
                )

                const differentCategory = otherEbooks.filter(
                    (item) => item.category !== ebookData.category,
                )

                setEbook(ebookData)
                setRecommendations(
                    [...sameCategory, ...differentCategory].slice(0, 3),
                )
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

        loadPage()

        return () => {
            controller.abort()
        }
    }, [slug])

    if (isLoading) {
        return (
            <main className="detail-page">
                <div className="detail-container">
                    <div className="state-message">
                        Memuat detail ebook...
                    </div>
                </div>
            </main>
        )
    }

    if (error || !ebook) {
        return (
            <main className="detail-page">
                <div className="detail-container">
                    <div className="state-message state-message--error">
                        <p>{error || 'Ebook tidak ditemukan'}</p>
                        <Link to="/">Kembali ke katalog</Link>
                    </div>
                </div>
            </main>
        )
    }

    return (
        <main className="detail-page">
            <div className="detail-container">
                <Link className="detail-back" to="/">
                    ← Kembali ke katalog
                </Link>

                <article className="ebook-detail">
                    <div className="ebook-detail__cover">
                        {ebook.cover_url ? (
                            <img
                                src={ebook.cover_url}
                                alt={`Cover ${ebook.title}`}
                            />
                        ) : (
                            <div className="ebook-detail__placeholder">
                                Belum ada cover
                            </div>
                        )}
                    </div>

                    <div className="ebook-detail__content">
                        <span className="ebook-detail__category">
                            {ebook.category}
                        </span>

                        <h1>{ebook.title}</h1>

                        <p className="ebook-detail__author">
                            Oleh {ebook.author}
                        </p>

                        <p className="ebook-detail__description">
                            {ebook.description}
                        </p>

                        <div className="ebook-detail__price">
                            {ebook.is_free
                                ? 'Gratis'
                                : formatRupiah(ebook.price)}
                        </div>

                        <button type="button" className="primary-button">
                            {ebook.is_free ? 'Dapatkan Gratis' : 'Beli Sekarang'}
                        </button>
                    </div>
                </article>

                {recommendations.length > 0 && (
                    <section className="recommendations">
                        <div className="recommendations__heading">
                            <span>Pilihan lainnya</span>
                            <h2>Rekomendasi ebook lainnya</h2>
                            <p>
                                Temukan bahan ajar lain yang mungkin lo butuhkan.
                            </p>
                        </div>

                        <div className="ebook-grid">
                            {recommendations.map((item) => (
                                <EbookCard key={item.id} ebook={item} />
                            ))}
                        </div>
                    </section>
                )}
            </div>
        </main>
    )
}