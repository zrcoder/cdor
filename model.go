package cdor

import (
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2target"
)

type Cdor struct {
	graph       *d2graph.Graph
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
	codeTag     string
	code        string
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
	icon  string
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
