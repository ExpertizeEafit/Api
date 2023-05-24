package database

import (
	"context"
	"database/sql"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/entities"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type SeniorityRepositoryDatabase interface {
	GetSeniorityRequestList(ctx context.Context) ([]entities.SeniorityRequest, error)
	CreateSeniorityRequest(ctx context.Context, userID, seniorityID int64) error
	UpdateStatusSeniorityRequest(ctx context.Context, id int64, status entities.Status) error
}

type seniorityRepositoryDatabase struct {
	database *sql.DB
}

func NewSeniorityRepositoryDatabase(container *dependencies.Container) SeniorityRepositoryDatabase {
	return &seniorityRepositoryDatabase{
		database: container.Database(),
	}
}
