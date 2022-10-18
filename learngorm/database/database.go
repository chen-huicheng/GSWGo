package database

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func InitDB(addr, dbname, user, passwd string) (*gorm.DB, error) {
	cfg := mysql.Config{
		User:      user,
		Passwd:    passwd,
		Net:       "tcp",
		Addr:      addr,
		DBName:    dbname,
		ParseTime: true,
	}
	// Get a database handle.
	var err error
	fmt.Println("dsn", cfg.FormatDSN())
	gormdb, err = gorm.Open(gmysql.Open(cfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormdb, nil
}
func ConnectGormDB() error {
	_, err := InitDB("192.168.1.6:3306", "blog", "user", "123456")
	if err != nil {
		return fmt.Errorf("connect %v", err)
	}
	fmt.Println("connected!!")
	return nil
}
