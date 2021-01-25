package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 *		   a
 *       /   \
 *		b     c
 *     / \   / \
 *   d    e f   g
 *
 */
var (
	left   = NewNode("b", NewNode("d", nil, nil), NewNode("e", nil, nil))
	right  = NewNode("c", NewNode("f", nil, nil), NewNode("g", nil, nil))
	right1 = NewNode("c", NewNode("f", nil, nil), NewNode("g", &Node{value: "h"}, nil))
	root   = NewNode("a", left, right)
	root1  = NewNode("a", left, right1)
)

func TestPreOrder(t *testing.T) {
	PreOrder(root)
}

func TestInOrder(t *testing.T) {
	InOrder(root)
}

func TestPostOrder(t *testing.T) {
	PostOrder(root)
}

func TestLevelOrder(t *testing.T) {
	LevelOrder(root)
}

func TestTreeHeight(t *testing.T) {
	assert.Equal(t, 3, TreeHeight(root))
}

func TestTreeNodeCount(t *testing.T) {
	assert.Equal(t, 7, TreeNodeCount(root))
}

func TestISEqual(t *testing.T) {
	assert.True(t, ISEqual(nil, nil))
	assert.False(t, ISEqual(nil, root))
	assert.False(t, ISEqual(root, nil))
	assert.True(t, ISEqual(root, root))
	assert.False(t, ISEqual(root, root1))
}
