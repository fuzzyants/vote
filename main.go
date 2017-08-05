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

	mux := http.NewServeMux()

	// serve static files from public folder
	files := http.FileServer(http.Dir("public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/save", SavePoll)
	mux.HandleFunc("/view", ViewPoll)

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	// go
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/create_form.html",
		"templates/steps.html"}
	templates := template.Must(template.ParseFiles(files...))

	// for now render without data
	templates.ExecuteTemplate(w, "layout", nil)

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
