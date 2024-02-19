package cdor

import "encoding/json"

type Option struct {
	Label string
	Shape string
	Style
}

type Style struct {
	Fill   string
	Stroke string
}

func (o *Option) class() string {
	data, _ := json.Marshal(o)
	return string(data)
}

func O() *Option {
	return &Option{}
}

func (o *Option) L(label string) *Option {
	o.Label = label
	return o
}

func (o *Option) Sh(shape string) *Option {
	o.Shape = shape
	return o
}

func (o *Option) Sty(style *Style) *Option {
	o.Style = *style
	return o
}

func (o *Option) F(fill string) *Option {
	o.Fill = fill
	return o
}

func (o *Option) S(stroke string) *Option {
	o.Stroke = stroke
	return o
}

func S() *Style {
	return &Style{}
}

func (s *Style) F(fill string) *Style {
	s.Fill = fill
	return s
}

func (s *Style) S(stroke string) *Style {
	s.Stroke = stroke
	return s
}
