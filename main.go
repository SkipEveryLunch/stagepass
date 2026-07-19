package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, stagepass",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
