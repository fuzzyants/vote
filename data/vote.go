package data

import (
	"time"
)

type Vote struct {
	Id        uint64
	MaxUsers  uint32
	Slug      string
	Expires   time.Time
	Done      bool
	Options   []*Option
	Ballots   []Ballot
	CreatedAt time.Time
}
