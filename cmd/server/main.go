package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// NewRouter builds the application HTTP router (Gin engine) with all routes
// registered. Extracted from main() so it can be exercised by tests without
// binding a TCP port.
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	r := NewRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
