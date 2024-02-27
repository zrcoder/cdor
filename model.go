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
	globalConOpt *conOption
	err          error
	built        bool
	config
}

type node struct {
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
