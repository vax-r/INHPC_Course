package model

import "gorm.io/gorm"

type CourseDetail struct {
	gorm.Model
	ID       string   `json:"id" gorm:"type:varchar(36);primary_key"`
	Name     string   `json:"name" gorm:"type:varchar(255);"`
	Year     int      `json:"years" gorm:"type:integer;"`
	Material []string `json:"material" gorm:"type:varchar(255)[]"`
}
