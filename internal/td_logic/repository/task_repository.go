package repository

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"
	dberrors "github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/utils/db_errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TaskRepository interface {
	GetTasks(userId int) ([]model.TaskH, error)
	GetTask(id int, userId int) (*model.TaskH, error)
	InsertTask(Task model.TaskH) error
	UpdateTask(TaskH model.TaskH) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db  *pgxpool.Pool
	lgr *logger.Logger
}

func NewTaskRepository(dsnStr string, lgr *logger.Logger) (TaskRepository, error) {
	dsn := fmt.Sprintf(dsnStr)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		lgr.ErrorLogger.Println("Failed to connect to the database:", err)
		return nil, err
	}
	lgr.InfoLogger.Println("SongRepository created successfully.")
	return &taskRepository{
		db:  db,
		lgr: lgr,
	}, nil
}

func (tr *taskRepository) GetTasks(userId int) ([]model.TaskH, error) {
	tr.lgr.DebugLogger.Printf("Getting all tasks from the database with userId: %v ...", userId)
	var tasks []model.TaskH
	query := `SELECT id, title, description, status, added_time, images, user_id FROM "Task" WHERE user_id=$1`
	rows, err := tr.db.Query(context.Background(), query, userId)
	if err != nil {
		tr.lgr.ErrorLogger.Println("Error querying tasks:", err)
		return nil, dberrors.UserError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task model.TaskH
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.AddedTime, &task.Images, &task.UserId)
		if err != nil {
			tr.lgr.ErrorLogger.Println("Error scanning task row:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if rows.Err() != nil {

		tr.lgr.ErrorLogger.Println("Row iteration error:", rows.Err())
		return nil, rows.Err()
	}
	tr.lgr.InfoLogger.Printf("Retrieved %d tasks from the database with user ID %v  \n", len(tasks), userId)
	return tasks, nil
}

func (tr *taskRepository) GetTask(id int, userId int) (*model.TaskH, error) {
	var TaskH model.TaskH
	query := `SELECT id, title, description, status, added_time, images, user_id FROM "Task" WHERE id=$1 AND user_id =$2 ;`
	err := tr.db.QueryRow(context.Background(), query, id, userId).Scan(&TaskH.ID, &TaskH.Title, &TaskH.Description, &TaskH.Status, &TaskH.AddedTime, &TaskH.Images, &TaskH.UserId)
	if err != nil {
		return nil, dberrors.UserError(err)
	}
	tr.lgr.InfoLogger.Printf("Retrieved task with ID %d.\n", id)
	return &TaskH, nil
}
func (tr *taskRepository) InsertTask(TaskH model.TaskH) error {

	query := `INSERT INTO "Task" (title, description, status, added_time, images, user_id) VALUES ($1, $2, $3, NOW(), $4, $5);`
	_, err := tr.db.Exec(context.Background(), query, TaskH.Title, TaskH.Description, TaskH.Status, TaskH.Images, TaskH.UserId)
	if err != nil {
		return dberrors.UserError(err)
	}

	return nil
}
func (tr *taskRepository) UpdateTask(TaskH model.TaskH) error {
	query := `UPDATE "Task" SET title=$1, description=$2, status=$3 WHERE id=$4;`
	_, err := tr.db.Exec(context.Background(), query, TaskH.Title, TaskH.Description, TaskH.Status, TaskH.ID)
	if err != nil {
		return dberrors.UserError(err)
	}
	return nil
}
func (tr *taskRepository) DeleteTask(id int) error {
	tr.lgr.DebugLogger.Printf("Deleting song with ID %d from the database.\n", id)
	query := `DELETE FROM "Task" WHERE id=$1`
	_, err := tr.db.Exec(context.Background(), query, id)
	if err != nil {

		return dberrors.UserError(err)
	}
	return nil
}
