package cdor

const GopPackage = true

type IMain interface {
	init()
	addWorker(IWorker)
}

type IWorker interface {
	init()
	Gen() ([]byte, error)
	getNodes() []*node
	getCons() []*connection
	ApplyOption(*option) *Cdor
	BaseOption(*option) *Cdor
	ApplyConfig(*config) *Cdor
	BaseConfig(*config) *Cdor
	ApplyConOption(opt *conOption) *Cdor
	BaseConOption(*conOption) *Cdor
}

var _ IMain = (*App)(nil)
var _ IWorker = (*Cdor)(nil)

func Gopt_App_Main(app IMain, workers ...IWorker) {
	app.init()
	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
		app.addWorker(worker)
	}
	app.(interface{ MainEntry() }).MainEntry()
}
