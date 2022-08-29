package main

import (
	"fmt"
	"os"

	db "github.com/chen-huicheng/GSWGo/learngorm/database"
)

func main() {
	fmt.Println("learning gorm!!!")
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage:%s sql/gorm\n", args[0])
		return
	}
	if args[1] == "sql" {
		fmt.Println("Accessing database by sql")
		db.RunSql()
	} else if args[1] == "gorm" {
		fmt.Println("Accessing database by gorm")
		db.RunGorm()
	} else {
		fmt.Printf("Usage:%s sql/gorm\n", args[0])
	}
}
