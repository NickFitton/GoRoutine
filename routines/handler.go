package routines

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"routine.nfitton.com/response"
)

func HandlePost(c *gin.Context) {
	userId, userIdError := response.GetIdFromPath(c, "userId")
	if userIdError != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("passed userId was not a uuid", userIdError))
		return
	}

	var routine Routine
	// Does this need an '&'
	bindErr := c.ShouldBindJSON(&routine)
	if bindErr != nil {
		log.Println(bindErr)
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("Failed to parse request body", bindErr))
		return
	}

	createdRoutine := Create(routine, userId)

	c.JSON(http.StatusCreated, gin.H{"data": createdRoutine})
}


func HandleGetAll(c *gin.Context) {
	userId, userIdError := response.GetIdFromPath(c, "userId")
	if userIdError != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("passed userId was not a uuid", userIdError))
		return
	}

	routines := GetAllByUser(userId, false)

	c.JSON(http.StatusOK, gin.H{"data": routines})
}

func HandleGetByRoutineId(c *gin.Context) {
	userId, userIdError := response.GetIdFromPath(c, "userId")
	if userIdError != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("passed userId was not a uuid", userIdError))
		return
	}

	routineId, rPathErr := response.GetIdFromPath(c, "routineId")
	if rPathErr != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("passed userId was not a uuid", rPathErr))
		return
	}

	routine, found := FindByUserIdAndId(userId, routineId)
	if !found {
		c.JSON(http.StatusNotFound, response.GenerateErrorResponse(fmt.Sprintf("Routine not found with given id [%s]", userId)))
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": routine})
}

func HandleDeleteByRoutineId(c *gin.Context) {
	userId, userIdError := response.GetIdFromPath(c, "userId")
	if userIdError != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("passed userId was not a uuid", userIdError))
		return
	}

	routineId, rPathErr := response.GetIdFromPath(c, "routineId")
	if rPathErr != nil {
		c.JSON(http.StatusBadRequest, response.GenerateErrorResponseWithError("passed userId was not a uuid", rPathErr))
		return
	}

	ok := DeleteByUserIdAndId(userId, routineId)
	if !ok {
		c.JSON(http.StatusNotFound, response.GenerateErrorResponse(fmt.Sprintf("User not found with given id [%s]", userId)))
		return
	}
	c.JSON(http.StatusAccepted, gin.H{})
}
