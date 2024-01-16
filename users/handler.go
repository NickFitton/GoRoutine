package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleGetAll(c *gin.Context) {
	user := User{
		// id: uuid.New(),
		email:       "nick.fitton@ovo.com",
		createdAt:   time.Now(),
		lastLoginAt: time.Now(),
	}
	c.JSON(http.StatusOK, gin.H{"data": []User{user}})
}

func HandlePost(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Failed to parse request body", "error": err})
		return
	}

	// createdUser := Create(user)
	// fmt.Println(createdUser)
	// log.Println(createdUser)
	// c.JSON(http.StatusCreated, gin.H{"data": createdUser, "other": user})
	c.JSON(http.StatusCreated, gin.H{"other": user})
}
