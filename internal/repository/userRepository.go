package repository

import (
	"gin-gonic-gorm-boilerplate/internal/model"
)

// UserRepository 는 사용자 정보에 대한 데이터베이스 작업을 관리합니다.
type UserRepository struct {
}

// NewUserRepository 는 새 UserRepository 인스턴스를 생성합니다.
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// FindAll 은 모든 사용자를 조회합니다.
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	var err error
	//result := db.RO[0].Find(&users)
	return users, err
}

// Save 는 사용자 정보를 저장합니다.
func (r *UserRepository) Save(user *model.User) error {
	var err error
	//return db.RW.Create(user).Error
	return err
}
