package controller

import (
	"errors"
	"main/pkg/bootstrap"
	"main/pkg/model"
	"main/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc model.UserService
	env     *bootstrap.Env
}

func NewUserController(svc model.UserService, env *bootstrap.Env) *UserController {
	return &UserController{
		userSvc: svc,
		env:     env,
	}
}

func (ctrl *UserController) Profile(c *gin.Context) {
	student_id, exist := c.Get("Student_ID")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{
			Msg: "Invalid student id",
		})
		return
	}

	student_id_str, _ := student_id.(string)
	profile, err := ctrl.userSvc.GetUserBySID(c, student_id_str)

	switch {
	case errors.Is(err, service.ErrUserNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, model.Response{
			Msg: "User not found",
		})
		return
	case err != nil:
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Msg:  "User found",
		Data: profile,
	})
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var request model.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}

	// TODO: add hash and salt for password

	user := &model.User{
		Student_ID: request.Student_ID,
		Password:   request.Password,
		Priviledge: request.Priviledge,
	}

	if err := ctrl.userSvc.CreateUser(c, user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "User Registration Success",
	})
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	var request model.UpdateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}

	if err := ctrl.userSvc.UpdateUser(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "User Update Success",
	})
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	var request model.DeleteUserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Msg: err.Error(),
		})
		return
	}

	if err := ctrl.userSvc.DeleteUser(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Msg: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Msg: "User Delete Success",
	})
}
