package service

import (
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
)

// mockgen -package mockdb -destination internal/infra/mock/service/user_service_mock.go github.com/livingdolls/go-paseto/internal/core/port/service UserPortService
type UserPortService interface {
	Register(user *request.RegisterUserRequest) (*response.RegisterUserResponse, error)
	ListUsers() (*[]response.RegisterUserResponse, error)
	GetUser(id *request.GetUserByIdRequest) (*response.GetUserByIdResponse, error)
}
