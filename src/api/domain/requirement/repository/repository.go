package repository

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/repository/database"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type RequirementRepository interface {
	UploadRequirement(ctx context.Context, data entities.RequirementFile) error
	GetRequirementHistory(ctx context.Context, id int64) ([]entities.UserRequirementStatus, error)
}

type requirementRepository struct {
	database.RequirementRepositoryDatabase
}

func NewRequirementRepository(container *dependencies.Container) RequirementRepository {
	return &requirementRepository{
		database.NewRequirementRepositoryDatabase(container),
	}
}
