package services

import (
	"henry-on-the-internet/go-react-auth-demo/backend/domain/users"

	restErrors "github.com/henry-on-the-internet/go-react-auth-demo/backend/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *restErrors.RestErr) {
	restErr := user.Validate()
	if restErr != nil {
		return nil, restErr
	}

	// encrpyt the password
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, restErrors.NewBadRequestError("failed to encrypt the password")
	}

	user.Password = string(pwSlice[:])
	restErr = user.Save()

	if err != nil {
		return nil, restErrors.NewInternalServerError("failed to save user")
	}
	return &user, nil
}

func GetUser(user users.User) (*users.User, *restErrors.RestErr) {
	result := &users.User{Email: user.Email}
	restErr := result.GetByEmail()
	if restErr != nil {
		return nil, restErr
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return nil, restErrors.NewBadRequestError("failed to decrypt")
	}

	resultWp := &users.User{ID: result.ID, FirstName: result.FirstName, LastName: result.LastName, Email: result.Email}
	return resultWp, nil
}

func GetUserById(userId int64) (*users.User, *restErrors.RestErr) {
	result := &users.User{ID: userId}
	err := result.GetByID()
	if err != nil {
		return nil, err
	}
	return result, nil
}
