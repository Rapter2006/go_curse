package nestedset

import (
	"gitlab.2gis.ru/i.yatsevich/go-course-1/product"
)

type Node struct {
	category product.Category
	children []*Node
	parent   *Node
}

func (n *Node) Add(child *Node) *Node {
	n.children = append(n.children)

	return n
}

func (n Node) Value() product.Category {
	return n.category
}

func (n *Node) ParentNode() *Node {
	return n.parent
}

func (n *Node) ChildrenNodes() []*Node {
	return n.children
}

func NewNode() *Node {

	return &Node{
		category: c,
		children: make([]*Node, 0),
	}
}
