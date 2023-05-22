package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/user/repository"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type UserUseCases interface {
	Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error)
	Register(ctx context.Context, users []entities.RegisterData) []entities.UserRegisterError
	UpdatePassword(ctx context.Context, data entities.UpdatePassword) error
}

type userUseCases struct {
	repository repository.UserRepository
}

func NewUserUseCases(container *dependencies.Container) UserUseCases {
	return &userUseCases{
		repository: repository.NewUserRepository(container),
	}
}
