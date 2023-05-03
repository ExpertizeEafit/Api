package http

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (handler *PathHTTPHandler) HandleGetPaths(ctx *gin.Context) {
	//path := entities.Path{
	//	1: {
	//		Name:        "Mid Backend",
	//		Description: "Test description",
	//		PriorTo:     []int{2},
	//	},
	//	2: {
	//		Name:        "Sr Backend",
	//		Description: "Test description",
	//		PriorTo:     []int{3, 4},
	//	},
	//	3: {
	//		Name:        "Project Leader",
	//		Description: "Test description",
	//		PriorTo:     []int{},
	//	},
	//	4: {
	//		Name:        "Technical Leader",
	//		Description: "Test description",
	//		PriorTo:     []int{},
	//	},
	//}
	path, err := handler.usecase.GetAllSeniority(ctx)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, path)
}

func (handler *PathHTTPHandler) HandlerGetCurrentAndNextSeniority(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	path, err := handler.usecase.GetCurrentAndNextSeniority(ctx, id)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, path)
}
