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
    if tree == nil {
        tree = MakeNode(num)
        return
    }
    if num < tree.Value {
        if tree.Left == nil {
            tree.Left = MakeNode(num)
        } else {
            tree.Left.Insert(num)
        }
    } else if num > tree.Value {
        if tree.Right == nil {
            tree.Right = MakeNode(num)
        } else {
            tree.Right.Insert(num)
        }
    }
 }

 func InOrder(n *Node) {
	if (n != nil) {
		InOrder(n.Left)
		fmt.Println(n.value, " ")
		InOrder(n.Right)
	}
 }
