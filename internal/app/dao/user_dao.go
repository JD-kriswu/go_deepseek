package dao

import (
	"github.com/kriswu/go_deepseek/internal/app/model"
	"gorm.io/gorm"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 创建用户DAO实例
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// Create 创建新用户
func (d *UserDAO) Create(user *model.User) error {
	return d.db.Create(user).Error
}

// GetByID 根据ID获取用户
func (d *UserDAO) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := d.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
func (d *UserDAO) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := d.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (d *UserDAO) Update(user *model.User) error {
	return d.db.Save(user).Error
}

// Delete 删除用户
func (d *UserDAO) Delete(id uint) error {
	return d.db.Delete(&model.User{}, id).Error
}

// List 获取用户列表
func (d *UserDAO) List(offset, limit int) ([]model.User, error) {
	var users []model.User
	err := d.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}