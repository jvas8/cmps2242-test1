package main

import (
	"encoding/json"
	"net/http"
)

type Subject struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (app *application) getSubjects(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, name, description FROM subjects")
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer rows.Close()

	var subjects []Subject

	for rows.Next() {
		var s Subject
		rows.Scan(&s.ID, &s.Name, &s.Description)
		subjects = append(subjects, s)
	}

	writeJSON(w, 200, subjects)
}

func (app *application) createSubject(w http.ResponseWriter, r *http.Request) {
	var s Subject

	json.NewDecoder(r.Body).Decode(&s)

	err := app.DB.QueryRow(
		"INSERT INTO subjects (name, description) VALUES ($1, $2) RETURNING id",
		s.Name, s.Description,
	).Scan(&s.ID)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 201, s)
}
