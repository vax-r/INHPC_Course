package main

import (
	"main/pkg/bootstrap"
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

	// Init router

	server := gin.Default()
	server.LoadHTMLGlob("./pkg/main/template/html/*")
	server.GET("/", test)
	server.Run(":8888")

	_ = app
}
