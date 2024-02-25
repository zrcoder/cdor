package cdor

import (
	"fmt"
)

func Example_hello() {
	c := Ctx()
	c.Con("a", "b").Label("hi") // or c.Con("a", "b", c.Opt().Label("hi"))
	fmt.Println(c.d2())
	// Output:
	// direction: down
	// a <-> b: hi
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
	// direction: down
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
	// direction: down
	// pg: PostgreSQL
	// Cloud: my cloud
}

func Example_shape() {
	c := Ctx()
	c.Node("cloud").Shape("cloud")
	fmt.Println(c.d2())
	// Output:
	// direction: down
	// cloud: {shape: cloud}
}

func Example_connection() {
	c := Ctx()
	c.Con("x", "y").Stroke("red")
	c.Con("x", "y").Stroke("blue")
	fmt.Println(c.d2())
	// Output:
	// direction: down
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
	// direction: down
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
	// direction: down
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
	// direction: down
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
	// direction: down
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
	// direction: down
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
	// direction: down
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
	// direction: down
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
	// direction: down
	// MyClass: {
	//   shape: class
	//   field: "[]string"
	//   method(a uint64): (x, y int)
	//   \# peekn(n int): (s string, eof bool)
	// }
}

func Example_sequece() {
	c := Ctx()
	c.Sequence()
	c.Scon("alice", "bob").Label("What does it mean\nto be well-adjusted?")
	c.Scon("bob", "alice").Label("The ability to play bridge or\ngolf as if they were games.")
	fmt.Println(c.d2())
	// Output:
	// direction: down
	// shape: sequence_diagram
	// alice -> bob: "What does it mean\nto be well-adjusted?"
	// bob -> alice: "The ability to play bridge or\ngolf as if they were games."
}
