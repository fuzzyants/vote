package data

import (
	"time"
)

type Option struct {
	Id        uint64
	Title     string
	Url       string
	CreatedAt time.Time
}
