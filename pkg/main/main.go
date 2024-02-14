package main

import (
	"main/pkg/bootstrap"
	"main/pkg/router"
	"main/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "INHPC Courses"
	c.HTML(http.StatusOK, "index.html", data)
}
func main() {

	// Init app
	app := bootstrap.App()

	// Init service
	userService := service.NewUserService(app.Conn)
	courseService := service.NewCourseService(app.Conn)

	services := &router.Services{
		UserService:   userService,
		CourseService: courseService,
	}

	// Init router
	router.RegisterRoutes(app, services)

	app.Run()
	// server := gin.Default()
	// server.LoadHTMLGlob("./pkg/main/template/html/*")
	// server.GET("/", test)
	// server.Run(":8888")
}
