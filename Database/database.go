package database

import (
	"database/sql"
	"fmt"
	"time"
)

func createConnection() (*sql.DB, error) {

	conn, err := sql.Open(DRIVER, USERNAME+":"+PASS+"@tcp("+HOST+":"+PORT+")/"+DB_NAME+"?parseTime=true")

	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)
	conn.SetConnMaxIdleTime(5 * time.Minute)
	conn.SetConnMaxLifetime(60 * time.Minute)

	fmt.Println("Success Connect")

	return conn, nil
}

func GetConnection() (*sql.DB, error) {
	conn, err := createConnection()
	return conn, err
}
