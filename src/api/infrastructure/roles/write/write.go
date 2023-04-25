package write

import (
	"net/http"

	"github.com/gin-gonic/gin"

	requirementHTTP "github.com/ExpertizeEafit/Api/src/api/domain/requirement/delivery/http"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

// Write is main structure of consumer functionality
type Write struct {
	container *dependencies.Container
}

// NewRead returns a concrete implementation of the Read scope
func NewWrite(container *dependencies.Container) *Write {
	return &Write{
		container: container,
	}
}

// RegisterRoutes contains the routes that the controller needs to register in order to work
func (writer *Write) RegisterRoutes(basePath string) func(*gin.RouterGroup) {
	requirementHandler := requirementHTTP.NewRequirementHTTPHandler(writer.container)
	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/writer")

		roleGroup.GET("", writer.NotImplementedHandler)
		roleGroup.POST("/upload", requirementHandler.HandlerUploadRequirement)
	}
}

// NotImplementedHandler handler for not implemented response
func (writer *Write) NotImplementedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"data": "writer"})
}
