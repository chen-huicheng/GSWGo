package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
