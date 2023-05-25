package repository

import (
	"context"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/repository/database"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type SeniorityRepository interface {
	GetSeniorityRequestList(ctx context.Context) ([]entities.SeniorityRequest, error)
	CreateSeniorityRequest(ctx context.Context, userID, seniorityID int64) error
	UpdateStatusSeniorityRequest(ctx context.Context, id int64, status entities.Status) error
	SelectUserAndUpdateSeniority(ctx context.Context, id int64) error
}

type seniorityRepository struct {
	database.SeniorityRepositoryDatabase
}

func NewSeniorityRepository(container *dependencies.Container) SeniorityRepository {
	return &seniorityRepository{
		database.NewSeniorityRepositoryDatabase(container),
	}
}
