package router

import (
	"main/pkg/bootstrap"
	"main/pkg/controller"
)

func RegisterCourseRoutes(app *bootstrap.Application, controller *controller.CourseController) {
	r := app.Engine.Group("/course")

	r.POST("/create", controller.CreateCourse)
	r.GET("/read", controller.GetCourseByName)
	r.GET("/all", controller.GetAllCourses)
	r.PATCH("/update", controller.UpdateCourse)
	r.DELETE("/delete", controller.DeleteCourse)
}
