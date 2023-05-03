package database

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
)

const (
	login = `SELECT id,username FROM user WHERE id_number = ? && password = ?;`
)

func (repository *userRepositoryDatabase) Login(ctx context.Context, idNumber int, password string) (entities.LoginBasicInfo, error) {
	row := repository.database.QueryRow(login, idNumber, password)
	res := entities.LoginBasicInfo{}
	err := row.Scan(&res.Id, &res.Username)
	if err != nil {
		return entities.LoginBasicInfo{}, err
	}
	return res, nil
}
