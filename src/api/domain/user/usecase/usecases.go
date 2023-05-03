package usecase

import (
	"context"
	
	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/user/repository"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type UserUseCases interface {
	Login(ctx context.Context, idNumber int, password string) (entities.LoginBasicInfo, error)
}

type userUseCases struct {
	repository repository.UserRepository
}

func NewUserUseCases(container *dependencies.Container) UserUseCases {
	return &userUseCases{
		repository: repository.NewUserRepository(container),
	}
}
