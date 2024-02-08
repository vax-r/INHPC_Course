package main

import (
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
	server := gin.Default()
	server.LoadHTMLGlob("./pkg/main/template/html/*")
	server.GET("/", test)
	server.Run(":8888")
}
