package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BasicUser struct {
	id          string `json:"id"`
	email       string
	createdAt   string
	lastLoginAt string
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/users", func(c *gin.Context) {
		user := BasicUser{
			id:          "id",
			email:       "email",
			createdAt:   "createdAt",
			lastLoginAt: "lastLoginAt",
		}
		fmt.Println(user)
		c.JSON(http.StatusOK, gin.H{"data": &user})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
