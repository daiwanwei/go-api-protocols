package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-api-protocols/business/services"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	FindUser(ctx *gin.Context)
}

type userController struct {
	user services.UserService
}

func NewUserController() (controller UserController, err error) {
	service, err := services.GetService()
	if err != nil {
		return
	}
	return &userController{
		user: service.User,
	}, nil
}

// CreateUser godoc
// @Summary 建立會員
// @Tags user
// @produce application/json
// @Param register body services.CreateUserDto true "租客註冊資料"
// @Success 200 {object}  rest.NonDataResp "成功後返回的值"
// @Router /rest/user/createUser [post]
func (controller *userController) CreateUser(ctx *gin.Context) {

	register := services.CreateUserDto{}
	if err := ctx.ShouldBindJSON(&register); err != nil {
		respond(ctx, err)
		return
	}
	user, err := controller.user.CreateUser(context.Background(), register)
	respondWithData(ctx, user, err)
}

// FindUser godoc
// @Summary 取得會員
// @Tags user
// @produce application/json
// @Param userId query string false "search by userId"
// @Success 200 {object}  rest.DataResp{data=services.UserDto} "成功後返回的值"
// @Router /rest/user/findUser [get]
func (controller *userController) FindUser(ctx *gin.Context) {
	user, err := controller.user.FindUserByID(context.TODO(), ctx.Query("userId"))
	respondWithData(ctx, user, err)
}
