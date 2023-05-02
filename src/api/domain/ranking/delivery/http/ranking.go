package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler *RankingHTTPHandler) HandlerGetRanking(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	ranking, err := handler.usecase.GetRanking(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusAccepted, ranking)
}
