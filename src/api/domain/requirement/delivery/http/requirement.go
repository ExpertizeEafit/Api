package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

func (handler *RequirementHTTPHandler) HandlerUploadFile(ctx *gin.Context) {
	var requirementFile entities.RequirementFile
	if err := ctx.ShouldBind(&requirementFile); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	fileType := requirementFile.File.Header.Get("Content-Type")
	if fileType == "application/pdf" {
		ctx.JSON(http.StatusAccepted, requirementFile.File)
		return
	}
	ctx.JSON(http.StatusBadRequest, fileType+" no accepted")
	//err := ctx.SaveUploadedFile(file.File, "assets/"+file.File.Filename)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	file, _ := requirementFile.File.Open()
	defer file.Close()

}
