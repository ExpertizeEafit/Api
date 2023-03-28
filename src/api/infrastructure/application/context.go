package application

import (
	"fmt"
	"os"
	"strings"

	"github.com/ExpertizeEafit/Api/src/api/infrastructure/application/entities"
)

type (
	// Contexter interface for application context
	Contexter interface {
		// Environment getter for environment
		Environment() entities.Environment
		// Role getter for role
		Role() entities.Role
	}
	context struct{}
)

var (
	contextInitialized bool

	environment entities.Environment
	role        entities.Role
)

func InitContext() {
	if !contextInitialized {
		scope := os.Getenv("SCOPE")

		parts := strings.Split(strings.ToLower(scope), "-")
		if len(parts) <= 1 {
			panic(fmt.Sprintf("invalid scope %s", scope))
		}

		environment = entities.Environment(parts[0])
		if !entities.PossibleEnvironments.Contains(environment) {
			panic(fmt.Sprintf("invalid environment %s", environment))
		}

		role = entities.Role(parts[1])
		if !entities.PossibleRoles.Contains(role) {
			panic(fmt.Sprintf("invalid role %s", role))
		}

		contextInitialized = true
	}
}

// Context getter for application context
func Context() Contexter {
	if !contextInitialized {
		panic("application context not initialized")
	}

	return &context{}
}

func (*context) Environment() entities.Environment {
	return environment
}

func (*context) Role() entities.Role {
	return role
}

// IsLocal used to allow local run
func IsLocal() bool {
	return strings.EqualFold(os.Getenv("LOCAL_RUN"), "true")
}
