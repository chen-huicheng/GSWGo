package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randnm(mini_num, max_num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()%(max_num+1-mini_num) + mini_num
}

func createNum() int {
	return randnm(1000, 9999)
}
func int2Slice(n int) (numSlice []int) {
	numSlice = []int{0, 0, 0, 0}
	if n > 9999 {
		return
	}
	for i := len(numSlice) - 1; i >= 0; i-- {
		numSlice[i] = n % 10
		n /= 10
	}
	return
}
func main() {
	randNum := createNum()
	fmt.Println(randNum)
	numSlice := int2Slice(randNum)
	fmt.Println(numSlice)
	guessNum := 0
	for {
		fmt.Printf("input num:")
		fmt.Scan(&guessNum)
		fmt.Println(guessNum)
		if guessNum == randNum {
			fmt.Println("bingo ")
			break
		}
		guessSlice := int2Slice(guessNum)
		fmt.Println(guessSlice)
		res := make([]int, 4)
		for i := 0; i < len(numSlice); i++ {
			if guessSlice[i] < numSlice[i] {
				res[i] = -1
			} else if guessSlice[i] > numSlice[i] {
				res[i] = 1
			}
		}
		fmt.Println("guessNum != randNum ", res)
	}

}
