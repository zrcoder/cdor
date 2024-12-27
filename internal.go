package cdor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/sloghuman"
	"github.com/BurntSushi/toml"
	"github.com/cloudwego/gjson"
	"gopkg.in/yaml.v3"
	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2layouts/d2elklayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/d2renderers/d2fonts"
	"oss.terrastruct.com/d2/d2renderers/d2svg"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/lib/imgbundler"
	"oss.terrastruct.com/d2/lib/log"
	"oss.terrastruct.com/d2/lib/simplelog"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

func (c *Cdor) init() {
	c.config = c.Cfg()
	c.globalOption = c.Opt()
}

func (c *Cdor) cdor() *Cdor {
	return c
}

func (c *Cdor) buildGraph() {
	if c.err != nil {
		return
	}

	_, c.graph, c.err = d2lib.Compile(context.Background(), "", nil, nil)

	c.set("", "direction", c.direction)

	if c.isSequence {
		c.set("", "shape", "sequence_diagram")
	}

	c.setInt("", "grid-rows", c.globalOption.gridRows, 1)
	c.setInt("", "grid-columns", c.globalOption.gridCols, 1)
	c.setInt("", "grid-gap", c.globalOption.gridGap)
	c.setInt("", "horizontal-gap", c.globalOption.horizontalGap)
	c.setInt("", "vertical-gap", c.globalOption.verticalGap)
	c.set("", "style.fill-pattern", c.globalOption.fillPattern)

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

	ctx := log.With(context.TODO(), slog.Make(sloghuman.Sink(os.Stdout)).Leveled(slog.LevelWarn))
	d2 := c.d2()
	var ruler *textmeasure.Ruler
	if ruler, c.err = textmeasure.NewRuler(); c.err != nil {
		return
	}
	layoutResolver := func(engine string) (d2graph.LayoutGraph, error) {
		if c.config.elkLayout {
			return d2elklayout.DefaultLayout, nil
		}
		return d2dagrelayout.DefaultLayout, nil
	}
	compileOpts := &d2lib.CompileOptions{
		LayoutResolver: layoutResolver,
		Ruler:          ruler,
	}
	renderOpt := &d2svg.RenderOpts{
		ThemeID:            c.cfg.ThemeID,
		DarkThemeID:        c.cfg.DarkThemeID,
		Pad:                c.cfg.Pad,
		Sketch:             c.cfg.Sketch,
		Center:             c.cfg.Center,
		ThemeOverrides:     c.cfg.ThemeOverrides,
		DarkThemeOverrides: c.cfg.DarkThemeOverrides,
	}
	if c.config.cfg.Sketch != nil && *(c.config.cfg.Sketch) {
		renderOpt.Font = string(d2fonts.HandDrawn)
	}
	var diagram *d2target.Diagram
	if diagram, c.graph, c.err = d2lib.Compile(ctx, d2, compileOpts, renderOpt); c.err != nil {
		return
	}

	if svg, c.err = d2svg.Render(diagram, renderOpt); c.err != nil {
		return
	}
	lg := simplelog.FromLibLog(ctx)
	if svg, c.err = imgbundler.BundleLocal(ctx, lg, svg, true); c.err != nil {
		return
	}
	svg, c.err = imgbundler.BundleRemote(ctx, lg, svg, true)
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
	c.applyGrid(key, option)
	return
}

func (c *Cdor) genCon(con *connection) (key string) {
	if c.err != nil {
		return
	}
	id := con.genKey()
	if c.graph, key, c.err = d2oracle.Create(c.graph, nil, id); c.err != nil {
		return
	}
	c.apply(key, &con.option)

	// filled must be set firstly than other fields of heads
	if con.srcHead.filledFlag {
		s := strconv.FormatBool(con.srcHead.filled)
		c.set(key, "source-arrowhead.style.filled", s)
	}
	if con.dstHead.filledFlag {
		s := strconv.FormatBool(con.dstHead.filled)
		c.set(key, "target-arrowhead.style.filled", s)
	}
	c.set(key, "source-arrowhead.label", con.srcHead.label)
	c.set(key, "source-arrowhead.shape", con.srcHead.shape)
	c.set(key, "target-arrowhead.label", con.dstHead.label)
	c.set(key, "target-arrowhead.shape", con.dstHead.shape)
	return
}

