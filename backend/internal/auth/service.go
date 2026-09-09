package auth

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"
)

type PasswordHasher interface {
	Hash(plainPassword string) (string, error)
	Verify(
		plainPassword string,
		encodedPassword string,
	) (bool, error)
}

type Service struct {
	repository Repository
	hasher     PasswordHasher
}

func NewService(
	repository Repository,
	hasher PasswordHasher,
) *Service {
	return &Service{
		repository: repository,
		hasher:     hasher,
	}
}

func (service *Service) Register(
	ctx context.Context,
	request RegisterRequest,
) (User, error) {
	name := strings.TrimSpace(request.Name)
	email := strings.ToLower(strings.TrimSpace(request.Email))

	if utf8.RuneCountInString(name) < 2 {
		return User{}, ErrInvalidName
	}

	passwordHash, err := service.hasher.Hash(request.Password)
	if err != nil {
		return User{}, fmt.Errorf("hash password: %w", err)
	}

	user := User{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         "customer",
		IsActive:     true,
	}

	return service.repository.Create(ctx, user)
}
