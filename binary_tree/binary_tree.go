package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (self *Node) HasChildren() bool {
	return self.CountChildren() > 0
}

func (self *Node) CountChildren() int {
	count := 0

	if self.Left != nil {
		count++
	}

	if self.Right != nil {
		count++
	}

	return count
}

type BinaryTree struct {
	Root *Node
}

func (self *BinaryTree) AddNode(Value int) {
	newNode := &Node{Value: Value}
	if self.Root == nil {
		self.Root = newNode
	} else {
		current := self.Root

		for {
			if current.Value > newNode.Value {
				if current.Left == nil {
					current.Left = newNode
					return
				}

				current = current.Left
			} else {
				if current.Right == nil {
					current.Right = newNode
					return
				}

				current = current.Right
			}
		}
	}
}

func (self *BinaryTree) doInOrderTraversal(current *Node, arr []int) []int {
	otherArr := make([]int, len(arr))
	copy(otherArr, arr)
	if current != nil {
		otherArr = self.doInOrderTraversal(current.Left, otherArr)

		otherArr = append(otherArr, current.Value)

		otherArr = self.doInOrderTraversal(current.Right, otherArr)
	}

	return otherArr
}

func (self *BinaryTree) InOrderSlice() []int {
	return self.doInOrderTraversal(self.Root, []int{})
}

func (self *BinaryTree) InOrderTraversal() {
	inOrder := self.InOrderSlice()
	for i := range inOrder {
		fmt.Println(i)
	}
}

func (self *BinaryTree) doPreOrderTraversal(current *Node) {
	if current != nil {
		fmt.Println(current.Value)

		self.doPreOrderTraversal(current.Left)
		self.doPreOrderTraversal(current.Right)
	}
}

func (self *BinaryTree) PreOrderTraversal() {
	self.doPreOrderTraversal(self.Root)
}

func (self *BinaryTree) doPostOrderTraversal(current *Node) {
	if current != nil {
		self.doPostOrderTraversal(current.Left)
		self.doPostOrderTraversal(current.Right)

		fmt.Println(current.Value)
	}
}

func (self *BinaryTree) PostOrderTraversal() {
	self.doPostOrderTraversal(self.Root)
}

func (self *BinaryTree) Find(value int) *Node {
	current := self.Root

	for current.Value != value {
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}

		if current == nil {
			return nil
		}
	}

	return current
}

func (self *BinaryTree) Min() int {
	current := self.Root
	for current.Left != nil {
		current = current.Left
	}

	return current.Value
}

func (self *BinaryTree) Delete(value int) bool {
	var parent *Node
	current := self.Root

	if current.Value == value {
		dummyNode := Node{Value: 10000000000, Left: current}
		dummyTree := BinaryTree{Root: &dummyNode}
		dummyTree.Delete(value)

		return true
	}

	for current.Value != value {
		parent = current
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}

		if current == nil {
			return false
		}
	}

	if !current.HasChildren() {
		if parent.Left.Value == value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	} else {
		if current.CountChildren() == 1 {
			if parent.Left.Value == value {
				if current.Left != nil {
					parent.Left = current.Left
				} else {
					parent.Left = current.Right
				}
			} else {
				if current.Left != nil {
					parent.Right = current.Left
				} else {
					parent.Right = current.Right
				}
			}
		} else {
			rightSubtree := BinaryTree{Root: current.Right}
			min := rightSubtree.Min()
			rightSubtree.Delete(min)

			current.Value = min
		}
	}

	return true
}

func main() {
	tree := BinaryTree{}

	tree.AddNode(10)
	tree.AddNode(1)
	tree.AddNode(2)
	tree.AddNode(15)
	tree.AddNode(20)
	tree.AddNode(18)
	tree.AddNode(14)
	tree.AddNode(17)

	fmt.Println("InOrder")
	tree.InOrderTraversal()

	fmt.Println("PreOrder")
	tree.PreOrderTraversal()

	fmt.Println("PostOrder")
	tree.PostOrderTraversal()

	fmt.Println("Find")
	node := tree.Find(15)
	fmt.Println(node.Value)
	fmt.Println(tree.Find(30))

	fmt.Println(tree.Delete(18))
	tree.PreOrderTraversal()

	fmt.Println(tree.Delete(1))
	tree.PreOrderTraversal()

	fmt.Println(tree.Delete(15))
	tree.PreOrderTraversal()

	fmt.Println(tree.Delete(10))
	tree.PreOrderTraversal()

	fmt.Println("Min")
	fmt.Println(tree.Min())
}
