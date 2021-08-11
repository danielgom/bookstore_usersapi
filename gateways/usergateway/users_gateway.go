package usergateway

import (
	"context"
	"fmt"
	"github.com/danielgom/bookstore_usersapi/datasource/postgresql/usersdb"
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/logger"
	"github.com/danielgom/bookstore_utils-go/errors"
)

const (
	queryInsertUser = `INSERT INTO
    users_db.users("firstName", "lastName", "email", "dateCreated", "password", "status")
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	queryGetUserById  = `SELECT "id", "firstName", "lastName", "email", "dateCreated", "status" FROM users_db.users WHERE id=$1`
	queryUpdateUser   = `UPDATE users_db.users SET "firstName"=$1, "lastName"=$2, "email"=$3 WHERE id=$4`
	queryDeleteUser   = `DELETE FROM users_db.users WHERE id=$1`
	queryFindByStatus = `SELECT "id", "firstName", "lastName", "email", "dateCreated", "status" FROM users_db.users WHERE status=$1`
	queryFindByEmail  = `SELECT "id", "firstName", "lastName", "email", "dateCreated", "status", "password" FROM users_db.users WHERE email=$1 AND status=$2`
)

func Get(uId int64) (*users.User, errors.RestErr) {

	var u users.User

	getErr := usersdb.Client.QueryRow(context.Background(), queryGetUserById, uId).Scan(&u.Id,
		&u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status)

	// Error when not found is not a postgres error therefore we change functionality to find only whenever it is a pg err

	if getErr != nil {
		logger.Error("Error when trying to get user by id", getErr.Error())
		return nil, errors.NewInternalServerError("Database error", getErr)
	}

	return &u, nil
}

func Save(u *users.User) errors.RestErr {

	// By design postgres will always generate a new Id even if the query fails

	saveErr := usersdb.Client.QueryRow(context.Background(),
		queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated, u.Password, u.Status).Scan(&u.Id)

	if saveErr != nil {
		logger.Error("Error when trying to save the user", saveErr.Error())
		return errors.NewInternalServerError("Database error", saveErr)
	}

	return nil
}

func Update(u *users.User) errors.RestErr {
	query, updateErr := usersdb.Client.Query(context.Background(),
		queryUpdateUser, u.FirstName, u.LastName, u.Email, u.Id)

	if updateErr != nil {
		logger.Error("Error when trying to update the user", updateErr.Error())
		return errors.NewInternalServerError("Database error", updateErr)
	}

	defer query.Close()
	return nil
}

func Delete(uId int64) errors.RestErr {
	query, deleteErr := usersdb.Client.Query(context.Background(), queryDeleteUser, uId)
	if deleteErr != nil {
		logger.Error("Error when trying to delete the user", deleteErr.Error())
		return errors.NewInternalServerError("Error when trying to delete the user", deleteErr)
	}

	defer query.Close()
	return nil
}

func FindByStatus(s string) (users.Users, errors.RestErr) {
	rows, findErr := usersdb.Client.Query(context.Background(), queryFindByStatus, s)
	if findErr != nil {
		logger.Error("Error when trying to find the users by status", findErr.Error())
		return nil, errors.NewInternalServerError("Database error", findErr)
	}

	defer rows.Close()

	uList := make([]users.User, 0)
	for rows.Next() {
		var u users.User
		if findErr = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); findErr != nil {
			logger.Error("Error when trying to scan the rows into users struct", findErr.Error())
			return nil, errors.NewInternalServerError("Database error", findErr)
		}
		uList = append(uList, u)
	}

	return uList, nil
}

func FindByEmail(email string) (*users.User, errors.RestErr) {

	var u users.User

	getErr := usersdb.Client.QueryRow(context.Background(), queryFindByEmail, email, users.StatusInactive).Scan(&u.Id,
		&u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status, &u.Password)

	if getErr != nil {
		logger.Error("Error when trying to get user by email and password", getErr.Error())
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with Email %s not found", email))
	}

	return &u, nil
}
