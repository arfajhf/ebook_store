package ebook

type Ebook struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	IsFree      bool   `json:"is_free"`
	CoverURL    string `json:"cover_url"`
}
