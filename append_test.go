package TestPerformance

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_append(t *testing.T) {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("2th of arr1 addr: %x\n", unsafe.Pointer(&arr1[1]))
	slice1 := arr1[1:2]
	fmt.Printf("slice1: len is %d, cap is %d\n", len(slice1), cap(slice1))
	fmt.Printf("slice1 data addr before append: %x\n", ((*reflect.SliceHeader)(unsafe.Pointer(&slice1))).Data)
	slice1 = append(slice1, 6, 7, 8)
	fmt.Printf("slice1 data addr after append: %x\n", ((*reflect.SliceHeader)(unsafe.Pointer(&slice1))).Data)
	fmt.Println("slice1:", slice1)
	fmt.Println("arr1:", arr1)

	println("----------------------------------------")

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("2th of arr2 addr: %x\n", unsafe.Pointer(&arr2[1]))
	slice2 := arr2[1:3]
	fmt.Printf("slice2 data addr before append: %x\n", ((*reflect.SliceHeader)(unsafe.Pointer(&slice2))).Data)
	slice2 = append(slice2, 6, 7, 8)
	fmt.Printf("slice2 data addr after append: %x\n", ((*reflect.SliceHeader)(unsafe.Pointer(&slice2))).Data)
	fmt.Println("slice2:", slice2)
	fmt.Println("arr2:", arr2)
}
