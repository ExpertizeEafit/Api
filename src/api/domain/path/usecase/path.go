package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
)

func (usecase *pathUseCases) GetAllSeniority(ctx context.Context) (entities.Path, error) {
	return usecase.pathRepository.GetAllSeniority(ctx)
}
