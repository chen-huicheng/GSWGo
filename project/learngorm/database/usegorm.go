package database

import (
	"fmt"

	"github.com/chen-huicheng/GSWGo/project/learngorm/entity"
	"gorm.io/gorm"
)

type GormDao struct {
	DB *gorm.DB
}

func NewGormDao() *GormDao {
	if gormdb == nil {
		ConnectGormDB()
	}
	return &GormDao{gormdb}
}

// select query all User
func (dao *GormDao) GormSelect() ([]entity.User, error) {
	res := make([]entity.User, 0)
	dao.DB.Find(&res)
	return res, nil
}

func (dao *GormDao) GormSave(users []entity.User) error {

	if err := dao.DB.Save(users).Error; err != nil {
		return fmt.Errorf("gormSave %v", err)
	}
	return nil
}
func RunGorm() {
	NewGormDao()
}
