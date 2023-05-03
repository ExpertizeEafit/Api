package write

import (
	"net/http"

	"github.com/gin-gonic/gin"

	pathHTTP "github.com/ExpertizeEafit/Api/src/api/domain/path/delivery/http"
	rankingHTTP "github.com/ExpertizeEafit/Api/src/api/domain/ranking/delivery/http"
	requirementHTTP "github.com/ExpertizeEafit/Api/src/api/domain/requirement/delivery/http"
	userHTTP "github.com/ExpertizeEafit/Api/src/api/domain/user/delivery/http"
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
	pathHandler := pathHTTP.NewPathDeliveryHTTPHandler(writer.container)
	rankingHandler := rankingHTTP.NewRankingHTTPHanlder(writer.container)
	userHandler := userHTTP.NewUserHTTPHandler(writer.container)

	return func(g *gin.RouterGroup) {
		v1Group := g.Group(basePath + "/v1")
		roleGroup := v1Group.Group("/writer")

		roleGroup.GET("", writer.NotImplementedHandler)
		roleGroup.POST("/upload", requirementHandler.HandlerUploadRequirement)

		// Reader
		roleGroup.GET("/getPaths", pathHandler.HandleGetPaths)
		roleGroup.GET("/getCurrentAndNextSeniority/:id", pathHandler.HandlerGetCurrentAndNextSeniority)
		roleGroup.GET("/getRequirementsHistory/:id", requirementHandler.HandlerGetRequirementHistory)
		roleGroup.GET("/ranking/:id", rankingHandler.HandlerGetRanking)
		roleGroup.POST("/login", userHandler.HandlerLogin)
	}
}

// NotImplementedHandler handler for not implemented response
func (writer *Write) NotImplementedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"data": "writer"})
}
