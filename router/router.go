package router

import (
	"github.com/gin-gonic/gin"

	"app/controllers"
	"app/database"
)

func Run() {
	router := SetupRouter()
	router.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	handler := &controllers.Handler{
		DB: database.GetDB(),
	}

	api := router.group("/api")
	v1 := api.group("/v1")
	addAuthRouter(v1, handler)

	return router
}
