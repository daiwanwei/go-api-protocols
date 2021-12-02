package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	routerInstance *gin.Engine
)

func GetRouter() (instance *gin.Engine, err error) {
	if routerInstance == nil {
		instance, err = newRouter()
		if err != nil {
			return nil, err
		}
		routerInstance = instance
	}
	return routerInstance, nil
}

func newRouter() (router *gin.Engine, err error) {
	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = InitUserRouter(engine)
	if err != nil {
		return
	}
	err = InitAuthRouter(engine)
	if err != nil {
		return
	}
	return engine, nil
}
