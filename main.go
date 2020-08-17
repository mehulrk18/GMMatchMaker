package main

import (

	// "testing"
	"time"

	match "github.com/mehulrk18/GMMatchMaker/match/match"
	models "github.com/mehulrk18/GMMatchMaker/match/models"
	deque "gopkg.in/karalabe/cookiejar.v1/collections/deque"

	// "math"
	"fmt"
)

var u1 models.User = models.User{
	ID:          1,
	Name:        "User 1",
	GamesRating: 8,
}

var u2 models.User = models.User{
	ID:          2,
	Name:        "User 2",
	GamesRating: 45,
}

var u3 models.User = models.User{
	ID:          3,
	Name:        "User 3",
	GamesRating: 20,
}

var u4 models.User = models.User{
	ID:          4,
	Name:        "User 4",
	GamesRating: 78,
}

var u5 models.User = models.User{
	ID:          5,
	Name:        "User 5",
	GamesRating: 68,
}

var g models.Game = models.Game{
	ID:       1,
	Name:     "Game 1",
	Players:  2,
	Category: models.ONEvONE,
}

func main() {
	var testChannel match.GamePool
	testChannel.PoolChannel = make(chan match.Pool)
	que := deque.New()

	que.PushLeft(match.Player{
		User:           u1,
		ExpiryWaitTime: time.Now().Add(1 * time.Second),
	})
	que.PushLeft(match.Player{
		User:           u2,
		ExpiryWaitTime: time.Now().Add(2 * time.Second),
	})
	que.PushLeft(match.Player{
		User:           u3,
		ExpiryWaitTime: time.Now().Add(4 * time.Second),
	})
	que.PushLeft(match.Player{
		User:           u4,
		ExpiryWaitTime: time.Now().Add(7 * time.Second),
	})
	// que.PushLeft(match.Player{
	// 	User:           u5,
	// 	ExpiryWaitTime: time.Now().Add(1 * time.Second),
	// })

	testPool := match.Pool{
		Game:    g,
		Players: que,
	}

	go func() {
		testChannel.PoolChannel <- testPool
	}()

	matches := testChannel.MakeMatches()
	fmt.Println("The Matches are as follows: ")
	for i := range matches {
		fmt.Println(matches[i].ID, " : ", matches[i].Players[0].Name, " v/s ", matches[i].Players[1].Name)
	}

}
