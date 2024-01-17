package response

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIdFromPath(c *gin.Context, name string) (uuid.UUID, error) {
	id := c.Params.ByName(name)
	return uuid.Parse(id)
}

func GenerateErrorResponse(reason string) gin.H {
	return gin.H{"error": gin.H{"reason": reason}}
}

func GenerateErrorResponseWithError(reason string, err error) gin.H {
	return gin.H{"error": gin.H{"reason": reason, "error": err}}
}
