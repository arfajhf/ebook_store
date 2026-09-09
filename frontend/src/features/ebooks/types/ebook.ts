export interface Ebook {
  id: number
  title: string
  slug: string
  author: string
  category: string
  description: string
  price: number
  is_free: boolean
  cover_url: string
}

export interface EbookListResponse {
  status: string
  data: Ebook[]
}

export interface EbookDetailResponse {
  status: string
  data: Ebook
}