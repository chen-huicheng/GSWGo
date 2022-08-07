package main

import "fmt"

type People struct {
	name string
}

func print(p *People) {
	fmt.Println(p)
}

func main() {
	p := &People{"kitten"}
	print(p)
}
