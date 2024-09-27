package routes

import (
	"github.com/YurcheuskiRadzivon/to_do_app/pkg/handlers"
	"github.com/gorilla/mux"
)

func NewMuxRoute(taskHandler handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.MainPage)
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST").Headers("Content-Type", "application/json")
	//r.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")
	return r
}
