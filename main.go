package main

import (
	"github.com/gin-gonic/gin"
	"routine.nfitton.com/users"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", users.HandleGetAll)
	r.POST("/users", users.HandlePost)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
