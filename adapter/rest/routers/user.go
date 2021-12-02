package routers

import (
	"github.com/gin-gonic/gin"
	"go-api-protocols/adapter/rest/controllers"
)

func InitUserRouter(engine *gin.Engine) (err error) {
	controller, err := controllers.GetController()
	if err != nil {
		return
	}
	app := engine.Group("rest")

	user := app.Group("user")
	user.POST("/createUser", controller.User.CreateUser)
	user.GET("/findUser", controller.User.FindUser)
	return
}
