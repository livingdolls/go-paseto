package repository

import (
	"context"

	"github.com/livingdolls/go-paseto/internal/core/dto"
	"github.com/livingdolls/go-paseto/internal/core/entity"
	"github.com/livingdolls/go-paseto/internal/core/model/request"
	"github.com/livingdolls/go-paseto/internal/core/model/response"
	"github.com/livingdolls/go-paseto/internal/core/port/repository"
)

type userRepository struct {
	db repository.Database
}

func NewUsersRepository(db repository.Database) repository.UserPortRepository {
	return &userRepository{
		db: db,
	}
}

const (
	createUser = "INSERT INTO Users (" +
		"`id`," +
		"`username`," +
		"`fullname`," +
		"`email`," +
		"`isemailverified`," +
		"`hashedpassword`," +
		"`role`," +
		"`passwordchangedat`," +
		"`createdat`" +
		") VALUES (?,?,?,?,?,?,?,?,?)"

	getUser     = "SELECT fullname,username,email,createdat FROM Users"
	getUserById = "SELECT id, username, fullname, email, isemailverified, role, passwordchangedat, createdat FROM Users WHERE id = ?"
	loginUser   = "SELECT id, username, fullname, email, hashedpassword, isemailverified, role, passwordchangedat, createdat FROM Users WHERE username = ?"
)

// GetListUser implements repository.UserPortRepository.
func (u *userRepository) GetListUser() (*[]response.RegisterUserResponse, error) {
	res, err := u.db.GetDB().QueryContext(context.Background(), getUser)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	users := []response.RegisterUserResponse{}

	for res.Next() {
		var i response.RegisterUserResponse

		if err := res.Scan(
			&i.FullName,
			&i.Username,
			&i.Email,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, i)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

// CreateUser implements repository.UserPortRepository.
func (u *userRepository) CreateUser(user *dto.UserDTO) (*dto.UserDTO, error) {
	res, err := u.db.GetDB().Exec(
		createUser,
		user.ID,
		user.Username,
		user.FullName,
		user.Email,
		user.IsEmailVerified,
		user.HashedPassword,
		user.Role,
		user.PasswordChangedAt,
		user.CreatedAt,
	)

	if err != nil {
		if entity.IsUniqueConstraintViolationError(err) {
			return nil, entity.ErrConflictingData
		}

		return nil, err
	}

	numbRow, err := res.RowsAffected()

	if err != nil {
		return nil, err
	}

	if numbRow != 1 {
		return nil, entity.ErrNoUpdatedData
	}

	return user, nil
}

// GetUserById implements repository.UserPortRepository.
func (u *userRepository) GetUserById(id *request.GetUserByIdRequest) (response.GetUserByIdResponse, error) {
	var user response.GetUserByIdResponse
	row := u.db.GetDB().QueryRow(getUserById, id.ID)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.IsEmailVerified,
		&user.Role,
		&user.PasswordChangedAt,
		&user.CreatedAt,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return user, entity.ErrDataNotFound
		}
		return user, err
	}

	return user, err
}

// Login implements repository.UserPortRepository.
func (u *userRepository) Login(user *request.LoginUserRequest) (*response.LoginUserResponseWithPassword, error) {
	var userRes response.LoginUserResponseWithPassword

	row := u.db.GetDB().QueryRow(loginUser, user.Username)

	err := row.Scan(
		&userRes.ID,
		&userRes.Username,
		&userRes.FullName,
		&userRes.Email,
		&userRes.HashedPassword,
		&userRes.IsEmailVerified,
		&userRes.Role,
		&userRes.PasswordChangedAt,
		&userRes.CreatedAt,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, entity.ErrDataNotFound
		}

		return nil, err
	}

	if user.Password != userRes.HashedPassword {
		return nil, entity.ErrNoMatchPassword
	}

	return &userRes, nil
}
