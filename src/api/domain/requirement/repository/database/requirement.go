package database

import (
	"bytes"
	"context"
	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
	"io"
)

const (
	InsertUserRequirement     = `INSERT INTO user_requirement(requirement_id, user_id, status,file) VALUES (?,?,?,?);`
	GetUserRequirementHistory = `SELECT u.id,r.name,u.status FROM user_requirement u INNER JOIN requirement r on u.requirement_id = r.id WHERE user_id = ?;`
	GetPendingRequirements    = `SELECT ur.id,r.name,status FROM user_requirement ur LEFT JOIN requirement r on r.id = ur.requirement_id WHERE ur.status = 'PENDING';`
	UpdateStatus              = `UPDATE user_requirement SET status = ? WHERE id = ?;`
	SelectFile                = `SELECT file FROM user_requirement WHERE id = ?;`
	CreateRequirement         = `INSERT INTO requirement (name, description, recommendation, points) VALUES (?,?,?,?);`
	SelectRequirements        = `SELECT id,name,description FROM requirement;`
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

func (repository *requirementRepositoryDatabase) GetPendingRequirements(ctx context.Context) ([]entities.UserRequirementStatus, error) {
	rows, err := repository.database.Query(GetPendingRequirements)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	requirements := make([]entities.UserRequirementStatus, 0)
	for rows.Next() {
		aux := entities.UserRequirementStatus{}
		rows.Scan(&aux.Id, &aux.Name, &aux.Status)
		requirements = append(requirements, aux)
	}
	return requirements, nil
}

func (repository *requirementRepositoryDatabase) UpdateRequirementStatus(ctx context.Context, status entities.Status, id int) error {
	_, err := repository.database.Exec(UpdateStatus, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *requirementRepositoryDatabase) DownloadRequirementFile(ctx context.Context, id int) (*bytes.Buffer, error) {
	row := repository.database.QueryRow(SelectFile, id)

	data := make([]byte, 0)
	err := row.Scan(&data)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBuffer(data)
	return buff, nil
}

func (repository *requirementRepositoryDatabase) CreateRequirement(ctx context.Context, data entities.RequirementData) (int64, error) {
	res, err := repository.database.Exec(CreateRequirement, data.Name, data.Description, data.Recommendation, data.Points)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *requirementRepositoryDatabase) GetAllRequirements(ctx context.Context) ([]entities.RequirementBasicData, error) {
	rows, err := repository.database.Query(SelectRequirements)
	if err != nil {
		return nil, err
	}
	requirements := make([]entities.RequirementBasicData, 0)
	for rows.Next() {
		aux := entities.RequirementBasicData{}
		rows.Scan(&aux.Id, &aux.Name, &aux.Description)
		requirements = append(requirements, aux)
	}
	return requirements, nil
}
