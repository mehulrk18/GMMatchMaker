package models

//GameCategory Type
type GameCategory string

//
const (
	ONEvONE GameCategory = "ONEvONE"
	ONEvN   GameCategory = "ONEvN"
	NvN     GameCategory = "NvN"
	NvTN    GameCategory = "NvTN"
)

// Game struct
type Game struct {
	ID       uint64
	Name     string
	Players  uint64
	Category GameCategory
}

// GamesRating rating of each player in game played.
// type GamesRating struct {
// 	ID     uint64
// 	UserID uint64
// 	User   User
// 	GameID uint64
// 	Game   Game
// 	Rating float32
// }
