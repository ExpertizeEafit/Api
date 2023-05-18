package database

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

type RequirementRepositoryDatabase interface {
	UploadRequirement(ctx context.Context, data entities.RequirementFile) error
	GetRequirementHistory(ctx context.Context, id int64) ([]entities.UserRequirementStatus, error)
	GetPendingRequirements(ctx context.Context) ([]entities.UserRequirementStatus, error)
	UpdateRequirementStatus(ctx context.Context, status entities.Status, id int) error
	DownloadRequirementFile(ctx context.Context, id int) (*bytes.Buffer, error)
}

type requirementRepositoryDatabase struct {
	database *sql.DB
}

func NewRequirementRepositoryDatabase(container *dependencies.Container) RequirementRepositoryDatabase {
	return &requirementRepositoryDatabase{
		database: container.Database(),
	}
}
