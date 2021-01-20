package sort

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckResult(in []int) bool {
	ret := make([]string, 0, len(in))
	for _, i := range in {
		ret = append(ret, fmt.Sprintf("%v", i))
	}
	return strings.Join(ret, ",") == "0,1,2,3,4,5,6,7,8,9"
}

func TestBubbleSort(t *testing.T) {
	list := []int{9, 3, 2, 8, 4, 0, 1, 5, 7, 6}
	BubbleSort(list)
	assert.Condition(t, func() (success bool) {
		return CheckResult(list)
	})
}

func TestInsertSort(t *testing.T) {
	list := []int{9, 3, 2, 8, 4, 0, 1, 5, 7, 6}
	InsertSort(list)
	assert.Condition(t, func() (success bool) {
		return CheckResult(list)
	})
}

func TestSelectSort(t *testing.T) {
	list := []int{9, 3, 2, 8, 4, 0, 1, 5, 7, 6}
	SelectSort(list)
	assert.Condition(t, func() (success bool) {
		return CheckResult(list)
	})
}

func TestMergeSort(t *testing.T) {
	list := []int{9, 3, 2, 8, 4, 0, 1, 5, 7, 6}
	MergeSort(list)
	assert.Condition(t, func() (success bool) {
		return CheckResult(list)
	})
}

func TestQuickSort(t *testing.T) {
	list := []int{9, 3, 2, 8, 4, 0, 1, 5, 7, 6}
	QuickSort(list)
	assert.Condition(t, func() (success bool) {
		return CheckResult(list)
	})
}

func TestFindKNumber(t *testing.T) {
	kNumberList := []int{5, 3, 1, 12, 8}
	tests := []struct {
		K        int
		Expected int
	}{
		{
			K:        0,
			Expected: -1,
		},
		{
			K:        1,
			Expected: 1,
		},
		{
			K:        2,
			Expected: 3,
		},
		{
			K:        3,
			Expected: 5,
		},
		{
			K:        4,
			Expected: 8,
		},
		{
			K:        5,
			Expected: 12,
		},
		{
			K:        6,
			Expected: -1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, FindKNumber(kNumberList, test.K))
	}
}

func TestCountingSort(t *testing.T) {
	list := []int{9, 3, 2, 8, 4, 0, 1, 5, 7, 6}
	CountingSort(list)
	assert.Condition(t, func() (success bool) {
		return CheckResult(list)
	})
}
