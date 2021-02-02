package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"

	"github.com/jefersonfdasilv/codepix-go/domain/model"
)

func TestModel_NewUser(t *testing.T) {
	email := "jeferson@teste.com.br"
	name := "jeferson"
	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser(name, "jeferson.com")
	require.NotNil(t, err)
}