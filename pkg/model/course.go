package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID    string `json:"id" gorm:"type:varchar(36);primary_key"`
	Name  string `json:"name" gorm:"type:varchar(255);"`
	Years []int  `json:"years" gorm:"type:integer[];"`
}
