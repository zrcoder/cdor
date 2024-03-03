package cdor

const GopPackage = true

type IMain interface {
	init([]IWorker)
}

type IWorker interface {
	init()
	Gen() ([]byte, error)
	getNodes() []*node
	getCons() []*connection
	getConfig() *config
	getBaseOption() *option
	ApplyOption(*option) *Cdor
	ApplyConfig(*config) *Cdor
}

type App = Mgr

var _ IMain = (*App)(nil)
var _ IWorker = (*Cdor)(nil)

func Gopt_App_Main(app IMain, workers ...IWorker) {
	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
	}
	app.init(workers)
	app.(interface{ MainEntry() }).MainEntry()
}
