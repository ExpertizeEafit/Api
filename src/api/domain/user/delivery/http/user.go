package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler *UserHTTPHandler) HandlerLogin(ctx *gin.Context) {
	idNumber, err := strconv.Atoi(ctx.Query("idNumber"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	password := ctx.Query("password")
	res, err := handler.usecase.Login(ctx, idNumber, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusAccepted, res)
}
