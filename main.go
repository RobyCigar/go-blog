package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/authors", Index)
	router.POST("/authors", CreateAuthor)
	router.POST("/posts", CreatePost)

	router.Run("localhost:8080")
}
