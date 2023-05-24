package usecase

import (
	"context"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/repository"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type SeniorityUseCases interface {
	GetSeniorityRequestList(ctx context.Context) ([]entities.SeniorityRequest, error)
	CreateSeniorityRequest(ctx context.Context, userID, seniorityID int64) error
	UpdateStatusSeniorityReques(ctx context.Context, id int64, status entities.Status) error
}

type seniorityUseCases struct {
	seniorityRepository repository.SeniorityRepository
}

func NewSeniorityUseCases(container *dependencies.Container) SeniorityUseCases {
	return &seniorityUseCases{
		seniorityRepository: repository.NewSeniorityRepository(container),
	}
}
