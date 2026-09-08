package ebook

import "context"

type Repository interface {
	FindAll(context.Context) ([]Ebook, error)
}
