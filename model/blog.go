package model

import (
	"fmt"

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

func (blog *Blog) Delete() error {
	if result := database.DB.Delete(&blog); result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", blog.ID)
	}
	return nil
}
