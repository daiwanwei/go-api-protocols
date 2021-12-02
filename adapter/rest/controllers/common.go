package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api-protocols/adapter/rest"
	"go-api-protocols/utils"
	"go-api-protocols/utils/page"
	"net/http"
	"strconv"
)

func getPageFromQuery(ctx *gin.Context) (pageable *page.Pageable, err error) {
	p, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	if err != nil {
		return
	}
	size, err := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		return
	}
	var sort map[string]int
	if sortType := ctx.Query("sort"); len(sortType) != 0 {
		order, err := strconv.Atoi(ctx.DefaultQuery("order", "1"))
		if err != nil {
			return nil, err
		}
		sort = map[string]int{
			sortType: order,
		}
	}
	return &page.Pageable{
		Size: size, Page: p, Sort: sort,
	}, nil
}

func respondWithData(ctx *gin.Context, data interface{}, err error) {
	if err != nil {
		if e, ok := err.(utils.CustomError); ok {
			ctx.JSON(http.StatusOK, rest.DataResp{Code: e.GetCode(), Msg: e.GetMsg(), Data: nil})
			return
		}
		ctx.JSON(http.StatusOK, rest.DataResp{Code: 500, Msg: err.Error(), Data: nil})
		return
	} else {
		ctx.JSON(http.StatusOK, rest.DataResp{Code: 200, Msg: "OK", Data: data})
	}
}

func respond(ctx *gin.Context, err error) {
	if err != nil {
		if e, ok := err.(utils.CustomError); ok {
			ctx.JSON(http.StatusOK, rest.NonDataResp{Code: e.GetCode(), Msg: e.GetMsg()})
			return
		}
		ctx.JSON(http.StatusOK, rest.NonDataResp{Code: 500, Msg: err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, rest.NonDataResp{Code: 200, Msg: "OK"})
	}
}
