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
	ID                string       `json:"id"`
	Username          string       `json:"username"`
	FullName          string       `json:"full_name"`
	Email             string       `json:"email"`
	IsEmailVerified   bool         `json:"isemail_verified"`
	Role              dto.UserRole `json:"role"`
	PasswordChangedAt time.Time    `json:"password_changedAt"`
	CreatedAt         time.Time    `json:"createdAt"`
}

type UserResponse struct {
	ID                string       `json:"id"`
	Username          string       `json:"username"`
	FullName          string       `json:"full_name"`
	Email             string       `json:"email"`
	IsEmailVerified   bool         `json:"isemail_verified"`
	Role              dto.UserRole `json:"role"`
	PasswordChangedAt time.Time    `json:"password_changedAt"`
	CreatedAt         time.Time    `json:"createdAt"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type LoginUserResponseWithPassword struct {
	ID                string       `json:"id"`
	Username          string       `json:"username"`
	FullName          string       `json:"full_name"`
	Email             string       `json:"email"`
	HashedPassword    string       `json:"password"`
	IsEmailVerified   bool         `json:"isemail_verified"`
	Role              dto.UserRole `json:"role"`
	PasswordChangedAt time.Time    `json:"password_changedAt"`
	CreatedAt         time.Time    `json:"createdAt"`
}
