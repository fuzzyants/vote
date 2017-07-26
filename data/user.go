package data

import (
	"time"
)

type User struct {
	Id        uint64
	Token     string
	Name      string
	Email     string
	LastSeen  time.Time
	CreatedAt time.Time
}
