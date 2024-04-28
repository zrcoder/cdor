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

// Node creats a node with default shape rectangle
func (c *Cdor) Node(id string) *node {
	node := &node{
		id:     id,
		option: c.Opt(),
		Cdor:   c,
	}
	c.nodes = append(c.nodes, node)
	return node
}

// Rectangle creates a rectangle node
func (c *Cdor) Rectangle(id string) *node {
	return c.Node(id).Shape("rectangle")
}

// Square creates a square node
func (c *Cdor) Square(id string) *node {
	return c.Node(id).Shape("square")
}

// Page creates a page node
func (c *Cdor) Page(id string) *node {
	return c.Node(id).Shape("page")
}

// Parallelogram creates a parallelogram node
func (c *Cdor) Parallelogram(id string) *node {
	return c.Node(id).Shape("parallelogram")
}

// Document creates a document node
func (c *Cdor) Document(id string) *node {
	return c.Node(id).Shape("document")
}

// Cylinder creates a cylinder node
func (c *Cdor) Cylinder(id string) *node {
	return c.Node(id).Shape("cylinder")
}

// Queue creates a queue node
func (c *Cdor) Queue(id string) *node {
	return c.Node(id).Shape("queue")
}

// Pkg creates a package node
func (c *Cdor) Pkg(id string) *node {
	return c.Node(id).Shape("package")
}

// Step creates a step node
func (c *Cdor) Step(id string) *node {
	return c.Node(id).Shape("step")
}

// Callout creates a callout node
func (c *Cdor) Callout(id string) *node {
	return c.Node(id).Shape("callout")
}

// StoredData creates a stored_data node
func (c *Cdor) StoredData(id string) *node {
	return c.Node(id).Shape("stored_data")
}

// Person creates a person node
func (c *Cdor) Person(id string) *node {
	return c.Node(id).Shape("person")
}

// Diamond creates a diamond node
func (c *Cdor) Diamond(id string) *node {
	return c.Node(id).Shape("diamond")
}

// Oval creates a oval node
func (c *Cdor) Oval(id string) *node {
	return c.Node(id).Shape("oval")
}

// Circle creates a circle node
func (c *Cdor) Circle(id string) *node {
	return c.Node(id).Shape("circle")
}

// Hexagon creates a hexagon node
func (c *Cdor) Hexagon(id string) *node {
	return c.Node(id).Shape("hexagon")
}

// Cloud creates a cloud node
func (c *Cdor) Cloud(id string) *node {
	return c.Node(id).Shape("cloud")
}

// SqlTable creates a sql_table node
func (c *Cdor) SqlTable(id string) *node {
	return c.Node(id).Shape("sql_table")
}

// Image creates an image node
func (c *Cdor) Image(id string) *node {
	return c.Node(id).Shape("image")
}

// Class creates a uml class node
func (c *Cdor) Class(id string) *node {
	return c.Node(id).Shape("class")
}

// Text creates a text node
func (c *Cdor) Text(id, content string) *node {
	return c.Node(id).Shape("text").Label(content)
}

