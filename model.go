package cdor

import (
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2target"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

type Cdor struct {
	graph       *d2graph.Graph
	ruler       *textmeasure.Ruler
	nodes       []*node
	connections []*connection
	option      *option
	arrow       *arrow
	err         error
	built       bool
	config
}

type node struct {
	children    []*node
	connections []*connection
	id          string
	idSolved    bool
	*option
}

type connection struct {
	*option
	*arrow
	src, dst string
}

type option struct {
	label string
	shape string
	style
}

type style struct {
	fill       string
	filled     bool
	filledFlag bool
	stroke     string
}

type arrow struct {
	srcHead option
	dstHead option
}

type config struct {
	cfg       d2target.Config
	elkLayout bool
	direction string
}

type ThemeOverrides = d2target.ThemeOverrides
