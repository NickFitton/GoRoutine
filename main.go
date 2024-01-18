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
	r.GET("/", app.HandleRoot)

	r.POST("/api/users", users.HandlePost)
	r.GET("/api/users", users.HandleGetAll)
	r.GET("/api/users/:userId", users.HandleGet)
	r.DELETE("/api/users/:userId", users.HandleDelete)

	r.POST("/api/users/:userId/routines", routines.HandlePost)
	r.GET("/api/users/:userId/routines", routines.HandleGetAll)
	r.GET("/api/users/:userId/routines/:routineId", routines.HandleGetByRoutineId)
	r.DELETE("/api/users/:userId/routines/:routineId", routines.HandleDeleteByRoutineId)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
