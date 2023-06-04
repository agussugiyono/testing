package service

import (
	"github.com/agussugiyono/coba/features/user"
	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// Create implements user.UserServiceInterface
func (*userService) Create(user *user.Core) error {
	panic("unimplemented")
}

// Create implements user.UserServiceInterface

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return newFunction(repo)
}

func newFunction(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}
