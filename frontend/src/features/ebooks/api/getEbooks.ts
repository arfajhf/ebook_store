import type {
  Ebook,
  EbookListResponse,
} from '../types/ebook'

export async function getEbooks(
  signal?: AbortSignal,
): Promise<Ebook[]> {
  const response = await fetch('/api/v1/ebooks', {
    method: 'GET',
    headers: {
      Accept: 'application/json',
    },
    signal,
  })

  if (!response.ok) {
    throw new Error('Gagal mengambil data ebook')
  }

  const result: EbookListResponse = await response.json()

  return result.data
}