package cdor

import "fmt"

func Example_hello() {
	print(Ctx().
		Cons(
			Con("x", "y", Opt().Label("hello, cdor!")),
		).
		d2())
	// Output:
	// x <-> y: hello, cdor!
}

func Example_id() {
	print(Ctx().
		Nodes(
			Node("imAShape"),
			Node("im_a_shape"),
			Node("im a shape"),
			Node("i'm a shape"),
			Node("a-shape"),
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
	print(Ctx().
		Nodes(
			Node("pg", Opt().Label("PostgreSQL")),
			Node("Cloud", Opt().Label("my cloud")),
		).
		d2())
	// Output:
	// pg: PostgreSQL
	// Cloud: my cloud
}

func Example_shape() {
	print(Ctx().
		Nodes(
			Node("cloud", Opt().Shape("cloud")),
		).
		d2())
	// Output:
	// cloud: {shape: cloud}
}

func Example_connection() {
	print(Ctx().
		Cons(
			Con("x", "y", Opt().Stroke("red")),
			Con("x", "y", Opt().Stroke("blue")),
		).
		d2())
	// Output:
	// x <-> y: {style.stroke: red}
	// x <-> y: {style.stroke: blue}
}

func Example_container() {
	print(Ctx().
		Nodes(
			Node("server.process"),
			Node("im a parent.im a child"),
		).
		Cons(
			Con("apartment.Bedroom.Bathroom", "office.Spare Room.Bathroom", Opt().Label("Portal")),
		).
		d2())
	// Output:
	// server.process
	// im a parent.im a child
	// apartment.Bedroom.Bathroom <-> office.Spare Room.Bathroom: Portal
}

func Example_nested_containers() {
	print(Ctx().
		Nodes(
			Node("clouds").
				Children(
					Node("aws").Cons(
						Con("load_balancer", "api"),
						Con("api", "db"),
					),
					Node("gcloud").Cons(
						Con("auth", "db"),
					),
				).
				Cons(
					Con("gcloud", "aws"),
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
	print(Ctx().
		Nodes(
			Node("christmas", Opt().Fill("#ACE1AF")),
		).
		Cons(
			Con("christmas.presents", "birthdays.presents", Opt().Label("regift")),
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
