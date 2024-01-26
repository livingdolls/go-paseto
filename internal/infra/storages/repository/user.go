package repository

import (
	"gopaseto/internal/core/dto"
	"gopaseto/internal/core/entity"
	"gopaseto/internal/core/port/repository"
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
)

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
