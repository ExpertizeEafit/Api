package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

func (handler *RequirementHTTPHandler) HandlerUploadRequirement(ctx *gin.Context) {
	var requirementFile entities.RequirementFile
	if err := ctx.ShouldBind(&requirementFile); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	fileType := requirementFile.File.Header.Get("Content-Type")
	if fileType != "application/pdf" {
		ctx.JSON(http.StatusBadRequest, fileType+" no accepted")
		return
	}
	err := handler.usecase.UploadRequirement(ctx, requirementFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, requirementFile)
		return
	}
	ctx.JSON(http.StatusCreated, requirementFile)
}

func (handler *RequirementHTTPHandler) HandlerGetRequirementHistory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	history, err := handler.usecase.GetRequirementHistory(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, history)
}
