package match

import (
	"fmt"
	"math"
	"time"

	models "github.com/mehulrk18/GMMatchMaker/match/models"
	deque "gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

const (
	ratingRange float64 = 10.0
)

func inRange(x, y float64) bool {
	if math.Abs(x-y) < ratingRange {
		return true
	}
	return false
}

// Player struct
type Player struct {
	User           models.User
	ExpiryWaitTime time.Time
}

// Pool contatins players
type Pool struct {
	Players *deque.Deque
	Game    models.Game
}

// GamePool channel of pool
type GamePool struct {
	PoolChannel chan Pool
}

// Init initialize GamePool
func (g *GamePool) Init() {
	g.PoolChannel = make(chan Pool)

	go func() {
		g.PoolChannel <- Pool{
			Game:    models.Game{},
			Players: deque.New(),
		}
	}()
}

//MakeMatches creating matches
func (g *GamePool) MakeMatches() []models.Match {
	pool := <-g.PoolChannel

	var playersReadyForMatch []Player
	var matches []models.Match
	tempDeque := deque.New()
	i := 0
	for !pool.Players.Empty() {
		if pool.Game.Category == models.ONEvONE {

			p := pool.Players.PopRight()

			playersReadyForMatch = append(playersReadyForMatch, p.(Player))
			var homePlayer, awayPlayer Player
			foundPlayer := false
			if !pool.Players.Empty() {
				awayPlayer = pool.Players.Right().(Player)
				pool.Players.PopRight()
				foundPlayer = true
			}

			homePlayer = playersReadyForMatch[i]

			nearestRating := math.Abs(homePlayer.User.GamesRating - awayPlayer.User.GamesRating)

			if inRange(playersReadyForMatch[i].User.GamesRating, awayPlayer.User.GamesRating) {
				foundPlayer = true
			} else if time.Now().Sub(playersReadyForMatch[i].ExpiryWaitTime) < 0 {
				var p1 interface{}
				pool.Players.PushLeft(awayPlayer)
				for !pool.Players.Empty() {
					p1 = pool.Players.PopRight()
					if nearestRating > math.Abs(homePlayer.User.GamesRating-p1.(Player).User.GamesRating) {
						awayPlayer = p1.(Player)

						nearestRating = math.Abs(homePlayer.User.GamesRating - p1.(Player).User.GamesRating)
						foundPlayer = true
					}
					tempDeque.PushLeft(p1)

				}
				for !tempDeque.Empty() {
					p1 = tempDeque.PopRight()
					if p1.(Player).User.ID != awayPlayer.User.ID {
						pool.Players.PushLeft(p1)
					}
				}
			} else {
				pool.Players.PushLeft(homePlayer)
				pool.Players.PushLeft(awayPlayer)
				foundPlayer = false
			}

			if foundPlayer {
				var finalPlayers []models.User
				finalPlayers = append(finalPlayers, homePlayer.User)
				finalPlayers = append(finalPlayers, awayPlayer.User)
				m := models.Match{
					ID:           fmt.Sprintf("match-%s-%d-%d-%d", time.Now().Format("2006-11-02T15:04:05"), homePlayer.User.ID, awayPlayer.User.ID, pool.Game.ID),
					GameID:       pool.Game.ID,
					Players:      finalPlayers,
					MaxCapacity:  2,
					AvgRating:    (homePlayer.User.GamesRating + awayPlayer.User.GamesRating) / 2,
					TotalPlayers: 2,
				}
				matches = append(matches, m)
				fmt.Println("created a match = ", m)
			}
			i++
		}
	}

	go func() {
		g.PoolChannel <- pool
	}()

	return matches
}
