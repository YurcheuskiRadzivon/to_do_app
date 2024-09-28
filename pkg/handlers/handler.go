package handlers

import (
	"net/http"
)

type TaskHandler interface {
	GetTasks(w http.ResponseWriter, req *http.Request)
	GetTask(w http.ResponseWriter, req *http.Request)
	CreateTask(w http.ResponseWriter, req *http.Request)
	DeleteTask(w http.ResponseWriter, req *http.Request)
}

//
/*
	curl -X POST http://localhost:8080/task \
	     -H "Content-Type: application/json" \
	     -d '{
	           "id": 1,
	           "title": "Learn Go",
	           "notes": "Complete the Go tutorial",
	           "completed": false,
	           "priority": 1
	         }'
*/
