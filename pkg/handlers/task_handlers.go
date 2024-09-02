package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func GetTasks(w http.ResponseWriter, req *http.Request) {

}
func MainPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..", "templates", "main_page.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
func GetTask(w http.ResponseWriter, req *http.Request) {

}
func CreateTask(w http.ResponseWriter, req *http.Request) {

}
func UpdateTask(w http.ResponseWriter, req *http.Request) {

}
func DeleteTask(w http.ResponseWriter, req *http.Request) {

}
