package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Printf("useage:%s filename\n", list[0])
		return
	}
	fileName := list[1]

	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("name\n", info.Name())
	fmt.Println("size\n", info.Size())

}
