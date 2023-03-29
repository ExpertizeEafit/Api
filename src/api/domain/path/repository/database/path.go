package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ExpertizeEafit/Api/src/api/domain/path/entities"
)

const (
	SelectAllSeniority = `SELECT id,name,description,prior_to FROM seniority`
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
		fmt.Println(aux)
	}
	return path, nil
}
