package repository

import (
	"errors"

	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"

	_ "github.com/golang/mock/mockgen/model"
)

var DuplicateUser = errors.New("Duplicate Users")
var NoPersonsFound = errors.New("Users Not Found")

// mockgen -package mockdb -destination internal/infra/mock/repository/user_repo_mock.go github.com/livingdolls/go-paseto/internal/core/port/repository UserPortRepository
type UserPortRepository interface {
	CreateUser(user *dto.UserDTO) (*dto.UserDTO, error)
	GetListUser() (*[]response.RegisterUserResponse, error)
	GetUserById(id *request.GetUserByIdRequest) (response.GetUserByIdResponse, error)
	Login(user *request.LoginUserRequest) (*response.LoginUserResponseWithPassword, error)
}
