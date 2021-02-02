package model_test

import (
	"fmt"
	"os"
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"

	"github.com/jefersonfdasilv/codepix-go/domain/model"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Coisa Amarela"
	bank, err := model.NewBank(code, name)
	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)
	kind := "email"
	key := "j@j.com"
	pixKey, err := model.NewPixKey(kind, account, key)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(kind, account, key)
	require.Nil(t, err)

	_, err = model.NewPixKey("nome", account, key)
	require.NotNil(t, err)
}
