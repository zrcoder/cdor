package cdor

import (
	"fmt"
	"os"
	"strings"
)

// --- Cdor ---

func Ctx() *Cdor {
	c := &Cdor{}
	c.init()
	return c
}

// Node creats a node
func (c *Cdor) Node(id string, opt ...*option) *node {
	node := &node{
		id:     id,
		option: c.Opt(),
		Cdor:   c,
	}
	if len(opt) > 0 {
		node.option = opt[0]
	}
	c.nodes = append(c.nodes, node)
	return node
}

// Con creats a connection
func (c *Cdor) Con(src, dst string, opt ...*conOption) *connection {
	con := &connection{
		src:       src,
		dst:       dst,
		conOption: &conOption{},
		Cdor:      c,
	}
	if len(opt) > 0 {
		con.conOption = opt[0]
	}
	c.connections = append(c.connections, con)
	return con
}

// Scon creats a single connection
func (c *Cdor) Scon(src, dst string, opt ...*conOption) *connection {
	con := c.Con(src, dst, opt...)
	con.isSingle = true
	return con
}

// Gen generates svg data for the diagram
func (c *Cdor) Gen() (svg []byte, err error) {
	if c.err != nil {
		return nil, c.err
	}
	svg = c.genSvg()
	return svg, c.err
}

// Cfg creates a default config
func (c *Cdor) Cfg() *config {
	res := &config{}
	return res.DarkTheme(DarkMauve)
}

// Opt creates a default option
func (c *Cdor) Opt() *option {
	return &option{
		gridGap:       -1,
		horizontalGap: -1,
		verticalGap:   -1,
	}
}

// ApplyConfig applys the diagram's config
func (c *Cdor) ApplyConfig(cfg *config) *Cdor {
	c.config.apply(cfg)
	return c
}

// ApplyOption applys the diagram's global option
func (c *Cdor) ApplyOption(opt *option) *Cdor {
	if opt == nil {
		return c
	}
	c.globalOption.Apply(opt)
	return c
}

// ConOpt creates a default connection option
func (c *Cdor) ConOpt() *conOption {
	return &conOption{}
}

// SaveFile saves the svg data to filename
func (c *Cdor) SaveFile(filename string) error {
	data, err := c.Gen()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0600)
}

// Clear clears the diagram
func (c *Cdor) Clear() {
	c.init()
	c.nodes = nil
	c.connections = nil
	c.isSequence = false
}

// MdCode creates a node with md code style
func (c *Cdor) MdCode(code string, id ...string) *node {
	const mdCodeTemp = "```\n%s\n```"
	key := "md"
	if len(id) > 0 {
		key = id[0]
	}
	return c.Node(key).Code("md", fmt.Sprintf(mdCodeTemp, code))
}

// Json creates nodes and connections for a json
func (c *Cdor) Json(json string) *Cdor {
	c.direction = "right"
	c.elkLayout = true
	c.json(json)
	return c
}

// Yaml creates nodes and connections for a yaml
func (c *Cdor) Yaml(yaml string) *Cdor {
	json := c.yaml2json(yaml)
	if c.err != nil {
		return c
	}
	return c.Json(json)
}

// Toml creates nodes and connections for a toml
func (c *Cdor) Toml(yaml string) *Cdor {
	json := c.toml2json(yaml)
	if c.err != nil {
		return c
	}
	return c.Json(json)
}

// Obj creates nodes and connections for an object
func (c *Cdor) Obj(x any) *Cdor {
	json := c.obj2json(x)
	if c.err != nil {
		return c
	}
	return c.Json(json)
}

// GridRows sets the global grid rows
func (c *Cdor) GridRows(rows int) *Cdor {
	c.globalOption.gridRows = rows
	return c
}

// GridCols sets the global grid columns
func (c *Cdor) GridCols(cols int) *Cdor {
	c.globalOption.gridCols = cols
	return c
}

// GridGap sets the global grid gap
func (c *Cdor) GridGap(gap int) *Cdor {
	c.globalOption.gridGap = gap
	return c
}

// VertivalGap sets the global vertical gap
func (c *Cdor) VerticalGap(gap int) *Cdor {
	c.globalOption.verticalGap = gap
	return c
}

// HorizontalGap sets the global horizontal gap
func (c *Cdor) HorizontalGap(gap int) *Cdor {
	c.globalOption.horizontalGap = gap
	return c
}

// --- node ---

// Children adds children for the node
func (n *node) Children(children ...*node) *node {
	n.children = append(n.children, children...)
	return n
}

// Cons adds connections for the node
func (n *node) Cons(cons ...*connection) *node {
	n.connections = append(n.connections, cons...)
	return n
}

// Opt sets the node's option
func (n *node) Opt(opt *option) *node {
	n.option.Apply(opt)
	return n
}

// Label sets the node's label
func (n *node) Label(label string) *node {
	if label == "" {
		n.blankLabel = true
	}
	n.label = label
	return n
}

