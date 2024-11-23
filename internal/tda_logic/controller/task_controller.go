package controller

import (
	"context"
	"encoding/json"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/model"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/repository"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/utils/jwt_service"
)

type TaskController interface {
	GetTasks(ctx context.Context, tokenString string) ([]model.Task, error)
	GetTask(ctx context.Context, id int) (*model.Task, error)
	InsertTask(ctx context.Context, Task model.Task, tokenString string) error
	UpdateTask(ctx context.Context, Task model.Task, tokenString string) error
	DeleteTask(ctx context.Context, id int) error
}
type taskController struct {
	repo repository.TaskRepository
}

func NewTaskController(repo repository.TaskRepository) TaskController {
	return &taskController{repo: repo}
}

func (tc *taskController) GetTasks(ctx context.Context, tokenString string) ([]model.Task, error) {
	UserId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
		return nil, err
	}
	TaskHList, err := tc.repo.GetTasks(UserId)
	if err != nil {
		return nil, err
	}
	var TaskList []model.Task
	for _, value := range TaskHList {
		var images []string
		err = json.Unmarshal([]byte(value.Images), &images)
		if err != nil {
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
		TaskList = append(TaskList, Task)
	}

	return TaskList, nil
}
func (tc *taskController) GetTask(ctx context.Context, id int) (*model.Task, error) {
	TaskH, err := tc.repo.GetTask(id)
	if err != nil {
		return nil, err
	}
	var images []string
	err = json.Unmarshal([]byte(TaskH.Images), &images)
	if err != nil {
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
	imagesJSON, err := json.Marshal(Task.Images)
	if err != nil {
		return err
	}
	UserId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
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

		return err
	}
	return nil
}
func (tc *taskController) UpdateTask(ctx context.Context, Task model.Task, tokenString string) error {
	imagesJSON, err := json.Marshal(Task.Images)
	if err != nil {
		return err
	}
	UserId, err := jwt_service.GetUserId(tokenString)
	if err != nil {
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
	if err = tc.repo.UpdateTask(TaskH); err != nil {
		return err
	}
	return nil

}
func (tc *taskController) DeleteTask(ctx context.Context, id int) error {
	if err := tc.repo.DeleteTask(id); err != nil {
		return err
	}
	return nil
}
