package database

import (
	"fmt"

	"github.com/chen-huicheng/GSWGo/learngorm/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

type GormDao struct {
	DB *gorm.DB
}

func NewGormDao() *GormDao {
	if gormdb == nil {
		ConnectGormDB()
	}
	return &GormDao{gormdb}
}
func ConnectGormDB() error {
	var err error
	gormdb, err = gorm.Open(mysql.Open("user:123456@tcp(192.168.1.6:3306)/blog?parseTime=true"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connect %v", err)
	}
	fmt.Println("connected!!")
	return nil
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

}
