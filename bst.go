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

func (tree *Node) Insert(num int) {
	if tree == nil {
        return // 노드가 비어 있으면 반환
    }

    if num < tree.Value {
        if tree.Left == nil {
            tree.Left = MakeNode(num) // 왼쪽 자식이 비어 있으면 새 노드 생성
        } else {
            tree.Left.Insert(num) // 왼쪽 자식에 삽입
        }
    } else if num > tree.Value {
        if tree.Right == nil {
            tree.Right = MakeNode(num) // 오른쪽 자식이 비어 있으면 새 노드 생성
        } else {
            tree.Right.Insert(num) // 오른쪽 자식에 삽입
        }
    }
}

func InOrder(n *Node) {
    if n != nil {
        InOrder(n.Left)
        fmt.Print(n.Value, " ") 
        InOrder(n.Right)
    }
}