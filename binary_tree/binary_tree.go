package binary_tree

import (
	"github.com/liyanbing/golang-data-structures/stack"
)

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func NewNode(value interface{}, left, right *Node) *Node {
	return &Node{
		value: value,
		left:  left,
		right: right,
	}
}

// 1、前序遍历
func PreOrder(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var values []interface{}
	values = append(values, root.value)
	values = append(values, PreOrder(root.left)...)
	values = append(values, PreOrder(root.right)...)
	return values
}

// 1、前序遍历（非递归）
func PreOrder2(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var values []interface{}
	s := stack.NewArrayStack()
	cur := root
	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			values = append(values, cur.value)
			s.Push(cur)
			cur = cur.left
		} else {
			cur = s.Pop().(*Node)
			cur = cur.right
		}
	}
	return values
}

// 2、中序遍历
func InOrder(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var values []interface{}
	values = append(values, InOrder(root.left)...)
	values = append(values, root.value)
	values = append(values, InOrder(root.right)...)
	return values
}

// 2、中序遍历（非递归）
func InOrder2(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var values []interface{}
	s := stack.NewArrayStack()
	cur := root
	for !s.IsEmpty() || cur != nil {
		if cur != nil {
			s.Push(cur)
			cur = cur.left
		} else {
			cur = s.Pop().(*Node)
			values = append(values, cur.value)
			cur = cur.right
		}
	}
	return values
}

// 3、后序遍历
func PostOrder(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	var values []interface{}
	values = append(values, PostOrder(root.left)...)
	values = append(values, PostOrder(root.right)...)
	values = append(values, root.value)
	return values
}

// 3、后序遍历（非递归）
func PostOrder2(root *Node) []interface{} {
	if root == nil {
		return nil
	}

	s1 := stack.NewArrayStack()
	s2 := stack.NewArrayStack()
	s1.Push(root)

	for !s1.IsEmpty() {
		cur := s1.Pop()
		s2.Push(cur)

		node := cur.(*Node)
		if node.left != nil {
			s1.Push(node.left)
		}

		if node.right != nil {
			s1.Push(node.right)
		}
	}

	var values []interface{}
	for !s2.IsEmpty() {
		cur := s2.Pop().(*Node)
		values = append(values, cur.value)
	}
	return values
}

// 4、层次遍历
func LevelOrder(root *Node) []interface{} {
	var tmp []*Node
	tmp = append(tmp, root)
	var values []interface{}
	for len(tmp) > 0 {
		cur := tmp[0]
		tmp = tmp[1:]
		if cur == nil {
			return nil
		}

		values = append(values, cur.value)
		if cur.left != nil {
			tmp = append(tmp, cur.left)
		}

		if cur.right != nil {
			tmp = append(tmp, cur.right)
		}
	}
	return values
}

// 5、树的高度
func TreeHeight(root *Node) int {
	if root == nil {
		return 0
	}

	lHeight := TreeHeight(root.left)
	rHeight := TreeHeight(root.right)
	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}

// 6、树节点个数
func TreeNodeCount(root *Node) int {
	if root == nil {
		return 0
	}
	return TreeNodeCount(root.left) + TreeNodeCount(root.right) + 1
}

// 7、比较两棵树是否相同
func ISEqual(t1, t2 *Node) bool {
	if t1 != nil && t2 != nil {
		if t1.value == t2.value {
			if ISEqual(t1.left, t2.left) && ISEqual(t1.right, t2.right) {
				return true
			}
		}
	} else {
		if t1 == nil && t2 == nil {
			return true
		}
	}
	return false
}
