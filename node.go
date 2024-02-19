package cdor

type Node struct {
	subNodes    []*Node
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

func (n *Node) Sub(nodes ...*Node) *Node {
	n.subNodes = append(n.subNodes, nodes...)
	return n
}

func (n *Node) Connections(cons ...*Connection) *Node {
	n.connections = append(n.connections, cons...)
	return n
}

func (n *Node) Connect(src, dst string) *Node {
	n.connections = append(n.connections, &Connection{src: src, dst: dst})
	return n
}
