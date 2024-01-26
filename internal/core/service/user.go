package service

import (
	"gopaseto/internal/core/dto"
	"gopaseto/internal/core/model/request"
	"gopaseto/internal/core/model/response"
	"gopaseto/internal/core/port/repository"
	"gopaseto/internal/core/port/service"
	"time"

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
