package main

import (
	"locally/goback/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDb()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GOBACK server is running",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
