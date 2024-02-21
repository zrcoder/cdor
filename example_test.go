package cdor

import "fmt"

func Example_hello() {
	c := Ctx()
	print(c.
		Cons(
			c.Con("a", "b").Label("hi"),
			c.Con("x", "y", c.Opt().Label("hello")),
		).
		d2())
	// Output:
	// a <-> b: hi
	// x <-> y: hello
}

func Example_id() {
	c := Ctx()
	print(c.
		Nodes(
			c.Node("imAShape"),
			c.Node("im_a_shape"),
			c.Node("im a shape"),
			c.Node("i'm a shape"),
			c.Node("a-shape"),
		).
		d2())
	// Output:
	// imAShape
	// im_a_shape
	// im a shape
	// i'm a shape
	// a-shape
}

func Example_label() {
	c := Ctx()
	print(c.
		Nodes(
			c.Node("pg").Label("PostgreSQL"),
			c.Node("Cloud").Label("my cloud"),
		).
		d2())
	// Output:
	// pg: PostgreSQL
	// Cloud: my cloud
}

func Example_shape() {
	c := Ctx()
	print(c.
		Nodes(
			c.Node("cloud").Shape("cloud"),
		).
		d2())
	// Output:
	// cloud: {shape: cloud}
}

func Example_connection() {
	c := Ctx()
	print(c.
		Cons(
			c.Con("x", "y").Stroke("red"),
			c.Con("x", "y").Stroke("blue"),
		).
		d2())
	// Output:
	// x <-> y: {style.stroke: red}
	// x <-> y: {style.stroke: blue}
}

func Example_container() {
	c := Ctx()
	print(c.
		Nodes(
			c.Node("server.process"),
			c.Node("im a parent.im a child"),
		).
		Cons(
			c.Con("apartment.Bedroom.Bathroom", "office.Spare Room.Bathroom").Label("Portal"),
		).
		d2())
	// Output:
	// server.process
	// im a parent.im a child
	// apartment.Bedroom.Bathroom <-> office.Spare Room.Bathroom: Portal
}

func Example_nested_containers() {
	c := Ctx()
	print(c.
		Nodes(
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
				),
		).
		d2())
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
	print(c.
		Nodes(
			c.Node("christmas").Fill("#ACE1AF"),
		).
		Cons(
			c.Con("christmas.presents", "birthdays.presents").Label("regift"),
		).
		d2())
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
