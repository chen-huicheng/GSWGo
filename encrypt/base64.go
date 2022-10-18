package main

import (
	"encoding/base64"
	"fmt"
)

func Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Decode(encrypt string) string {
	res, _ := base64.StdEncoding.DecodeString(encrypt)
	return string(res)
}
func DeEnCode() {
	text := "test base64.StdEncoding"
	encrypt := Encode(text)
	fmt.Println(encrypt)
	fmt.Println(Decode(encrypt))

}
func main() {
	DeEnCode()
}
