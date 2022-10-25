package domain

import (
	"fmt"

	"github.com/chen-huicheng/GSWGo/project/package/driver"
)

func init() {
	fmt.Println("package domain ")
	driver.InitMysql()
}

func Run() {
	fmt.Println("run")
}
