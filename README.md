# Study Group Planner API

**Name:** Jeimy Vasquez

## Description

This project is a REST API built in Go as part of my final project.
It serves as the backend for a Study Group Planner system where users can manage study groups, members, subjects, and study sessions.

A graphical user interface (GUI) will be developed in the future to interact with this API.

## Features

* Manage users
* Manage study groups
* Manage subjects
* Add members to groups
* View group members
* Manage study sessions

## How to Run

1. Ensure PostgreSQL is running
2. Update the database connection string in `main.go`
3. Run the API:

   ```
   go run ./cmd/api
   ```
4. The server will start on:

   ```
   http://localhost:8080
   ```

## Testing Endpoints (Using cURL)

### Get all users

```
curl http://localhost:8080/users
```

### Create a user

```
curl -X POST http://localhost:8080/users/create \
-H "Content-Type: application/json" \
-d '{"name":"Jeimy Vasquez","email":"jeimy@test.com"}'
```

---

### Get all subjects

```
curl http://localhost:8080/subjects
```

### Create a subject

```
curl -X POST http://localhost:8080/subjects/create \
-H "Content-Type: application/json" \
-d '{"name":"Physics","description":"Study of motion"}'
```

---

### Get all study groups

```
curl http://localhost:8080/groups
```

### Create a study group

```
curl -X POST http://localhost:8080/groups/create \
-H "Content-Type: application/json" \
-d '{"name":"Exam Prep","description":"Finals review","creator_id":1,"subject_id":1}'
```

---

### Add member to group

```
curl -X POST http://localhost:8080/groups/add-member \
-H "Content-Type: application/json" \
-d '{"user_id":1,"group_id":1}'
```

### Get members in a group

```
curl "http://localhost:8080/groups/members?group_id=1"
```

---

### Get all study sessions

```
curl http://localhost:8080/sessions
```

### Create a study session

```
curl -X POST http://localhost:8080/sessions/create \
-H "Content-Type: application/json" \
-d '{"group_id":1,"title":"Final Review","session_date":"2026-04-10 15:00:00","location":"Library","notes":"Bring notes"}'
```

## Output

All responses are returned in JSON format.
