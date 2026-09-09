package auth

import "time"

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
}
