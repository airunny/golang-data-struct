package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	node := NewNodeWithNext(1,NewNodeWithNext(2,NewNodeWithNext(3,NewNodeWithNext(4,NewNode(5)))))
	assert.Equal(t,"1->2->3->4->5->nil", node.String())

	// 单链表反转
	newNode := Traverse(node)
	assert.Equal(t,"5->4->3->2->1->nil", newNode.String())
}

func TestTraverse2(t *testing.T) {
	node := NewNodeWithNext(1,NewNodeWithNext(2,NewNodeWithNext(3,NewNodeWithNext(4,NewNode(5)))))
	t.Log( node.String())

	// 单链表相邻节点两两反转
	newNode := Traverse2(node)
	assert.Equal(t,"2->1->4->3->5->nil",newNode.String())

	node = NewNode(1)
	assert.Equal(t, "1->nil",Traverse2(node).String())

	node = NewNodeWithNext(1,NewNode(2))
	assert.Equal(t,"2->1->nil", Traverse2(node).String())
}

func TestHasCircle(t *testing.T) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNodeWithNext(3,NewNodeWithNext(4,NewNodeWithNext(5,NewNodeWithNext(6,node2))))
	node1.Next = node2
	node2.Next = node3

	assert.Equal(t,true,HasCircle(node1))
	assert.Equal(t,false,HasCircle(NewNodeWithNext(1,NewNodeWithNext(2,NewNodeWithNext(3,NewNode(4))))))
}

func TestGetMiddleNode(t *testing.T) {
	assert.Equal(t,3,GetMiddleNode(NewNodeWithNext(1,NewNodeWithNext(2,NewNodeWithNext(3,NewNodeWithNext(4,NewNode(5)))))).Value)
	assert.Equal(t,2,GetMiddleNode(NewNodeWithNext(1,NewNodeWithNext(2,NewNodeWithNext(3,NewNode(4))))).Value)
}

func TestRDelNode(t *testing.T) {
	node := NewNodeWithNext(1,NewNodeWithNext(2,NewNodeWithNext(3,NewNodeWithNext(4,NewNode(5)))))
	t.Log(node.String())

	delNode := RDelNode(node,3)
	assert.Equal(t,3,delNode.Value)
}

func TestMerge(t *testing.T) {
	//node1 := NewNodeWithNext(1,NewNodeWithNext(3,NewNode(5)))
	//node2 := NewNodeWithNext(2,NewNode(4))
	//node := Merge(node1,node2)
	//assert.Equal(t,"1->2->3->4->5->nil",node.String())
	//
	//node1 = NewNodeWithNext(2,NewNodeWithNext(3,NewNode(4)))
	//node2 = NewNodeWithNext(1,NewNode(5))
	//node = Merge(node1,node2)
	//assert.Equal(t,"1->2->3->4->5->nil",node.String())

	node1 := NewNodeWithNext(2,NewNodeWithNext(2,NewNode(2)))
	node2 := NewNodeWithNext(2,NewNode(2))
	node := Merge(node1,node2)
	assert.Equal(t,"2->2->2->2->2->nil",node.String())
}

func BenchmarkNewLeastRecentlyUsed(b *testing.B) {
	lru := NewLeastRecentlyUsed(5)

	for i := 0;i < b.N;i ++  {
		lru.Use(1)
		lru.Head.String()

		lru.Use(2)
		lru.Head.String()

		lru.Use(3)
		lru.Head.String()

		lru.Use(4)
		lru.Head.String()

		lru.Use(5)
		lru.Head.String()

		lru.Use(6)
		lru.Head.String()

		lru.Use(6)
		lru.Head.String()

		lru.Use(5)
		lru.Head.String()
	}
}
