package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/repository"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/utils/jwt_service"
	"github.com/jung-kurt/gofpdf"
	"sort"
	"strings"
)

type TaskController interface {
	GetTasks(ctx context.Context, tokenString string, sortParam string) ([]model.Task, error)
	GetTask(ctx context.Context, id int, tokenString string) (*model.Task, error)
	InsertTask(ctx context.Context, Task model.Task, tokenString string) error
	UpdateTask(ctx context.Context, Task model.Task, tokenString string, id int) error
	DeleteTask(ctx context.Context, id int) error
	ExportTasks(ctx context.Context, tokenString string) ([]byte, error)
}

type taskController struct {
	repo repository.TaskRepository
	lgr  *logger.Logger
}

func NewTaskController(repo repository.TaskRepository, lgr *logger.Logger) TaskController {
	return &taskController{
		repo: repo,
		lgr:  lgr,
	}
}

func (tc *taskController) GetTasks(ctx context.Context, tokenString string, sortParam string) ([]model.Task, error) {
	tc.lgr.DebugLogger.Printf("GetTasks called with sortParam: %s\n", sortParam)
	UserId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to get user ID from token: %v\n", err)
		return nil, err
	}
	TaskHList, err := tc.repo.GetTasks(UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %v", err)
	}
	var taskList []model.Task
	for _, value := range TaskHList {
		var images []string
		err = json.Unmarshal([]byte(value.Images), &images)
		if err != nil {
			tc.lgr.ErrorLogger.Printf("Failed to unmarshal images: %v\n", err)
			return nil, err
		}
		Task := model.Task{
			ID:          value.ID,
			Description: value.Description,
			Title:       value.Title,
			Status:      value.Status,
			AddedTime:   value.AddedTime,
			Images:      images,
		}

		taskList = append(taskList, Task)
	}

	allowedSorts := map[string]bool{
		"status": true,
		"name":   true,
		"date":   true,
	}
	if !allowedSorts[sortParam] {
		sortParam = "date"
		tc.lgr.DebugLogger.Printf("Invalid sort parameter: %s, defaulting to date\n", sortParam)
	}
	tc.lgr.DebugLogger.Printf("Sorting tasks by %s\n", sortParam)
	switch sortParam {
	case "date":
		sort.Slice(taskList, func(i, j int) bool {
			return taskList[i].ID < taskList[j].ID
		})
	case "status":
		sort.Slice(taskList, func(i, j int) bool {
			return taskList[i].Status == false && taskList[j].Status == true
		})
	case "name":
		sort.Slice(taskList, func(i, j int) bool {
			return strings.ToLower(taskList[i].Title) < strings.ToLower(taskList[j].Title)
		})
	}
	tc.lgr.DebugLogger.Printf("Sorting tasks by %s\n", sortParam)
	return taskList, nil
}

func (tc *taskController) GetTask(ctx context.Context, id int, tokenString string) (*model.Task, error) {
	userId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to get user ID from token: %v\n", err)
		return nil, err
	}
	tc.lgr.DebugLogger.Printf("GetTask called with id: %d\n", id)
	TaskH, err := tc.repo.GetTask(id, userId)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to retrieve task with ID %d: %v\n", id, err)
		return nil, err
	}
	var images []string
	err = json.Unmarshal([]byte(TaskH.Images), &images)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to unmarshal images for task ID %d: %v\n", id, err)
		return nil, err
	}
	Task := &model.Task{
		ID:          TaskH.ID,
		Description: TaskH.Description,
		Title:       TaskH.Title,
		Status:      TaskH.Status,
		AddedTime:   TaskH.AddedTime,
		Images:      images,
	}

	return Task, nil
}

func (tc *taskController) InsertTask(ctx context.Context, Task model.Task, tokenString string) error {
	tc.lgr.DebugLogger.Printf("InsertTask called with task: %+v\n", Task)
	imagesJSON, err := json.Marshal([]string{})
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to marshal images: %v\n", err)
		return err
	}
	UserId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to get user ID from token: %v\n", err)
		return err
	}
	TaskH := model.TaskH{
		ID:          Task.ID,
		Title:       Task.Title,
		Description: Task.Description,
		Status:      Task.Status,
		Images:      imagesJSON,
		UserId:      UserId,
	}
	if err = tc.repo.InsertTask(TaskH); err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to insert task: %v\n", err)
		return err
	}
	tc.lgr.InfoLogger.Printf("Inserted task with Tittle %s for user ID %d\n", Task.Title, UserId)
	return nil
}

func (tc *taskController) UpdateTask(ctx context.Context, Task model.Task, tokenString string, id int) error {
	tc.lgr.DebugLogger.Printf("UpdateTask called with task: %+v\n", Task)

	imagesJSON, err := json.Marshal([]string{})
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to marshal images: %v\n", err)
		return err
	}
	UserId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to get user ID from token: %v\n", err)
		return err
	}
	TaskH := model.TaskH{
		ID:          id,
		Title:       Task.Title,
		Description: Task.Description,
		Status:      Task.Status,
		Images:      imagesJSON,
		UserId:      UserId,
	}
	if err = tc.repo.UpdateTask(TaskH); err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to update task with ID %d: %v\n", Task.ID, err)
		return err
	}
	tc.lgr.InfoLogger.Printf("Updated task with ID %d for user ID %d\n", Task.ID, UserId)
	return nil
}

func (tc *taskController) DeleteTask(ctx context.Context, id int) error {
	tc.lgr.DebugLogger.Printf("DeleteTask called with id: %d\n", id)
	if err := tc.repo.DeleteTask(id); err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to delete task with ID %d: %v\n", id, err)
		return err
	}
	tc.lgr.InfoLogger.Printf("Deleted task with ID %d\n", id)
	return nil
}
func (tc *taskController) ExportTasks(ctx context.Context, tokenString string) ([]byte, error) {
	tc.lgr.DebugLogger.Println("ExportTasks called")

	// Получаем список задач
	tasks, err := tc.GetTasks(ctx, tokenString, "date")
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to retrieve tasks: %v\n", err)
		return nil, err
	}

	// Создаем PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Exported Tasks")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	for _, task := range tasks {
		pdf.Cell(40, 10, fmt.Sprintf("Title: %s", task.Title))
		pdf.Ln(6)
		pdf.Cell(40, 10, fmt.Sprintf("Description: %s", task.Description))
		pdf.Ln(6)
		pdf.Cell(40, 10, fmt.Sprintf("Status: %t", task.Status))
		pdf.Ln(10)
	}

	var pdfBytes bytes.Buffer
	err = pdf.Output(&pdfBytes)
	if err != nil {
		tc.lgr.ErrorLogger.Printf("Failed to generate PDF: %v\n", err)
		return nil, err
	}

	return pdfBytes.Bytes(), nil
}
