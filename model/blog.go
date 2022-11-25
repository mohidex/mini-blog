package model

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}
