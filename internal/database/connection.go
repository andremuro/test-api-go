package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "changeme"
const DbName = "test_api_go"
const TableName = "test_api_1"

func OpenConnection() (*sql.DB, error) {

	var DataSourceName = fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DbName)

	fmt.Printf("Accessing %s ... ", DbName)

	conn, err := sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	err = conn.Ping()

	return conn, err
}
