package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// http://example.com/hello GET -> JSON / Status Code
	// {"key": value} => {"greeting": "Hello Server"}
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello Server"})
	})

	r.Run()
}
