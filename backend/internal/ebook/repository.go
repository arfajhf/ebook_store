package ebook

import "context"

type Repository interface {
	FindAll(context.Context) ([]Ebook, error)
	FindBySlug(ctx context.Context, slug string) (Ebook, error)
}
