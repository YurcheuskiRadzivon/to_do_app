package handlers

import (
	"net/http"
)

type TaskHandler interface {
	GetTasks(w http.ResponseWriter, req *http.Request)
	//GetTask(w http.ResponseWriter, req *http.Request)
	//CreateTask(w http.ResponseWriter, req *http.Request)
	//PutTask(w http.ResponseWriter, req *http.Request)
	//DeleteTask(w http.ResponseWriter, req *http.Request)
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, req *http.Request)
	LoginUser(w http.ResponseWriter, req *http.Request)
	//UpdateUser() error
	//DeleteUser() error
	//ReadUser() error
}
