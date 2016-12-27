package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

// Handler for `HTTP GET /notes/edit/{id}`
func editNote(w http.ResponseWriter, r *http.Request) {
	var viewModel EditNote
	vars := mux.Vars(r)
	k := vars["id"]
	if note, ok := noteStore[k]; ok {
		viewModel = EditNote{note, k}
	} else {
		http.Error(w, "Could not find resource to edit.", http.StatusBadRequest)
	}
	renderTemplate(w, "edit", "base", viewModel)
}

// Handler for `HTTP GET /notes/update/{id}`
func updateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	if note, ok := noteStore[k]; ok {
		noteStore[k] = Note{r.PostFormValue("title"), r.PostFormValue("description"), note.CreatedOn}
	} else {
		http.Error(w, "Could note find the resource to update.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// Handler for `HTTP GET /notes/delete/{id}`
func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		http.Error(w, "Could not find the resource to delete.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
