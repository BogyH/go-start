package main

import (
	"fmt"
)

type path []byte

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
	newSlice := make([]int, len(slice), 2*cap(slice))
	copy(newSlice, slice)
	slice = newSlice
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}
