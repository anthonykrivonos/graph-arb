package graph

import "fmt"

type Node struct {
	value interface{}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}
