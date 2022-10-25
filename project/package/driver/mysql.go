package driver

import (
	"fmt"
	"sync"
)

func init() {
	fmt.Println("package driver")
}

var once sync.Once

func InitMysql() {
	once.Do(
		func() {
			fmt.Println("init mysql")
		},
	)
}
