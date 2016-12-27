package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html", "templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
