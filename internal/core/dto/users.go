package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	Admin UserRole = "admin"
	User  UserRole = "user"
)

type UserDTO struct {
	ID                uuid.UUID
	Username          string
	FullName          string
	Email             string
	IsEmailVerified   bool
	HashedPassword    string
	Role              UserRole
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}
