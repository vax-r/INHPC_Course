package service

import (
	"context"
	"errors"
	"main/pkg/model"

	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

func NewUserService(db *gorm.DB) model.UserService {
	return &UserServiceImpl{
		db: db,
	}
}

type UserServiceImpl struct {
	db *gorm.DB
}

func (us *UserServiceImpl) CreateUser(ctx context.Context, user *model.User) error {
	return us.db.Create(user).Error
}

func (us *UserServiceImpl) GetUserBySID(ctx context.Context, student_id string) (*model.User, error) {
	user := &model.User{}
	err := us.db.Where(&model.User{Student_ID: student_id}).First(user).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrUserNotFound
	case err != nil:
		return nil, err
	}
	return user, nil
}

func (us *UserServiceImpl) UpdateUser(ctx context.Context, request *model.UpdateUserRequest) error {
	user := &model.User{}
	err := us.db.Where(&model.User{Student_ID: request.Student_ID}).First(user).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrUserNotFound
	case err != nil:
		return err
	}

	user.Student_ID = request.Student_ID
	user.Password = request.Password
	user.Priviledge = request.Priviledge

	if err := us.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (us *UserServiceImpl) DeleteUser(ctx context.Context, request *model.DeleteUserRequest) error {
	user := &model.User{}
	if err := us.db.Where(&model.User{Student_ID: request.Student_ID}).First(user).Error; err != nil {
		return err
	}

	if err := us.db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}
