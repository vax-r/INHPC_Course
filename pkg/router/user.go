package router

import (
	"main/pkg/bootstrap"
	"main/pkg/controller"
)

func RegisterUserRoutes(app *bootstrap.Application, controller *controller.UserController) {
	r := app.Engine.Group("/users")

	r.GET("/profile", controller.Profile)
	r.POST("/register", controller.CreateUser)
	r.PATCH("/update", controller.UpdateUser)
	r.DELETE("/delete", controller.DeleteUser)
}
