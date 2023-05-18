package repository

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
	"github.com/ExpertizeEafit/Api/src/api/domain/user/repository/database"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type UserRepository interface {
	Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error)
	Register(ctx context.Context, data entities.RegisterData) error
	UpdatePassword(ctx context.Context, data entities.UpdatePassword) error
}

type userRepository struct {
	database.UserRepositoryDatabase
}

func NewUserRepository(container *dependencies.Container) UserRepository {
	return &userRepository{
		database.NewUserRepositoryDatabase(container),
	}
}
