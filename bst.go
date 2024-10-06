package bst

import "fmt"

type Node struct {
    Value int
    Left  *Node
    Right *Node
}

func MakeNode(num int) *Node {
    return &Node{Value: num}
}

func Insert(tree *Node, num int) *Node {
    if tree == nil {
        return MakeNode(num) 
    }

    if num < tree.Value {
        tree.Left = Insert(tree.Left, num) 
    } else if num > tree.Value {
        tree.Right = Insert(tree.Right, num) 
    }

    return tree 
}

func InOrder(n *Node) {
    if n != nil {
        InOrder(n.Left)
        fmt.Print(n.Value, " ") 
        InOrder(n.Right)
    }
}