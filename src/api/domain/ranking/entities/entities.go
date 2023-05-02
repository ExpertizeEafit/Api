package entities

type UserRank struct {
	Rank     int
	Username string
	Points   int
}

type RankList struct {
	Ranking []UserRank
	MyRank  UserRank
}
