package controller

import (
	"context"
	"fmt"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/model"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/repository"
	"github.com/YurcheuskiRadzivon/to_do_app/internal/tda_logic/utils/jwt_service"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	GetUser(ctx context.Context, nickname, email string) (*model.User, error)
	InsertUser(ctx context.Context, User model.User) error
	UpdateUser(ctx context.Context, id int, UserUp model.User) (string, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserPassword(ctx context.Context, id int) ([]byte, error)
	LoginUser(ctx context.Context, User *model.User) (string, error)
}
type userController struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) UserController {
	return &userController{repo: repo}
}
func (uc *userController) GetUser(ctx context.Context, nickname, email string) (*model.User, error) {
	User, err := uc.repo.GetUser(nickname, email)
	if err != nil {
		return nil, err
	}
	return User, err
}
func (uc *userController) InsertUser(ctx context.Context, User model.User) error {
	var UserH model.UserHash

	UserH.Email, UserH.Nickname = User.Email, User.Nickname
	var pass []byte
	pass, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	UserH.Password = pass
	if err = uc.repo.InsertUser(UserH); err != nil {
		return err
	}

	return nil
}
func (uc *userController) UpdateUser(ctx context.Context, id int, UserUp model.User) (string, error) {

	if err := uc.repo.UpdateUser(id, UserUp); err != nil {
		return "", err
	}
	payload := jwt.MapClaims{

		"email":  UserUp.Email,
		"name":   UserUp.Nickname,
		"sub_id": UserUp.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	t, err := jwt_service.CreateToken(payload)
	if err != nil {
		return "", err
	}
	return t, nil
}
func (uc *userController) DeleteUser(ctx context.Context, id int) error {

	if err := uc.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
func (uc *userController) GetUserPassword(ctx context.Context, id int) ([]byte, error) {
	pass, err := uc.repo.GetUserPassword(id)
	if err != nil {
		return nil, err
	}
	return pass, err
}
func (uc *userController) LoginUser(ctx context.Context, User *model.User) (string, error) {
	U, err := uc.repo.GetUser(User.Nickname, User.Email)

	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	pass, err := uc.GetUserPassword(ctx, U.ID)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	if err := bcrypt.CompareHashAndPassword(pass, []byte(User.Password)); err != nil {
		return "", fmt.Errorf("%s", err)
	}
	payload := jwt.MapClaims{

		"email":  User.Email,
		"name":   User.Nickname,
		"sub_id": U.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
	t, err := jwt_service.CreateToken(payload)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	return t, nil

}
