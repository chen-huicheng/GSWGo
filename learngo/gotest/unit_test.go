package gotest_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/chen-huicheng/GSWGo/learngo/gotest"
	"github.com/chen-huicheng/GSWGo/stl"
)

func TestMySort1(t *testing.T) {
	arr := make([]int, 10)
	rand.Seed(time.Now().UnixMicro() % 1000)
	for i := range arr {
		arr[i] = rand.Int() % 100
	}
	fmt.Println(arr)
	gotest.MySort1(arr)
	fmt.Println(arr)
}

func TestMySort(t *testing.T) {
	arr := make([]int, 10)
	rand.Seed(time.Now().UnixMicro() % 1000)
	for i := range arr {
		arr[i] = rand.Int() % 100
	}
	fmt.Println(arr)
	stl.MySort(arr)
	fmt.Println(arr)
}
