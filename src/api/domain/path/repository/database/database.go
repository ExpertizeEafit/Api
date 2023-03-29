package database

import (
	"context"
	"database/sql"

	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type PathRepositoryDatabase interface {
	GetAllSeniority(ctx context.Context) (entities.Path, error)
	GetCurrentAndNextSeniority(ctx context.Context, seniorityID int) (entities.PathExtended, error)
}

type pathRepositoryDatabase struct {
	database *sql.DB
}

func NewPathRepositoryDatabase(container *dependencies.Container) PathRepositoryDatabase {
	return &pathRepositoryDatabase{
		database: container.Database(),
	}
}
