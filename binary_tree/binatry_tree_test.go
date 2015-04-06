package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasChildrenWhitoutChildren(t *testing.T) {
	node := Node{Value: 1}
	assert.False(t, node.HasChildren())
}

func TestHasChildrenWithLeftChild(t *testing.T) {
	node := Node{
		Value: 1,
		Left:  &Node{Value: 0},
	}

	assert.True(t, node.HasChildren())
}

func TestHasChildrenWithRightChild(t *testing.T) {
	node := Node{
		Value: 1,
		Right: &Node{Value: 2},
	}

	assert.True(t, node.HasChildren())
}

func TestHasChildrenWithAllChidren(t *testing.T) {
	node := Node{
		Value: 1,
		Left:  &Node{Value: 0},
		Right: &Node{Value: 2},
	}

	assert.True(t, node.HasChildren())
}

func TestAddNodeWhenEmptyTree(t *testing.T) {
	tree := BinaryTree{}
	tree.AddNode(10)

	assert.Equal(t, 10, tree.Root.Value)
}

func TestAddNodeWhenLessThanRoot(t *testing.T) {
	tree := BinaryTree{&Node{Value: 10}}
	tree.AddNode(9)

	assert.Equal(t, 9, tree.Root.Left.Value)
}

func TestAddNodeWhenGraterThanRoot(t *testing.T) {
	tree := BinaryTree{&Node{Value: 10}}
	tree.AddNode(12)

	assert.Equal(t, 12, tree.Root.Right.Value)
}

func TestInOrderSlice(t *testing.T) {
	tree := BinaryTree{}
	tree.AddNode(12)
	tree.AddNode(15)
	tree.AddNode(20)
	tree.AddNode(1)
	tree.AddNode(5)
	tree.AddNode(2)
	tree.AddNode(13)

	assert.Equal(t, []int{1, 2, 5, 12, 13, 15, 20}, tree.InOrderSlice())
}
