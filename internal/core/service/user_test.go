package service

import (
	"testing"

	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/port/repository"

	"github.com/stretchr/testify/require"
)

type mockUserRepository struct{}

func (m *mockUserRepository) CreateUser(user *dto.UserDTO) (*dto.UserDTO, error) {
	if user.FullName == "yurina" {
		return nil, repository.DuplicateUser
	}

	return user, nil
}

func TestUserService_Register_Success(t *testing.T) {
	userRepo := &mockUserRepository{}

	userService := NewUserService(userRepo)

	req := &request.RegisterUserRequest{
		Username: "yurinaHirate",
		Password: "myyurina",
		FullName: "Yurina Hirate",
		Email:    "yurina@gmail.com",
	}

	res, err := userService.Register(req)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, req.Email, res.Email)
	require.Equal(t, req.FullName, res.FullName)
	require.Equal(t, req.Username, res.Username)

	require.NotZero(t, res.CreatedAt)
}
