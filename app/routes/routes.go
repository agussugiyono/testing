package routes

import (
	_userData "github.com/agussugiyono/coba/features/user/data"
	_userHandler "github.com/agussugiyono/coba/features/user/handler"
	_userService "github.com/agussugiyono/coba/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	userData := _userData.New(db)

	userService := _userService.New(userData)

	userHandlerAPI := _userHandler.New(userService)

	// // Register middleware
	// jwtMiddleware := middlewares.JWTMiddleware()

	// // User Routes

	e.POST("/users", userHandlerAPI.CreateUser)

}
