package main

import (
	"time"
)

type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

type EditNote struct {
	Note
	Id string
}

var noteStore = make(map[string]Note)

var id int = 0
