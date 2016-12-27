package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("notes/add", addNote)
	r.HandleFunc("notes/save", saveNote)
	r.HandleFunc("notes/edit/{id}", editNote)
	r.HandleFunc("notes/update/{id}", updateNote)
	r.HandleFunc("notes/delete/{id}", deleteNote)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
