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
	getBaseConOption() *conOption
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
	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
	}
	app.init(workers)
	app.(interface{ MainEntry() }).MainEntry()
}
