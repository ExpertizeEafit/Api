package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/path/repository"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type PathUseCases interface {
	GetAllSeniority(ctx context.Context) (entities.Path, error)
	GetCurrentAndNextSeniority(ctx context.Context, userID int) (entities.PathExtended, error)
}

type pathUseCases struct {
	pathRepository repository.PathRepository
}

func NewPathUseCases(container *dependencies.Container) PathUseCases {
	return &pathUseCases{
		pathRepository: repository.NewPathRepository(container),
	}
}
