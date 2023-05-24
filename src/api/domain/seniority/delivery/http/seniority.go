package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *SeniorityHTTPHandler) HandlerGetSeniorityRequestList(ctx *gin.Context) {
	requests, err := handler.usecase.GetSeniorityRequestList(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, requests)
}

func (handler *SeniorityHTTPHandler) HandlerCreateSeniorityRequest(ctx *gin.Context) {
	data := struct {
		UserID      int64 `json:"user_id"`
		SeniorityID int64 `json:"seniority_id"`
	}{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := handler.usecase.CreateSeniorityRequest(ctx, data.UserID, data.SeniorityID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (handler *SeniorityHTTPHandler) HandlerUpdateSeniorityRequest(ctx *gin.Context) {
	data := struct {
		Id     int64  `json:"id"`
		Status string `json:"status"`
	}{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	status := entities.Status(data.Status)
	if !entities.PossibleStatus.Contains(status) {
		ctx.JSON(http.StatusBadRequest, "Invalid status")
		return
	}
	err := handler.usecase.UpdateStatusSeniorityReques(ctx, data.Id, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}
