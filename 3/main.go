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
	if nt == nil {
		return
	}

	nt.Print()
	nt.Left.Traverse()
	nt.Right.Traverse()
}
func main() {

	root := CreateNodeTree(3)
	root.Left = &NodeTree{}
	root.Right = &NodeTree{5, nil, nil}
	root.Right.Left = new(NodeTree)
	root.Left.Right = CreateNodeTree(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
}
