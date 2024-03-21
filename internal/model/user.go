package model

import "gorm.io/gorm"

// User 모델 정의
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
}
