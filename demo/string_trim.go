package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "\nhello \t"
	res := strings.Trim(str, "")
	fmt.Printf("<%s>\n", res)
	res = strings.Trim(str, " ")
	fmt.Printf("<%s>\n", res)
	res = strings.Trim(str, " \t\r\n")
	fmt.Printf("<%s>\n", res)
	res = strings.TrimSpace(res)
	fmt.Printf("<%s>\n", res)
	/*output
	  <
	  hello   >
	  <
	  hello   >
	  <hello>
	  <hello>
	*/
}
