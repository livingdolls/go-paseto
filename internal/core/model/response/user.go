package response

import (
	"time"

	"github.com/livingdolls/go-paseto/internal/core/dto"
)

type RegisterUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUserByIdResponse struct {
	ID                string
	Username          string
	FullName          string
	Email             string
	IsEmailVerified   bool
	Role              dto.UserRole
	PasswordChangedAt time.Time
	CreatedAt         time.Time
}
