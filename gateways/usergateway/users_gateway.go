package usergateway

import (
	"context"
	"github.com/danielgom/bookstore_usersapi/datasource/postgresql/usersdb"
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/utils/errors"
	"github.com/danielgom/bookstore_usersapi/utils/pgsqlutils"
)

const (
	queryInsertUser = `INSERT INTO
    users_db.users("firstName", "lastName", "email", "dateCreated", "password", "status")
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	queryGetUserById      = `SELECT "id", "firstName", "lastName", "email", "dateCreated" FROM users_db.users WHERE id=$1`
	queryUpdateUser       = `UPDATE users_db.users SET "firstName"=$1, "lastName"=$2, "email"=$3 WHERE id=$4`
	queryDeleteUser       = `DELETE FROM users_db.users WHERE id=$1`
	queryFindUserByStatus = `SELECT "firstName", "lastName", "email", "dateCreated", "status" FROM users_db.users WHERE status=$1`
)

func Get(uId int64) (*users.User, *errors.RestErr) {

	var u users.User

	err := usersdb.Client.QueryRow(context.Background(), queryGetUserById, uId).Scan(&u.Id,
		&u.FirstName, &u.LastName, &u.Email, &u.DateCreated)

	// Error when not found is not a postgres error therefore we change functionality to find only whenever it is a pg err

	if err != nil {
		return nil, pgsqlutils.ParseError(err)
	}

	return &u, nil
}

func Save(u *users.User) *errors.RestErr {

	// By design postgres will always generate a new Id even if the query fails

	err := usersdb.Client.QueryRow(context.Background(),
		queryInsertUser, u.FirstName, u.LastName, u.Email, u.DateCreated, u.Password, u.Status).Scan(&u.Id)

	if err != nil {
		return pgsqlutils.ParseError(err)
	}

	return nil
}

func Update(u *users.User) *errors.RestErr {
	query, err := usersdb.Client.Query(context.Background(),
		queryUpdateUser, u.FirstName, u.LastName, u.Email, u.Id)

	if err != nil {
		return pgsqlutils.ParseError(err)
	}

	defer query.Close()
	return nil
}

func Delete(uId int64) *errors.RestErr {
	query, err := usersdb.Client.Query(context.Background(), queryDeleteUser, uId)
	if err != nil {
		return pgsqlutils.ParseError(err)
	}

	defer query.Close()
	return nil
}

func FindByStatus(s string) ([]users.User, *errors.RestErr) {
	rows, err := usersdb.Client.Query(context.Background(), queryFindUserByStatus, s)
	if err != nil {
		return nil, pgsqlutils.ParseError(err)
	}

	defer rows.Close()

	uList := make([]users.User, 0)
	for rows.Next() {
		var u users.User
		if err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); err != nil {
			return nil, pgsqlutils.ParseError(err)
		}
		uList = append(uList, u)
	}

	return uList, nil
}
