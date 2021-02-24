package pgsqlutils

import (
	errs "errors"
	"fmt"
	"github.com/danielgom/bookstore_usersapi/utils/errors"
	"github.com/jackc/pgconn"
	"strconv"
	"strings"
)

const errorNoRows = "no rows in result set"

func ParseError(err error) *errors.RestErr {
	var pgErr *pgconn.PgError
	if !errs.As(err, &pgErr) {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error parsing database result: %s", err.Error()))
	}

	switch a, _ := strconv.Atoi(pgErr.Code); a {
	case 23505:
		return errors.NewBadRequestError("Invalid data")
	case 23502:
		return errors.NewBadRequestError("Not null error")
	default:
		return errors.NewInternalServerError(fmt.Sprintf("Error processing request: %s", pgErr))
	}
}
