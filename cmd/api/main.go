package main

import (
	"database/sql"
	"fmt"

	dbConfig "github.com/andremuro/test-api-go/postgres/dbconfig"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//Definindo objeto do gin

	fmt.Printf("Accessing %s ... ", dbConfig.DbName)
	db, err = sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}
	defer db.Close()

	g := gin.Default()
	//Definindo rota (endpoint)
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{

			"error":   "null",
			"message": "Hello World!",
		})
	})
	g.Run(":3002")
}
