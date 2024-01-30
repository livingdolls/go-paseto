package service

import (
	"time"

	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
	"github.com/livingdolls/go-paseto/internal/core/port/repository"
	"github.com/livingdolls/go-paseto/internal/core/port/service"

	"github.com/google/uuid"
)

type userService struct {
	repo repository.UserPortRepository
}

func NewUserService(repo repository.UserPortRepository) service.UserPortService {
	return &userService{
		repo: repo,
	}
}

// Register implements service.UserPortService.
func (u *userService) Register(user *request.RegisterUserRequest) (*response.RegisterUserResponse, error) {
	id := uuid.New()

	userDTO := &dto.UserDTO{
		ID:                id,
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		IsEmailVerified:   false,
		HashedPassword:    user.Password,
		Role:              dto.Admin,
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}

	res, err := u.repo.CreateUser(userDTO)

	if err != nil {
		return nil, err
	}

	userResponse := &response.RegisterUserResponse{
		Username:  res.Username,
		FullName:  res.FullName,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}

	return userResponse, nil
}
