package main

import (
	"github.com/gin-gonic/gin"
	"routine.nfitton.com/app"
	"routine.nfitton.com/routines"
	"routine.nfitton.com/users"
)

func setupRouter() *gin.Engine {
	r := gin.Default()


	r.HTMLRender = &app.TemplRender{}
	component := app.Hello("John")
	r.GET("/", app.HtmlResponse(component))
	r.POST("/users", users.HandlePost)
	r.GET("/users", users.HandleGetAll)
	r.GET("/users/:userId", users.HandleGet)
	r.DELETE("/users/:userId", users.HandleDelete)

	r.POST("/users/:userId/routines", routines.HandlePost)
	r.GET("/users/:userId/routines", routines.HandleGetAll)
	r.GET("/users/:userId/routines/:routineId", routines.HandleGetByRoutineId)
	r.DELETE("/users/:userId/routines/:routineId", routines.HandleDeleteByRoutineId)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
