package handler

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo, h *UserHandler) {
	app.GET("/", h.BasePage)
	app.GET("/fragments/users/form", h.AddUserForm)
	app.POST("/fragments/users/form", h.SubmitUserForm)
	app.GET("/fragments/users/table", h.GetUserTable)
	app.GET("/fragments/users", h.ShowUsers)
}
