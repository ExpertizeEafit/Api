package reader

import (
	pathHTTP "github.com/ExpertizeEafit/Api/src/api/domain/path/delivery/http"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Read is main structure of consumer functionality
type Read struct {
	container *dependencies.Container
}

// NewRead returns a concrete implementation of the Read scope
func NewRead(container *dependencies.Container) *Read {
	return &Read{
		container: container,
	}
}

// RegisterRoutes contains the routes that the controller needs to register in order to work
func (reader *Read) RegisterRoutes(basePath string) func(*gin.RouterGroup) {
	pathHandler := pathHTTP.NewPathDeliveryHTTPHandler(reader.container)
	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/reader")

		roleGroup.GET("", reader.NotImplementedHandler)
		roleGroup.GET("/getPaths", pathHandler.HandleGetPaths)
		roleGroup.GET("/getCurrentAndNextSeniority", pathHandler.HandlerGetCurrentAndNextSeniority)
	}
}

// NotImplementedHandler handler for not implemented response
func (reader *Read) NotImplementedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"data": "reader"})
}
