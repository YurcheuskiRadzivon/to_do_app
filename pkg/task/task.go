package task

import (
	json "encoding/json"
	"log"
	"os"
	"path/filepath"
)

//easyjson:json
type (
	Task struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Notes     string `json:"notes"`
		Completed bool   `json:"completed"`
		Priority  int    `json:"priority"`
	}
	TaskInput struct {
		Title     string `json:"title"`
		Notes     string `json:"notes"`
		Completed bool   `json:"completed"`
		Priority  int    `json:"priority"`
	}
	ID struct {
		Value int `json:"id"`
	}

	Tasks struct {
		List []Task `json:"list"`
	}
)

func getIDFromFile() (int, error) {
	var idStruct ID
	tmplPath := filepath.Join("..", "pkg", "task", "id.json")
	data, err := os.ReadFile(tmplPath)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(data, &idStruct)
	if err != nil {
		return 0, err
	}
	newID := idStruct.Value + 1
	log.Println(newID)
	idStruct.Value = newID
	newData, err := json.Marshal(idStruct)
	if err != nil {
		return 0, err
	}
	err = os.WriteFile(tmplPath, newData, 0644)
	if err != nil {
		return 0, err
	}
	return newID, nil
}
func CreateT(input TaskInput) Task {
	id, err := getIDFromFile()
	if err != nil {
		log.Println(err)
	}
	return Task{
		ID:        id,
		Title:     input.Title,
		Notes:     input.Notes,
		Completed: input.Completed,
		Priority:  input.Priority,
	}
}
