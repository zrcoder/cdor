package cdor

import "fmt"

func Example_hello() {
	c := Ctx()
	c.Con("a", "b").Label("hi") // or c.Con("a", "b", c.Opt().Label("hi"))
	print(c.d2())
	// Output:
	// a <-> b: hi
}

func Example_id() {
	c := Ctx()
	c.Node("imAShape")
	c.Node("im_a_shape")
	c.Node("im a shape")
	c.Node("i'm a shape")
	c.Node("a-shape")
	print(c.d2())
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
	print(c.d2())
	// Output:
	// pg: PostgreSQL
	// Cloud: my cloud
}

func Example_shape() {
	c := Ctx()
	c.Node("cloud").Shape("cloud")
	print(c.d2())
	// Output:
	// cloud: {shape: cloud}
}

func Example_connection() {
	c := Ctx()
	c.Con("x", "y").Stroke("red")
	c.Con("x", "y").Stroke("blue")
	print(c.d2())
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
	print(c.d2())
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
	print(c.d2())
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
	print(c.d2())
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
	print(c.d2())
	// Output:
	// christmas: {style.fill: "#ACE1AF"}
	// christmas.presents <-> birthdays.presents: regift
}

func print(d2 string, err error) {
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(d2)
	}
}
