package models

import (
	"github.com/andremuro/test-api-go/internal/database"
)

func Insert(post Post) (id int64, err error) {

	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO test_api_1 (username, body, created_at) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(sql, post.Username, post.Body, post.CreatedAt).Scan(&id)

	return
}
