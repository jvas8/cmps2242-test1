package main

import (
	"encoding/json"
	"net/http"
)

type StudyGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatorID   int    `json:"creator_id"`
	SubjectID   int    `json:"subject_id"`
	CreatedAt   string `json:"created_at"`
}

func (app *application) getGroups(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query(`
		SELECT id, name, description, creator_id, subject_id, created_at 
		FROM study_groups
	`)
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer rows.Close()

	var groups []StudyGroup

	for rows.Next() {
		var g StudyGroup
		rows.Scan(&g.ID, &g.Name, &g.Description, &g.CreatorID, &g.SubjectID, &g.CreatedAt)
		groups = append(groups, g)
	}

	writeJSON(w, 200, groups)
}

func (app *application) createGroup(w http.ResponseWriter, r *http.Request) {
	var g StudyGroup

	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		errorJSON(w, err)
		return
	}

	err = app.DB.QueryRow(`
		INSERT INTO study_groups (name, description, creator_id, subject_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`, g.Name, g.Description, g.CreatorID, g.SubjectID).
		Scan(&g.ID, &g.CreatedAt)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 201, g)
}
