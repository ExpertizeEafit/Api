package database

import (
	"context"
	"database/sql"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
	"github.com/ExpertizeEafit/Api/src/api/infrastructure/dependencies"
)

type UserRepositoryDatabase interface {
	Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error)
}

type userRepositoryDatabase struct {
	database *sql.DB
}

func NewUserRepositoryDatabase(container *dependencies.Container) UserRepositoryDatabase {
	return &userRepositoryDatabase{
		database: container.Database(),
	}
}
