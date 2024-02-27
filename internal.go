package cdor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"oss.terrastruct.com/d2/d2exporter"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func (c *Cdor) init() {
	c.globalOption = &option{}
	c.globalConOpt = &conOption{}
	c.config = &config{}
	c.direction = "down"
	c.config.DarkTheme(DarkMauve)
	_, c.graph, c.err = d2lib.Compile(context.Background(), "", nil, nil)
}

func (c *Cdor) buildGraph() {
	if c.built || c.err != nil {
		return
	}

	c.built = true

	if c.graph, c.err = d2oracle.Set(c.graph, nil, "direction", nil, &c.direction); c.err != nil {
		return
	}

	if c.isSequence {
		seq := "sequence_diagram"
		if c.graph, c.err = d2oracle.Set(c.graph, nil, "shape", nil, &seq); c.err != nil {
			return
		}
	}

	for _, n := range c.nodes {
		c.soveID(n)
	}

	for _, n := range c.nodes {
		if c.gen(n.id, n.option); c.err != nil {
			return
		}
		if c.setCode(n.id, n.codeTag, n.code); c.err != nil {
			return
		}
		for _, item := range n.sqlFields {
			if c.set(n.id, item.key, item.value); c.err != nil {
				return
			}
			if c.set(n.id, combinID(item.key, "constraint"), item.constraint); c.err != nil {
				return
			}
		}
	}
	for _, con := range c.connections {
		if c.genCon(con); c.err != nil {
			return
		}
	}
}

func (c *Cdor) soveID(node *node) {
	if node.idSolved {
		return
	}
	node.idSolved = true
	for _, child := range node.children {
		child.id = combinID(node.id, child.id)
		c.soveID(child)
	}
	for _, con := range node.connections {
		con.src = combinID(node.id, con.src)
		con.dst = combinID(node.id, con.dst)
	}
}

func (c *Cdor) genSvg() (svg []byte) {
	if c.err != nil {
		return
	}

	if c.direction == "" {
		c.direction = "down"
	}

	var ruler *textmeasure.Ruler
	if ruler, c.err = textmeasure.NewRuler(); c.err != nil {
		return
	}

	if c.err = c.graph.SetDimensions(nil, ruler, nil); c.err != nil {
		return
	}

	ctx := context.Background()
	if c.config.elkLayout {
		if c.err = d2elklayout.Layout(ctx, c.graph, nil); c.err != nil {
			return
		}
	} else {
		if c.err = d2dagrelayout.Layout(ctx, c.graph, nil); c.err != nil {
			return
		}
	}

	var diagram *d2target.Diagram
	diagram, c.err = d2exporter.Export(ctx, c.graph, nil)
	if c.err != nil {
		return
	}
	svg, c.err = d2svg.Render(diagram, &d2svg.RenderOpts{
		ThemeID:            c.cfg.ThemeID,
		DarkThemeID:        c.cfg.DarkThemeID,
		Pad:                c.cfg.Pad,
		Sketch:             c.cfg.Sketch,
		Center:             c.cfg.Center,
		ThemeOverrides:     c.cfg.ThemeOverrides,
		DarkThemeOverrides: c.cfg.DarkThemeOverrides,
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
	if key = c.gen(con.genKey(), &con.option); c.err != nil {
		return
	}

	c.set(key, "source-arrowhead.label", con.srcHead.label)
	c.set(key, "source-arrowhead.shape", con.srcHead.shape)
	c.set(key, "target-arrowhead.label", con.dstHead.label)
	c.set(key, "target-arrowhead.shape", con.dstHead.shape)
	/* FIX ME

	Error failed to set "(x <-> y)[0].target-arrowhead.style.filled" to "\"true\"": malformed style setting, expected 2 part path
	*/
	if con.srcHead.filledFlag {
		s := strconv.FormatBool(con.srcHead.filled)
		c.set(key, "source-arrowhead.style.filled", s)
	}
	if con.dstHead.filledFlag {
		s := strconv.FormatBool(con.dstHead.filled)
		c.set(key, "target-arrowhead.style.filled", s)
	}

	return
}

func (c *Cdor) apply(id string, o *option) {
	if o == nil {
		return
	}

	c.set(id, "icon", o.icon) // shoul set icon befor shape
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

func (c *Cdor) setCode(id, tag, code string) {
	if code == "" || c.err != nil {
		return
	}
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, &tag, &code)
}

func (c *Cdor) d2() string {
	c.buildGraph()
	if c.err != nil {
		return ""
	}
	return d2format.Format(c.graph.AST)
}

func (c *Cdor) getCons() []*connection {
	return c.connections
}
func (c *Cdor) getNodes() []*node {
	return c.nodes
}

func combinID(parts ...string) string {
	return strings.Join(parts, ".")
}

func (c *connection) genKey() string {
	if c.isSingle {
		return fmt.Sprintf("%s -> %s", c.src, c.dst)
	}
	return fmt.Sprintf("%s <-> %s", c.src, c.dst)
}

func (c *config) apply(cfg *config) *config {
	if cfg == nil {
		return c
	}

	c.direction = defaultStr(c.direction, cfg.direction)
	c.elkLayout = cfg.elkLayout || c.elkLayout
	c.cfg.Center = defaultBoolPoint(c.cfg.Center, cfg.cfg.Center)
	c.cfg.ThemeID = defaultInt64Point(c.cfg.ThemeID, cfg.cfg.ThemeID)
	c.cfg.DarkThemeID = defaultInt64Point(c.cfg.DarkThemeID, cfg.cfg.DarkThemeID)
	if cfg.cfg.ThemeOverrides != nil {
		c.cfg.ThemeOverrides = cfg.cfg.ThemeOverrides
	}
	if cfg.cfg.DarkThemeOverrides != nil {
		c.cfg.DarkThemeOverrides = cfg.cfg.DarkThemeOverrides
	}
	c.cfg.Pad = defaultInt64Point(c.cfg.Pad, cfg.cfg.Pad)
	c.cfg.Sketch = defaultBoolPoint(c.cfg.Sketch, cfg.cfg.Sketch)
	return c
}

func defaultStr(src, dst string) string {
	if dst == "" {
		return src
	}
	return dst
}

func defaultBoolPoint(src, dst *bool) *bool {
	if dst == nil {
		return src
	}
	return dst
}

func defaultInt64Point(src, dst *int64) *int64 {
	if dst == nil {
		return src
	}
	return dst
}
