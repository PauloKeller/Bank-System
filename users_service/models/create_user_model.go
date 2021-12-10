package models

import (
	"net/mail"
	"strconv"
	"strings"
	"users_service/utils"
)

type CreateUserModel struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func (model *CreateUserModel) IsValid() (bool, string) {
	if !(len(strings.Trim(model.FirstName, "")) > 0) {
		return false, utils.InvalidFirstNameMessage.Value()
	}

	if !(len(strings.Trim(model.LastName, "")) > 0) {
		return false, utils.InvalidLastNameMessage.Value()
	}

	if !(len(strings.Trim(model.Username, "")) > 4) {
		return false, utils.InvalidLastNameMessage.Value()
	}

	_, err := mail.ParseAddress(model.Email)
	if err != nil {
		return false, utils.InvalidEmailMessage.Value()
	}

	if !(len(strings.Trim(model.Password, "")) > 0) || !(len(strings.Trim(model.Password, "")) < 5) {
		return false, utils.InvalidPasswordMessage.Value()
	}

	if _, err := strconv.ParseInt(model.Password, 10, 64); err != nil {
		return false, utils.InvalidPasswordMessage.Value()
	}

	return true, ""
}
