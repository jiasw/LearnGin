package repositories

import (
	"gorm.io/gorm"
	"visiontest/models"
)

type UserInfoRepository struct {
	*BaseRepository[models.UserInfo]
}

func NewUserInfoRepository(db *gorm.DB) *UserInfoRepository {
	return &UserInfoRepository{
		BaseRepository: NewBaseRepository[models.UserInfo](db),
	}
}

// 自定义查询方法
func (r *UserInfoRepository) FindActiveUsers() ([]models.UserInfo, error) {
	var users []models.UserInfo
	err := r.db.Where("is_active = ?", true).Find(&users).Error
	return users, err
}

// 自定义业务逻辑
func (r *UserInfoRepository) DeactivateUser(userID uint) error {
	return r.UpdateFields(userID, map[string]interface{}{
		"is_active": false,
	})
}
