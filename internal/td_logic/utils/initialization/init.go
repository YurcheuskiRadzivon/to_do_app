package initia

import (
	"fmt"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/controller"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/handler"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/repository"

	"github.com/pkg/errors"
)

func InitializeComponentsUser(dsnStr string) (handler.UserHandler, error) {
	userRepo, err := repository.NewUserRepository(dsnStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error DB connection: %v", err))
	}
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	return userHandler, nil
}
func InitializeComponentsTask(dsnStr string) (handler.TaskHandler, error) {
	taskRepo, err := repository.NewTaskRepository(dsnStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error DB connection: %v", err))
	}
	taskController := controller.NewTaskController(taskRepo)
	taskHandler := handler.NewTaskHandler(taskController)
	return taskHandler, nil
}
