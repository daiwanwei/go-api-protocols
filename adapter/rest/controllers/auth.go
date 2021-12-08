package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-api-protocols/business/services"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	auth services.AuthService
}

func NewAuthController() (controller AuthController, err error) {
	service, err := services.GetService()
	if err != nil {
		return
	}
	return &authController{service.Auth}, nil
}

// Login godoc
// @Summary 登入
// @Tags auth
// @produce application/json
// @Param login body services.LoginDto true "登入Dto"
// @Success 200 {object}  rest.DataResp{data=services.PassportDto} "成功後返回的值"
// @router /rest/auth/login [post]
func (controller *authController) Login(ctx *gin.Context) {
	login := services.LoginDto{}
	if err := ctx.ShouldBindJSON(&login); err != nil {
		respondWithData(ctx, nil, err)
		return
	}
	passportDto, err := controller.auth.Login(context.TODO(), login)
	respondWithData(ctx, passportDto, err)

}
