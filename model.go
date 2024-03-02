package cdor

import (
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2target"
)

type Cdor struct {
	graph        *d2graph.Graph
	nodes        []*node
	connections  []*connection
	globalOption *option
	err          error
	*config
}

type node struct {
	*Cdor
	children    []*node
	connections []*connection
	sqlFields   []sqlField
	id          string
	codeTag     string
	code        string
	idSolved    bool
	*option
}

type connection struct {
	*Cdor
	*conOption
	isSingle bool
	src, dst string
}

type option struct {
	label         string
	blankLabel    bool
	shape         string
	icon          string
	width         int
	height        int
	gridRows      int
	gridCols      int
	gridGap       int
	verticalGap   int
	horizontalGap int
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

type conOption struct {
	arrow
	option
}

type sqlField struct {
	key        string
	value      string
	constraint string
}

type config struct {
	cfg        d2target.Config
	elkLayout  bool
	isSequence bool
	direction  string
}

type ThemeOverrides = d2target.ThemeOverrides
