package main

import (
	"net/http"
	"strconv"
	"time"
)

// Handler for `HTTP GET /`
func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", noteStore)
}

// Handler for `HTTP GET /notes/add`
func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

// Handler for `HTTP POST /notes/save`
func saveNote(w http.ResponseWriter, r *http.Request) {
	note := Note{r.PostFormValue("title"), r.PostFormValue("description"), time.Now()}
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	http.Redirect(w, r, "/", http.StatusFound)
}

func editNote(w http.ResponseWriter, r *http.Request)   {}
func updateNote(w http.ResponseWriter, r *http.Request) {}
func deleteNote(w http.ResponseWriter, r *http.Request) {}