// Shape sets the node's shape
func (n *node) Shape(shape string) *node {
	n.shape = shape
	return n
}

// Fill sets the node's fill style
func (n *node) Fill(fill string) *node {
	n.fill = fill
	return n
}

// Stroke sets the node's stroke style
func (n *node) Stroke(stroke string) *node {
	n.stroke = stroke
	return n
}

// Code creates a code node
func (n *node) Code(tag, code string) *node {
	if tag == "latex" || tag == "tex" {
		code = strings.ReplaceAll(code, `\`, `\\`)
	}
	n.codeTag = tag
	n.code = code
	return n
}

// Icon creates an icon node
func (n *node) Icon(icon string) *node {
	n.icon = icon
	return n
}

// Field sets a field for the node, who's shape is sql table or uml class
func (n *node) Field(s ...string) *node {
	if len(s) < 2 {
		return n
	}
	key := s[0]
	if key[0] == '#' {
		key = `\` + key
	}
	if len(s) == 2 {
		n.sqlFields = append(n.sqlFields, sqlField{key: key, value: s[1]})
	} else if len(s) > 2 {
		n.sqlFields = append(n.sqlFields, sqlField{key: key, value: s[1], constraint: strings.Join(s[2:], " ")})
	}
	return n
}

// Json creates children nodes and connections as the json's layout in the node
func (n *node) Json(json string) *node {
	nodes, cons := n.Cdor.json(json)
	n.Children(nodes...)
	n.Cons(cons...)
	return n
}

// Yaml creates children nodes and connections as the yaml's layout in the node
func (n *node) Yaml(yaml string) *node {
	json := n.Cdor.yaml2json(yaml)
	if n.Cdor.err != nil {
		return n
	}
	return n.Json(json)
}

// Toml creates children nodes and connections as the toml's layout in the node
func (n *node) Toml(toml string) *node {
	json := n.Cdor.toml2json(toml)
	if n.Cdor.err != nil {
		return n
	}
	return n.Json(json)
}

// Obj creates children nodes and connections as the object's layout in the node
func (n *node) Obj(x any) *node {
	json := n.Cdor.obj2json(x)
	if n.Cdor.err != nil {
		return n
	}
	return n.Json(json)
}

// Width sets the node's width
func (n *node) Width(w int) *node {
	n.width = w
	return n
}

// Height sets the node's height
func (n *node) Height(h int) *node {
	n.height = h
	return n
}

// GridRows sets the node's grid rows
func (n *node) GridRows(r int) *node {
	n.option.gridRows = r
	return n
}

// GridCols sets the node's grid columns
func (n *node) GridCols(c int) *node {
	n.option.gridCols = c
	return n
}

// GridGap sets the node's grid gap
func (n *node) GridGap(g int) *node {
	n.option.gridGap = g
	return n
}

// VerticalGap sets the node's vertical gap
func (n *node) VerticalGap(gap int) *node {
	n.option.verticalGap = gap
	return n
}

// HorizontalGap sets the node's horizontal gap
func (n *node) HorizontalGap(gap int) *node {
	n.option.horizontalGap = gap
	return n
}

// Sequence sets the node's shape as sequence
func (n *node) Sequence() *node {
	n.shape = "sequence_diagram"
	return n
}

// --- connection ---

// Label sets the connection's label
func (c *connection) Label(label string) *connection {
	c.label = label
	return c
}

// Fill sets the connection's fill
func (c *connection) Fill(fill string) *connection {
	c.fill = fill
	return c
}

// Stroke sets the connection's stroke
func (c *connection) Stroke(stroke string) *connection {
	c.stroke = stroke
	return c
}

// SrcHeadShape sets the connection's source head shape
func (c *connection) SrcHeadShape(shape string) *connection {
	c.srcHead.shape = shape
	return c
}

// SrcHeadLabel sets the connection's source head label
func (c *connection) SrcHeadLabel(label string) *connection {
	c.srcHead.label = label
	return c
}

// DstHeadShape sets the connection's target head shape
func (c *connection) DstHeadShape(shape string) *connection {
	c.dstHead.shape = shape
	return c
}

// DstHeadLabel sets the connection's target head label
func (c *connection) DstHeadLabel(label string) *connection {
	c.dstHead.label = label
	return c
}

// Opt applies the connection's option
func (c *connection) Opt(opt *conOption) *connection {
	c.conOption.Apply(opt)
	return c
}

// Con create a connection from current connection's destination to dst
func (c *connection) Con(dst string, opt ...*conOption) *connection {
	return c.Cdor.Con(c.dst, dst, opt...)
}

// --- option ---

// Copy copies the option
func (o *option) Copy() *option {
	res := *o
	return &res
}

// Label sets the option's label
func (o *option) Label(label string) *option {
	if label == "" {
		o.blankLabel = true
	}
	o.label = label
	return o
}

// Shape sets the option's shape
func (o *option) Shape(shape string) *option {
	o.shape = shape
	return o
}

// Fill sets the option's fill style
func (o *option) Fill(fill string) *option {
	o.fill = fill
	return o
}

// Stroke sets the option's stroke style
func (o *option) Stroke(stroke string) *option {
	o.stroke = stroke
	return o
}

// Width sets the option's width
func (o *option) Width(w int) *option {
	o.width = w
	return o
}

// Height sets the option's height
func (o *option) Height(h int) *option {
	o.height = h
	return o
}

// Apply applies the option
func (o *option) Apply(opt *option) *option {
	if opt == nil {
		return o
	}
	o.blankLabel = opt.blankLabel
	o.fill = defaultStr(o.fill, opt.fill)
	o.stroke = defaultStr(o.stroke, opt.stroke)
	o.shape = defaultStr(o.shape, opt.shape)
	o.label = defaultStr(o.label, opt.label)
	o.gridGap = defaultGap(o.gridGap, opt.gridGap)
	o.horizontalGap = defaultGap(o.horizontalGap, opt.horizontalGap)
	o.verticalGap = defaultGap(o.verticalGap, opt.verticalGap)
	return o
}

// --- style ---

// --- arrowOption ---

// Copy copies the connection option
func (a *conOption) Copy() *conOption {
	res := *a
	return &res
}

// SrcHeadLabel sets the option's source head label
func (a *conOption) SrcHeadLabel(label string) *conOption {
	a.srcHead.label = label
	return a
}

// SrcHeadShape sets the option's source head shape
func (a *conOption) SrcHeadShape(shape string) *conOption {
	a.srcHead.shape = shape
	return a
}

// DstHeadLabel sets the option's target head label
func (a *conOption) DstHeadLabel(label string) *conOption {
	a.dstHead.label = label
	return a
}

// DstHeadShape sets the option's target head shape
func (a *conOption) DstHeadShape(shape string) *conOption {
	a.dstHead.shape = shape
	return a
}

// Label sets the option's label
func (o *conOption) Label(label string) *conOption {
	o.label = label
	return o
}

// Fill sets the option's fill style
func (o *conOption) Fill(fill string) *conOption {
	o.fill = fill
	return o
}

// Stroke sets the option's stroke style
func (o *conOption) Stroke(stroke string) *conOption {
	o.stroke = stroke
	return o
}

// Apply applies the option
func (o *conOption) Apply(opt *conOption) *conOption {
	if opt == nil {
		return o
	}
	o.arrow.srcHead.Apply(&opt.arrow.srcHead)
	o.arrow.dstHead.Apply(&opt.arrow.dstHead)
	o.option.Apply(&opt.option)
	return o
}

// --- config ---

// Sketch sets if the diagram is sketch style
func (c *config) Sketch(b ...bool) *config {
	c.cfg.Sketch = c.solveBool(b)
	return c
}

// ElkLayout sets if the diagram is elk layouted
func (c *config) ElkLayout(b ...bool) *config {
	c.elkLayout = *c.solveBool(b)
	return c
}

// Center sets if the diagram is centered
func (c *config) Center(b ...bool) *config {
	c.cfg.Center = c.solveBool(b)
	return c
}

// Pad sets the diagram's pad
func (c *config) Pad(pad int) *config {
	p := int64(pad)
	c.cfg.Pad = &p
	return c
}

// Theme sets the diagram's theme id
func (c *config) Theme(theme int) *config {
	t := int64(theme)
	c.cfg.ThemeID = &t
	return c
}

// DarkTheme sets the diagram's dark theme id
func (c *config) DarkTheme(theme int) *config {
	t := int64(theme)
	c.cfg.DarkThemeID = &t
	return c
}

// ThemeOverrides sets the diagram's theme over rides
func (c *config) ThemeOverrides(ov *ThemeOverrides) *config {
	c.cfg.ThemeOverrides = ov
	return c
}

// DarkThemeOverrides sets the diagram's dark theme over rides
func (c *config) DarkThemeOverrides(ov *ThemeOverrides) *config {
	c.cfg.DarkThemeOverrides = ov
	return c
}

// Direction sets the diagram's direction
func (c *config) Direction(dir string) *config {
	c.direction = dir
	return c
}

// Sequence sets if the diagram is a sequence
func (c *config) Sequence(b ...bool) *config {
	c.isSequence = *c.solveBool(b)
	return c
}

const (
	// Light themes:
	Neutral           = 0
	NeutralGrey       = 1
	Flagship          = 3
	Cool              = 4
	MixedBerryBlue    = 5
	GrapeSoda         = 6
	Aubergine         = 7
	ColorblindClear   = 8
	VanillaNitroCola  = 100
	OrangeCreamsicle  = 101
	ShirleyTemple     = 102
	EarthTones        = 103
	EvergladeGreen    = 104
	ButteredToast     = 105
	Terminal          = 300
	TerminalGrayscale = 301
	Origami           = 302
	// Dark themes:
	DarkMauve               = 200
	DarkFlagshipTerrastruct = 201
)
