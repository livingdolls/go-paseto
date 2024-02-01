package service

import (
	"time"

	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/entity"
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

// ListUsers implements service.UserPortService.
func (u *userService) ListUsers() (*[]response.RegisterUserResponse, error) {
	res, err := u.repo.GetListUser()

	if err != nil {
		return nil, err
	}

	return res, err
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
		HashedPassword:    user.Password, //TODO :: Hash Password
		Role:              dto.Admin,
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}

	res, err := u.repo.CreateUser(userDTO)

	if err != nil {
		if err == entity.ErrConflictingData {
			return nil, entity.ErrConflictingData
		}
		return nil, err
	}

	// TODO :: Create Message Broker Send Mail

	userResponse := &response.RegisterUserResponse{
		Username:  res.Username,
		FullName:  res.FullName,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}

	return userResponse, nil
}

// GetUser implements service.UserPortService.
func (u *userService) GetUser(id *request.GetUserByIdRequest) (*response.GetUserByIdResponse, error) {
	newId := &request.GetUserByIdRequest{
		ID: id.ID,
	}
	res, err := u.repo.GetUserById(newId)

	if err != nil {
		if err == entity.ErrDataNotFound {
			return nil, err
		}

		return nil, entity.ErrInternal
	}

	return &res, nil
}

// Login implements service.UserPortService.
func (u *userService) Login(user *request.LoginUserRequest) (*response.LoginUserResponse, error) {
	res, err := u.repo.Login(user)
	var result response.LoginUserResponse

	if err != nil {
		if err == entity.ErrDataNotFound {
			return nil, entity.ErrDataNotFound
		} else if err == entity.ErrNoMatchPassword {
			return nil, entity.ErrNoMatchPassword
		}

		return nil, entity.ErrInternal
	}

	// TODO :: Make Paseto Token

	result = response.LoginUserResponse{
		ID:                res.ID,
		Username:          res.Username,
		FullName:          res.FullName,
		Email:             res.Email,
		IsEmailVerified:   res.IsEmailVerified,
		Role:              res.Role,
		PasswordChangedAt: result.PasswordChangedAt,
		CreatedAt:         res.CreatedAt,
	}

	return &result, nil

}
