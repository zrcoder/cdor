package cdor

import (
	"fmt"
)

func Example_hello() {
	c := Ctx()
	c.Direction("right")
	c.Con("Go+", "Go").Label("cdor")
	fmt.Println(c.d2())
	// Output:
	// direction: right
	// Go+ <-> Go: cdor
}

func Example_id() {
	c := Ctx()
	c.Node("imAShape")
	c.Node("im_a_shape")
	c.Node("im a shape")
	c.Node("i'm a shape")
	c.Node("a-shape")
	fmt.Println(c.d2())
	// Output:
	// imAShape
	// im_a_shape
	// im a shape
	// i'm a shape
	// a-shape
}

func Example_label() {
	c := Ctx()
	c.Node("pg").Label("PostgreSQL")
	c.Node("Cloud").Label("my cloud")
	fmt.Println(c.d2())
	// Output:
	// pg: PostgreSQL
	// Cloud: my cloud
}

func Example_shape() {
	c := Ctx()
	c.Node("cloud").Shape("cloud")
	fmt.Println(c.d2())
	// Output:
	// cloud: {shape: cloud}
}

func Example_connection() {
	c := Ctx()
	c.Con("x", "y").Stroke("red")
	c.Con("x", "y").Stroke("blue")
	fmt.Println(c.d2())
	// Output:
	// x <-> y: {style.stroke: red}
	// x <-> y: {style.stroke: blue}
}

func Example_connection_arrow() {
	c := Ctx()
	c.Con("x", "y").
		Stroke("red").
		SrcHeadLabel("from").SrcHeadShape("arrow").
		DstHeadLabel("to").DstHeadShape("diamond")
	fmt.Println(c.d2())
	// Output:
	// x <-> y: {
	//   style.stroke: red
	//   source-arrowhead.label: from
	//   source-arrowhead.shape: arrow
	//   target-arrowhead.label: to
	//   target-arrowhead.shape: diamond
	// }
}

func Example_container() {
	c := Ctx()
	c.Node("server.process")
	c.Node("im a parent.im a child")
	c.Con("apartment.Bedroom.Bathroom", "office.Spare Room.Bathroom").Label("Portal")
	fmt.Println(c.d2())
	// Output:
	// server.process
	// im a parent.im a child
	// apartment.Bedroom.Bathroom <-> office.Spare Room.Bathroom: Portal
}

func Example_nested_containers() {
	c := Ctx()
	c.Node("clouds").
		Children(
			c.Node("aws").Cons(
				c.Con("load_balancer", "api"),
				c.Con("api", "db"),
			),
			c.Node("gcloud").Cons(
				c.Con("auth", "db"),
			),
		).
		Cons(
			c.Con("gcloud", "aws"),
		)
	fmt.Println(c.d2())
	// Output:
	// clouds: {
	//   aws: {
	//     load_balancer <-> api
	//     api <-> db
	//   }
	//   gcloud: {
	//     auth <-> db
	//   }
	//   gcloud <-> aws
	// }
}

func Example_same_name_sub_containers() {
	c := Ctx()
	c.Node("christmas").Fill("#ACE1AF")
	c.Con("christmas.presents", "birthdays.presents").Label("regift")
	fmt.Println(c.d2())
	// Output:
	// christmas: {style.fill: "#ACE1AF"}
	// christmas.presents <-> birthdays.presents: regift
}

func Example_code() {
	c := Ctx()
	c.Node("markdown").Code("md", `# Hi cdor
	- Go+
	- Go
	`)
	c.Node("latex").Code("latex", `\lim_{h \rightarrow 0 } \frac{f(x+h)-f(x)}{h}`)
	fmt.Println(c.d2())
	// Output:
	// markdown: |md
	//   # Hi cdor
	//   	- Go+
	//   	- Go
	// |
	// latex: |latex \\lim_{h \\rightarrow 0 } \\frac{f(x+h)-f(x)}{h} |
}

func Example_icon() {
	c := Ctx()
	c.Node("github").Icon("https://icons.terrastruct.com/dev/github.svg")
	c.Node("gg").Icon("https://icons.terrastruct.com/dev/github.svg").Shape("image")
	fmt.Println(c.d2())
	// Output:
	// github: {icon: https://icons.terrastruct.com/dev/github.svg}
	// gg: {
	//   icon: https://icons.terrastruct.com/dev/github.svg
	//   shape: image
	// }

}

func Example_sql_table() {
	c := Ctx()
	c.Node("table").Shape("sql_table").
		Field("id", "int", "primary_key").
		Field("last_updated", "timestamp with time zone")
	fmt.Println(c.d2())
	// Output:
	// table: {
	//   shape: sql_table
	//   id: int
	//   id.constraint: primary_key
	//   last_updated: timestamp with time zone
	// }
}

func Example_class() {
	c := Ctx()
	c.Node("MyClass").
		Shape("class").
		Field("field", "[]string").
		Field("method(a uint64)", "(x, y int)").
		Field(`# peekn(n int)`, "(s string, eof bool)")
	fmt.Println(c.d2())
	// Output:
	// MyClass: {
	//   shape: class
	//   field: "[]string"
	//   method(a uint64): (x, y int)
	//   \# peekn(n int): (s string, eof bool)
	// }
}

func Example_json() {
	c := Ctx()
	json := `[13, {"37": 37}]`
	c.Json(json)
	fmt.Println(c.d2())
	// Output:
	// direction: right
	// 0: {
	//   shape: sql_table
	//   13: " "
	// }
	// 1: {
	//   shape: sql_table
	//   37: 37
	// }
	// 0.1 -> 1
}

func Example_node_json() {
	json := `{
		"fruit":"Apple"
	 }`

	c := Ctx()
	c.Node("ttt").Json(json)

	fmt.Println(c.d2())
	// Output:
	// ttt: {
	//   0: {
	//     shape: sql_table
	//     fruit: Apple
	//   }
	// }
}

func Example_grid() {
	c := Ctx()
	c.GridRows(2).GridCols(4).GridGap(0)

	c.Node("Element")
	c.Node("Atomic Number")
	c.Node("Atomic Mass")
	c.Node("Melting Point")

	c.Node("Hydrogen")
	c.Node(`"1"`)
	c.Node(`"1.008"`)
	c.Node(`"-259.16"`)

	c.Node("Carbon")
	c.Node(`"6"`)
	c.Node(`"12.011"`)
	c.Node(`"3500"`)

	c.Node("Oxygen")
	c.Node(`"8"`)
	c.Node(`"15.999"`)
	c.Node(`"-218.79"`)

	fmt.Println(c.d2())

	// Output:
	// grid-rows: 2
	// grid-columns: 4
	// grid-gap: 0
	// Element
	// Atomic Number
	// Atomic Mass
	// Melting Point
	// Hydrogen
	// "1"
	// "1.008"
	// "-259.16"
	// Carbon
	// "6"
	// "12.011"
	// "3500"
	// Oxygen
	// "8"
	// "15.999"
	// "-218.79"
}
func Example_sequece() {
	c := Ctx()
	c.Sequence()
	c.Scon("alice", "bob").Label("What does it mean\nto be well-adjusted?")
	c.Scon("bob", "alice").Label("The ability to play bridge or\ngolf as if they were games.")
	fmt.Println(c.d2())
	// Output:
	// shape: sequence_diagram
	// alice -> bob: "What does it mean\nto be well-adjusted?"
	// bob -> alice: "The ability to play bridge or\ngolf as if they were games."
}
