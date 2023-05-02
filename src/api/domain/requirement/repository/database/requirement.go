package database

import (
	"bytes"
	"context"
	"io"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

const (
	InsertUserRequirement     = `INSERT INTO user_requirement(requirement_id, user_id, status,file) VALUES (?,?,?,?);`
	GetUserRequirementHistory = `SELECT u.id,r.name,u.status FROM user_requirement u INNER JOIN requirement r on u.requirement_id = r.id WHERE user_id = ?;`
)

func (repository *requirementRepositoryDatabase) UploadRequirement(ctx context.Context, data entities.RequirementFile) error {
	file, err := data.File.Open()
	defer file.Close()
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, file); err != nil {
		return err
	}
	_, err = repository.database.Exec(InsertUserRequirement, data.RequirementId, data.UserId, "PENDING", buff.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (repository *requirementRepositoryDatabase) GetRequirementHistory(ctx context.Context, id int64) ([]entities.UserRequirementStatus, error) {
	rows, err := repository.database.Query(GetUserRequirementHistory, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	history := make([]entities.UserRequirementStatus, 0)
	for rows.Next() {
		aux := entities.UserRequirementStatus{}
		rows.Scan(&aux.Id, &aux.Name, &aux.Status)
		history = append(history, aux)
	}
	return history, nil
}
