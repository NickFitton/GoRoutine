package app

import (
	"github.com/gin-gonic/gin"
	"routine.nfitton.com/app/templates"
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
