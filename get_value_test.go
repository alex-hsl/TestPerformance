package api

import (
	"testing"
)

// 一级指针
type Count1 struct {
	num int
}

var c1 = &Count1{100}

func Benchmark_var1(b *testing.B) {
	num := c1.num

	for i := 0; i < b.N; i++ {
		_ = num
	}
}

func Benchmark_pointer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = c1.num
	}
}

// 二级指针
type Count2 struct {
	C1 *Count1
}

var c2 = &Count2{&Count1{100}}

func Benchmark_var2(b *testing.B) {
	num := c2.C1.num

	for i := 0; i < b.N; i++ {
		_ = num
	}
}

func Benchmark_pointer2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = c2.C1.num
	}
}

type Count3 struct {
	C2 *Count2
}

// 三级指针
var c3 = &Count3{&Count2{&Count1{100}}}

func Benchmark_var3(b *testing.B) {
	num := c3.C2.C1.num

	for i := 0; i < b.N; i++ {
		_ = num
	}
}

func Benchmark_pointer3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = c3.C2.C1.num
	}
}
