package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}

// 写一个context 带值的函数
func add(a, b int) int {
	return a + b
}

// 写一个快速排序算法
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	result := append(quickSort(left), pivot)
	return append(result, quickSort(right)...)
}
