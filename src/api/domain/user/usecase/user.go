package usecase

import (
	"context"
	
	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
)

func (usecase *userUseCases) Login(ctx context.Context, idNumber int, password string) (entities.LoginBasicInfo, error) {
	return usecase.repository.Login(ctx, idNumber, password)
}
