package cdor

// --- Cdor ---

func Ctx() *Cdor {
	c := &Cdor{}
	c.init()
	return c
}

func (c *Cdor) Cfg() *Cdor {
	// todo
	return c
}

func (c *Cdor) Nodes(nodes ...*node) *Cdor {
	if c.err != nil {
		return c
	}

	c.nodes = append(c.nodes, nodes...)
	return c
}

func (c *Cdor) Cons(cons ...*connection) *Cdor {
	if c.err != nil {
		return c
	}

	c.connections = append(c.connections, cons...)
	return c
}

func (c *Cdor) Node(id string, opt ...*option) *node {
	node := &node{id: id}
	if len(opt) > 0 {
		node.option = opt[0]
	} else {
		node.option = c.option.Copy()
	}
	return node
}

func (c *Cdor) Con(src, dst string, opt ...*option) *connection {
	con := &connection{src: src, dst: dst}
	if len(opt) > 0 {
		con.option = opt[0]
	} else {
		con.option = c.option.Copy()
	}
	con.arrow = c.arrow.Copy()
	return con
}

func (c *Cdor) Gen() (svg []byte, err error) {
	if c.err != nil {
		return nil, c.err
	}

	c.buildGraph()
	svg = c.genSvg()
	return svg, c.err
}

func (c *Cdor) Opt() *option {
	return &option{}
}

// --- node ---

func (n *node) Children(children ...*node) *node {
	n.children = append(n.children, children...)
	return n
}

func (n *node) Cons(cons ...*connection) *node {
	n.connections = append(n.connections, cons...)
	return n
}

func (n *node) Label(label string) *node {
	n.label = label
	return n
}

func (n *node) Shape(shape string) *node {
	n.shape = shape
	return n
}
func (n *node) Fill(fill string) *node {
	n.fill = fill
	return n
}

func (n *node) Stroke(stroke string) *node {
	n.stroke = stroke
	return n
}

// --- connection ---

func (c *connection) Label(label string) *connection {
	c.label = label
	return c
}

func (c *connection) Shape(shape string) *connection {
	c.shape = shape
	return c
}
func (c *connection) Fill(fill string) *connection {
	c.fill = fill
	return c
}

func (c *connection) Stroke(stroke string) *connection {
	c.stroke = stroke
	return c
}

func (c *connection) SrcHeadShape(shape string) *connection {
	c.arrow.srcHead.shape = shape
	return c
}

func (c *connection) SrcHeadLabel(label string) *connection {
	c.srcHead.label = label
	return c
}

func (c *connection) DstHeadShape(shape string) *connection {
	c.dstHead.shape = shape
	return c
}

func (c *connection) DstHeadLabel(label string) *connection {
	c.dstHead.label = label
	return c
}

// --- option ---

func (o *option) Copy() *option {
	res := *o
	return &res
}

func (o *option) Label(label string) *option {
	o.label = label
	return o
}

func (o *option) Shape(shape string) *option {
	o.shape = shape
	return o
}

func (o *option) Fill(fill string) *option {
	o.fill = fill
	return o
}

func (o *option) Stroke(stroke string) *option {
	o.stroke = stroke
	return o
}

// --- style ---

// --- arrow ---

func (a *arrow) Copy() *arrow {
	res := *a
	return &res
}
