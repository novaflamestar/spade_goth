package handler

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupStyles(app *echo.Echo, styleDir embed.FS) {
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "styles",
		Filesystem: http.FS(styleDir),
	}))
}

func SetupRoutes(app *echo.Echo, h *UserHandler) {
	app.GET("/", h.BasePage)
	app.GET("/fragments/users/form", h.AddUserForm)
	app.POST("/fragments/users/form", h.SubmitUserForm)
	app.GET("/fragments/users/table", h.GetUserTable)
	app.GET("/fragments/users", h.ShowUsers)
}
