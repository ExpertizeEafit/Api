package entities

type UserRank struct {
	Rank     int    `json:"rank"`
	Username string `json:"username"`
	Points   int    `json:"points"`
}

type RankList struct {
	Ranking []UserRank `json:"ranking"`
	MyRank  UserRank   `json:"myrank"`
}
