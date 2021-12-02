package routers

import (
	"github.com/gin-gonic/gin"
	"go-api-protocols/adapter/rest/controllers"
)

func InitAuthRouter(engine *gin.Engine) (err error) {
	controller, err := controllers.GetController()
	if err != nil {
		return
	}
	app := engine.Group("rest")

	auth := app.Group("auth")
	auth.POST("/login", controller.Auth.Login)
	return
}
