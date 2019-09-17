package dto

import (
	"errors"
	"reflect"
	"strings"

	"gopkg.in/validator.v2"
)

func validateEmail(v interface{}, param string) error {
	str := reflect.ValueOf(v).String()
	if !strings.Contains(str, "@") {
		return errors.New("the value typed isn't an email")
	}
	return nil
}

func validatePasswordAndUser(v interface{}, param string) error {
	str := reflect.ValueOf(v).String()
	if strings.Contains(str, " ") {
		return errors.New("the value can't have any white spaces")
	}
	return nil
}

func main() {
	validator.SetValidationFunc("validateEmail", validateEmail)
	validator.SetValidationFunc("validatePasswordAndUser", validatePasswordAndUser)
}

type UserCreateDto struct {
	ID       int
	Username string `validate:"min=3,max=30,regexp=[a-zA-Z]-_,nonzero,validatePasswordAndUser"`
	Password string `validate:"min=8,max=15,nonzero,regexp=\S+,validatePasswordAndUser"`
	Email    string `validate:"min=8,max=15,nonzero,regexp=[a-zA-Z@.]+\S+,validateEmail"`
}

type UserDto struct {
	ID       int
	Username string `validate:"min=3,max=30,regexp=[a-zA-Z]-_,nonzero,validatePasswordAndUser"`
	Password string `validate:"min=8,max=15,nonzero,validatePasswordAndUser"`
	Email    string `validate:"min=8,max=15,nonzero,regexp=[a-zA-Z@.]+\S+,validateEmail"`
	// [a-zA-Z@.] - matches any characters between a-z or A-Z, @ or .
	//\S+ don't accept white spaces
}
