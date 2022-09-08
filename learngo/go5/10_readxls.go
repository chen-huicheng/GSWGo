package main

import (
	"fmt"
	"os"
)

func Read(path string) (data []byte) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	buf := make([]byte, 1024*4*8)
	n, err1 := f.Read(buf)
	if err1 != nil {
		fmt.Println(err1)
		return buf[:n]
	}

	return buf[:n]
}
func main() {

	data := Read("./test.xlsx")
	fmt.Println(data)

}
