package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Poll struct {
	Id         uint
	Title      string
	Expiration uint
	Options    []Option
}

func (p Poll) save() error {
	polls[p.Id] = p
	return nil
}

type Option struct {
	Name string
}

var polls map[uint]Poll

func init() {
	polls = make(map[uint]Poll)
}

func main() {
	// Can never be removed
	fmt.Println("Flo ist super!")

	http.HandleFunc("/save", SavePoll)
	http.HandleFunc("/view", ViewPoll)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func SavePoll(w http.ResponseWriter, r *http.Request) {
	poll := Poll{Id: 1,
		Title: r.FormValue("Title")}
	polls[1] = poll
	http.Redirect(w, r, "/view", http.StatusFound)
}

func ViewPoll(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/poll.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, polls[1]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
