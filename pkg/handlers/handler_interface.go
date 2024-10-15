package handlers

import (
	"net/http"
)

type TaskHandler interface {
	GetTasks(w http.ResponseWriter, req *http.Request)
	GetTask(w http.ResponseWriter, req *http.Request)
	CreateTask(w http.ResponseWriter, req *http.Request)
	PutTask(w http.ResponseWriter, req *http.Request)
	DeleteTask(w http.ResponseWriter, req *http.Request)
}

type AccountHandler interface {
	CreateAccount(w http.ResponseWriter, req *http.Request)
	//UpdateUser() error
	//DeleteUser() error
	//ReadUser() error
}
