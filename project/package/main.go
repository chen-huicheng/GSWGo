package main

import (
	"github.com/chen-huicheng/GSWGo/project/package/domain"
	"github.com/chen-huicheng/GSWGo/project/package/driver"
)

func main() {
	driver.InitMysql()
	domain.Run()
}
