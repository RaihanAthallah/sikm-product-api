package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"type:varchar(255);not null;"`
	Role       string `json:"role" gorm:"type:varchar(50);not null;"`
	Email      string `json:"email" gorm:"type:varchar(255);unique;not null;"`
	Password   string `json:"-" gorm:"type:varchar(255);not null;"`
}

type Product struct {
	gorm.Model `json:"-"`
	ID         uint     `json:"id" gorm:"primarykey;"`
	Name       string   `json:"name" gorm:"type:varchar(255);not null;" binding:"required"`
	Price      int      `json:"price" gorm:"type:numeric(10);not null;" binding:"required,gte=1"`
	CategoryID int      `json:"category_id" binding:"required"`
	PostedBy   uint     `json:"posted_by" gorm:"type:varchar(255);not null;" binding:"required"`
	User       User     `json:"-" gorm:"foreignKey:PostedBy"`          // product belong to user
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"` // product belong to category
}

type Category struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"type:varchar(255);not null;"`
}
