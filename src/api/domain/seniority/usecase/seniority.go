package usecase

import (
	"context"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/entities"
)

func (usecase *seniorityUseCases) GetSeniorityRequestList(ctx context.Context) ([]entities.SeniorityRequest, error) {
	return usecase.seniorityRepository.GetSeniorityRequestList(ctx)
}

func (usecase *seniorityUseCases) CreateSeniorityRequest(ctx context.Context, userID, seniorityID int64) error {
	//validate seniority
	return usecase.seniorityRepository.CreateSeniorityRequest(ctx, userID, seniorityID)
}

func (usecase *seniorityUseCases) UpdateStatusSeniorityReques(ctx context.Context, id int64, status entities.Status) error {
	return usecase.seniorityRepository.UpdateStatusSeniorityRequest(ctx, id, status)
}
