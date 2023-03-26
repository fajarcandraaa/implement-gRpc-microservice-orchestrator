package app

import (
	"context"
	"database/sql"
	"database/sql/driver"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
)

// Error is a pgsqldb database error.
type Error string

// Error implement error interface.
func (e Error) Error() string {
	return string(e)
}

// Err are known errors for PostgreSQ.
const (
	ErrUniqueViolation     = Error("unique_violation")
	ErrNullValueNotAllowed = Error("null_value_not_allowed")
	ErrorUndefinedTable    = Error("undefined_table")
	ErrNoRowsFound         = Error("no rows found")
)

// canceledMessage is an error that occurs when deadline exceeded.
const canceledMessage = "pq: canceling statement due to user request"

// parseSQLError converts the error from pq driver using postgreSQL error codes.
// https://www.postgresql.org/docs/9.3/errcodes-appendix.html
func ParsePostgreSQLError(err error) error {
	// Parse by value
	switch err {
	case sql.ErrNoRows:
		return ErrNoRowsFound
	case driver.ErrBadConn:
		return context.DeadlineExceeded
	}

	// Parse by type
	switch et := err.(type) {
	case *pq.Error:
		switch et.Code {
		case "23505":
			return ErrUniqueViolation
		case "42P01":
			return ErrorUndefinedTable
		case "22004":
			return ErrNullValueNotAllowed
		}
	}

	// Parse by message
	switch err.Error() {
	case canceledMessage:
		return context.DeadlineExceeded
	}

	return err
}

// parseSQLError converts the error from mysql driver using error codes.
// https://mariadb.com/kb/en/mariadb-error-codes/
func ParseMysqlSQLError(err error) error {
	// Parse by value
	switch err {
	case sql.ErrNoRows:
		return ErrNoRowsFound
	case driver.ErrBadConn:
		return context.DeadlineExceeded
	}

	// Parse by type
	switch et := err.(type) {
	case *mysqlDriver.MySQLError:
		switch et.Number {
		case 1062:
			return ErrUniqueViolation
		case 1138:
			return ErrNullValueNotAllowed
		}
	}

	// Parse by message
	switch err.Error() {
	case canceledMessage:
		return context.DeadlineExceeded
	}

	return err
}