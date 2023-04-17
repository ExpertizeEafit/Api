package http

import "github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"

type RequirementHTTPHandler struct {
}

func NewRequirementHTTPHandler(container *dependencies.Container) *RequirementHTTPHandler {
	return &RequirementHTTPHandler{}
}
