package main

import (
	"log"
	"os"
)

func InitLog() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("http.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.SetFlags(log.Llongfile | log.LstdFlags)
}
