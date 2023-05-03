package database

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
)

const (
	SelectAllSeniority     = `SELECT id,name,description,prior_to FROM seniority`
	SelectCurrentSeniority = `SELECT id,name,description,prior_to,requirements FROM seniority WHERE id = ?`
	SelectRequirements     = `SELECT r.id,r.name,r.description,r.recommendation,IFNULL(ur.status, 'LOCK') as status FROM requirement r
    							LEFT JOIN
    							(SELECT status,requirement_id FROM user_requirement WHERE user_id = ?) AS ur on r.id = ur.requirement_id WHERE r.id IN('%s');`
	SelectNextSeniority = `SELECT id,name,description,requirements FROM seniority WHERE id IN ('%s');`
	SelectSeniorityID   = `SELECT seniority_id FROM user where id = ?;`
)

func (repository *pathRepositoryDatabase) GetAllSeniority(ctx context.Context) (entities.Path, error) {
	rows, err := repository.database.Query(SelectAllSeniority)
	defer rows.Close()
	if err != nil {
		return entities.Path{}, err
	}
	path := entities.Path{}
	for rows.Next() {
		seniority := entities.SeniorityInfo{}
		var aux []byte
		var id int
		err := rows.Scan(&id, &seniority.Name, &seniority.Description, &aux)
		if err != nil {
			return entities.Path{}, err
		}
		err = json.Unmarshal(aux, &seniority.PriorTo)
		if err != nil {
			return entities.Path{}, err
		}
		path[id] = seniority
	}
	return path, nil
}

func (repository *pathRepositoryDatabase) GetCurrentAndNextSeniority(ctx context.Context, userID int) (entities.PathExtended, error) {

	seniorityID := 0
	row := repository.database.QueryRow(SelectSeniorityID, userID)
	err := row.Scan(&seniorityID)
	if err != nil {
		return entities.PathExtended{}, err
	}

	result := entities.PathExtended{}
	row = repository.database.QueryRow(SelectCurrentSeniority, seniorityID)
	var priorTo []byte
	var requirementsRaw []byte
	var requirements []int
	var auxId int
	auxCurr := entities.SeniorityInfoExtended{}
	auxCurr.Requirements = map[int]entities.RequirementInfo{}
	err = row.Scan(&auxId, &auxCurr.Name, &auxCurr.Description, &priorTo, &requirementsRaw)
	if err != nil {
		return entities.PathExtended{}, err
	}
	err = json.Unmarshal(priorTo, &auxCurr.PriorTo)
	if err != nil {
		return entities.PathExtended{}, err
	}
	result[auxId] = auxCurr
	err = json.Unmarshal(requirementsRaw, &requirements)
	if err != nil {
		return entities.PathExtended{}, err
	}
	stringRequirements := strings.Trim(strings.Replace(fmt.Sprint(requirements), " ", "','", -1), "[]")
	query := fmt.Sprintf(SelectRequirements, stringRequirements)
	rows, err := repository.database.Query(query, userID)
	for rows.Next() {
		req := entities.RequirementInfo{}
		err := rows.Scan(&auxId, &req.Name, &req.Description, &req.Recommendation, &req.Status)
		if err != nil {
			return entities.PathExtended{}, err
		}
		result[seniorityID].Requirements[auxId] = req
	}

	stringPriorTo := strings.Trim(strings.Replace(fmt.Sprint(result[seniorityID].PriorTo), " ", "','", -1), "[]")
	query = fmt.Sprintf(SelectNextSeniority, stringPriorTo)
	rows, err = repository.database.Query(query)
	for rows.Next() {
		nextSeniority := entities.SeniorityInfoExtended{}
		nextSeniority.Requirements = map[int]entities.RequirementInfo{}
		err = rows.Scan(&auxId, &nextSeniority.Name, &nextSeniority.Description, &requirementsRaw)
		if err != nil {
			return entities.PathExtended{}, err
		}
		result[auxId] = nextSeniority
		err = json.Unmarshal(requirementsRaw, &requirements)
		if err != nil {
			return entities.PathExtended{}, err
		}
		stringRequirements = strings.Trim(strings.Replace(fmt.Sprint(requirements), " ", "','", -1), "[]")
		query = fmt.Sprintf(SelectRequirements, stringRequirements)
		rows2, err := repository.database.Query(query, userID)
		for rows2.Next() {
			var auxIdReq int
			req := entities.RequirementInfo{}
			err = rows2.Scan(&auxIdReq, &req.Name, &req.Description, &req.Recommendation, &req.Status)
			if err != nil {
				return entities.PathExtended{}, err
			}
			result[auxId].Requirements[auxIdReq] = req
		}
	}
	return result, nil
}
