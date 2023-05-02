package repository

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/repository/database"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RankingRepository interface {
	GetRanking(ctx context.Context, userID int) (entities.RankList, error)
}

type rankingRepository struct {
	database.RankingRepositoryDatabase
}

func NewRankingRepository(container *dependencies.Container) RankingRepository {
	return &rankingRepository{
		database.NewRankingRepositoryDatabase(container),
	}
}
