package graph

import "fmt"

type Node struct {
	value interface{}
}

func NewNode(value interface{}) *Node {
	n := new(Node)
	n.value = value
	return n
}

func (n *Node) String() string {
	return fmt.Sprint(n.value)
}
