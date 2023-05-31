package router

import (
	_userData "be17/cleanarch/features/user/data"
	_userHandler "be17/cleanarch/features/user/handler"
	_userService "be17/cleanarch/features/user/service"
	"net/http"

	"be17/cleanarch/app/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	// userData := _userData.NewRaw(dbMysql)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.GET("/users", userHandlerAPI.GetAllUser, middlewares.JWTMiddleware())
	e.POST("/users", userHandlerAPI.CreateUser)
	e.POST("/login", userHandlerAPI.Login)

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello world",
		})
	})

	e.GET("/world", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "world",
		})
	})
}
