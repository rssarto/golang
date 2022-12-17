package repository

import (
	"github.com/rssarto/gin-tutorial/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video) error
	Update(video entity.Video) error
	Delete(video entity.Video) error
	FindAll() []entity.Video
}

type videoRepository struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &videoRepository{
		connection: db,
	}
}

func (db *videoRepository) Save(video entity.Video) error {
	err := db.connection.Create(&video).Error
	if err != nil {
		panic("Failed to create record")
	}
	return nil
}
func (db *videoRepository) Update(video entity.Video) error {
	err := db.connection.Save(&video).Error
	if err != nil {
		panic("Failed to update record")
	}
	return nil
}
func (db *videoRepository) Delete(video entity.Video) error {
	err := db.connection.Delete(&video).Error
	if err != nil {
		panic("Failed to delete record")
	}
	return nil
}
func (db *videoRepository) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Author").Find(&videos)
	return videos
}
