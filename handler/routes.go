package handler

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo, h *UserHandler) {
	app.GET("/", h.BasePage)
	app.GET("/users/form", h.AddUserForm)
	app.POST("/users/form", h.SubmitUserForm)
	app.GET("/components/users/table", h.GetUserTable)
	app.GET("/components/users", h.ShowUsers)
}
