package main

import (
	"encoding/json"
	"net/http"
)

type GroupMember struct {
	ID      int `json:"id"`
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
}

func (app *application) addMember(w http.ResponseWriter, r *http.Request) {
	var gm GroupMember

	json.NewDecoder(r.Body).Decode(&gm)

	_, err := app.DB.Exec(
		"INSERT INTO group_members (user_id, group_id) VALUES ($1, $2)",
		gm.UserID, gm.GroupID,
	)

	if err != nil {
		errorJSON(w, err)
		return
	}

	writeJSON(w, 201, "Member added")
}

func (app *application) getGroupMembers(w http.ResponseWriter, r *http.Request) {
	groupID := r.URL.Query().Get("group_id")

	rows, err := app.DB.Query(`
		SELECT u.id, u.name, u.email
		FROM users u
		JOIN group_members gm ON u.id = gm.user_id
		WHERE gm.group_id = $1
	`, groupID)

	if err != nil {
		errorJSON(w, err)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	writeJSON(w, 200, users)
}
