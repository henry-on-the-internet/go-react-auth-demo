package services

import (
	"henry-on-the-internet/go-react-auth-demo/backend/domain/users"
	"henry-on-the-internet/go-react-auth-demo/backend/utils/errors"
	"os/user"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*user.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	// encrpyt the password
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the password")
	}

	user.Password = string(pwSlice[:])
	user.Save()

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}
	err := result.GetByEmail()
	if err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.NewBadRequestError("failed to decrypt")
	}

	resultWp := &users.User{ID: result.ID, FirstName: result.FirstName, LastName: result.LastName, Email: result.Email}
	return resultWp, nil
}

func GetUserById(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	err := result.GetByID()
	if err != nil {
		return nil, err
	}
	return result, nil
}
