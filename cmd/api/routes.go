package main

import "net/http"

func (app *application) routes(mux *http.ServeMux) {
	// Users
	mux.HandleFunc("/users", app.getUsers)
	mux.HandleFunc("/users/create", app.createUser)

	// Subjects
	mux.HandleFunc("/subjects", app.getSubjects)
	mux.HandleFunc("/subjects/create", app.createSubject)

	// Groups
	mux.HandleFunc("/groups", app.getGroups)
	mux.HandleFunc("/groups/create", app.createGroup)

	// Group Members
	mux.HandleFunc("/groups/add-member", app.addMember)
	mux.HandleFunc("/groups/members", app.getGroupMembers)

	// Study Sessions
	mux.HandleFunc("/sessions", app.getSessions)
	mux.HandleFunc("/sessions/create", app.createSession)
}
