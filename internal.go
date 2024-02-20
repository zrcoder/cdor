package cdor

import (
	"context"
	"fmt"

	"oss.terrastruct.com/d2/d2exporter"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
)

func (c *Cdor) add(id string) {
	if c.err != nil {
		return
	}
	c.graph, _, c.err = d2oracle.Create(c.graph, nil, id)
}

func (c *Cdor) set(id, key, val string) {
	if val == "" || c.err != nil {
		return
	}
	id = fmt.Sprintf("%s.%s", id, key)
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &val)
}

func (c *Cdor) con(src, dst string) (id string) {
	if c.err != nil {
		return
	}
	i := fmt.Sprintf("%s -> %s", src, dst)
	c.graph, id, c.err = d2oracle.Create(c.graph, nil, i)
	return
}

func (c *Cdor) buildGraph() {
	if c.built || c.err != nil {
		return
	}

	c.built = true

	var nodes []*Node
	var flatten func(node *Node)
	flatten = func(node *Node) {
		nodes = append(nodes, node)
		c.connections = append(c.connections, node.connections...)
		for _, child := range node.children {
			flatten(child)
		}
	}
	for _, n := range c.nodes {
		flatten(n)
	}

	for _, n := range nodes {
		if c.err = c.gen(n.id, n.option); c.err != nil {
			return
		}
	}
	for _, con := range c.connections {
		id := fmt.Sprintf("%s -> %s", con.src, con.dst)
		if c.err = c.gen(id, con.option); c.err != nil {
			return
		}
	}
}

func (c *Cdor) genSvg() (svg []byte) {
	if c.err != nil {
		return
	}

	c.graph.ApplyTheme(d2themescatalog.NeutralDefault.ID)
	if c.err = c.graph.SetDimensions(nil, c.ruler, nil); c.err != nil {
		return
	}
	if c.err = d2dagrelayout.Layout(context.Background(), c.graph, nil); c.err != nil {
		return
	}
	var diagram *d2target.Diagram
	diagram, c.err = d2exporter.Export(context.Background(), c.graph, nil)
	if c.err != nil {
		return
	}
	diagram.Config = &d2target.Config{}
	svg, c.err = d2svg.Render(diagram, &d2svg.RenderOpts{
		ThemeID: &d2themescatalog.NeutralDefault.ID,
	})
	return
}

func (c *Cdor) gen(id string, option *Option) (err error) {
	if c.graph, id, err = d2oracle.Create(c.graph, nil, id); err != nil {
		return
	}
	c.apply(id, option)
	return
}

func (c *Cdor) apply(id string, o *Option) {
	if o == nil {
		return
	}

	c.set(id, "shape", o.Shape)
	c.set(id, "label", o.Label)
	c.set(id, "style.fill", o.Fill)
	c.set(id, "style.stroke", o.Stroke)
}

func (c *Cdor) D2() (d2 string, err error) {
	c.buildGraph()
	if c.err != nil {
		return "", c.err
	}
	return d2format.Format(c.graph.AST), nil
}
