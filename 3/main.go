package main

import "fmt"

//二叉数
type NodeTree struct {
	Value int
	Right *NodeTree
	Left  *NodeTree
}

func CreateNodeTree(value int) *NodeTree {
	return &NodeTree{
		Value: value,
	}
}
func (nt *NodeTree) Print() {
	fmt.Println(nt.Value)
}
func (nt *NodeTree) SetValue(value int) {
	nt.Value = value
}
func (nt *NodeTree) Traverse() {
	nt.TraverseFun(func(node *NodeTree) {
		node.Print()
	})

}

func (nt *NodeTree) TraverseFun(f func(*NodeTree)) {
	if nt == nil {
		return
	}
	nt.Left.TraverseFun(f)
	f(nt)
	nt.Right.TraverseFun(f)

}
func main() {

	root := CreateNodeTree(3)
	root.Left = &NodeTree{}
	root.Right = &NodeTree{5, nil, nil}
	root.Right.Left = new(NodeTree)
	root.Left.Right = CreateNodeTree(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	nodeCount := 0
	root.TraverseFun(func(*NodeTree) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
