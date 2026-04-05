package main

import (
	"encoding/json"
	"net/http"
)

type Member struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (app *application) getMembers(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, name, email FROM members")
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer rows.Close()

	var members []Member

	for rows.Next() {
		var m Member
		rows.Scan(&m.ID, &m.Name, &m.Email)
		members = append(members, m)
	}

	writeJSON(w, 200, members)
}

func (app *application) createMember(w http.ResponseWriter, r *http.Request) {
	var m Member

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		errorJSON(w, err)
		return
	}

	err = app.DB.QueryRow(
		"INSERT INTO members (name, email) VALUES ($1, $2) RETURNING id",
		m.Name, m.Email,
	).Scan(&m.ID)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 201, m)
}
func (app *application) getMemberByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var m Member

	err := app.DB.QueryRow(
		"SELECT id, name, email FROM members WHERE id=$1",
		id,
	).Scan(&m.ID, &m.Name, &m.Email)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 200, m)
}

func (app *application) updateMember(w http.ResponseWriter, r *http.Request) {
	var m Member

	json.NewDecoder(r.Body).Decode(&m)

	_, err := app.DB.Exec(
		"UPDATE members SET name=$1, email=$2 WHERE id=$3",
		m.Name, m.Email, m.ID,
	)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 200, m)
}

func (app *application) deleteMember(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, err := app.DB.Exec("DELETE FROM members WHERE id=$1", id)
	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 200, "Deleted")
}
