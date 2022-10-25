package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"hello", "world", "health", "mid", "happy"}
	sort.Strings(strs)
	sort.Reverse()
	fmt.Println(strs)
	sort.Stable()
}
