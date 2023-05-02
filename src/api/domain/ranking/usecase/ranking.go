package usecase

import (
	"context"

	"github.com/ExpertizeEafit/Api/src/api/domain/ranking/entities"
)

func (usecase *rankingUseCases) GetRanking(ctx context.Context, usedID int) (entities.RankList, error) {
	ranking, err := usecase.repository.GetRanking(ctx, usedID)
	if err != nil {
		return entities.RankList{}, err
	}
	return ranking, nil
}
