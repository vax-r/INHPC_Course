package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string `json:"id" gorm:"type:varchar(36);primary_key"`
	Student_ID string `json:"student_id" gorm:"type:varchar(255);index"`
	Password   string `json:"password" gorm:"type:varchar(255);"`
	Priviledge uint   `json:"priviledge"`
}

type CreateUserRequest struct {
	Student_ID string `json:"student_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Priviledge uint   `json:"priviledge,default=3" binding:"required"`
}

type UpdateUserRequest struct {
	Student_ID string `json:"student_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Priviledge uint   `json:"priviledge" binding:"required"`
}

type DeleteUserRequest struct {
	Student_ID string `json:"student_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserBySID(ctx context.Context, student_id string) (*User, error)
	UpdateUser(ctx context.Context, request *UpdateUserRequest) error
	DeleteUser(ctx context.Context, request *DeleteUserRequest) error
}
