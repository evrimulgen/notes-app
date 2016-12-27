package main

import (
	"net/http"
)

// Handler for "/"
func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", noteStore)
}

// Handler for "/notes/add"
func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

func saveNote(w http.ResponseWriter, r *http.Request)   {}
func editNote(w http.ResponseWriter, r *http.Request)   {}
func updateNote(w http.ResponseWriter, r *http.Request) {}
func deleteNote(w http.ResponseWriter, r *http.Request) {}
