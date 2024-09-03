package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/task"
	"github.com/mailru/easyjson"
)

var (
	tasks task.Tasks
	mu    sync.Mutex
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
func CreateTask(w http.ResponseWriter, req *http.Request) {
	var task task.Task

	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmplPath := filepath.Join("..", "pkg", "task", "tasks.json")
	file, err := os.ReadFile(tmplPath)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		return
	}
	if err := easyjson.Unmarshal(file, &tasks); err != nil {
		log.Println("Ошибка при декодировании JSON:", err)
		return
	}
	tasks.List = append(tasks.List, task)
	updatedData, err := easyjson.Marshal(tasks)
	if err != nil {
		log.Println("Ошибка при кодировании JSON:", err)
		return
	}
	err = os.WriteFile(tmplPath, updatedData, 0644)
	if err != nil {
		log.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Файл успешно обновлен")

}
func UpdateTask(w http.ResponseWriter, req *http.Request) {

}
func DeleteTask(w http.ResponseWriter, req *http.Request) {

}