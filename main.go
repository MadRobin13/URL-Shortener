package main

import (
	"fmt"

	// "github.com/go-redis/redis/v8"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "Hey Go URL Shortener",
		})
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server - Error %v", err))
	}
}