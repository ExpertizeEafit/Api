package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

func (usecase *requirementUseCases) UploadRequirement(ctx context.Context, data entities.RequirementFile) error {
	return usecase.repository.UploadRequirement(ctx, data)
}

func (usecase *requirementUseCases) GetRequirementHistory(ctx context.Context, id int) ([]entities.UserRequirementStatus, error) {
	return usecase.repository.GetRequirementHistory(ctx, int64(id))
}
