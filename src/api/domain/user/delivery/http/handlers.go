package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/user/usecase"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type UserHTTPHandler struct {
	usecase usecase.UserUseCases
}

func NewUserHTTPHandler(container *dependencies.Container) *UserHTTPHandler {
	return &UserHTTPHandler{
		usecase: usecase.NewUserUseCases(container),
	}
}
