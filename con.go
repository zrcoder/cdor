package cdor

type Connection struct {
	option   *Option
	src, dst string
}

func C(src, dst string, opt *Option) *Connection {
	return &Connection{src: src, dst: dst, option: opt}
}
