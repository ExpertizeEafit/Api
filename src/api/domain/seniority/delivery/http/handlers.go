package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/usecase"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type SeniorityHTTPHandler struct {
	usecase usecase.SeniorityUseCases
}

func NewSeniorityDeliveryHTTPHandler(container *dependencies.Container) *SeniorityHTTPHandler {
	return &SeniorityHTTPHandler{
		usecase: usecase.NewSeniorityUseCases(container),
	}
}
