package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(c *gin.Context) {
	// H is a shortcut for map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
