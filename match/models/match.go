package models


// Match of a game
type Match struct{
	ID				string
	GameID			uint64
	Players			[]User
	MaxCapacity		uint32
	TotalPlayers	uint32
	AvgRating		float64
}


// Result of the Match.
type Result struct{
	MatchID string
	Match	Match
	Winner 	User
}

