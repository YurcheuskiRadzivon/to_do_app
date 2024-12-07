package initia

import (
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/controller"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/handler"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/repository"

	"github.com/pkg/errors"
)

func InitializeComponentsUser(dsnStr string, lgr *logger.Logger) (handler.UserHandler, error) {
	userRepo, err := repository.NewUserRepository(dsnStr, lgr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error DB connection: %v", err))
	}
	userController := controller.NewUserController(userRepo, lgr)
	userHandler := handler.NewUserHandler(userController, lgr)
	return userHandler, nil
}
func InitializeComponentsTask(dsnStr string, lgr *logger.Logger) (handler.TaskHandler, error) {
	taskRepo, err := repository.NewTaskRepository(dsnStr, lgr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DB connection has failed: %v", err))
	}
	taskController := controller.NewTaskController(taskRepo, lgr)
	taskHandler := handler.NewTaskHandler(taskController, lgr)
	return taskHandler, nil
}
