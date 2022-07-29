package users

import (
	"strings"

	restErrors "github.com/henry-on-the-internet/go-react-auth-demo/backend/utils/errors"
)

type User struct {
	ID        int64  `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email`
}

func (user *User) Validate() *restErrors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return restErrors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return restErrors.NewBadRequestError("invalid password")
	}
	return nil
}
