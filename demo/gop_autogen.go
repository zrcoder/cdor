// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"bytes"
	"github.com/zrcoder/cdor"
	"os"
	"strconv"
)

const _ = true

type class struct {
	cdor.Cdor
	*App
}
type connections struct {
	cdor.Cdor
	*App
}
type containers struct {
	cdor.Cdor
	*App
}
type hello struct {
	cdor.Cdor
	*App
}
type id struct {
	cdor.Cdor
	*App
}
type latex struct {
	cdor.Cdor
	*App
}
type App struct {
	cdor.App
}
type md struct {
	cdor.Cdor
	*App
}
type shape struct {
	cdor.Cdor
	*App
}
type sql_table struct {
	cdor.Cdor
	*App
}
//line demo/main_cdor.gox:8
func (this *App) MainEntry() {
//line demo/main_cdor.gox:8:1
	buf := bytes.NewBuffer(nil)
//line demo/main_cdor.gox:10:1
	i := 0
//line demo/main_cdor.gox:11:1
	this.RangeDiagrams(func(data []byte, err error) error {
//line demo/main_cdor.gox:12:1
		if err != nil {
//line demo/main_cdor.gox:13:1
			panic(err)
//line demo/main_cdor.gox:14:1
			return err
		}
//line demo/main_cdor.gox:16:1
		err = os.WriteFile("demo/"+strconv.Itoa(i)+".svg", data, 0600)
//line demo/main_cdor.gox:17:1
		if err != nil {
//line demo/main_cdor.gox:18:1
			panic(err)
//line demo/main_cdor.gox:19:1
			return err
		}
//line demo/main_cdor.gox:21:1
		buf.WriteString("![" + strconv.Itoa(i) + "](" + strconv.Itoa(i) + ".svg)\n")
//line demo/main_cdor.gox:22:1
		i++
//line demo/main_cdor.gox:23:1
		return nil
	})
//line demo/main_cdor.gox:38:1
	if
//line demo/main_cdor.gox:38:1
	err := os.WriteFile("demo/main.md", buf.Bytes(), 0600); err != nil {
//line demo/main_cdor.gox:39:1
		panic(err)
	}
}
func main() {
	cdor.Gopt_App_Main(new(App), new(class), new(connections), new(containers), new(hello), new(id), new(latex), new(md), new(shape), new(sql_table))
}
//line demo/class_cdor.gox:1
func (this *class) Main() {
//line demo/class_cdor.gox:1:1
	this.MdCode(`node("MyClass").
	shape("class").
	field("Age", "int").
	field("+ Field", "[]string").
	field("- method(a uint64)", "(x, y int)").
	field("# peekn(n int)", "(s string, eof bool)")`)
//line demo/class_cdor.gox:7:1
	this.Node("MyClass").Shape("class").Field("Age", "int").Field("+ Field", "[]string").Field("- method(a uint64)", "(x, y int)").Field("# peekn(n int)", "(s string, eof bool)")
}
//line demo/connections_cdor.gox:1
func (this *connections) Main() {
//line demo/connections_cdor.gox:1:1
	this.MdCode(`con("x", "y").srcHeadShape("none").stroke("orange")
con("x", "y").dstHeadShape("none").stroke("green")
con("x", "y").srcHeadShape("none").dstHeadShape("none")
con("a", "b").srcHeadShape("circle")
con("b", "c").srcHeadShape("circle")
con("c", "a")`)
//line demo/connections_cdor.gox:7:1
	this.Con("x", "y").SrcHeadShape("none").Stroke("orange")
//line demo/connections_cdor.gox:8:1
	this.Con("x", "y").DstHeadShape("none").Stroke("green")
//line demo/connections_cdor.gox:9:1
	this.Con("x", "y").SrcHeadShape("none").DstHeadShape("none")
//line demo/connections_cdor.gox:10:1
	this.Con("a", "b").SrcHeadShape("circle")
//line demo/connections_cdor.gox:11:1
	this.Con("b", "c").SrcHeadShape("circle")
//line demo/connections_cdor.gox:12:1
	this.Con("c", "a")
}
//line demo/containers_cdor.gox:1
func (this *containers) Main() {
//line demo/containers_cdor.gox:1:1
	this.Direction("right")
//line demo/containers_cdor.gox:2:1
	this.MdCode(`opt := arrowOpt().srcHeadShape("none")
node("clouds").children(
	node("aws").label("AWS").cons(
		con("load_balancer", "api"),
		con("api", "db"),
	),
	node("gcloud").label("Google Cloud").cons(
		con("auth", "db").arrowOpt(opt),
	)).
	cons(
		con("gcloud", "aws").arrowOpt(opt),
	)
con("users", "clouds.aws.load_balancer").arrowOpt(opt)
con("users", "clouds.gcloud.auth").arrowOpt(opt)
con("ci.deploys", "clouds").arrowOpt(opt)
`)
//line demo/containers_cdor.gox:18:1
	opt := this.ArrowOpt().SrcHeadShape("none")
//line demo/containers_cdor.gox:19:1
	this.Node("clouds").Children(this.Node("aws").Label("AWS").Cons(this.Con("load_balancer", "api"), this.Con("api", "db")), this.Node("gcloud").Label("Google Cloud").Cons(this.Con("auth", "db").ArrowOpt(opt))).Cons(this.Con("gcloud", "aws").ArrowOpt(opt))
//line demo/containers_cdor.gox:30:1
	this.Con("users", "clouds.aws.load_balancer").ArrowOpt(opt)
//line demo/containers_cdor.gox:31:1
	this.Con("users", "clouds.gcloud.auth").ArrowOpt(opt)
//line demo/containers_cdor.gox:32:1
	this.Con("ci.deploys", "clouds").ArrowOpt(opt)
}
//line demo/hello_cdor.gox:1
func (this *hello) Main() {
//line demo/hello_cdor.gox:1:1
	this.Direction("right")
//line demo/hello_cdor.gox:2:1
	this.MdCode(`con("Go+", "Go").label("cdor")`)
//line demo/hello_cdor.gox:3:1
	this.Con("Go+", "Go").Label("cdor")
}
//line demo/id_cdor.gox:1
func (this *id) Main() {
//line demo/id_cdor.gox:1:1
	this.MdCode(`node "imAShape"
node "im_a_shape"
node "im a shape"
node "i'm a shape"
node "a-shape"
`)
//line demo/id_cdor.gox:7:1
	this.Node("imAShape")
//line demo/id_cdor.gox:8:1
	this.Node("im_a_shape")
//line demo/id_cdor.gox:9:1
	this.Node("im a shape")
//line demo/id_cdor.gox:10:1
	this.Node("i'm a shape")
//line demo/id_cdor.gox:11:1
	this.Node("a-shape")
}
//line demo/latex_cdor.gox:1
func (this *latex) Main() {
//line demo/latex_cdor.gox:1:1
	tex := `\lim_{h \rightarrow 0 } \frac{f(x+h)-f(x)}{h}`
//line demo/latex_cdor.gox:2:1
	this.MdCode("tex := `" + tex + "`\nnode(\"tex\").code(\"latex\", tex)")
//line demo/latex_cdor.gox:3:1
	this.Node("tex").Code("latex", tex)
}
//line demo/md_cdor.gox:1
func (this *md) Main() {
//line demo/md_cdor.gox:1:1
	mdContent := `# I can do headers
  - lists
  - lists

  And other normal markdown stuff
`
//line demo/md_cdor.gox:7:1
	this.MdCode("mdContent := `" + mdContent + "`\nnode(\"markdown\").code(\"md\", mdContent)")
//line demo/md_cdor.gox:8:1
	this.Node("markdown").Code("md", mdContent)
}
//line demo/shape_cdor.gox:1
func (this *shape) Main() {
//line demo/shape_cdor.gox:1:1
	this.MdCode(`node("pg").label("PostgreSQL")
node("Cloud").label("my cloud").shape("cloud")`)
//line demo/shape_cdor.gox:3:1
	this.Node("pg").Label("PostgreSQL")
//line demo/shape_cdor.gox:4:1
	this.Node("Cloud").Label("my cloud").Shape("cloud")
}
//line demo/sql_table_cdor.gox:1
func (this *sql_table) Main() {
//line demo/sql_table_cdor.gox:1:1
	this.MdCode(`node("table").shape("sql_table").
	field("id", "int", "primary_key").
	field("last_updated", "timestamp with time zone")`)
//line demo/sql_table_cdor.gox:4:1
	this.Node("table").Shape("sql_table").Field("id", "int", "primary_key").Field("last_updated", "timestamp with time zone")
}