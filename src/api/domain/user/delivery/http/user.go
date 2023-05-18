package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *UserHTTPHandler) HandlerLogin(ctx *gin.Context) {
	var data entities.LoginData
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	res, err := handler.usecase.Login(ctx, data.DNI, data.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusAccepted, res)
}

func (handler *UserHTTPHandler) HandlerRegister(ctx *gin.Context) {
	data := [][]string{}
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	err := handler.usecase.Register(ctx, data)
	ctx.JSON(http.StatusCreated, err)
}

func (handler *UserHTTPHandler) HandlerUpdatePassword(ctx *gin.Context) {
	data := entities.UpdatePassword{}
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	err := handler.usecase.UpdatePassword(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}
