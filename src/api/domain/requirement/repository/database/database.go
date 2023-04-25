package database

import (
	"context"
	"database/sql"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

type RequirementRepositoryDatabase interface {
	UploadRequirement(ctx context.Context, data entities.RequirementFile) error
}

type requirementRepositoryDatabase struct {
	database *sql.DB
}

func NewRequirementRepositoryDatabase(container *dependencies.Container) RequirementRepositoryDatabase {
	return &requirementRepositoryDatabase{
		database: container.Database(),
	}
}
