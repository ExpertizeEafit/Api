package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
)

func (usecase *userUseCases) Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error) {
	return usecase.repository.Login(ctx, idNumber, password)
}

func (usecase *userUseCases) Register(ctx context.Context, users []entities.RegisterData) []entities.UserRegisterError {
	userErrors := []entities.UserRegisterError{}
	for _, user := range users {
		err := usecase.repository.Register(ctx, user)
		if err != nil {
			userErrors = append(userErrors, entities.UserRegisterError{
				User: user,
				Err:  err,
			})
		}
	}
	return userErrors
}

func (usecase *userUseCases) UpdatePassword(ctx context.Context, data entities.UpdatePassword) error {
	return usecase.repository.UpdatePassword(ctx, data)
}
