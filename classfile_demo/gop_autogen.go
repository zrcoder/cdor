// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/zrcoder/cdor"
	"os"
)

const _ = true

type hello struct {
	cdor.Cdor
}
//line classfile_demo/hello_cdor.gox:3
func (this *hello) MainEntry() {
//line classfile_demo/hello_cdor.gox:3:1
	this.Cons(this.Con("Go+", "Go").Label("cdor"))
//line classfile_demo/hello_cdor.gox:7:1
	data, err := this.Gen()
//line classfile_demo/hello_cdor.gox:8:1
	if err != nil {
//line classfile_demo/hello_cdor.gox:9:1
		fmt.Println("err:", err)
//line classfile_demo/hello_cdor.gox:10:1
		return
	}
//line classfile_demo/hello_cdor.gox:13:1
	os.WriteFile("hello.svg", data, 0600)
}
func main() {
	cdor.Gopt_App_Main(new(hello))
}
