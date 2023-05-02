package database

import (
	"context"
	"database/sql"

	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/entities"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RankingRepositoryDatabase interface {
	GetRanking(ctx context.Context, userID int) (entities.RankList, error)
}

type rankingRepositoryDatabase struct {
	database *sql.DB
}

func NewRankingRepositoryDatabase(container *dependencies.Container) RankingRepositoryDatabase {
	return &rankingRepositoryDatabase{
		database: container.Database(),
	}
}
