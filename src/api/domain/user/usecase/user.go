package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
)

func (usecase *userUseCases) Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error) {
	return usecase.repository.Login(ctx, idNumber, password)
}

func (usecase *userUseCases) Register(ctx context.Context, users [][]string) []entities.UserRegisterError {
	userErrors := []entities.UserRegisterError{}
	for _, user := range users {
		auxUser := entities.RegisterData{
			DNI:         user[0],
			Name:        user[1],
			LastName:    user[2],
			SeniorityID: user[3],
			AreaID:      user[4],
		}
		err := usecase.repository.Register(ctx, auxUser)
		if err != nil {
			userErrors = append(userErrors, entities.UserRegisterError{
				User: auxUser,
				Err:  err,
			})
		}
	}
	return userErrors
}

func (usecase *userUseCases) UpdatePassword(ctx context.Context, data entities.UpdatePassword) error {
	return usecase.repository.UpdatePassword(ctx, data)
}
