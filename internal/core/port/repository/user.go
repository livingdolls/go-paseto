package repository

import (
	"errors"
	"gopaseto/internal/core/dto"
)

var DuplicateUser = errors.New("Duplicate Users")
var NoPersonsFound = errors.New("Users Not Found")

type UserPortRepository interface {
	CreateUser(user *dto.UserDTO) (*dto.UserDTO, error)
}
