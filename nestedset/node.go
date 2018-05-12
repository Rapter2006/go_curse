package nestedset

import (
	"gitlab.2gis.ru/i.yatsevich/go-course-1/product"
)

type Node struct {
	value    ValueGetter
	children []*Node
	parent   *Node
}

func (n *Node) Add(child *Node) *Node {
	child.parent = n
	n.children = append(n.children, child)
	return n
}

func (n Node) Value() ValueGetter {
	return n.value
}

func (n *Node) ParentNode() *Node {
	return n.parent
}

func (n *Node) ChildrenNodes() []*Node {
	return n.children
}

type ValueGetter interface {
	GetName() string
}

func NewNode(ng ValueGetter) *Node {
	return &Node{
		value:    ng,
		children: make([]*Node, 0),
	}
}

func (n *Node) Items() []product.Item {
	acc := make([]product.Item, 0)
	items(n, &acc)
	return acc
}

func items(tree *Node, acc *[]product.Item) {
	value := tree.value
	item, ok := value.(product.Item)
	if ok {
		*acc = append(*acc, item)
	}

	if len(tree.children) == 0 {
		return
	}

	for _, subTree := range tree.children {
		items(subTree, acc)
	}
}
