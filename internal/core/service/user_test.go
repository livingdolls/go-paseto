package service

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/entity"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
	mockdb "github.com/livingdolls/go-paseto/internal/infra/mock/repository"
	mocksv "github.com/livingdolls/go-paseto/internal/infra/mock/service"
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
			tokenRepo *mocksv.MockMaker,
		)
		input    createUserInput
		output   createUserOutput
		expected outputService
	}{
		{
			desc: "Success",
			mocks: func(userRepo *mockdb.MockUserPortRepository, tokenMaker *mocksv.MockMaker) {
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
			mocks: func(userRepo *mockdb.MockUserPortRepository, tokenMaker *mocksv.MockMaker) {
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
		{
			desc: "FAIL_DUPLICATED",
			mocks: func(userRepo *mockdb.MockUserPortRepository, tokenMaker *mocksv.MockMaker) {
				userRepo.EXPECT().
					CreateUser(gomock.Any()).
					Times(1).
					Return(nil, entity.ErrConflictingData)
			},
			input: createUserInput{
				user: repoInput,
			},
			output: createUserOutput{
				user: nil,
				err:  entity.ErrConflictingData,
			},
			expected: outputService{
				user: nil,
				err:  entity.ErrConflictingData,
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
			tokenRepo := mocksv.NewMockMaker(ctrl)

			tc.mocks(userRepo, tokenRepo)

			userService := NewUserService(userRepo, tokenRepo)

			users := &request.RegisterUserRequest{
				FullName: tc.input.user.FullName,
				Username: tc.input.user.Username,
				Password: tc.input.user.HashedPassword,
				Email:    tc.input.user.Email,
			}

			user, err := userService.Register(users)
			assert.Equal(t, tc.expected.err, err)
			assert.Equal(t, tc.expected.user, user)

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
			tokenRepo *mocksv.MockMaker,
		)
		expected listUserseExpected
	}{
		{
			desc: "SUCCESS GET USERS",
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenRepo *mocksv.MockMaker) {
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
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenRepo *mocksv.MockMaker) {
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
			tokenRepo := mocksv.NewMockMaker(ctrl)

			tc.mock(userRepo, tokenRepo)

			userService := NewUserService(userRepo, tokenRepo)

			result, err := userService.ListUsers()

			assert.Equal(t, tc.expected.err, err)
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
			tokenRepo *mocksv.MockMaker,
		)
		input    request.GetUserByIdRequest
		expected getUserByIdExpected
	}{
		{
			desc: "SUCCESS",
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenRepo *mocksv.MockMaker) {
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
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenRepo *mocksv.MockMaker) {
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
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenRepo *mocksv.MockMaker) {
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
			tokenRepo := mocksv.NewMockMaker(ctrl)

			tc.mock(userRepo, tokenRepo)

			userService := NewUserService(userRepo, tokenRepo)
			result, err := userService.GetUser(&tc.input)
			assert.Equal(t, tc.expected.err, err)
			assert.Equal(t, tc.expected.user, result)
		})
	}
}

type loginSvResponse struct {
	user *response.LoginUserResponseWithPassword
	err  error
}

type expectedResponse struct {
	user *response.LoginUserResponse
	err  error
}

func TestLogin(t *testing.T) {
	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, true, false, 8)

	token := gofakeit.UUID()

	reqLogin := request.LoginUserRequest{
		Username: username,
		Password: password,
	}

	serviceRespon := &response.LoginUserResponseWithPassword{
		ID:                gofakeit.UUID(),
		Username:          username,
		FullName:          gofakeit.Name(),
		Email:             gofakeit.Email(),
		HashedPassword:    password,
		PasswordChangedAt: time.Time{},
		Role:              dto.Admin,
		IsEmailVerified:   false,
		CreatedAt:         time.Now(),
	}

	// serviceResponMiss := &response.LoginUserResponseWithPassword{
	// 	ID:                gofakeit.UUID(),
	// 	Username:          "Yurina",
	// 	FullName:          gofakeit.Name(),
	// 	Email:             gofakeit.Email(),
	// 	HashedPassword:    password,
	// 	PasswordChangedAt: time.Time{},
	// 	Role:              dto.Admin,
	// 	IsEmailVerified:   false,
	// 	CreatedAt:         time.Now(),
	// }

	userRes := &response.UserResponse{
		ID:                gofakeit.UUID(),
		Username:          username,
		FullName:          gofakeit.Name(),
		Email:             gofakeit.Email(),
		PasswordChangedAt: time.Time{},
		Role:              dto.Admin,
		IsEmailVerified:   false,
		CreatedAt:         time.Now(),
	}

	testCases := []struct {
		desc string
		mock func(
			userRepo *mockdb.MockUserPortRepository,
			tokenSvc *mocksv.MockMaker,
		)
		input    request.LoginUserRequest
		output   loginSvResponse
		expected expectedResponse
	}{
		{
			desc: "SUCCESS",
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenSvc *mocksv.MockMaker) {
				userRepo.EXPECT().Login(gomock.Eq(&reqLogin)).Times(1).Return(serviceRespon, nil)
				tokenSvc.EXPECT().CreateToken(gomock.Eq(serviceRespon.Username), gomock.Any()).Times(1).Return(token, nil)
			},
			input: reqLogin,
			output: loginSvResponse{
				user: serviceRespon,
				err:  nil,
			},
			expected: expectedResponse{
				user: &response.LoginUserResponse{
					User:        *userRes,
					AccessToken: token,
				},
				err: nil,
			},
		},
		{
			desc: "INTERNAL_ERROR",
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenSvc *mocksv.MockMaker) {
				userRepo.EXPECT().Login(gomock.Eq(&reqLogin)).Times(1).Return(nil, entity.ErrInternal)
			},
			input: reqLogin,
			output: loginSvResponse{
				user: nil,
				err:  entity.ErrInternal,
			},
			expected: expectedResponse{
				user: &response.LoginUserResponse{
					User:        *userRes,
					AccessToken: "",
				},
				err: entity.ErrInternal,
			},
		},
		{
			desc: "USER_NOT_FOUND",
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenSvc *mocksv.MockMaker) {
				userRepo.EXPECT().Login(gomock.Eq(&reqLogin)).Times(1).Return(nil, entity.ErrDataNotFound)
			},
			input: reqLogin,
			output: loginSvResponse{
				user: nil,
				err:  entity.ErrDataNotFound,
			},
			expected: expectedResponse{
				user: &response.LoginUserResponse{
					User:        *userRes,
					AccessToken: "",
				},
				err: entity.ErrDataNotFound,
			},
		},
		{
			desc: "FAIL_TOKEN_CREATION",
			mock: func(userRepo *mockdb.MockUserPortRepository, tokenSvc *mocksv.MockMaker) {
				userRepo.EXPECT().Login(gomock.Eq(&reqLogin)).Times(1).Return(serviceRespon, nil)
				tokenSvc.EXPECT().CreateToken(gomock.Eq(reqLogin.Username), time.Duration(15)).Times(1).Return("", entity.ErrTokenCreation)
			},
			input: reqLogin,
			output: loginSvResponse{
				user: nil,
				err:  entity.ErrTokenCreation,
			},
			expected: expectedResponse{
				user: &response.LoginUserResponse{
					User:        *userRes,
					AccessToken: "",
				},
				err: entity.ErrTokenCreation,
			},
		},
		// {
		// 	desc: "FAIL_PASSWORD_MISSMATCH",
		// 	mock: func(userRepo *mockdb.MockUserPortRepository, tokenSvc *mocksv.MockMaker) {
		// 		userRepo.EXPECT().Login(gomock.Eq(&reqLogin)).Times(1).Return(serviceResponMiss, nil)
		// 	},
		// 	input: reqLogin,
		// 	output: loginSvResponse{
		// 		user: serviceResponMiss,
		// 		err:  nil,
		// 	},
		// 	expected: expectedResponse{
		// 		user: &response.LoginUserResponse{
		// 			User:        *userRes,
		// 			AccessToken: "",
		// 		},
		// 		err: entity.ErrNoMatchPassword,
		// 	},
		// },
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mockdb.NewMockUserPortRepository(ctrl)
			tokenSv := mocksv.NewMockMaker(ctrl)

			tc.mock(userRepo, tokenSv)

			userService := NewUserService(userRepo, tokenSv)

			token, err := userService.Login(&tc.input)
			if err != tc.expected.err {
				t.Errorf("[case: %s] expected to get %q; got %q", tc.desc, tc.expected.err, err)
			}

			if tc.desc == "SUCCESS" {
				if token.AccessToken != tc.expected.user.AccessToken {
					t.Errorf("[case: %s] expected to get %q; got %q", tc.desc, tc.expected.user.AccessToken, token.AccessToken)
				}
			}
		})
	}

}
