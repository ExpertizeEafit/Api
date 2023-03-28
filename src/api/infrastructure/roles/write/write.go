package write

import (
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	"net/http"
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
func (reader *Write) RegisterRoutes(basePath string) func(*gin.RouterGroup) {

	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/writer")

		roleGroup.GET("", reader.NotImplementedHandler)
	}
}

// NotImplementedHandler handler for not implemented response
func (reader *Write) NotImplementedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"data": "writer"})
}
