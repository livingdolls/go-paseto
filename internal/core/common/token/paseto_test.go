package token

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/livingdolls/go-paseto/internal/core/entity"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("asefabckgpsoterzlskdferpwcdwflre")

	require.NoError(t, err)

	username := gofakeit.Username()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)

}

func TestExpiredToken(t *testing.T) {
	maker, err := NewPasetoMaker("asefabckgpsoterzlskdferpwcdwflre")

	require.NoError(t, err)

	token, err := maker.CreateToken(gofakeit.Username(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.Error(t, err)
	require.EqualError(t, err, entity.ErrExpiredToken.Error())

	require.Nil(t, payload)

}
