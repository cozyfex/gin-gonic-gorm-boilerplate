package repository

import (
	"gin-gonic-gorm-boilerplate/internal/db"
	"gin-gonic-gorm-boilerplate/internal/model"
)

// UserRepository 는 사용자 정보에 대한 데이터베이스 작업을 관리합니다.
type UserRepository struct {
	dbManager *db.Manager
}

// NewUserRepository 는 새 UserRepository 인스턴스를 생성합니다.
func NewUserRepository(dbManager *db.Manager) *UserRepository {
	return &UserRepository{dbManager: dbManager}
}

// FindAll 은 모든 사용자를 조회합니다.
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.dbManager.Reader().Find(&users).Error
	return users, err
}

// Save 는 사용자 정보를 저장합니다.
func (r *UserRepository) Save(user *model.User) error {
	err := r.dbManager.Writer().Create(user).Error
	return err
}
