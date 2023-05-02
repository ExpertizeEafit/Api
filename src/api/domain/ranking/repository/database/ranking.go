package database

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/entities"
)

const (
	GetRanking = `SELECT r.user_id,RANK() over (ORDER BY r.points DESC) AS ranking,u.username,r.points FROM ranking r INNER JOIN user u on r.user_id = u.id ORDER BY r.points DESC;
`
)

func (repository *rankingRepositoryDatabase) GetRanking(ctx context.Context, userID int) (entities.RankList, error) {
	rows, err := repository.database.Query(GetRanking)
	defer rows.Close()
	if err != nil {
		return entities.RankList{}, err
	}
	ranking := make([]entities.UserRank, 0)
	myRank := entities.UserRank{}
	for rows.Next() {
		aux := entities.UserRank{}
		auxId := 0
		rows.Scan(&auxId, &aux.Rank, &aux.Username, &aux.Points)
		ranking = append(ranking, aux)
		if auxId == userID {
			myRank = aux
		}
	}
	res := entities.RankList{
		Ranking: ranking,
		MyRank:  myRank,
	}
	return res, nil
}
