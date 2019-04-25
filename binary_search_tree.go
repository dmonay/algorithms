package main

import (
	"fmt"
	"log"
)

const (
	// ErrNilTree is returned when an operation is attempted upon  a null tree
	ErrNilTree = Error("you got a null tree there buddy. slow down")
)

// Error is a way to turn an error into a const
type Error string

func (e Error) Error() string {
	return string(e)
}

//                     ---[ 5
//              ---[ 1			---[ 7
//        ---[ 0     ---[ 3            ---[ 8
// 						  ---[ 2

func bst() {
	var err error
	var value string

	leafNode1 := &Node{Data: 0, Value: "blah"}
	leafNode2 := &Node{Data: 3, Value: "bleh"}
	leafNode3 := &Node{Data: 8, Value: "woah"}
	innerNode1 := &Node{Data: 1, Value: "huh", Left: leafNode1, Right: leafNode2}
	innerNode2 := &Node{Data: 7, Value: "why", Right: leafNode3}
	rootNode := &Node{Data: 5, Value: "I am root", Left: innerNode1, Right: innerNode2}

	tree := Tree{Root: rootNode}
	tree.String()

	if value, err = tree.Find(2); err != nil {
		log.Printf("unable to find %d in tree: %+v (correct)\n", 2, err)
	}

	if err = tree.Insert(2, "another one"); err != nil {
		log.Printf("unable to insert into tree: %+v\n", err)
	}

	if value, err = tree.Find(2); err != nil {
		log.Printf("unable to find %d in tree: %+v\n", 2, err)
	}
	fmt.Printf("The value of 2 in tree is %s\n", value)

	if err = tree.Delete(1); err != nil {
		log.Printf("unable to delete 1 in tree: %+v\n", err)
	}

	if value, err = tree.Find(1); err != nil {
		log.Printf("unable to find 1 in tree: %+v (correct)\n", err)
	}

	tree.String()
}

// Node is the basic node of the BST
type Node struct {
	Data  int
	Value string
	Left  *Node
	Right *Node
}

// Tree holds the entire BST
type Tree struct {
	Root *Node
}

// Insert inserts an item into the proper place in the BST
// Basic algorithm:
// 1. Start at root node. Check if value is equal. If not:
// 2. If value is less than, compare against left child. If no left child, insert.
// 3. If value is greater than, compare against right child. If no right child, insert.
// NOTE: this does not attempt to keep the tree balanced.
func (n *Node) Insert(data int, value string) (err error) {
	if n.Data == data {
		fmt.Println("element already in tree")
		return
	}

	if data < n.Data {
		if n.Left == nil {
			newNode := &Node{Data: data, Value: value}
			n.Left = newNode
			return
		}

		return n.Left.Insert(data, value)
	}

	if data > n.Data {
		if n.Right == nil {
			newNode := &Node{Data: data, Value: value}
			n.Right = newNode
			return
		}

		return n.Right.Insert(data, value)
	}

	return
}

// Find locates an element in the tree
func (n *Node) Find(data int) (value string, haveFound bool) {
	if n == nil {
		return
	}

	if n.Data == data {
		value = n.Value
		haveFound = true
		return
	}

	if data < n.Data {
		return n.Left.Find(data)
	}

	if data > n.Data {
		return n.Right.Find(data)
	}

	return
}

// Delete removes the provided node.
// Leaf node (2 in our sample tree above) -> set parent's pointer to the node to nil
// Half lead node (3 in our sample tree) -> replace node by its child node
// Internal node (1 in our sample tree) ->
//       	* if node is right child of parent, find the MAX node in the left subtree, call it Node B
//			* replace node with Node B
//			* call Delete on Node B
//
//			* if node is left child of parent, find the MAX node in the right subtree, call it Node B
//			* replace node with Node B
//			* call Delete on Node B
//
// Assume:
// * n and parent are both the root node
// * n is not nil
// * if n is the one to be deleted, the root node has at least one child node
func (n *Node) Delete(data int, parent *Node) {
	if data < n.Data {
		n.Left.Delete(data, n)
		return
	}

	if data > n.Data {
		n.Right.Delete(data, n)
		return
	}

	// Leaf node scenario:
	if n.Left == nil && n.Right == nil {
		if parent.Left == n {
			parent.Left = nil
			return
		}

		if parent.Right == n {
			parent.Right = nil
			return
		}
	}

	// Half leaf node scenario:
	if n.Left == nil {
		n = n.Right
		return
	}

	if n.Right == nil {
		n = n.Left
		return
	}

	// Internal node scenario

	// node is right child of parent
	if parent.Right == n {
		maxNode, parentOfMaxNode := n.Left.FindMax(n)

		n.Data = maxNode.Data
		n.Value = maxNode.Value

		maxNode.Delete(maxNode.Data, parentOfMaxNode)
	}

	// node is left child of parent
	if parent.Left == n {
		maxNode, parentOfMaxNode := n.Right.FindMax(n)

		n.Data = maxNode.Data
		n.Value = maxNode.Value

		maxNode.Delete(maxNode.Data, parentOfMaxNode)
	}
}

// FindMax returns the node with the max key in the subtree of the node n
// Note: need to return parent b/c the calling function will need to delete
// the max node
func (n *Node) FindMax(ogParent *Node) (maxNode, parentOfMaxNode *Node) {
	maxNode = n
	parentOfMaxNode = ogParent

	if n.Left != nil {
		leftMax, leftMaxParent := n.Left.FindMax(n)
		if maxNode.Data < leftMax.Data {
			maxNode = leftMax
			parentOfMaxNode = leftMaxParent
		}
	}

	if n.Right != nil {
		rightMax, rightMaxParent := n.Right.FindMax(n)
		if maxNode.Data < rightMax.Data {
			maxNode = rightMax
			parentOfMaxNode = rightMaxParent
		}
	}

	return
}

// Insert element into tree
func (t *Tree) Insert(data int, value string) (err error) {
	if t.Root == nil {
		err = ErrNilTree
		return
	}

	return t.Root.Insert(data, value)
}

// Find returns an element from the tree
func (t *Tree) Find(data int) (value string, err error) {
	var haveFound bool

	if t.Root == nil {
		err = ErrNilTree
		return
	}

	fmt.Printf("\n about to look for data\n")
	if value, haveFound = t.Root.Find(data); haveFound {
		return
	}

	fmt.Printf("%d is not in the tree\n", data)
	return

}

// Delete node from tree
func (t *Tree) Delete(data int) (err error) {
	if t.Root == nil {
		err = ErrNilTree
		return
	}

	// if the node to deleted is the root node AND that's the only node,
	// we set it to nil and return
	if t.Root.Data == data && t.Root.Left == nil && t.Root.Right == nil {
		t.Root = nil
		return
	}

	// we can now assume that
	// 1. the tree is not nil
	// 2. if the root node is to be deleted, it's not the only node
	t.Root.Delete(data, t.Root)

	return
}

func (t *Tree) String() {
	fmt.Println("------------------------------------------------")
	stringify(t.Root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Data)
		stringify(n.Right, level)
	}
}

