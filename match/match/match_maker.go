package match


import (

	deque "gopkg.in/karalabe/cookiejar.v1/collections/deque"
	"time"
)

type Player struct {
	User	models.User
	ExpireTime	time.Time
}

type Pool struct{
	Players	deque.Deque

}


func makeMatches() {

}