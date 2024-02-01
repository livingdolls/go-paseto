package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/entity"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
	mockdb "github.com/livingdolls/go-paseto/internal/infra/mock/repository"
	"github.com/stretchr/testify/assert"
)

type createUserInput struct {
	user *dto.UserDTO
}

type createUserOutput struct {
	user *dto.UserDTO
	err  error
}

type createUserService struct {
	user *request.RegisterUserRequest
}

type outputService struct {
	user *response.RegisterUserResponse
	err  error
}

func TestCreateUser(t *testing.T) {
	username := gofakeit.Username()
	fullname := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, false, false, 8)
	newuuid := uuid.New()

	repoInput := &dto.UserDTO{
		Username:          username,
		FullName:          fullname,
		Email:             email,
		HashedPassword:    password,
		ID:                newuuid,
		Role:              dto.Admin,
		PasswordChangedAt: time.Time{},
		CreatedAt:         time.Time{},
		IsEmailVerified:   false,
	}

	repoOutput := &dto.UserDTO{
		Username:          repoInput.Username,
		FullName:          repoInput.FullName,
		Email:             repoInput.Email,
		HashedPassword:    repoInput.HashedPassword,
		ID:                newuuid,
		Role:              repoInput.Role,
		PasswordChangedAt: repoInput.PasswordChangedAt,
		CreatedAt:         repoInput.CreatedAt,
		IsEmailVerified:   repoInput.IsEmailVerified,
	}

	userOutput := &response.RegisterUserResponse{
		Username:  repoOutput.Username,
		FullName:  repoOutput.FullName,
		Email:     repoOutput.Email,
		CreatedAt: repoOutput.CreatedAt,
	}

	testCases := []struct {
		desc  string
		mocks func(
			userRepo *mockdb.MockUserPortRepository,
		)
		input    createUserInput
		output   createUserOutput
		expected outputService
	}{
		{
			desc: "Success",
			mocks: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					CreateUser(gomock.Any()).
					Times(1).
					Return(repoOutput, nil)
			},
			input: createUserInput{
				user: repoInput,
			},
			output: createUserOutput{
				user: repoOutput,
				err:  nil,
			},
			expected: outputService{
				user: userOutput,
				err:  nil,
			},
		},
		{
			desc: "FAIL_INTERNALERROR",
			mocks: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					CreateUser(gomock.Any()).
					Times(1).
					Return(nil, entity.ErrInternal)
			},
			input: createUserInput{
				user: repoInput,
			},
			output: createUserOutput{
				user: repoOutput,
				err:  nil,
			},
			expected: outputService{
				user: nil,
				err:  entity.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mockdb.NewMockUserPortRepository(ctrl)

			tc.mocks(userRepo)

			userService := NewUserService(userRepo)

			users := &request.RegisterUserRequest{
				FullName: tc.input.user.FullName,
				Username: tc.input.user.Username,
				Password: tc.input.user.HashedPassword,
				Email:    tc.input.user.Email,
			}

			user, err := userService.Register(users)
			assert.Equal(t, tc.expected.err, err, "Error Mismatch")
			assert.Equal(t, tc.expected.user, user, "User missmatch")

		})
	}

}

type listUserseExpected struct {
	user *[]response.RegisterUserResponse
	err  error
}

func TestGetList(t *testing.T) {
	var users []response.RegisterUserResponse

	for i := 0; i < 1; i++ {
		users = append(users, response.RegisterUserResponse{
			Username:  gofakeit.Username(),
			FullName:  gofakeit.Name(),
			Email:     gofakeit.Email(),
			CreatedAt: time.Now(),
		})
	}

	testCases := []struct {
		desc string
		mock func(
			userRepo *mockdb.MockUserPortRepository,
		)
		expected listUserseExpected
	}{
		{
			desc: "SUCCESS GET USERS",
			mock: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					GetListUser().
					Times(1).
					Return(&users, nil)
			},
			expected: listUserseExpected{
				user: &users,
				err:  nil,
			},
		},
		{
			desc: "FAIL INTERNAL ERROR",
			mock: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					GetListUser().
					Times(1).
					Return(nil, entity.ErrInternal)
			},
			expected: listUserseExpected{
				user: nil,
				err:  entity.ErrInternal,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mockdb.NewMockUserPortRepository(ctrl)

			tc.mock(userRepo)

			userService := NewUserService(userRepo)

			result, err := userService.ListUsers()

			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.user, result)

		})
	}

}

type getUserById struct {
	ID string
}

type getUserByIdExpected struct {
	user *response.GetUserByIdResponse
	err  error
}

// type requestId *request.GetUserByIdRequest

func TestGetUserById(t *testing.T) {
	id := gofakeit.UUID()

	requ := &request.GetUserByIdRequest{
		ID: id,
	}

	user := response.GetUserByIdResponse{
		ID:                id,
		Username:          gofakeit.Username(),
		FullName:          gofakeit.Name(),
		Email:             gofakeit.Email(),
		IsEmailVerified:   false,
		Role:              dto.Admin,
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}

	userNil := response.GetUserByIdResponse{}

	testCases := []struct {
		desc string
		mock func(
			userRepo *mockdb.MockUserPortRepository,
		)
		input    request.GetUserByIdRequest
		expected getUserByIdExpected
	}{
		{
			desc: "SUCCESS",
			mock: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					GetUserById(gomock.Eq(requ)).
					Times(1).
					Return(user, nil)
			},
			input: request.GetUserByIdRequest{
				ID: id,
			},
			expected: getUserByIdExpected{
				user: &user,
				err:  nil,
			},
		},
		{
			desc: "INTERNAL ERROR",
			mock: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					GetUserById(gomock.Eq(requ)).
					Times(1).
					Return(userNil, entity.ErrInternal)
			},
			input: request.GetUserByIdRequest{
				ID: id,
			},
			expected: getUserByIdExpected{
				user: nil,
				err:  entity.ErrInternal,
			},
		},
		{
			desc: "NOTFOUND",
			mock: func(userRepo *mockdb.MockUserPortRepository) {
				userRepo.EXPECT().
					GetUserById(gomock.Eq(requ)).
					Times(1).
					Return(userNil, entity.ErrDataNotFound)
			},
			input: request.GetUserByIdRequest{
				ID: id,
			},
			expected: getUserByIdExpected{
				user: nil,
				err:  entity.ErrDataNotFound,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mockdb.NewMockUserPortRepository(ctrl)

			tc.mock(userRepo)

			userService := NewUserService(userRepo)
			result, err := userService.GetUser(&tc.input)
			fmt.Println("Result >>")
			fmt.Println(result)
			assert.Equal(t, tc.expected.err, err, "Error mismatch")
			assert.Equal(t, tc.expected.user, result)
		})
	}
}
