package users

import (
	"../../utils/errors"
	"strings"
)

type User struct {
	 
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_ceated"`
}

func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == ""{
		message := "Invalid email address"
		return errors.NewBadrequestError(message)
	}
	return nil
}
