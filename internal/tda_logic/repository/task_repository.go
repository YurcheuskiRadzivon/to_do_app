package repository

import (
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TaskRepository interface {
	GetTasks(userId int) ([]model.Task, error)
	GetTask(id int) (*model.Task, error)
	InsertTask(Task model.Task) error
	UpdateTask(id int, Task model.Task) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db *pgxpool.Pool
}
