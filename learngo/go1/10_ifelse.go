package main

import "fmt"

func main() {
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("a == 10")
	}

	switch a {
	case 10:
		fmt.Println("a == 10")
		// break//默认不加break  具有break效果  使用fallthrough 不跳出switch
	case 1, 2, 3, 4:
		fmt.Println("1234")
	}

	b := 5
	switch {
	case b < 5:
		fmt.Println("b < 5")
	case b == 5:
		fmt.Println("b == 5")
	case b > 5:
		fmt.Println("b > 5")
	}
}
