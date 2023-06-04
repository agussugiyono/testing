package handler

import (
	"net/http"
	"strings"

	"github.com/agussugiyono/coba/features/user"
	"github.com/agussugiyono/coba/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := UserRequest{}

	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	// Pengecekan peran pengguna
	if userInput.Role != "admin" && userInput.Role != "manager" {
		return c.JSON(http.StatusForbidden, helper.FailedResponse("Only admin and manager can add users"))
	}

	userCore := &user.Core{
		Name:     userInput.Name,
		Phone:    userInput.Phone,
		Email:    userInput.Email,
		Password: userInput.Password,
		Role:     user.UserRole(userInput.Role),
		Status:   user.UserStatus(userInput.Status), Team: user.UserTeam(userInput.Team),
	}

	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}
