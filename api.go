package cdor

import (
	"context"

	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func Ctx() *ctx {
	c := &ctx{}
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

func (c *ctx) Cfg() *ctx {
	// todo
	return c
}

func (c *ctx) Nodes(nodes ...*node) *ctx {
	if c.err != nil {
		return c
	}

	c.nodes = append(c.nodes, nodes...)
	return c
}

func (c *ctx) Cons(cons ...*connection) *ctx {
	if c.err != nil {
		return c
	}

	c.connections = append(c.connections, cons...)
	return c
}

func (c *ctx) Gen() (svg []byte, err error) {
	if c.err != nil {
		return nil, c.err
	}

	c.buildGraph()
	svg = c.genSvg()
	return svg, c.err
}

func Con(src, dst string, opt ...*option) *connection {
	con := &connection{src: src, dst: dst}
	if len(opt) > 0 {
		con.option = opt[0]
	}
	return con
}

func Node(id string, opt ...*option) *node {
	node := &node{id: id}
	if len(opt) > 0 {
		node.option = opt[0]
	}
	return node
}

func (n *node) Children(children ...*node) *node {
	n.children = append(n.children, children...)
	return n
}

func (n *node) Cons(cons ...*connection) *node {
	n.connections = append(n.connections, cons...)
	return n
}

func Opt() *option {
	return &option{}
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

func Style() *style {
	return &style{}
}

func (s *style) Fill(fill string) *style {
	s.fill = fill
	return s
}

func (s *style) Stroke(stroke string) *style {
	s.stroke = stroke
	return s
}
