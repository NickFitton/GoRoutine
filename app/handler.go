package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"routine.nfitton.com/app/templates"
	"routine.nfitton.com/response"
	"routine.nfitton.com/routines"
	"routine.nfitton.com/users"
)

func HandleRoot(c *gin.Context) {
	metadata := templates.Metadata{
		Title:       "GoRoutines",
		Description: "A routine management project written in Golang",
	}
	users := users.GetAll()
	content := templates.Root(users)
	component := templates.Wrapper(metadata, content)
	HtmlResponse(component)(c)
}

func HandleRoutines(c *gin.Context) {
	metadata := templates.Metadata{
		Title:       "Your routines",
		Description: "A routine management project written in Golang",
	}
	userId, err := response.GetIdFromPath(c, "userId")
	if err != nil {
		c.HTML(http.StatusNotFound, "", templates.Wrapper(metadata, templates.NotFound("User ID invalid", err)))
		return
	}

	user, found := users.FindById(userId)
	if !found {
		c.HTML(http.StatusNotFound, "", templates.Wrapper(metadata, templates.NotFound("User not found", err)))
		return
	}

	userRoutines := routines.GetAllByUser(userId, false)
	component := templates.Routines(user, userRoutines)
	HtmlResponse(component)(c)
}
