package main

import (
	"fmt"
	"net/http"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	rootPath := "root"
	filepath := rootPath + r.URL.Path
	fmt.Println(filepath)
	if !Exists(filepath) {
		fmt.Fprintln(w, "hello world")
	}
	fp, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 1024*8)
	fp.Read(buf)
	w.Write(buf)
}
func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