// Code creates a code node
func (c *Cdor) Code(id, tag, code string) *node {
	if tag == "latex" || tag == "tex" {
		code = strings.ReplaceAll(code, `\`, `\\`)
	}
	node := c.Node(id)
	node.codeTag = tag
	node.code = code
	return node
}

// MdCode creates a node with md code style
func (c *Cdor) MdCode(code string, id ...string) *node {
	const mdCodeTemp = "```\n%s\n```"
	key := "md"
	if len(id) > 0 {
		key = id[0]
	}
	return c.Markdown(key, fmt.Sprintf(mdCodeTemp, code))
}

// Markdown creates a markdown node
func (c *Cdor) Markdown(id, content string) *node {
	return c.Code(id, "md", content)
}

// Latex creates a latex node
func (c *Cdor) Latex(id, content string) *node {
	return c.Code(id, "latex", content)
}

// Con creats a connection
func (c *Cdor) Con(src, dst string) *connection {
	con := &connection{
		src:       src,
		dst:       dst,
		conOption: &conOption{},
		Cdor:      c,
	}
	con.conOption.option = *c.Opt()
	con.conOption.arrow.srcHead = *c.Opt()
	con.conOption.arrow.dstHead = *c.Opt()
	c.connections = append(c.connections, con)
	return con
}

// Scon creats a single connection
func (c *Cdor) Scon(src, dst string) *connection {
	con := c.Con(src, dst)
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
	return os.WriteFile(filename, data, 0o600)
}

// Clear clears the diagram
func (c *Cdor) Clear() {
	c.init()
	c.nodes = nil
	c.connections = nil
	c.isSequence = false
}

// Json creates nodes and connections for a json
func (c *Cdor) Json(json string) *Cdor {
	c.direction = "right"
	c.elkLayout = true
	c.json(json)
	return c
}

// Jsonn creates a node with id id and children with json
func (c *Cdor) Jsonn(id, json string) *node {
	node := c.Node(id)
	children, cons := c.json(json)
	node.Children(children...)
	node.Cons(cons...)
	return node
}

// Yaml creates nodes and connections for a yaml
func (c *Cdor) Yaml(yaml string) *Cdor {
	json := c.yaml2json(yaml)
	return c.Json(json)
}

// Yamln creats a node with id id and children with yaml
func (c *Cdor) Yamln(id, yaml string) *node {
	json := c.yaml2json(yaml)
	return c.Jsonn(id, json)
}

// Toml creates nodes and connections for a toml
func (c *Cdor) Toml(toml string) *Cdor {
	json := c.toml2json(toml)
	return c.Json(json)
}

// Tomln creates a node with id id and children with toml
func (c *Cdor) Tomln(id, toml string) *node {
	json := c.toml2json(toml)
	return c.Jsonn(id, json)
}

// Obj creates nodes and connections for an object
func (c *Cdor) Obj(x any) *Cdor {
	json := c.obj2json(x)
	return c.Json(json)
}

// Objn creates a node with id id and children with obj
func (c *Cdor) Objn(id string, obj any) *node {
	json := c.obj2json(obj)
	return c.Jsonn(id, json)
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

// FillPattern sets the global fill pattern
func (c *Cdor) FillPattern(pattern string) *Cdor {
	c.globalOption.fillPattern = pattern
	return c
}

// D2 append d2 script as part of the diagram
func (c *Cdor) D2(d2script string) *Cdor {
	c.d2s = append(c.d2s, d2script)
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

func (n *node) Tooltip(tip string) *node {
	n.tooltip = tip
	return n
}

func (n *node) Link(link string) *node {
	n.link = link
	return n
}

func (n *node) Near(position string) *node {
	n.position = position
	return n
}

func (n *node) LabelNear(position string) *node {
	n.labelPosition = position
	return n
}

func (n *node) IconNear(position string) *node {
	n.iconPosition = position
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
func (c *connection) Con(dst string) *connection {
	return c.Cdor.Con(c.dst, dst)
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

// Sequence sets the diagram as a sequence
func (c *config) Sequence() *config {
	c.isSequence = true
	return c
}

// Sequencen creates a node with a sequence shape
func (c *Cdor) Sequencen(id string) *node {
	return c.Node(id).Sequence()
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

// arrow head shapes

func (c *Cdor) HeadNone() string {
	return "none"
}

func (c *Cdor) HeadTriangle() string {
	return "triangle"
}

func (c *Cdor) HeadArrow() string {
	return "arrow"
}

func (c *Cdor) HeadDiamond() string {
	return "diamond"
}

func (c *Cdor) HeadCfOne() string {
	return "cf-one"
}

func (c *Cdor) HeadCfOneRequired() string {
	return "cf-one-required"
}

func (c *Cdor) HeadCfMany() string {
	return "cf-many"
}

func (c *Cdor) HeadCfManyRequired() string {
	return "cf-many-required"
}

// fill patterns

func (c *Cdor) Dots() string {
	return "dots"
}

func (c *Cdor) Lines() string {
	return "lines"
}

func (c *Cdor) Grain() string {
	return "grain"
}
