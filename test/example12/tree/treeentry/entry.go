package treeentry

import "fmt"

//定义二叉树

type Node struct {
	Value int
	Left, Right *Node
}

func (node *Node) Print()  {
	fmt.Print(node.Value, " ")
}

func (node *Node)SetValue(value int)  {
	if node == nil {
		fmt.Println("setting value to nil" + "node.Ignored")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node  {
	return &Node{Value: value}
}