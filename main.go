package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/novaflamestar/spade_goth/handler"
	"github.com/novaflamestar/spade_goth/service"
	"github.com/novaflamestar/spade_goth/storage"
)

func main() {
	app := echo.New()

	userStore := storage.InitStore()
	userService := service.InitService(userStore)
	userHandler := handler.InitHandler(userService)

	handler.SetupRoutes(app, userHandler)

	fmt.Println("Listening on :3000")
	app.Start(":3000")
}
