package database

import (
	"context"
	"errors"
	"math/rand"

	"github.com/ExpertizeEafit/Api/src/api/domain/user/entities"
)

const (
	login           = `SELECT id,username,change_password,rol FROM user WHERE id_number = ? && password = ?;`
	PasswordCorrect = `SELECT id FROM user WHERE id = ? && password = ?;`
	UpdatePassword  = `UPDATE user SET password = ? WHERE id = ?;`
	InsertUser      = `INSERT INTO user(area_id, seniority_id, name, last_name,id_number, username, password) VALUES (?,?,?,?,?,?,?);`
)

func (repository *userRepositoryDatabase) Login(ctx context.Context, idNumber string, password string) (entities.LoginBasicInfo, error) {
	row := repository.database.QueryRow(login, idNumber, password)
	res := entities.LoginBasicInfo{}
	err := row.Scan(&res.Id, &res.Username, &res.ChangePassword, &res.Rol)
	if err != nil {
		return entities.LoginBasicInfo{}, err
	}
	return res, nil
}

func (repository *userRepositoryDatabase) Register(ctx context.Context, data entities.RegisterData) error {
	random1 := rand.Intn(len(entities.Names))
	random2 := rand.Intn(len(entities.Lastnames))
	username := entities.Names[random1] + " " + entities.Lastnames[random2]
	params := []interface{}{
		data.AreaID,
		data.SeniorityID,
		data.Name,
		data.LastName,
		data.DNI,
		username,
		data.DNI,
	}
	_, err := repository.database.Exec(InsertUser, params...)
	if err != nil {
		return err
	}
	return nil
}

func (repository *userRepositoryDatabase) UpdatePassword(ctx context.Context, data entities.UpdatePassword) error {
	row := repository.database.QueryRow(PasswordCorrect, data.Id, data.OldPassword)
	id := -1
	row.Scan(&id)
	if id == data.Id {
		_, err := repository.database.Exec(UpdatePassword, data.NewPassword, data.Id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("incorrect password")
}
