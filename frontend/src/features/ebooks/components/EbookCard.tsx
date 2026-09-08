import { formatRupiah } from '../../../shared/utils/currency'
import type { Ebook } from '../types/ebook'

interface EbookCardProps {
  ebook: Ebook
}

export function EbookCard({ ebook }: EbookCardProps) {
  return (
    <article className="ebook-card">
      <div className="ebook-card__cover">
        {ebook.cover_url ? (
          <img src={ebook.cover_url} alt={`Cover ${ebook.title}`} />
        ) : (
          <div className="ebook-card__placeholder">
            <span>Belum ada cover</span>
          </div>
        )}

        <span
          className={
            ebook.is_free
              ? 'ebook-card__badge ebook-card__badge--free'
              : 'ebook-card__badge'
          }
        >
          {ebook.is_free ? 'Gratis' : 'Berbayar'}
        </span>
      </div>

      <div className="ebook-card__content">
        <span className="ebook-card__category">
          {ebook.category}
        </span>

        <h2>{ebook.title}</h2>
        <p className="ebook-card__author">Oleh {ebook.author}</p>
        <p className="ebook-card__description">
          {ebook.description}
        </p>

        <div className="ebook-card__footer">
          <strong>
            {ebook.is_free ? 'Gratis' : formatRupiah(ebook.price)}
          </strong>

          <button type="button">Lihat Detail</button>
        </div>
      </div>
    </article>
  )
}