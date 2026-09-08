package ebook

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) GetAll(
	ctx context.Context,
) ([]Ebook, error) {
	return service.repository.FindAll(ctx)
}

func (service *Service) GetBySlug(
	ctx context.Context,
	slug string,
) (Ebook, error) {
	return service.repository.FindBySlug(ctx, slug)
}
