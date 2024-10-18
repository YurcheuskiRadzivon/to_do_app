package handlers

import (
	"net/http"
	"path/filepath"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..", "templates", "home.html")
	tmpl := ParseFiles(w, tmplPath)
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..", "templates", "login.html")
	tmpl := ParseFiles(w, tmplPath)
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
func Registration(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..", "templates", "registration.html")
	tmpl := ParseFiles(w, tmplPath)
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
func Tasks(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..", "templates", "tasks.html")
	tmpl := ParseFiles(w, tmplPath)
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
