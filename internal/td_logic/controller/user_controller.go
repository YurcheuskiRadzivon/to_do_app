package controller

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/online_music_library/pkg/logger"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/model"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/repository"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/td_logic/utils/jwt_service"

	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	GetUser(ctx context.Context, tokenStr string) (*model.User, error)
	InsertUser(ctx context.Context, User model.User) error
	UpdateUser(ctx context.Context, UserUp model.User, tokenStr string) (string, error)
	DeleteUser(ctx context.Context, tokenStr string) error
	GetUserPassword(ctx context.Context, id int) ([]byte, error)
	LoginUser(ctx context.Context, User *model.User) (string, error)
}

type userController struct {
	repo repository.UserRepository
	lgr  *logger.Logger
}

func NewUserController(repo repository.UserRepository, lgr *logger.Logger) UserController {
	return &userController{
		repo: repo,
		lgr:  lgr,
	}
}

func (uc *userController) GetUser(ctx context.Context, tokenStr string) (*model.User, error) {
	nickname, err := jwt_service.GetUserNickname(tokenStr)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve nickname from token: %s\n", err)
		return nil, err
	}
	email, err := jwt_service.GetEmailNickname(tokenStr)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve email from token: %s\n", err)
		return nil, err
	}
	uc.lgr.DebugLogger.Printf("GetUser called with nickname: %s, email: %s\n", nickname, email)
	User, err := uc.repo.GetUser(nickname, email)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve user with nickname: %s, email: %s: %v\n", nickname, email, err)
		return nil, err
	}
	uc.lgr.InfoLogger.Printf("Retrieved user with nickname: %s, email: %s\n", nickname, email)
	return User, nil
}

func (uc *userController) InsertUser(ctx context.Context, User model.User) error {
	uc.lgr.DebugLogger.Printf("InsertUser called with user: %+v\n", User)
	var UserH model.UserHash

	UserH.Email, UserH.Nickname = User.Email, User.Nickname
	pass, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to hash password: %v\n", err)
		return err
	}
	UserH.Password = pass
	if err = uc.repo.InsertUser(UserH); err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to insert user: %v\n", err)
		return err
	}
	uc.lgr.InfoLogger.Printf("Inserted user with email: %s, nickname: %s\n", User.Email, User.Nickname)
	return nil
}

func (uc *userController) UpdateUser(ctx context.Context, UserUp model.User, tokenStr string) (string, error) {
	id, err := jwt_service.GetUserId(tokenStr)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve userId from token: %s\n", err)
		return tokenStr, err
	}
	uc.lgr.DebugLogger.Printf("UpdateUser called with id: %d, user: %+v\n", id, UserUp)
	if err := uc.repo.UpdateUser(id, UserUp); err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to update user with id: %d: %v\n", id, err)
		return tokenStr, err
	}
	payload := jwt.MapClaims{
		"email":  UserUp.Email,
		"name":   UserUp.Nickname,
		"sub_id": id,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	t, err := jwt_service.CreateToken(payload)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to create token for user with id: %d: %v\n", id, err)
		return tokenStr, err
	}
	uc.lgr.InfoLogger.Printf("Updated user with id: %d\n", id)
	return t, nil
}

func (uc *userController) DeleteUser(ctx context.Context, tokenStr string) error {
	id, err := jwt_service.GetUserId(tokenStr)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve userId from token: %s\n", err)
		return err
	}
	uc.lgr.DebugLogger.Printf("DeleteUser called with id: %d\n", id)
	if err := uc.repo.DeleteUser(id); err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to delete user with id: %d: %v\n", id, err)
		return err
	}
	uc.lgr.InfoLogger.Printf("Deleted user with id: %d\n", id)
	return nil
}

func (uc *userController) GetUserPassword(ctx context.Context, id int) ([]byte, error) {
	uc.lgr.DebugLogger.Printf("GetUserPassword called with id: %d\n", id)
	pass, err := uc.repo.GetUserPassword(id)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve password for user with id: %d: %v\n", id, err)
		return nil, err
	}
	uc.lgr.InfoLogger.Printf("Retrieved password for user with id: %d\n", id)
	return pass, nil
}

func (uc *userController) LoginUser(ctx context.Context, User *model.User) (string, error) {
	uc.lgr.DebugLogger.Printf("LoginUser called with user: %+v\n", User)
	U, err := uc.repo.GetUser(User.Nickname, User.Email)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve user with nickname: %s, email: %s: %v\n", User.Nickname, User.Email, err)
		return "", fmt.Errorf("failed to retrieve user: %v", err)
	}
	pass, err := uc.GetUserPassword(ctx, U.ID)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to retrieve password for user with id: %d: %v\n", U.ID, err)
		return "", fmt.Errorf("failed to retrieve password: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword(pass, []byte(User.Password)); err != nil {
		uc.lgr.ErrorLogger.Printf("Password mismatch for user with id: %d: %v\n", U.ID, err)
		return "", fmt.Errorf("password mismatch: %v", err)
	}
	payload := jwt.MapClaims{
		"email":  User.Email,
		"name":   User.Nickname,
		"sub_id": U.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	t, err := jwt_service.CreateToken(payload)
	if err != nil {
		uc.lgr.ErrorLogger.Printf("Failed to create token for user with id: %d: %v\n", U.ID, err)
		return "", fmt.Errorf("failed to create token: %v", err)
	}
	uc.lgr.InfoLogger.Printf("User logged in with id: %d\n", U.ID)
	return t, nil
}
