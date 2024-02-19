package cdor

import (
	"context"
	"fmt"

	"oss.terrastruct.com/d2/d2exporter"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type Cdor struct {
	graph       *d2graph.Graph
	ruler       *textmeasure.Ruler
	nodes       []*Node
	connections []*Connection
}

func New() *Cdor {
	ruler, _ := textmeasure.NewRuler()
	layoutResolver := func(engine string) (d2graph.LayoutGraph, error) {
		return d2dagrelayout.DefaultLayout, nil
	}
	compileOpts := &d2lib.CompileOptions{
		LayoutResolver: layoutResolver,
		Ruler:          ruler,
	}
	_, graph, _ := d2lib.Compile(context.Background(), "", compileOpts, nil)
	return &Cdor{graph: graph, ruler: ruler}
}

func (c *Cdor) Cfg() *Cdor {
	// todo
	return c
}

func (c *Cdor) Nodes(nodes ...*Node) *Cdor {
	c.nodes = append(c.nodes, nodes...)
	return c
}

func (c *Cdor) Cons(cons ...*Connection) *Cdor {
	c.connections = append(c.connections, cons...)
	return c
}

func (c *Cdor) Gen() (svg []byte, err error) {
	var nodes []*Node
	var flatten func(node *Node)
	flatten = func(node *Node) {
		nodes = append(nodes, node)
		for _, n := range node.subNodes {
			n.id = fmt.Sprintf("%s.%s", node.id, n.id)
			flatten(n)
		}
	}
	for _, n := range c.nodes {
		flatten(n)
	}
	for _, n := range nodes {
		fmt.Println("id:", n.id)
		if err = c.gen(n.id, n.option); err != nil {
			return
		}
	}
	for _, con := range c.connections {
		id := fmt.Sprintf("%s -> %s", con.src, con.dst)
		if err = c.gen(id, con.option); err != nil {
			return
		}
	}
	c.graph.ApplyTheme(d2themescatalog.NeutralDefault.ID)
	if err = c.graph.SetDimensions(nil, c.ruler, nil); err != nil {
		return
	}
	if err = d2dagrelayout.Layout(context.Background(), c.graph, nil); err != nil {
		return
	}
	var diagram *d2target.Diagram
	diagram, err = d2exporter.Export(context.Background(), c.graph, nil)
	if err != nil {
		return
	}
	diagram.Config = &d2target.Config{}
	return d2svg.Render(diagram, &d2svg.RenderOpts{
		ThemeID: &d2themescatalog.NeutralDefault.ID,
	})
}

func (c *Cdor) gen(id string, option *Option) (err error) {
	if c.graph, _, err = d2oracle.Create(c.graph, nil, id); err != nil {
		return
	}
	return c.apply(id, option)
}

func (c *Cdor) apply(key string, o *Option) (err error) {
	// fix
	return
	if o == nil {
		return
	}
	set := func(tag, val *string) {
		c.graph, err = d2oracle.Set(c.graph, nil, key, tag, val)
	}

	switch {
	case o.Fill != "":
		fill := "fill"
		if set(&fill, &o.Fill); err != nil {
			return
		}
	case o.Stroke != "":
		stroke := "stroke"
		if set(&stroke, &o.Stroke); err != nil {
			return
		}
	case o.Shape != "":
		shape := "shape"
		if set(&shape, &o.Shape); err != nil {
			return
		}
	case o.Label != "":
		label := "label"
		if set(&label, &o.Label); err != nil {
			return
		}
	}
	return
}
