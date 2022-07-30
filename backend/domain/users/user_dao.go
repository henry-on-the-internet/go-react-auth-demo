package users

import (
	"github.com/henry-on-the-internet/go-react-auth-demo/backend/datasource/mysql/users_db"

	restErrors "github.com/henry-on-the-internet/go-react-auth-demo/backend/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users(first_name, last_name, email, password VALUES(?, ?, ?, ?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email=?;"
	queryGetUserByID    = "SELECT id, first_name, last_name, email FROM users WHERE id=?"
)

func (user *User) Save() *restErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return restErrors.NewBadRequestError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if saveErr != nil {
		return restErrors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return restErrors.NewInternalServerError("database error")
	}

	user.ID = userID

	return nil
}

func (user *User) GetByEmail() *restErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return restErrors.NewInternalServerError("invalid email")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if getErr != nil {
		return restErrors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) GetByID() *restErrors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByID)
	if err != nil {
		return restErrors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if getErr != nil {
		return restErrors.NewInternalServerError("database error")
	}
	return nil
}
