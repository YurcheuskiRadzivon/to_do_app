package routes

import (
	"github.com/YurcheuskiRadzivon/to_do_app/pkg/handlers"
	"github.com/gorilla/mux"
)

func NewMuxRoute() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.MainPage)
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.GetTask).Methods("GET")
	r.HandleFunc("/task", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id:[0-9]+}/update", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/task/{id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")
	return r
}
