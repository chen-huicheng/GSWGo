package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("status:", resp.Status)
	fmt.Println("body :", resp.Body)
	buf := make([]byte, 1024*8)
	for {

		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf))
	}
}
