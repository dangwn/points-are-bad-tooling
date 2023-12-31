package pabsqldriver

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type SqlDriver struct {
	DB *sql.DB
}

func NewSqlDriverV2(user, password, host, port, db string) (*SqlDriver, error) {
	url := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + db + "?sslmode=disable"
	if database, err := sql.Open("postgres", url); err != nil {
		return nil, err
	} else {
		return &SqlDriver{DB: database}, nil
	}
}

func NewSqlDriver(user string, password string, db string) (*SqlDriver, error) {
	if database, err := sql.Open(
		"postgres",
		"user=" + user + " password=" + password + " dbname=" + db + " sslmode=disable",
	); err != nil {
		return nil, err
	} else {
		return &SqlDriver{DB: database}, nil
	}
}

func (d *SqlDriver) Exec(statement string, args ...any) (sql.Result, error) {
	return d.DB.Exec(statement, args...)
}

func (d *SqlDriver) Query(queryStatement string, args ...any) (*sql.Rows, error) {
	return d.DB.Query(queryStatement, args...)
}

func (d *SqlDriver) QueryRow(queryStatement string, args ...any) *sql.Row {
	return d.DB.QueryRow(queryStatement, args...)
}

func (d *SqlDriver) ValueExists(table string, column string, value any) (bool, error) {
	var exists bool 
	
	if err := d.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM " + table + " WHERE " + column + " = $1)",
		value,
	).Scan(&exists); err != nil {
		return false, err
	}
	
	return exists, nil
}