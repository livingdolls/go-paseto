package service

import (
	"gopaseto/internal/core/model/request"
	"gopaseto/internal/core/model/response"
)

type UserPortService interface {
	Register(user *request.RegisterUserRequest) (*response.RegisterUserResponse, error)
}
