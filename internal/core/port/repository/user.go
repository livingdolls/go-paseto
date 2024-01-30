package repository

import (
	"errors"

	"github.com/livingdolls/go-paseto/internal/core/dto"

	_ "github.com/golang/mock/mockgen/model"
)

var DuplicateUser = errors.New("Duplicate Users")
var NoPersonsFound = errors.New("Users Not Found")

type UserPortRepository interface {
	CreateUser(user *dto.UserDTO) (*dto.UserDTO, error)
}
