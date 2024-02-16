package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	pass := "fsc0ciety"

	hashed, err := HashPassword(pass)

	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	fmt.Println(hashed)

	err = CheckPassword(pass, hashed)

	require.NoError(t, err)

	wrongPass := "helloWorld"

	err = CheckPassword(wrongPass, hashed)

	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	pass2 := "myyurina"

	hashed2, err := HashPassword(pass2)

	require.NoError(t, err)
	require.NotEmpty(t, hashed2)
	require.NotEqual(t, hashed, hashed2)
}
