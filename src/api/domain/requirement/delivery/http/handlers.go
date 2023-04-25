package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/usecase"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RequirementHTTPHandler struct {
	usecase usecase.RequirementUseCases
}

func NewRequirementHTTPHandler(container *dependencies.Container) *RequirementHTTPHandler {
	return &RequirementHTTPHandler{
		usecase: usecase.NewRequirementUseCases(container),
	}
}
