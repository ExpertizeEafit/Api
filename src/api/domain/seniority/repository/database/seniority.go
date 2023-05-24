package database

import (
	"context"
	"github.com/ExpertizeEafit/Api/src/api/domain/seniority/entities"
)

const (
	GetSeniorityRequestList = `SELECT lj.id,lj.name,lj.last_name,s.name FROM
    (SELECT sr.id,u.name,u.last_name,sr.seniority_id FROM (SELECT id,user_id,seniority_id FROM seniority_request WHERE status = 'PENDING') sr
        LEFT JOIN user u on sr.user_id = u.id) lj LEFT JOIN seniority s on lj.seniority_id = s.id;`
	CreateSeniorityRequest = `INSERT INTO seniority_request(user_id, seniority_id) VALUES (?,?);`
	UpdateStatus           = `UPDATE seniority_request SET status = ? WHERE id = ?;`
)

func (repository *seniorityRepositoryDatabase) GetSeniorityRequestList(ctx context.Context) ([]entities.SeniorityRequest, error) {
	rows, err := repository.database.Query(GetSeniorityRequestList)
	if err != nil {
		return nil, err
	}
	ans := make([]entities.SeniorityRequest, 0)
	for rows.Next() {
		request := entities.SeniorityRequest{}
		rows.Scan(&request.Id, &request.Name, &request.LastName, &request.Seniority)
		ans = append(ans, request)
	}
	return ans, nil
}

func (repository *seniorityRepositoryDatabase) CreateSeniorityRequest(ctx context.Context, userID, seniorityID int64) error {
	_, err := repository.database.Exec(CreateSeniorityRequest, userID, seniorityID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *seniorityRepositoryDatabase) UpdateStatusSeniorityRequest(ctx context.Context, id int64, status entities.Status) error {
	_, err := repository.database.Exec(UpdateStatus, status, id)
	if err != nil {
		return err
	}
	return nil
}
