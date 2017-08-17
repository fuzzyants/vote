package data

import (
	"fmt"
	"log"
	"time"
)

type Vote struct {
	Id        int64
	MaxUsers  uint32
	Slug      string
	Expires   time.Time
	Done      bool
	Options   []Option
	Ballots   []Ballot
	CreatedAt time.Time
}

// SetName receives a pointer to Foo so it can modify it.
func (vote *Vote) SetId(Id int64) {
	vote.Id = Id
}

func (vote *Vote) Create() (err error) {

	stmt, err := Db.Prepare("INSERT INTO Votes(maxUsers, slug, expires, done, createdAt) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(vote.MaxUsers, vote.Slug, vote.Expires, vote.Done, vote.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	// Not sure this is idiomatic.
	if vote.Id == 0 {
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		} else {
			vote.SetId(lastId)
		}
	}

	return
}
