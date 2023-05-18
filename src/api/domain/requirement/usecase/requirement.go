package usecase

import (
	"bytes"
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

func (usecase *requirementUseCases) UploadRequirement(ctx context.Context, data entities.RequirementFile) error {
	return usecase.repository.UploadRequirement(ctx, data)
}

func (usecase *requirementUseCases) GetRequirementHistory(ctx context.Context, id int) ([]entities.UserRequirementStatus, error) {
	return usecase.repository.GetRequirementHistory(ctx, int64(id))
}

func (usecase *requirementUseCases) GetPendingRequirements(ctx context.Context) ([]entities.UserRequirementStatus, error) {
	return usecase.repository.GetPendingRequirements(ctx)
}

func (usecase *requirementUseCases) UpdateRequirementStatus(ctx context.Context, status entities.Status, id int) error {
	return usecase.repository.UpdateRequirementStatus(ctx, status, id)
}

func (usecase *requirementUseCases) DownloadRequirementFile(ctx context.Context, id int) (*bytes.Buffer, error) {
	return usecase.repository.DownloadRequirementFile(ctx, id)
}
