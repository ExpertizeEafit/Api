package repository

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/path/repository/database"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type PathRepository interface {
	GetAllSeniority(ctx context.Context) (entities.Path, error)
	GetCurrentAndNextSeniority(ctx context.Context, seniorityID int) (entities.PathExtended, error)
}

type pathRepository struct {
	database.PathRepositoryDatabase
}

func NewPathRepository(container *dependencies.Container) PathRepository {
	return &pathRepository{
		database.NewPathRepositoryDatabase(container),
	}
}
