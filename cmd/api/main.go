package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Definindo objeto do gin
	g := gin.Default()
	//Definindo rota (endpoint)
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{

			"error":   "null",
			"message": "Hello World!",
		})
	})
	g.Run(":3003")
}
