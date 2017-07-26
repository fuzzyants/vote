package data

import (
	"time"
)

type Ballot struct {
	Id        uint64
	User      *User
	Vote      *Vote
	Options   map[uint8]*Option
	CreatedAt time.Time
}
