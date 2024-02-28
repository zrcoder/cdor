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

func (c *Cdor) Cfg() *config {
	return &config{}
}

func (c *Cdor) BaseConfig(cfg *config) *Cdor {
	if cfg == nil {
		return c
	}

	tmp := *cfg // copy cfg
	c.config = tmp.apply(c.config)
	return c
}

func (c *Cdor) ApplyConfig(cfg *config) *Cdor {
	c.config.apply(cfg)
	return c
}

func (c *Cdor) BaseOption(opt *option) *Cdor {
	if opt == nil {
		return c
	}

	c.baseOption = opt.Copy()

	for _, node := range c.nodes {
		node.option = opt.Copy().Apply(node.option)
		for _, con := range node.connections {
			con.option = *opt.Copy().Apply(opt)
		}
	}
	for _, con := range c.connections {
		con.option = *opt.Copy().Apply(opt)
	}
	return c
}

func (c *Cdor) ApplyOption(opt *option) *Cdor {
	if opt == nil {
		return c
	}

	for _, node := range c.nodes {
		node.option.Apply(opt)
		for _, con := range node.connections {
			con.option.Apply(opt)
		}
	}
	for _, con := range c.connections {
		con.option.Apply(opt)
	}
	return c
}

func (c *Cdor) BaseConOption(opt *conOption) *Cdor {
	if opt == nil {
		return c
	}

	c.baseConOption = opt

	for _, con := range c.connections {
		con.conOption = opt.Copy().Apply(con.conOption)
	}
	return c
}

func (c *Cdor) ApplyConOption(opt *conOption) *Cdor {
	if opt == nil {
		return c
	}

	for _, con := range c.connections {
		con.conOption.Apply(opt)
	}
	return c
}

// Node creat a node
func (c *Cdor) Node(id string, opt ...*option) *node {
	node := &node{id: id, option: c.baseOption.Copy()}
	if len(opt) > 0 {
		node.option = opt[0]
	}
	c.nodes = append(c.nodes, node)
	return node
}

// Con creat a connection
func (c *Cdor) Con(src, dst string, opt ...*conOption) *connection {
	con := &connection{
		src:       src,
		dst:       dst,
		conOption: c.baseConOption.Copy(),
		Cdor:      c,
	}
	if len(opt) > 0 {
		con.conOption = opt[0]
	}
	c.connections = append(c.connections, con)
	return con
}

// Scon creat a single connection
func (c *Cdor) Scon(src, dst string, opt ...*conOption) *connection {
	con := &connection{
		src:       src,
		dst:       dst,
		isSingle:  true,
		conOption: c.baseConOption.Copy(),
		Cdor:      c,
	}
	if len(opt) > 0 {
		con.conOption = opt[0]
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

func (c *Cdor) ConOpt() *conOption {
	return &conOption{}
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
	c.isSequence = false
}

func (c *Cdor) MdCode(code string, id ...string) *node {
	const mdCodeTemp = "```\n%s\n```"
	key := "md"
	if len(id) > 0 {
		key = id[0]
	}
	return c.Node(key).Code("md", fmt.Sprintf(mdCodeTemp, code))
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

func (n *node) Opt(opt *option) *node {
	n.option.Apply(opt)
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
	c.srcHead.shape = shape
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

func (c *connection) Opt(opt *conOption) *connection {
	c.conOption.Apply(opt)
	return c
}

func (c *connection) Con(dst string, opt ...*conOption) *connection {
	return c.Cdor.Con(c.dst, dst, opt...)
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

func (o *option) Base(opt *option) *option {
	if opt == nil {
		return o
	}
	if o == nil {
		return opt
	}
	return opt.Copy().Apply(o)
}

func (o *option) Apply(opt *option) *option {
	if opt == nil {
		return o
	}
	o.fill = defaultStr(o.fill, opt.fill)
	o.stroke = defaultStr(o.stroke, opt.stroke)
	o.shape = defaultStr(o.shape, opt.shape)
	o.label = defaultStr(o.label, opt.label)
	return o
}

// --- style ---

// --- arrowOption ---

func (a *conOption) Copy() *conOption {
	res := *a
	return &res
}

func (a *conOption) SrcHeadLabel(label string) *conOption {
	a.srcHead.label = label
	return a
}

func (a *conOption) SrcHeadShape(shape string) *conOption {
	a.srcHead.shape = shape
	return a
}

func (a *conOption) DstHeadLabel(label string) *conOption {
	a.dstHead.label = label
	return a
}

func (a *conOption) DstHeadShape(shape string) *conOption {
	a.dstHead.shape = shape
	return a
}

func (o *conOption) Label(label string) *conOption {
	o.label = label
	return o
}

func (o *conOption) Shape(shape string) *conOption {
	o.shape = shape
	return o
}

func (o *conOption) Fill(fill string) *conOption {
	o.fill = fill
	return o
}

func (o *conOption) Stroke(stroke string) *conOption {
	o.stroke = stroke
	return o
}

func (o *conOption) Base(opt *conOption) *conOption {
	if o == nil {
		return opt
	}
	if opt == nil {
		return o
	}

	return opt.Copy().Apply(o)
}

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

func (c *config) Sequence(b ...bool) *config {
	c.isSequence = *c.solveBool(b)
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
