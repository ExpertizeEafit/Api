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
	res, err := handler.usecase.Login(ctx, data.Id, data.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusAccepted, res)
}
