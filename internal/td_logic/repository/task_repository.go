package repository

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type TaskRepository interface {
	GetTasks(userId int) ([]model.TaskH, error)
	GetTask(id int) (*model.TaskH, error)
	InsertTask(Task model.TaskH) error
	UpdateTask(TaskH model.TaskH) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(dsnStr string) (TaskRepository, error) {
	dsn := fmt.Sprintf(dsnStr)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
		return nil, err
	}
	return &taskRepository{db: db}, nil
}

func (tr *taskRepository) GetTasks(userId int) ([]model.TaskH, error) {
	var tasks []model.TaskH

	query := `SELECT id, title, description, status, added_time, images, user_id FROM "Task" WHERE user_id=$1`
	rows, err := tr.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.TaskH
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AddedTime, &task.Images, &task.UserId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

func (tr *taskRepository) GetTask(id int) (*model.TaskH, error) {
	var TaskH model.TaskH
	query := `SELECT id, title, description, status, added_time, images, user_id FROM "Task" WHERE id=$1;`
	err := tr.db.QueryRow(context.Background(), query, id).Scan(&TaskH.ID, &TaskH.Title, &TaskH.Status, &TaskH.AddedTime, &TaskH.Images, &TaskH.UserId)
	if err != nil {
		return nil, err
	}
	return &TaskH, nil
}
func (tr *taskRepository) InsertTask(TaskH model.TaskH) error {
	query := `INSERT INTO "Task" (title, description, status, added_time, images, user_id) VALUES ($1, $2, $3, NOW(), $4, $5);`
	_, err := tr.db.Exec(context.Background(), query, TaskH.Title, TaskH.Description, TaskH.Status, TaskH.Images, TaskH.UserId)
	if err != nil {
		return err
	}
	return nil
}
func (tr *taskRepository) UpdateTask(TaskH model.TaskH) error {
	query := `UPDATE "Task" SET title=$1, description=$2, status=$3, added_time=$4, images=$5 WHERE id=$7;`
	_, err := tr.db.Exec(context.Background(), query, TaskH.Title, TaskH.Description, TaskH.Status, TaskH.AddedTime, TaskH.Images, TaskH.ID)
	if err != nil {
		return err
	}
	return nil
}
func (tr *taskRepository) DeleteTask(id int) error {
	query := `DELETE FROM "Task" WHERE id=$1`
	_, err := tr.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
