package handlers

import (
	"html/template"
	"net/http"
)

func ParseFiles(w http.ResponseWriter, filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return nil
	}
	return tmpl
}
