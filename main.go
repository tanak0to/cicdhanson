package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, makeGreeting(name))
	})
	router.Run(":8080")
	// test
}

func makeGreeting(name string) string {
	// Create a greeting message
	return fmt.Sprintf("Hello, %s", name)

}
