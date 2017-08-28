package main

import (
	"fmt"
	"github.com/fuzzyants/vote/data"
	"html/template"
	"log"
	"net/http"
	"time"
)

func init() {
	data.InitDb(Config.DbHost, Config.DbPort, Config.DbUser, Config.DbPwd)
}

func main() {
	// Can never be removed
	fmt.Println("Flo ist super!")

	mux := http.NewServeMux()

	// serve static files from public folder
	files := http.FileServer(http.Dir("public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", NewVote)
	mux.HandleFunc("/create", CreateVote)
	mux.HandleFunc("/view", ViewVote)

	server := &http.Server{
		Addr:    Config.Address,
		Handler: mux,
	}
	// go
	server.ListenAndServe()
}

// NewVote responds with the empty vote creation form
func NewVote(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/create_form.html",
		"templates/steps.html"}
	templates := template.Must(template.ParseFiles(files...))

	// for now render without data
	templates.ExecuteTemplate(w, "layout", nil)

}

// CreateVote accepts formdata or querydata, creates all structs, stores them
// and redirects the user to the confirmation view
func CreateVote(w http.ResponseWriter, r *http.Request) {
	// TODO: create and save each option struct
	var options []data.Option

	vote := data.Vote{
		MaxUsers:  5,
		Expires:   time.Now().Add(time.Hour * 24 * 7),
		Done:      false,
		Options:   options,
		CreatedAt: time.Now(),
		Title:     r.FormValue("Title")} // TODO: sanitize to prevent injection

	id, err := vote.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: confirmation view, for now just redirect to ballot view
	http.Redirect(w, r, "/view/"+id, http.StatusFound)
	return

}

// ConfirmVote displays all settings for the vote, and allows the user to
// go back to editing or to start the vote (after which it becomes immutable)
// func ConfirmVote(w http.ResponseWriter, r *http.Request) {

// }

// ViewVote displays a ballot form if the vote is still active or the results
// if it is expired or MaxUsers have voted
func ViewVote(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/vote.html")

	// TODO: load vote from database by id
	var vote data.Vote

	if err != nil {
		log.Fatal("Error parsing template vote.html: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	if err := tmpl.Execute(w, vote); err != nil {
		log.Fatal("Error executing template vote.html: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
