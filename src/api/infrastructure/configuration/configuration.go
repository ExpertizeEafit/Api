package configuration

import (
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application/entities"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/configuration/develop"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/configuration/production"
)

type Configurationer interface {
	BasePath() string
	DBDriverName() string
	DatabaseConnectionString() string
}

func Configuration() Configurationer {
	switch application.Context().Environment() {
	case entities.ProductionEnvironment:
		configuration := struct {
			baseConfiguration
			production.Configuration
		}{}
		return &configuration
	case entities.TestEnvironment:
		fallthrough
	case entities.DevelopEnvironment:
		configuration := struct {
			baseConfiguration
			develop.Configuration
		}{}
		return &configuration
	default:
		panic("default_environment_not_implemented")
	}
}
