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

func (c *Cdor) Style() *style {
	return &style{}
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
	n.option.label = label
	return n
}

func (n *node) Shape(shape string) *node {
	n.option.shape = shape
	return n
}
func (n *node) Fill(fill string) *node {
	n.option.fill = fill
	return n
}

func (n *node) Stroke(stroke string) *node {
	n.option.stroke = stroke
	return n
}

// --- connection ---

func (c *connection) Label(label string) *connection {
	c.option.label = label
	return c
}

func (c *connection) Shape(shape string) *connection {
	c.option.shape = shape
	return c
}
func (c *connection) Fill(fill string) *connection {
	c.option.fill = fill
	return c
}

func (c *connection) Stroke(stroke string) *connection {
	c.option.stroke = stroke
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

func (o *option) Style(style *style) *option {
	o.style = *style
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

func (s *style) Fill(fill string) *style {
	s.fill = fill
	return s
}

func (s *style) Stroke(stroke string) *style {
	s.stroke = stroke
	return s
}
