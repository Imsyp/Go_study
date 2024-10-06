package bst
 type Node struct {
	Value int
	Left *Node
	Right *Node
 }

 func MakeNode(num int) *Node {
	return &Node{Value: num}
 }

 func (tree *Node) Insert(num int) {
	if (tree == nil) {
		return MakeNode(num)
	}
	if (num < tree.Value) {
		tree.Left = tree.Left.Insert(num)
	} else if (num > tree.Value) {
		tree.Right = tree.Right.Insert(num)
	}

	return tree
 }

 func InOrder(n *Node) {
	if (n != nil) {
		InOrder(n.Left)
		fmt.Println(n.value, " ")
		InOrder(n.Right)
	}
 }
