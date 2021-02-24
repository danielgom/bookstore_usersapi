package usersdb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const (
	postgresqlUsersUsername = "POSTGRES_DB_USERNAME"
	postgresqlUsersPassword = "POSTGRES_DB_PASSWORD"
	postgresqlUsersHost     = "localhost"
	postgresqlUsersName     = "POSTGRES_DB_NAME"
	postgresqlUsersSchema   = "POSTGRES_DB_SCHEMA"
)

var (
	Client *pgxpool.Pool

	host     = postgresqlUsersHost
	port     = 5433
	user     = os.Getenv(postgresqlUsersUsername)
	password = os.Getenv(postgresqlUsersPassword)
	dbname   = os.Getenv(postgresqlUsersName)
	schema   = os.Getenv(postgresqlUsersSchema)
)

func init() {

	var err error
	Client, err = pgxpool.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=%s&sslmode=disable", user, password, host, port, dbname, schema))
	if err != nil {
		panic(err)
	}

}
