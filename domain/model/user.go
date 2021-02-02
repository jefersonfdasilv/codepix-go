package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/badoux/checkmail"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base `valid:"required"`
	Name     string    `json:"name,omitempty" valid:"notnull"`
	Email    string    `json:"email,omitempty" valid:"notnull"`
}

func NewUser(name string, email string) (*User, error) {
	user := User {
		Name: name,
		Email: email,
	}
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}


func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	return nil
}