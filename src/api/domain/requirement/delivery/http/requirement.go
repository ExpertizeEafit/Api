package http

import (
	"io"
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
	requirementFile.File.Filename = "requirement"
	err := handler.usecase.UploadRequirement(ctx, requirementFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, requirementFile)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
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

func (handler *RequirementHTTPHandler) HandlerGetPendingRequirements(ctx *gin.Context) {
	requirements, err := handler.usecase.GetPendingRequirements(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusAccepted, requirements)
}

func (handler *RequirementHTTPHandler) HandlerUpdateRequirementStatus(ctx *gin.Context) {
	data := struct {
		id     int    `json:"id"`
		status string `json:"status"`
	}{}
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	status := entities.Status(data.status)
	if !entities.PossibleStatus.Contains(status) {
		ctx.JSON(http.StatusBadRequest, "Invalid status")
		return
	}
	err := handler.usecase.UpdateRequirementStatus(ctx, status, data.id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

func (handler *RequirementHTTPHandler) HandlerDownloadRequirementFile(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	buff, err := handler.usecase.DownloadRequirementFile(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Writer.Header().Set("Content-Type", "application/pdf")
	io.Copy(ctx.Writer, buff)
}
