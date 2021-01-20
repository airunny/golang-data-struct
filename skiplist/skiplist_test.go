package skiplist

import (
	"fmt"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	skipL := NewSkipList()
	skipL.Insert(1, 1)
	skipL.Insert(2, 2)
	skipL.Insert(4, 4)
	skipL.Insert(3, 3)
	skipL.Insert(5, 5)
	skipL.Insert(6, 6)
	skipL.Insert(7, 7)
	skipL.Insert(8, 8)
	skipL.Insert(9, 9)
	skipL.Insert(10, 10)
	skipL.Print()
}

func TestSkipList_Find(t *testing.T) {
	skipL := NewSkipList()
	skipL.Insert(1, 1)
	skipL.Insert(2, 2)
	skipL.Insert(4, 4)
	skipL.Insert(3, 3)
	skipL.Insert(5, 5)
	skipL.Insert(6, 6)
	skipL.Insert(7, 7)
	skipL.Insert(8, 8)
	skipL.Insert(9, 9)
	skipL.Insert(10, 10)
	skipL.Print()
	node := skipL.Find(1, 1)
	fmt.Println(node.Value, node.Score)
}

func TestSkipList_Delete(t *testing.T) {
	skipL := NewSkipList()
	skipL.Insert(1, 1)
	skipL.Insert(2, 2)
	skipL.Insert(4, 4)
	skipL.Insert(3, 3)
	skipL.Insert(5, 5)
	skipL.Insert(6, 6)
	skipL.Insert(7, 7)
	skipL.Insert(8, 8)
	skipL.Insert(9, 9)
	skipL.Insert(10, 10)
	skipL.Print()
	fmt.Println(skipL.Delete(1, 1))
	skipL.Print()
	fmt.Println(skipL.Delete(6, 6))
	skipL.Print()
}
