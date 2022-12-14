// Code generated by COMMENTS_BUILD_TOOLS 2.0.73. DO NOT EDIT.
package entity

import (
	"time"
)

func (User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}

type User struct {
	Id         int64     `gorm:"column:id;primary_key" json:"id"`
	AddTime    time.Time `gorm:"column:add_time" json:"add_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Name       string    `gorm:"column:name" json:"name"`
	Birthday   string    `gorm:"column:birthday" json:"birthday"`
}
