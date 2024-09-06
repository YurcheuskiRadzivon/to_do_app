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
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

var (
	tasks task.Tasks
	mu    sync.Mutex
)

func ParseFiles(w http.ResponseWriter, filename string) *template.Template {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return nil
	}
	return tmpl
}
func OpenJson() {
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
}
func CloseJson() {
	tmplPath := filepath.Join("..", "pkg", "task", "tasks.json")
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
}
func GetTasks(w http.ResponseWriter, req *http.Request) {
	tmplPathHtml := filepath.Join("..", "templates", "gettask_page.html")
	tmpl := ParseFiles(w, tmplPathHtml)
	OpenJson()
	err := tmpl.Execute(w, tasks.List)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}

}
func MainPage(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("..", "templates", "main_page.html")
	tmpl := ParseFiles(w, tmplPath)
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
func GetTask(w http.ResponseWriter, req *http.Request) {
	var (
		t task.Task
		b bool = false
	)

	vars := mux.Vars(req)
	id := vars["id"]
	OpenJson()
	for i, val := range tasks.List {
		if fmt.Sprintf("%v", val.ID) == id {
			t = tasks.List[i]
			b = true

		}
	}
	CloseJson()
	if b == true {
		tmplPathHtml := filepath.Join("..", "templates", "gettaskid_page.html")
		tmpl := ParseFiles(w, tmplPathHtml)
		err := tmpl.Execute(w, t)
		if err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		}
	} else {
		tmplPathHtml := filepath.Join("..", "templates", "gettaskiderror_page.html")
		tmpl := ParseFiles(w, tmplPathHtml)
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		}
	}

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
	var partTask task.TaskInput
	var taskVal task.Task

	if err := json.NewDecoder(req.Body).Decode(&partTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	OpenJson()
	taskVal = task.CreateT(partTask)
	mu.Lock()
	tasks.List = append(tasks.List, taskVal)
	mu.Unlock()
	CloseJson()
	fmt.Println("Файл успешно обновлен")

}
func DeleteTask(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	OpenJson()
	for i, val := range tasks.List {
		if fmt.Sprintf("%v", val.ID) == id {
			mu.Lock()
			tasks.List = append(tasks.List[:i], tasks.List[i+1:]...)
			mu.Unlock()
		}
	}
	CloseJson()

	fmt.Println("Файл успешно обновлен")

}
