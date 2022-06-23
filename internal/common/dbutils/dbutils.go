package dbutils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
)

func IsUniqueViolation(err error) bool {
	switch v := err.(type) {
	case *mysql.MySQLError:
		return v.Number == MYSQL_UNIQUE_VIOLATION_NUMBER
	case sqlite3.Error:
		return v.Code == sqlite3.ErrConstraint && (v.ExtendedCode == sqlite3.ErrConstraintUnique || v.ExtendedCode == sqlite3.ErrConstraintPrimaryKey)
	default:
		return false
	}
}
