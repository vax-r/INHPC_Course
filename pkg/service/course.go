package service

import (
	"context"
	"main/pkg/model"

	"gorm.io/gorm"
)

func NewCourseService(db *gorm.DB) model.CourseService {
	return &CourseServiceImpl{
		db: db,
	}
}

type CourseServiceImpl struct {
	db *gorm.DB
}

func (cs *CourseServiceImpl) CreateCourse(ctx context.Context, course *model.Course) error {
	return cs.db.Create(course).Error
}

func (cs *CourseServiceImpl) GetAllCourses(ctx context.Context) (courses []*model.Course, err error) {
	if err = cs.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return
}

func (cs *CourseServiceImpl) GetCourseByName(ctx context.Context, course_name string) (*model.Course, error) {
	course := &model.Course{}
	err := cs.db.Where(&model.Course{Name: course_name}).First(course).Error
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cs *CourseServiceImpl) UpdateCourse(ctx context.Context, course *model.Course) error {
	err := cs.db.Where(&model.Course{Name: course.Name}).Updates(course).Error
	if err != nil {
		return err
	}
	return nil
}

func (cs *CourseServiceImpl) DeleteCourse(ctx context.Context, course *model.Course) error {
	if err := cs.db.Where(&model.Course{Name: course.Name}).Delete(course).Error; err != nil {
		return err
	}
	return nil
}
