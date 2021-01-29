package binary_tree

import (
	"reflect"
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

// [a b d e c f g]
func TestPreOrder(t *testing.T) {
	assert.True(t, reflect.DeepEqual(PreOrder(root), []interface{}{"a", "b", "d", "e", "c", "f", "g"}))
}

// [a b d e c f g]
func TestPreOrder2(t *testing.T) {
	assert.True(t, reflect.DeepEqual(PreOrder2(root), []interface{}{"a", "b", "d", "e", "c", "f", "g"}))
}

// [d b e a f c g]
func TestInOrder(t *testing.T) {
	assert.True(t, reflect.DeepEqual(InOrder(root), []interface{}{"d", "b", "e", "a", "f", "c", "g"}))
}

// [d b e a f c g]
func TestInOrder2(t *testing.T) {
	assert.True(t, reflect.DeepEqual(InOrder2(root), []interface{}{"d", "b", "e", "a", "f", "c", "g"}))
}

// [d e b f g c a]
func TestPostOrder(t *testing.T) {
	assert.True(t, reflect.DeepEqual(PostOrder(root), []interface{}{"d", "e", "b", "f", "g", "c", "a"}))
}

// [d e b f g c a]
func TestPostOrder2(t *testing.T) {
	assert.True(t, reflect.DeepEqual(PostOrder2(root), []interface{}{"d", "e", "b", "f", "g", "c", "a"}))
}

// [a b c d e f g]
func TestLevelOrder(t *testing.T) {
	assert.True(t, reflect.DeepEqual(LevelOrder(root), []interface{}{"a", "b", "c", "d", "e", "f", "g"}))
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
