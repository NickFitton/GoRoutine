package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": GetAll()})
}

func HandlePost(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Failed to parse request body", "error": err})
		return
	}

	createdUser := Create(user)
	// fmt.Println(createdUser)
	// log.Println(createdUser)
	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}
