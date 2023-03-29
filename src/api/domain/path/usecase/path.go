package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
)

func (usecase *pathUseCases) GetAllSeniority(ctx context.Context) (entities.Path, error) {
	return usecase.pathRepository.GetAllSeniority(ctx)
}

func (usecase *pathUseCases) GetCurrentAndNextSeniority(ctx context.Context, seniorityID int) (entities.PathExtended, error) {
	return usecase.pathRepository.GetCurrentAndNextSeniority(ctx, seniorityID)
}