func (c *Cdor) apply(id string, o *option) {
	if o == nil {
		return
	}

	c.set(id, "icon", o.icon) // shoul set icon befor shape
	c.set(id, "shape", o.shape)
	if o.blankLabel {
		c.setBlank(id, "label")
	} else {
		c.set(id, "label", o.label)
	}
	c.set(id, "tooltip", o.tooltip)
	c.set(id, "link", o.link)
	c.set(id, "near", o.position)
	c.set(id, "icon.near", o.iconPosition)
	// c.set(id, "label.near", o.labelPosition) // TODO

	c.set(id, "style.fill", o.fill)
	c.set(id, "style.stroke", o.stroke)
	c.setInt(id, "width", o.width, 1)
	c.setInt(id, "height", o.height, 1)
	c.setFloat(id, "style.opacity", o.opacity)
	c.set(id, "style.fill-pattern", o.fillPattern)
	c.setInt(id, "style.stroke-width", o.strokeWidth)
	c.setInt(id, "style.stroke-dash", o.strokeDash)
	c.setInt(id, "style.border-radius", o.borderRadius)
	c.setBool(id, "style.shadow", o.shadow)
	c.setBool(id, "style.3d", o.is3d)
	c.setBool(id, "style.multiple", o.multiple)
	c.setBool(id, "style.double-border", o.doubleBorder)
	if o.fontSize >= 8 && o.fontSize <= 100 {
		c.setInt(id, "style.font-size", o.fontSize)
	}
	c.set(id, "style.font-color", o.fontColor)
	c.set(id, "style.font", o.font)
	c.setBool(id, "style.animated", o.animated)
	c.setBool(id, "style.bold", o.bold)
	c.setBool(id, "style.italic", o.italic)
	c.setBool(id, "style.underline", o.underline)
}

func (c *Cdor) applyGrid(id string, o *option) {
	if o == nil {
		return
	}
	c.setInt(id, "grid-rows", o.gridRows, 1)
	c.setInt(id, "grid-columns", o.gridCols, 1)
	c.setInt(id, "grid-gap", o.gridGap)
	c.setInt(id, "horizontal-gap", o.horizontalGap)
	c.setInt(id, "vertical-gap", o.verticalGap)
}

func (c *Cdor) set(id, key, val string) {
	if val == "" || c.err != nil {
		return
	}
	id = combinID(id, key)
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &val)
}

func (c *Cdor) setInt(id, key string, val int, mins ...int) {
	min := 0
	if len(mins) > 0 {
		min = mins[0]
	}
	if val < min || c.err != nil {
		return
	}
	id = combinID(id, key)
	sval := strconv.Itoa(val)
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &sval)
}

func (c *Cdor) setFloat(id, key string, val float64) {
	if val < 0 || c.err != nil {
		return
	}
	id = combinID(id, key)
	sval := strconv.FormatFloat(val, 'f', 1, 64)
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &sval)
}

func (c *Cdor) setBool(id, key string, val bool) {
	if !val || c.err != nil {
		return
	}
	id = combinID(id, key)
	sval := "true"
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &sval)
}

func (c *Cdor) setBlank(id, key string) {
	blank := ""
	id = combinID(id, key)
	c.graph, c.err = d2oracle.Set(c.graph, nil, id, nil, &blank)
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
	main := d2format.Format(c.graph.AST)
	c.d2s = append(c.d2s, main)
	return strings.Join(c.d2s, "\n")
}

func (c *Cdor) json(json string) (nodes []*node, cons []*connection) {
	if !gjson.Valid(json) {
		c.err = errors.New("invalid json")
		return
	}

	type Info struct {
		id  string
		obj gjson.Result
	}

	_id := -1
	genID := func() string {
		_id++
		return strconv.Itoa(_id)
	}

	genNode := func(id string) *node {
		node := c.Node(id)
		node.shape = "sql_table"
		nodes = append(nodes, node)
		return node
	}
	genCon := func(pid, key, sid string) {
		con := c.Scon(combinID(pid, key), sid)
		cons = append(cons, con)
	}

	q := []Info{{id: genID(), obj: gjson.Parse(json)}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		switch {
		case cur.obj.IsObject():
			node := genNode(cur.id)
			for key, val := range cur.obj.Map() {
				if val.IsObject() || val.IsArray() {
					node.Field(key, " ")
					sonID := genID()
					genCon(cur.id, key, sonID)
					q = append(q, Info{id: sonID, obj: val})
					continue
				}
				node.Field(key, val.String())
			}
		case cur.obj.IsArray():
			node := genNode(cur.id)
			for i, val := range cur.obj.Array() {
				if val.IsObject() || val.IsArray() {
					sonID := genID()
					genCon(cur.id, strconv.Itoa(i), sonID)
					q = append(q, Info{id: sonID, obj: val})
					continue
				}
				node.Field(val.String(), " ")
			}
		default:
		}

	}
	return
}

func (c *Cdor) yaml2json(input string) string {
	var obj any
	yamlData := []byte(input)
	if err := yaml.Unmarshal(yamlData, &obj); err != nil {
		c.err = err
		return ""
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		c.err = err
		return ""
	} else {
		return string(data)
	}
}

func (c *Cdor) toml2json(input string) string {
	var obj any
	if _, err := toml.Decode(input, &obj); err != nil {
		c.err = err
		return ""
	}
	if data, err := json.MarshalIndent(obj, "", "  "); err != nil {
		c.err = err
		return ""
	} else {
		return string(data)
	}
}

func (c *Cdor) obj2json(obj any) string {
	data, err := json.Marshal(obj)
	if err != nil {
		c.err = err
		return ""
	}
	return string(data)
}

func combinID(a, b string) string {
	if a == "" {
		return b
	}
	return a + "." + b
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

func defaultGap(src, dst int) int {
	if dst == -1 {
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

func solveBool(b []bool) *bool {
	if len(b) > 0 {
		return &b[0]
	}
	res := true
	return &res
}
