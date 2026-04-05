package main

import (
	"encoding/json"
	"net/http"
)

type StudySession struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	Title       string `json:"title"`
	SessionDate string `json:"session_date"`
	Location    string `json:"location"`
	Notes       string `json:"notes"`
}

func (app *application) getSessions(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query(`
		SELECT id, group_id, title, session_date, location, notes 
		FROM study_sessions
	`)
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer rows.Close()

	var sessions []StudySession

	for rows.Next() {
		var s StudySession
		rows.Scan(&s.ID, &s.GroupID, &s.Title, &s.SessionDate, &s.Location, &s.Notes)
		sessions = append(sessions, s)
	}

	writeJSON(w, 200, sessions)
}

func (app *application) createSession(w http.ResponseWriter, r *http.Request) {
	var s StudySession

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		errorJSON(w, err)
		return
	}

	err = app.DB.QueryRow(`
		INSERT INTO study_sessions (group_id, title, session_date, location, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, s.GroupID, s.Title, s.SessionDate, s.Location, s.Notes).
		Scan(&s.ID)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 201, s)
}
