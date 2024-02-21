package cdor

import (
	"context"
	"fmt"
	"strings"

	"oss.terrastruct.com/d2/d2exporter"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/d2themes/d2themescatalog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func (c *Cdor) init() {
	if c.ruler, c.err = textmeasure.NewRuler(); c.err != nil {
		return
	}
	layoutResolver := func(engine string) (d2graph.LayoutGraph, error) {
		return d2dagrelayout.DefaultLayout, nil
	}
	compileOpts := &d2lib.CompileOptions{
		LayoutResolver: layoutResolver,
		Ruler:          c.ruler,
	}
	c.option = &option{}
	c.arrow = &arrow{}
	_, c.graph, c.err = d2lib.Compile(context.Background(), "", compileOpts, nil)
}

func (c *Cdor) buildGraph() {
	if c.built || c.err != nil {
		return
	}

	c.built = true

	var nodes []*node
	var flatten func(node *node)
	flatten = func(node *node) {
		nodes = append(nodes, node)
		for _, child := range node.children {
			child.id = combinID(node.id, child.id)
			flatten(child)
		}
		for _, con := range node.connections {
			con.src = combinID(node.id, con.src)
			con.dst = combinID(node.id, con.dst)
			c.connections = append(c.connections, con)
		}
	}
	for _, n := range c.nodes {
		flatten(n)
	}

	for _, n := range nodes {
		if c.gen(n.id, n.option); c.err != nil {
			return
		}
	}
	for _, con := range c.connections {
		if c.genCon(con); c.err != nil {
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

func (c *Cdor) gen(id string, option *option) (key string) {
	if c.err != nil {
		return
	}

	if c.graph, key, c.err = d2oracle.Create(c.graph, nil, id); c.err != nil {
		return
	}
	c.apply(key, option)
	return
}

func (c *Cdor) genCon(con *connection) (key string) {
	if key = c.gen(con.genKey(), con.option); c.err != nil {
		return
	}

	c.set(key, "source-arrowhead.label", con.srcOpt.label)
	c.set(key, "source-arrowhead.style.fill", con.srcOpt.fill)
	c.set(key, "source-arrowhead.style.stroke", con.srcOpt.stroke)
	c.set(key, "source-arrowhead.shape", con.srcOpt.shape)
	c.set(key, "target-arrowhead.label", con.dstOpt.label)
	c.set(key, "target-arrowhead.style.fill", con.dstOpt.fill)
	c.set(key, "target-arrowhead.style.stroke", con.dstOpt.stroke)
	c.set(key, "target-arrowhead.shape", con.dstOpt.shape)

	return
}

func (c *Cdor) apply(id string, o *option) {
	if o == nil {
		return
	}

	c.set(id, "shape", o.shape)
	c.set(id, "label", o.label)
	c.set(id, "style.fill", o.fill)
	c.set(id, "style.stroke", o.stroke)
}

func (c *Cdor) set(id, key, val string) {
	if val == "" || c.err != nil {
		return
	}
	id = combinID(id, key)
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &val)
}

func (c *Cdor) d2() (d2 string, err error) {
	c.buildGraph()
	if c.err != nil {
		return "", c.err
	}
	return d2format.Format(c.graph.AST), nil
}

func combinID(parts ...string) string {
	return strings.Join(parts, ".")
}

func (c *connection) genKey() string {
	return fmt.Sprintf("%s <-> %s", c.src, c.dst)
}
