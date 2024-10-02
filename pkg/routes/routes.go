package routes

import (
	"net/http"
	"path/filepath"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/handlers"
	"github.com/gorilla/mux"
)

func NewMuxRoute(taskHandler handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	tmplPath := filepath.Join("..", "static")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(tmplPath))))
	r.HandleFunc("/", handlers.MainPage)
	r.HandleFunc("/login", handlers.Login).Methods("GET")
	r.HandleFunc("/registration", handlers.Registration).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST").Headers("Content-Type", "application/json")
	//r.HandleFunc("/tasks/{id:[0-9]+}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")
	return r
}
