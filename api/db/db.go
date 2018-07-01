package db

import (
	"database/sql"
	"os"
)

// Connect connect mysql
func Connect() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	pass := os.Getenv("MYSQL_PASSWORD")
	name := os.Getenv("MYSQL_DATABASE")
	port := os.Getenv("MYSQL_PORT")

	dbconf := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name
	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		return nil, err
	}
	return db, nil
}
