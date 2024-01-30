package service

import (
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
)

type UserPortService interface {
	Register(user *request.RegisterUserRequest) (*response.RegisterUserResponse, error)
}
