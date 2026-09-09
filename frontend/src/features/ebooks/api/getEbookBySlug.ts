import type {
  Ebook,
  EbookDetailResponse,
} from '../types/ebook'

export async function getEbookBySlug(
  slug: string,
  signal?: AbortSignal,
): Promise<Ebook> {
  const response = await fetch(
    `/api/v1/ebooks/${encodeURIComponent(slug)}`,
    {
      headers: {
        Accept: 'application/json',
      },
      signal,
    },
  )

  if (response.status === 404) {
    throw new Error('Ebook tidak ditemukan')
  }

  if (!response.ok) {
    throw new Error('Gagal mengambil detail ebook')
  }

  const result: EbookDetailResponse = await response.json()

  return result.data
}