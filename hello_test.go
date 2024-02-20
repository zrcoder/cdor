package cdor

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"oss.terrastruct.com/d2/d2format"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2layouts/d2dagrelayout"
	"oss.terrastruct.com/d2/d2lib"
	"oss.terrastruct.com/d2/d2oracle"
	"oss.terrastruct.com/d2/lib/textmeasure"
)

const d2wanted = `
cat: {
  meow: {style.fill: green}
}
dog: ddd {shape: circle}
cat.meow -> dog: haha {style.stroke: red}
`

func common() (string, error) {
	res, err := New().
		Nodes(
			N("cat").
				Children(
					N("meow", O().F("green")),
				),
			N("dog", O().Sh("circle").L("ddd")),
		).
		Cons(
			C("cat.meow", "dog", O().L("haha").S("red")),
		).
		D2()
	return strings.TrimSpace(res), err
}

func low() (string, error) {
	c := New()
	c.add("cat")
	c.add("cat.meow")
	c.set("cat.meow", "style.fill", "green")
	c.add("dog")
	c.set("dog", "shape", "circle")
	c.set("dog", "label", "ddd")
	key := c.con("cat.meow", "dog")
	c.set(key, "label", "haha")
	c.set(key, "style.stroke", "red")
	res, err := c.D2()
	return strings.TrimSpace(res), err
}

func d2() (string, error) {
	ruler, _ := textmeasure.NewRuler()
	layoutResolver := func(engine string) (d2graph.LayoutGraph, error) {
		return d2dagrelayout.DefaultLayout, nil
	}
	compileOpts := &d2lib.CompileOptions{
		LayoutResolver: layoutResolver,
		Ruler:          ruler,
	}
	_, graph, _ := d2lib.Compile(context.Background(), "", compileOpts, nil)

	graph, _, _ = d2oracle.Create(graph, nil, "cat")
	graph, _, _ = d2oracle.Create(graph, nil, "cat.meow")
	color := "green"
	graph, _ = d2oracle.Set(graph, nil, "cat.meow.style.fill", nil, &color)

	graph, _, _ = d2oracle.Create(graph, nil, "dog")
	circle := "circle"
	graph, _ = d2oracle.Set(graph, nil, "dog.shape", nil, &circle)
	ddd := "ddd"
	graph, _ = d2oracle.Set(graph, nil, "dog.label", nil, &ddd)

	key := "cat.meow -> dog"
	newKey := ""
	graph, newKey, _ = d2oracle.Create(graph, nil, key)
	haha := "haha"
	graph, _ = d2oracle.Set(graph, nil, fmt.Sprintf("%s.label", newKey), nil, &haha)
	red := "red"
	graph, _ = d2oracle.Set(graph, nil, fmt.Sprintf("%s.style.stroke", newKey), nil, &red)

	return strings.TrimSpace(d2format.Format(graph.AST)), nil
}

func TestHello(t *testing.T) {
	want := strings.TrimSpace(d2wanted)
	if res, err := d2(); err != nil || res != want {
		t.Errorf("err: %v, got: %s\n", err, res)
	}
	if res, err := low(); err != nil || res != want {
		t.Errorf("err: %v, got: %s\n", err, res)
	}
	if res, err := common(); err != nil || res != want {
		t.Errorf("err: %v, got: %s\n", err, res)
	}
}
