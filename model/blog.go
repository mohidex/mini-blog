package model

import (
	"github.com/mohidex/mini-blog/database"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (blog *Blog) Save() (*Blog, error) {
	if err := database.DB.Create(&blog).Error; err != nil {
		return &Blog{}, err
	}
	return blog, nil
}

func FindBlogById(id string) (Blog, error) {
	var blog Blog
	if err := database.DB.Where("ID=?", id).First(&blog).Error; err != nil {
		return Blog{}, err
	}
	return blog, nil
}
