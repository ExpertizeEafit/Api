package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/repository"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RankingUseCases interface {
	GetRanking(ctx context.Context, usedID int) (entities.RankList, error)
}

type rankingUseCases struct {
	repository repository.RankingRepository
}

func NewRankingUseCases(container *dependencies.Container) RankingUseCases {
	return &rankingUseCases{
		repository: repository.NewRankingRepository(container),
	}
}
