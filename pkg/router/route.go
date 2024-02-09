package router

import (
	"main/pkg/bootstrap"
	"main/pkg/controller"
	"main/pkg/model"
)

type Services struct {
	UserService model.UserService
}

func RegisterRoutes(app *bootstrap.Application, services *Services) {

	userController := controller.NewUserController(services.UserService, app.Env)
	RegisterUserRoutes(app, userController)
}
