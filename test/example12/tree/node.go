package tree

import "test/example12/tree/treeentry"

type myTreeNode struct {
	node *treeentry.Node
}

func (myNode *myTreeNode) posterOrder()  {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.posterOrder()
	right.posterOrder()
	myNode.node.Print()
}

func main()  {
	var root tree.Node
	root = tree.Node{Value: 3}
}