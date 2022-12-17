package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `json:"title" binding:"required,min=2,max=100" validate:"is-cool" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200"  gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url"  gorm:"type:varchar(256);unique"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
