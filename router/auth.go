package router

import (
	"github.com/gin-gonic/gin"

	"app/controllers"
)

func addAuthRouter(rg *gin.RouterGroup, h *controllers.Handler) {
	auth := rg.group("/auth")
	{
		auth.POST("/signup", handler.SignUpHandler)
	}
}
