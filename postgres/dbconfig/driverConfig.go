package dbconfig

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

const PostgresDriver = "postgres"
const User = "postgres"
const Host = "localhost"
const Port = "5432"
const Password = "changeme"
const DbName = "test_api_go"
const TableName = "test_api_1"

var DataSourceName = fmt.Sprintf(
	"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
	Host, Port, User, Password, DbName)
