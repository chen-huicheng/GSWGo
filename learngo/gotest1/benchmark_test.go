package gotest1_test

import (
	"testing"

	"github.com/chen-huicheng/GSWGo/learngo/gotest1"
)

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest1.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest1.MakeSliceWithPreAlloc()
	}
}
