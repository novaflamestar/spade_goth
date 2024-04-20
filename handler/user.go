package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/novaflamestar/spade_goth/service"
	"github.com/novaflamestar/spade_goth/view"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) View(c echo.Context, component templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func InitHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (handler *UserHandler) BasePage(c echo.Context) error {
	userTable := view.UserPage()
	index := view.Index(userTable)
	return handler.View(c, index)
}

func (handler *UserHandler) ShowUsers(c echo.Context) error {
	users := handler.UserService.UserStore.GetUsers()
	table := view.UserList(users)
	return handler.View(c, table)
}

func (handler *UserHandler) AddUserForm(c echo.Context) error {
	userForm := view.AddUserForm()
	return handler.View(c, userForm)
}

func (handler *UserHandler) SubmitUserForm(c echo.Context) error {
	name := c.FormValue("name")
	handler.UserService.UserStore.CreateUser(name)
	table := view.Table()
	return handler.View(c, table)
}

func (handler *UserHandler) GetUserTable(c echo.Context) error {
	table := view.Table()
	return handler.View(c, table)
}

func (handler *UserHandler) AddUser(c echo.Context) error    { return nil }
func (handler *UserHandler) EditUser(c echo.Context) error   { return nil }
func (handler *UserHandler) DeleteUser(c echo.Context) error { return nil }
