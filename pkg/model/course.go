package model

import (
	"context"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID    string        `json:"id" gorm:"type:varchar(36);primary_key"`
	Name  string        `json:"name" gorm:"type:varchar(255);"`
	Years pq.Int64Array `json:"years" gorm:"type:integer[];"`
}

type CreateCourseRequest struct {
	Course_name  string        `json:"course_name" binding:"required"`
	Course_years pq.Int64Array `json:"course_years" binding:"required"`
}

type UpdateCourseRequest struct {
	Course_name  string        `json:"course_name" binding:"required"`
	Course_years pq.Int64Array `json:"course_years" binding:"required"`
}

type DeleteCourseRequest struct {
	Course_name string `json:"course_name" binding:"required"`
}

type CourseService interface {
	CreateCourse(ctx context.Context, course *Course) error
	GetAllCourses(ctx context.Context) ([]*Course, error)
	GetCourseByName(ctx context.Context, course_name string) (*Course, error)
	UpdateCourse(ctx context.Context, course *Course) error
	DeleteCourse(ctx context.Context, course *Course) error
}
