package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"routine.nfitton.com/response"
)

func HandleGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": GetAll()})
}

func HandleGet(c *gin.Context) {
	userId, err := response.GetIdFromPath(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": "passed id was not a uuid"}})
		return
	}
	user, found := FindById(userId)
	if !found {
		c.JSON(http.StatusNotFound, response.GenerateErrorResponse(fmt.Sprintf("User not found with given id [%s]", userId)))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func HandlePost(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("Failed to parse request body", err))
		return
	}

	createdUser := Create(user)
	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

func HandleDelete(c *gin.Context) {
	userId, err := response.GetIdFromPath(c, "userId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": "passed id was not a uuid"}})
		return
	}

	ok := Delete(userId)
	if !ok {
		c.JSON(http.StatusNotFound, response.GenerateErrorResponse(fmt.Sprintf("User not found with given id [%s]", userId)))
		return
	}
	c.JSON(http.StatusAccepted, gin.H{})
}


