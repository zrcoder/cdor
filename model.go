package cdor

import (
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type Cdor struct {
	graph       *d2graph.Graph
	ruler       *textmeasure.Ruler
	nodes       []*node
	connections []*connection
	option      *option
	err         error
	built       bool
}

type node struct {
	children    []*node
	connections []*connection
	id          string
	*option
}

type connection struct {
	*option
	src, dst string
}

// todo: arrow options
type option struct {
	label string
	shape string
	style
}

type style struct {
	fill   string
	stroke string
}
