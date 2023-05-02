package http

import (
	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/usecase"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RankingHTTPHandler struct {
	usecase usecase.RankingUseCases
}

func NewRankingHTTPHanlder(container *dependencies.Container) *RankingHTTPHandler {
	return &RankingHTTPHandler{
		usecase: usecase.NewRankingUseCases(container),
	}
}
