package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/path/usecase"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type PathHTTPHandler struct {
	usecase usecase.PathUseCases
}

func NewPathDeliveryHTTPHandler(container *dependencies.Container) *PathHTTPHandler {
	return &PathHTTPHandler{
		usecase: usecase.NewPathUseCases(container),
	}
}
