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
	con.conOption.option = *c.Opt()
	con.conOption.arrow.srcHead = *c.Opt()
	con.conOption.arrow.dstHead = *c.Opt()
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
	return res.DarkTheme(c.DarkMauve())
}

// Opt creates a default option
func (c *Cdor) Opt() *option {
	res := &option{
		gridGap:       -1,
		horizontalGap: -1,
		verticalGap:   -1,
	}
	res.opacity = -1
	res.strokeWidth = -1
	res.strokeDash = -1
	res.borderRadius = -1
	return res
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

func (c *Cdor) FillPattern(pattern string) *Cdor {
	c.globalOption.fillPattern = pattern
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

// Code creates a code node
func (n *node) Code(tag, code string) *node {
	if tag == "latex" || tag == "tex" {
		code = strings.ReplaceAll(code, `\`, `\\`)
	}
	n.codeTag = tag
	n.code = code
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

// Sequence sets the node's shape as sequence
func (n *node) Sequence() *node {
	n.shape = "sequence_diagram"
	return n
}

func (n *node) Label(label string) *node {
	if label == "" {
		n.blankLabel = true
	}
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

func (n *node) Width(w int) *node {
	n.width = w
	return n
}

func (n *node) Height(h int) *node {
	n.height = h
	return n
}

func (n *node) GridRows(r int) *node {
	n.gridRows = r
	return n
}

func (n *node) GridCols(c int) *node {
	n.gridCols = c
	return n
}

func (n *node) GridGap(g int) *node {
	n.gridGap = g
	return n
}

func (n *node) VerticalGap(gap int) *node {
	n.verticalGap = gap
	return n
}

func (n *node) HorizontalGap(gap int) *node {
	n.horizontalGap = gap
	return n
}

func (n *node) Icon(i string) *node {
	n.icon = i
	return n
}

func (n *node) Opacity(op float64) *node {
	n.opacity = op
	return n
}

func (n *node) FillPattern(p string) *node {
	n.fillPattern = p
	return n
}
func (n *node) StrokeWidth(w int) *node {
	n.strokeWidth = w
	return n
}
func (n *node) StrokeDash(d int) *node {
	n.strokeDash = d
	return n
}
func (n *node) BorderRadius(r int) *node {
	n.borderRadius = r
	return n
}
func (n *node) Shadow(s ...bool) *node {
	b := solveBool(s)
	n.shadow = *b
	return n
}
func (n *node) Is3d(d ...bool) *node {
	b := solveBool(d)
	n.is3d = *b
	return n
}
func (n *node) Multiple(m ...bool) *node {
	b := solveBool(m)
	n.multiple = *b
	return n
}
func (n *node) DoubleBorder(d ...bool) *node {
	b := solveBool(d)
	n.doubleBorder = *b
	return n
}
func (n *node) FontSize(s int) *node {
	n.fontSize = s
	return n
}
func (n *node) FontColor(c string) *node {
	n.fontColor = c
	return n
}
func (n *node) Font(f string) *node {
	n.font = f
	return n
}
func (n *node) Bold(bold ...bool) *node {
	b := solveBool(bold)
	n.bold = *b
	return n
}
func (n *node) Italic(i ...bool) *node {
	b := solveBool(i)
	n.italic = *b
	return n
}
func (n *node) Underline(u ...bool) *node {
	b := solveBool(u)
	n.underline = *b
	return n
}

// --- connection ---

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

// SrcHeadFilled sets if the connection's source head is filled
func (c *connection) SrcHeadFilled(f ...bool) *connection {
	c.srcHead.filledFlag = true
	c.srcHead.filled = *solveBool(f)
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

// DstHeadFilled sets if the connection's target head is filled
func (c *connection) DstHeadFilled(f ...bool) *connection {
	c.dstHead.filledFlag = true
	c.dstHead.filled = *solveBool(f)
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

func (c *connection) Label(label string) *connection {
	if label == "" {
		c.blankLabel = true
	}
	c.label = label
	return c
}
func (c *connection) Shape(shape string) *connection {
	c.shape = shape
	return c
}
func (c *connection) Stroke(stroke string) *connection {
	c.stroke = stroke
	return c
}
func (c *connection) Icon(i string) *connection {
	c.icon = i
	return c
}
func (c *connection) Opacity(op float64) *connection {
	c.opacity = op
	return c
}
func (c *connection) StrokeWidth(w int) *connection {
	c.strokeWidth = w
	return c
}
func (c *connection) StrokeDash(d int) *connection {
	c.strokeDash = d
	return c
}
func (c *connection) BorderRadius(r int) *connection {
	c.borderRadius = r
	return c
}
func (c *connection) FontSize(s int) *connection {
	c.fontSize = s
	return c
}
func (c *connection) FontColor(color string) *connection {
	c.fontColor = color
	return c
}
func (c *connection) Font(f string) *connection {
	c.font = f
	return c
}
func (c *connection) Bold(bold ...bool) *connection {
	b := solveBool(bold)
	c.bold = *b
	return c
}
func (c *connection) Italic(i ...bool) *connection {
	b := solveBool(i)
	c.italic = *b
	return c
}
func (c *connection) Underline(u ...bool) *connection {
	b := solveBool(u)
	c.underline = *b
	return c
}
func (c *connection) Animated(a ...bool) *connection {
	b := solveBool(a)
	c.animated = *b
	return c
}

// --- option ---

// Copy copies the option
func (o *option) Copy() *option {
	res := *o
	return &res
}

func (o *option) Label(label string) *option {
	if label == "" {
		o.blankLabel = true
	}
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

func (o *option) Width(w int) *option {
	o.width = w
	return o
}

func (o *option) Height(h int) *option {
	o.height = h
	return o
}

func (o *option) GridRows(r int) *option {
	o.gridRows = r
	return o
}

func (o *option) GridCols(c int) *option {
	o.gridCols = c
	return o
}

func (o *option) GridGap(g int) *option {
	o.gridGap = g
	return o
}

func (o *option) VerticalGap(gap int) *option {
	o.verticalGap = gap
	return o
}

func (o *option) HorizontalGap(gap int) *option {
	o.horizontalGap = gap
	return o
}

func (o *option) Icon(i string) *option {
	o.icon = i
	return o
}

func (o *option) Opacity(op float64) *option {
	o.opacity = op
	return o
}

func (o *option) FillPattern(p string) *option {
	o.fillPattern = p
	return o
}
func (o *option) StrokeWidth(w int) *option {
	o.strokeWidth = w
	return o
}
func (o *option) StrokeDash(d int) *option {
	o.strokeDash = d
	return o
}
func (o *option) BorderRadius(r int) *option {
	o.borderRadius = r
	return o
}
func (o *option) Shadow(s ...bool) *option {
	b := solveBool(s)
	o.shadow = *b
	return o
}
func (o *option) Is3d(d ...bool) *option {
	b := solveBool(d)
	o.is3d = *b
	return o
}
func (o *option) Multiple(m ...bool) *option {
	b := solveBool(m)
	o.multiple = *b
	return o
}
func (o *option) DoubleBorder(d ...bool) *option {
	b := solveBool(d)
	o.doubleBorder = *b
	return o
}
func (o *option) FontSize(s int) *option {
	o.fontSize = s
	return o
}
func (o *option) FontColor(c string) *option {
	o.fontColor = c
	return o
}
func (o *option) Bold(bold ...bool) *option {
	b := solveBool(bold)
	o.bold = *b
	return o
}
func (o *option) Italic(i ...bool) *option {
	b := solveBool(i)
	o.italic = *b
	return o
}
func (o *option) Underline(u ...bool) *option {
	b := solveBool(u)
	o.underline = *b
	return o
}
func (o *option) Animated(a ...bool) *option {
	b := solveBool(a)
	o.animated = *b
	return o
}

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
	// todo
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
	c.cfg.Sketch = solveBool(b)
	return c
}

// ElkLayout sets if the diagram is elk layouted
func (c *config) ElkLayout(b ...bool) *config {
	c.elkLayout = *solveBool(b)
	return c
}

// Center sets if the diagram is centered
func (c *config) Center(b ...bool) *config {
	c.cfg.Center = solveBool(b)
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
	c.isSequence = *solveBool(b)
	return c
}

// const

// directions

func (c *Cdor) Right() string {
	return "right"
}
func (c *Cdor) Left() string {
	return "left"
}
func (c *Cdor) Up() string {
	return "up"
}
func (c *Cdor) Down() string {
	return "down"
}

// themes

func (c *Cdor) Neutral() int {
	return 0
}
func (c *Cdor) NeutralGrey() int {
	return 1
}
func (c *Cdor) Flagship() int {
	return 3
}
func (c *Cdor) Cool() int {
	return 4
}
func (c *Cdor) MixedBerryBlue() int {
	return 5
}
func (c *Cdor) GrapeSoda() int {
	return 6
}
func (c *Cdor) Aubergine() int {
	return 7
}
func (c *Cdor) ColorblindClear() int {
	return 8
}
func (c *Cdor) VanillaNitroCola() int {
	return 100
}
func (c *Cdor) OrangeCreamsicle() int {
	return 101
}
func (c *Cdor) ShirleyTemple() int {
	return 102
}
func (c *Cdor) EarthTones() int {
	return 103
}
func (c *Cdor) EvergladeGreen() int {
	return 104
}
func (c *Cdor) ButteredToast() int {
	return 105
}
func (c *Cdor) Terminal() int {
	return 300
}
func (c *Cdor) TerminalGrayscale() int {
	return 301
}
func (c *Cdor) Origami() int {
	return 302
}
func (c *Cdor) DarkMauve() int {
	return 200
}
func (c *Cdor) DarkFlagshipTerrastruct() int {
	return 201
}

// shapes

func (c *Cdor) Rectangle() string {
	return "rectangle"
}
func (c *Cdor) Square() string {
	return "square"
}
func (c *Cdor) Page() string {
	return "page"
}
func (c *Cdor) Parallelogram() string {
	return "parallelogram"
}
func (c *Cdor) Document() string {
	return "document"
}
func (c *Cdor) Cylinder() string {
	return "cylinder"
}
func (c *Cdor) Queue() string {
	return "queue"
}
func (c *Cdor) Pkg() string {
	return "package"
}
func (c *Cdor) Step() string {
	return "step"
}
func (c *Cdor) Callout() string {
	return "callout"
}
func (c *Cdor) StoredData() string {
	return "stored_data"
}
func (c *Cdor) Person() string {
	return "person"
}
func (c *Cdor) Diamond() string {
	return "diamond"
}
func (c *Cdor) Oval() string {
	return "oval"
}
func (c *Cdor) Circle() string {
	return "circle"
}
func (c *Cdor) Hexagon() string {
	return "hexagon"
}
func (c *Cdor) Cloud() string {
	return "cloud"
}
