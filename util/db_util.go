package util

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsDuplicate(err error) bool {
	var myErr *mysql.MySQLError
	if errors.As(err, &myErr) {
		return myErr.Number == 1062
	}
	return false
}

func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
