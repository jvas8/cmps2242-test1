package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (app *application) getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := app.DB.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		users = append(users, u)
	}

	writeJSON(w, 200, users)
}

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var u User

	json.NewDecoder(r.Body).Decode(&u)

	err := app.DB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at",
		u.Name, u.Email,
	).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 201, u)
}
