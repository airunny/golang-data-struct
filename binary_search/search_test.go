package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := BinarySearch(list, 0)
	fmt.Println(res)
}

func TestBinarySearchFirst(t *testing.T) {
	list := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}
	res := BinarySearchFirst(list, 3)
	fmt.Println(res)
}

func TestBinarySearchLast(t *testing.T) {
	list := []int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}
	res := BinarySearchLast(list, 3)
	fmt.Println(res)
}

func TestBinarySearchLastLessThan(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := BinarySearchLast(list, 3)
	fmt.Println(res)
}

func TestBinarySearchFirstGreaterThan(t *testing.T) {
	list := []int{1, 2, 3, 4, 7, 8, 9, 10}
	res := BinarySearchFirstGreaterThan(list, 6)
	fmt.Println(res)
}
