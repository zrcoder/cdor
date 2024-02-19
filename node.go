package cdor

import "fmt"

type Node struct {
	children    []*Node
	connections []*Connection
	id          string
	option      *Option
}

func N(id string, opt ...*Option) *Node {
	node := &Node{id: id}
	if len(opt) > 0 {
		node.option = opt[0]
	}
	return node
}

func (n *Node) Children(children ...*Node) *Node {
	for _, c := range children {
		c.id = fmt.Sprintf("%s.%s", n.id, c.id)
		n.children = append(n.children, c)
	}
	return n
}

func (n *Node) Connections(cons ...*Connection) *Node {
	for _, c := range cons {
		c.src = fmt.Sprintf("%s.%s", n.id, c.src)
		c.dst = fmt.Sprintf("%s.%s", n.id, c.dst)
		n.connections = append(n.connections, c)
	}
	return n
}
