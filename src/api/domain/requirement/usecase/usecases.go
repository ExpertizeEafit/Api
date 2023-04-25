package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/repository"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RequirementUseCases interface {
	UploadRequirement(ctx context.Context, data entities.RequirementFile) error
}

type requirementUseCases struct {
	repository repository.RequirementRepository
}

func NewRequirementUseCases(container *dependencies.Container) RequirementUseCases {
	return &requirementUseCases{
		repository: repository.NewRequirementRepository(container),
	}
}
