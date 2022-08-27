package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage:%s sql/gorm\n", args[0])
		return
	}
	if args[1] == "sql" {
		fmt.Println("Accessing database by sql")
		RunSql()
	} else if args[1] == "gorm" {
		fmt.Println("Accessing database by gorm")
		RunGorm()
	} else {
		fmt.Printf("Usage:%s sql/gorm\n", args[0])
	}
}
