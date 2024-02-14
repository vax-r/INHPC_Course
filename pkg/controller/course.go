package controller

import (
	"main/pkg/bootstrap"
	"main/pkg/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	courseSvc model.CourseService
	env       *bootstrap.Env
}

func NewCourseController(svc model.CourseService, env *bootstrap.Env) *CourseController {
	return &CourseController{
		courseSvc: svc,
		env:       env,
	}
}

func string2arr(years string) []int {
	tmp := strings.Split(years, ", ")
	var result []int
	for _, ele := range tmp {
		val, _ := strconv.Atoi(ele)
		result = append(result, val)
	}
	return result
}

func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	var request model.CreateCourseRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}

	course := &model.Course{
		Name:  request.Course_name,
		Years: request.Course_years,
	}

	if err := ctrl.courseSvc.CreateCourse(c, course); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "Course Creation Success",
	})
}

func (ctrl *CourseController) GetAllCourses(c *gin.Context) {
	courses, err := ctrl.courseSvc.GetAllCourses(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Msg:  "Get All Courses Success",
		Data: courses,
	})
}

func (ctrl *CourseController) GetCourseByName(c *gin.Context) {
	if len(c.Query("Course_name")) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: "Invalid parameters",
		})
		return
	}

	course, err := ctrl.courseSvc.GetCourseByName(c, c.Query("Course_name"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg:  "Get Course Success",
		Data: course,
	})
}

func (ctrl *CourseController) UpdateCourse(c *gin.Context) {
	var request model.UpdateCourseRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}

	course := &model.Course{
		Name:  request.Course_name,
		Years: request.Course_years,
	}

	if err := ctrl.courseSvc.UpdateCourse(c, course); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "Update Course Success",
	})
}

func (ctrl *CourseController) DeleteCourse(c *gin.Context) {
	var request model.DeleteCourseRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}

	course, err := ctrl.courseSvc.GetCourseByName(c, request.Course_name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	if err = ctrl.courseSvc.DeleteCourse(c, course); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "Delete Course Success",
	})

}
