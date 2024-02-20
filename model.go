package cdor

import (
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type Cdor struct {
	graph       *d2graph.Graph
	ruler       *textmeasure.Ruler
	nodes       []*Node
	connections []*Connection
	err         error
	built       bool
}

type Node struct {
	children    []*Node
	connections []*Connection
	id          string
	option      *Option
}

type Connection struct {
	option   *Option
	src, dst string
}

type Option struct {
	Label string
	Shape string
	Style
}

type Style struct {
	Fill   string
	Stroke string
}
