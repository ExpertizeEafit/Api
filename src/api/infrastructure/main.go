package infrastructure

import (
	"fmt"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application/entities"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/configuration"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/roles/reader"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/roles/write"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application"
)

func main() {
	scope := os.Getenv("SCOPE")
	if scope == "" {
		panic(fmt.Errorf("application initialization error - No scope defined"))
	}

	// start application context
	application.InitContext()

	// services initialization
	container := dependencies.StartDependencies()

	// server engine initialization
	var server *gin.Engine
	server = initializeEngine(scope, container)
	server.Run(":8080")
}

// build validate and initialize gordic engine
func initializeEngine(scope string, container *dependencies.Container) *gin.Engine {
	config := configuration.Configuration()
	basePath := config.BasePath()

	router := gin.New()
	group := router.Group("")

	routes := map[entities.Role]func(*gin.RouterGroup){
		entities.ReadRole:  reader.NewRead(container).RegisterRoutes(basePath),
		entities.WriteRole: write.NewWrite(container).RegisterRoutes(basePath),
	}

	routes[application.Context().Role()](group)

	return router
}
