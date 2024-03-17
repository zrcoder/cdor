// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"bytes"
	"fmt"
	"github.com/qiniu/x/stringutil"
	"github.com/zrcoder/cdor"
	"os"
	"path/filepath"
	"strings"
)

const _ = true
const cdorSuffix = "_cdor.gox"

type c01hello struct {
	cdor.Cdor
	*App
}
type c06md struct {
	cdor.Cdor
	*App
}
type c07latex struct {
	cdor.Cdor
	*App
}
type c08sql_table struct {
	cdor.Cdor
	*App
}
type c10jsonn struct {
	cdor.Cdor
	*App
}
type c14icon struct {
	cdor.Cdor
	*App
}
type c16shapes struct {
	cdor.Cdor
	*App
}
type App struct {
	cdor.Mgr
}
//line doc/main_cdor.gox:9
func (this *App) MainEntry() {
//line doc/main_cdor.gox:9:1
	this.ApplyConfig(this.Cfg().Sketch().ElkLayout().Theme(this.ButteredToast()).DarkTheme(this.DarkMauve()))
//line doc/main_cdor.gox:11:1
	buf := bytes.NewBuffer(nil)
//line doc/main_cdor.gox:13:1
	this.RangeCdors(func(name string, cdor *cdor.Cdor, err error) error {
//line doc/main_cdor.gox:14:1
		if err != nil {
//line doc/main_cdor.gox:15:1
			fmt.Println(err)
//line doc/main_cdor.gox:16:1
			return err
		}
//line doc/main_cdor.gox:19:1
		var code []byte
//line doc/main_cdor.gox:20:1
		code, err = os.ReadFile(filepath.Join("doc", name+cdorSuffix))
//line doc/main_cdor.gox:21:1
		if err != nil {
//line doc/main_cdor.gox:22:1
			fmt.Println(err)
//line doc/main_cdor.gox:23:1
			return err
		}
//line doc/main_cdor.gox:25:1
		cdor.MdCode(string(code))
//line doc/main_cdor.gox:26:1
		var data []byte
//line doc/main_cdor.gox:27:1
		data, err = cdor.Gen()
//line doc/main_cdor.gox:28:1
		if err != nil {
//line doc/main_cdor.gox:29:1
			fmt.Println(err)
//line doc/main_cdor.gox:30:1
			return err
		}
//line doc/main_cdor.gox:32:1
		name = name[3:]
//line doc/main_cdor.gox:33:1
		err = os.WriteFile(stringutil.Concat("doc/examples/", name, ".svg"), data, 0600)
//line doc/main_cdor.gox:34:1
		if err != nil {
//line doc/main_cdor.gox:35:1
			fmt.Println(err)
//line doc/main_cdor.gox:36:1
			return err
		}
//line doc/main_cdor.gox:38:1
		buf.WriteString(stringutil.Concat("![", name, "](doc/examples/", name, ".svg)\n"))
//line doc/main_cdor.gox:39:1
		return nil
	})
//line doc/main_cdor.gox:42:1
	readmeTemp, err := os.ReadFile("doc/readmeTemp.md")
//line doc/main_cdor.gox:43:1
	if err != nil {
//line doc/main_cdor.gox:44:1
		panic(err)
	}
//line doc/main_cdor.gox:47:1
	readmeContent := string(readmeTemp)
//line doc/main_cdor.gox:49:1
	readmeContent = strings.Replace(readmeContent, "{{ .Examples }}", buf.String(), 1)
//line doc/main_cdor.gox:50:1
	if
//line doc/main_cdor.gox:50:1
	err := os.WriteFile("README.md", []byte(readmeContent), 0600); err != nil {
//line doc/main_cdor.gox:51:1
		panic(err)
	}
}
func main() {
	cdor.Gopt_App_Main(new(App), new(c01hello), new(c06md), new(c07latex), new(c08sql_table), new(c10jsonn), new(c14icon), new(c16shapes))
}
//line doc/c01hello_cdor.gox:1
func (this *c01hello) Main() {
//line doc/c01hello_cdor.gox:1:1
	this.Direction(this.Right())
//line doc/c01hello_cdor.gox:2:1
	this.Con("Go+", "Go").Label("cdor")
}
//line doc/c06md_cdor.gox:1
func (this *c06md) Main() {
//line doc/c06md_cdor.gox:1:1
	mdContent := `# I can do headers
  - lists
  - lists

  And other normal markdown stuff
`
//line doc/c06md_cdor.gox:7:1
	this.Node("markdown").Code("md", mdContent)
}
//line doc/c07latex_cdor.gox:1
func (this *c07latex) Main() {
//line doc/c07latex_cdor.gox:1:1
	tex := `\lim_{h \rightarrow 0 } \frac{f(x+h)-f(x)}{h}`
//line doc/c07latex_cdor.gox:2:1
	this.Node("tex").Code("latex", tex)
}
//line doc/c08sql_table_cdor.gox:1
func (this *c08sql_table) Main() {
//line doc/c08sql_table_cdor.gox:1:1
	this.Node("table").Shape("sql_table").Field("id", "int", "primary_key").Field("last_updated", "timestamp with time zone")
}
//line doc/c10jsonn_cdor.gox:1
func (this *c10jsonn) Main() {
//line doc/c10jsonn_cdor.gox:1:1
	data := `{
   "fruit":"Apple", 
   "colors": ["Red", "Green"]
}`
//line doc/c10jsonn_cdor.gox:6:1
	this.Node("obj").Json(data)
//line doc/c10jsonn_cdor.gox:7:1
	this.Scon("root", "obj.0")
//line doc/c10jsonn_cdor.gox:8:1
	this.Direction(this.Right())
}
//line doc/c14icon_cdor.gox:3
func (this *c14icon) Main() {
//line doc/c14icon_cdor.gox:3:1
	iconPath := filepath.Join("doc", "examples", "icons", "github.svg")
//line doc/c14icon_cdor.gox:4:1
	this.Node("github").Icon(iconPath)
//line doc/c14icon_cdor.gox:5:1
	this.Node("gg").Icon(iconPath).Shape("image")
}
//line doc/c16shapes_cdor.gox:1
func (this *c16shapes) Main() {
//line doc/c16shapes_cdor.gox:1:1
	this.GridCols(1).GridGap(0)
//line doc/c16shapes_cdor.gox:3:1
	op := this.Opt().Label("").Fill("transparent").Stroke("transparent")
//line doc/c16shapes_cdor.gox:5:1
	this.Node("1").Opt(op).Children(this.Node("rectangle").Shape(this.Rectangle()), this.Node("square").Shape(this.Square()), this.Node("page").Shape(this.Page()), this.Node("parallelogram").Shape(this.Parallelogram()))
//line doc/c16shapes_cdor.gox:11:1
	this.Node("2").Opt(op).Children(this.Node("document").Shape(this.Document()), this.Node("cylinder").Shape(this.Cylinder()), this.Node("queue").Shape(this.Queue()), this.Node("package").Shape(this.Pkg()), this.Node("step").Shape(this.Step()))
//line doc/c16shapes_cdor.gox:18:1
	this.Node("3").Opt(op).Children(this.Node("callout").Shape(this.Callout()), this.Node("stored_data").Shape(this.StoredData()), this.Node("person").Shape(this.Person()), this.Node("diamond").Shape(this.Diamond()))
//line doc/c16shapes_cdor.gox:24:1
	this.Node("4").Opt(op).Children(this.Node("oval").Shape(this.Oval()), this.Node("circle").Shape(this.Circle()), this.Node("hexagon").Shape(this.Hexagon()), this.Node("cloud").Shape(this.Cloud()))
}
