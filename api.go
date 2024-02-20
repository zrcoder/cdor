package cdor

import (
	"context"
	"fmt"

	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func New() *Cdor {
	c := &Cdor{}
	c.ruler, c.err = textmeasure.NewRuler()
	layoutResolver := func(engine string) (d2graph.LayoutGraph, error) {
		return d2dagrelayout.DefaultLayout, nil
	}
	compileOpts := &d2lib.CompileOptions{
		LayoutResolver: layoutResolver,
		Ruler:          c.ruler,
	}
	_, c.graph, c.err = d2lib.Compile(context.Background(), "", compileOpts, nil)
	return c
}

func (c *Cdor) Cfg() *Cdor {
	// todo
	return c
}

func (c *Cdor) Nodes(nodes ...*Node) *Cdor {
	if c.err != nil {
		return c
	}

	c.nodes = append(c.nodes, nodes...)
	return c
}

func (c *Cdor) Cons(cons ...*Connection) *Cdor {
	if c.err != nil {
		return c
	}

	c.connections = append(c.connections, cons...)
	return c
}

func (c *Cdor) Gen() (svg []byte, err error) {
	if c.err != nil {
		return nil, c.err
	}

	c.buildGraph()
	svg = c.genSvg()
	return svg, c.err
}

func C(src, dst string, opt *Option) *Connection {
	return &Connection{src: src, dst: dst, option: opt}
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

func O() *Option {
	return &Option{}
}

func (o *Option) L(label string) *Option {
	o.Label = label
	return o
}

func (o *Option) Sh(shape string) *Option {
	o.Shape = shape
	return o
}

func (o *Option) Sty(style *Style) *Option {
	o.Style = *style
	return o
}

func (o *Option) F(fill string) *Option {
	o.Fill = fill
	return o
}

func (o *Option) S(stroke string) *Option {
	o.Stroke = stroke
	return o
}

func S() *Style {
	return &Style{}
}

func (s *Style) F(fill string) *Style {
	s.Fill = fill
	return s
}

func (s *Style) S(stroke string) *Style {
	s.Stroke = stroke
	return s
}
