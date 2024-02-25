package cdor

import (
	"os"
	"strings"
)

// --- Cdor ---

func Ctx() *Cdor {
	c := &Cdor{}
	c.init()
	return c
}

func (c *Cdor) Cfg() *config {
	return &c.config
}

func (c *Cdor) Nodes(nodes ...*node) *Cdor {
	return c
}

func (c *Cdor) Cons(cons ...*connection) *Cdor {
	return c
}

func (c *Cdor) Node(id string, opt ...*option) *node {
	node := &node{id: id, option: c.option.Copy()}
	if len(opt) > 0 {
		node.option = opt[0]
	}
	c.nodes = append(c.nodes, node)
	return node
}

func (c *Cdor) Con(src, dst string, opt ...*option) *connection {
	con := &connection{src: src, dst: dst, option: c.option.Copy(), arrow: c.arrow.Copy()}
	if len(opt) > 0 {
		con.option = opt[0]
	}
	c.connections = append(c.connections, con)
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

func (c *Cdor) ArrowOpt() *arrow {
	return &arrow{}
}

func (c *Cdor) SaveFile(filename string) error {
	data, err := c.Gen()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0600)
}

func (c *Cdor) Clear() {
	c.init()
	c.nodes = nil
	c.connections = nil
	c.built = false
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

func (n *node) Code(tag, code string) *node {
	if tag == "latex" || tag == "tex" {
		code = strings.ReplaceAll(code, `\`, `\\`)
	}
	n.codeTag = tag
	n.code = code
	return n
}

func (n *node) Icon(icon string) *node {
	n.icon = icon
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

func (c *connection) ArrowOpt(opt *arrow) *connection {
	c.arrow = opt
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

func (a *arrow) SrcHeadLabel(label string) *arrow {
	a.srcHead.label = label
	return a
}

func (a *arrow) SrcHeadShape(shape string) *arrow {
	a.srcHead.shape = shape
	return a
}

func (a *arrow) DstHeadLabel(label string) *arrow {
	a.dstHead.label = label
	return a
}

func (a *arrow) DstHeadShape(shape string) *arrow {
	a.dstHead.shape = shape
	return a
}

// --- config ---

func (c *config) Sketch(b ...bool) *config {
	c.cfg.Sketch = c.solveBool(b)
	return c
}

func (c *config) ElkLayout(b ...bool) *config {
	c.elkLayout = *c.solveBool(b)
	return c
}

func (c *config) Center(b ...bool) *config {
	c.cfg.Center = c.solveBool(b)
	return c
}

func (c *config) Pad(pad int) *config {
	p := int64(pad)
	c.cfg.Pad = &p
	return c
}

func (c *config) Theme(theme int) *config {
	t := int64(theme)
	c.cfg.ThemeID = &t
	return c
}

func (c *config) DarkTheme(theme int) *config {
	t := int64(theme)
	c.cfg.DarkThemeID = &t
	return c
}

func (c *config) ThemeOverrides(ov *ThemeOverrides) *config {
	c.cfg.ThemeOverrides = ov
	return c
}

func (c *config) DarkThemeOverrides(ov *ThemeOverrides) *config {
	c.cfg.DarkThemeOverrides = ov
	return c
}

func (c *config) Direction(dir string) *config {
	c.direction = dir
	return c
}

func (c *config) solveBool(b []bool) *bool {
	if len(b) > 0 {
		return &b[0]
	}
	res := true
	return &res
}

const (
	// Light:
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
	// Dark:
	DarkMauve               = 200
	DarkFlagshipTerrastruct = 201
)
