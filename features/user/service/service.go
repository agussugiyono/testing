package service

import (
	"fmt"

	"github.com/agussugiyono/coba/features/user"
	"github.com/agussugiyono/coba/helper"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Create implements user.UserServiceInterface
func (service *userService) Create(user *user.Core) error {
	// Lakukan validasi jika hanya admin atau manager yang dapat menambahkan pengguna
	if user.Role != helper.NewUserRole("admin") && user.Team != helper.NewUserTeam("manager") {
		return fmt.Errorf("only admin and manager can add users")
	}

	// Insert the user data into the database
	err := service.userData.Insert(user)
	if err != nil {
		return err
	}

	// Implementasi logika pembuatan pengguna
	// ...

	return nil
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return newFunction(repo)
}

func newFunction(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
