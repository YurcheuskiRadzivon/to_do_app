package handlers
import (
	"encoding/json"
	"fmt"
	
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/YurcheuskiRadzivon/to_do_app/pkg/task"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)
type FileTaskService struct {
	tasks task.Tasks
	mu    sync.Mutex
}
func (fts *FileTaskService) OpenJson() {
	tmplPath := filepath.Join("..", "pkg", "task", "tasks.json")
	file, err := os.ReadFile(tmplPath)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		return
	}
	if err := easyjson.Unmarshal(file, &fts.tasks); err != nil {
		log.Println("Ошибка при декодировании JSON:", err)
		return
	}
}
func (fts *FileTaskService)CloseJson() {
	tmplPath := filepath.Join("..", "pkg", "task", "tasks.json")
	updatedData, err := easyjson.Marshal(fts.tasks)
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
func(fts *FileTaskService) GetTasks(w http.ResponseWriter, req *http.Request) {
	tmplPathHtml := filepath.Join("..", "templates", "gettask_page.html")
	tmpl := ParseFiles(w, tmplPathHtml)
	fts.OpenJson()
	err := tmpl.Execute(w, fts.tasks.List)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}

}
func (fts *FileTaskService) GetTask(w http.ResponseWriter, req *http.Request) {
	var (
		t task.Task
		b bool = false
	)

	vars := mux.Vars(req)
	id := vars["id"]
	fts.OpenJson()
	for i, val := range fts.tasks.List {
		if fmt.Sprintf("%v", val.ID) == id {
			t = fts.tasks.List[i]
			b = true

		}
	}
	fts.CloseJson()
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


func (fts *FileTaskService) CreateTask(w http.ResponseWriter, req *http.Request) {
	var partTask task.TaskInput
	var taskVal task.Task

	if err := json.NewDecoder(req.Body).Decode(&partTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fts.OpenJson()
	taskVal = task.CreateT(partTask)
	fts.mu.Lock()
	fts.tasks.List = append(fts.tasks.List, taskVal)
	fts.mu.Unlock()
	fts.CloseJson()
	fmt.Println("Файл успешно обновлен")

}
func (fts *FileTaskService) DeleteTask(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fts.OpenJson()
	for i, val := range fts.tasks.List {
		if fmt.Sprintf("%v", val.ID) == id {
			fts.mu.Lock()
			fts.tasks.List = append(fts.tasks.List[:i], fts.tasks.List[i+1:]...)
			fts.mu.Unlock()
		}
	}
	fts.CloseJson()

	fmt.Println("Файл успешно обновлен")

}