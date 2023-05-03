package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
)

func (usecase *userUseCases) Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error) {
	return usecase.repository.Login(ctx, idNumber, password)
}
