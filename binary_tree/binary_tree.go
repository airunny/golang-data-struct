package binary_tree

import "fmt"

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
func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v->", root.value)
	PreOrder(root.left)
	PreOrder(root.right)
}

// 2、中序遍历
func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Printf("%v->", root.value)
	InOrder(root.right)
}

// 3、后序遍历
func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Printf("%v->", root.value)
}

// 4、层次遍历
func LevelOrder(root *Node) {
	var tmp []*Node
	tmp = append(tmp, root)
	for len(tmp) > 0 {
		cur := tmp[0]
		tmp = tmp[1:]
		if cur == nil {
			return
		}

		fmt.Printf("%v->", cur.value)
		if cur.left != nil {
			tmp = append(tmp, cur.left)
		}

		if cur.right != nil {
			tmp = append(tmp, cur.right)
		}
	}
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
