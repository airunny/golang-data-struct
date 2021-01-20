package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cases := []struct {
		Number   int
		Expected int
	}{
		{
			Number:   0,
			Expected: -1,
		},
		{
			Number:   1,
			Expected: 0,
		},
		{
			Number:   2,
			Expected: 1,
		},
		{
			Number:   3,
			Expected: 2,
		},
		{
			Number:   4,
			Expected: 3,
		},
		{
			Number:   5,
			Expected: 4,
		},
		{
			Number:   6,
			Expected: -1,
		},
	}

	for _, ins := range cases {
		assert.Equal(t, ins.Expected, BinarySearch(list, ins.Number))
	}
}

func TestBinarySearchFirst(t *testing.T) {
	list := []int{1, 1, 2, 2, 3, 3, 4, 4}
	cases := []struct {
		Number   int
		Expected int
	}{
		{
			Number:   0,
			Expected: -1,
		},
		{
			Number:   1,
			Expected: 0,
		},
		{
			Number:   2,
			Expected: 2,
		},
		{
			Number:   3,
			Expected: 4,
		},
		{
			Number:   4,
			Expected: 6,
		},
		{
			Number:   5,
			Expected: -1,
		},
	}

	for _, ins := range cases {
		assert.Equal(t, ins.Expected, BinarySearchFirst(list, ins.Number))
	}
}

func TestBinarySearchLast(t *testing.T) {
	list := []int{1, 1, 2, 2, 3, 3, 4, 4}
	cases := []struct {
		Number   int
		Expected int
	}{
		{
			Number:   0,
			Expected: -1,
		},
		{
			Number:   1,
			Expected: 1,
		},
		{
			Number:   2,
			Expected: 3,
		},
		{
			Number:   3,
			Expected: 5,
		},
		{
			Number:   4,
			Expected: 7,
		},
		{
			Number:   5,
			Expected: -1,
		},
	}

	for _, ins := range cases {
		assert.Equal(t, ins.Expected, BinarySearchLast(list, ins.Number))
	}
}

func TestBinarySearchLastLessThan(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cases := []struct {
		Number   int
		Expected int
	}{
		{
			Number:   0,
			Expected: -1,
		},
		{
			Number:   1,
			Expected: 0,
		},
		{
			Number:   2,
			Expected: 1,
		},
		{
			Number:   3,
			Expected: 2,
		},
		{
			Number:   4,
			Expected: 3,
		},
		{
			Number:   5,
			Expected: 4,
		},
		{
			Number:   6,
			Expected: 4,
		},
	}

	for _, ins := range cases {
		assert.Equal(t, ins.Expected, BinarySearchLastLessThan(list, ins.Number))
	}
}

func TestBinarySearchFirstGreaterThan(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	cases := []struct {
		Number   int
		Expected int
	}{
		{
			Number:   0,
			Expected: 0,
		},
		{
			Number:   1,
			Expected: 0,
		},
		{
			Number:   2,
			Expected: 1,
		},
		{
			Number:   3,
			Expected: 2,
		},
		{
			Number:   4,
			Expected: 3,
		},
		{
			Number:   5,
			Expected: 4,
		},
		{
			Number:   6,
			Expected: -1,
		},
	}
	for _, ins := range cases {
		assert.Equal(t, ins.Expected, BinarySearchFirstGreaterThan(list, ins.Number))
	}
}
