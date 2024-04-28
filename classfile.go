package cdor

const GopPackage = true

type IMain interface {
	init([]IWorker)
}

type IWorker interface {
	init()
	cdor() *Cdor
}

type App = Mgr

var (
	_ IMain   = (*App)(nil)
	_ IWorker = (*Cdor)(nil)
)

func Gopt_App_Main(app IMain, workers ...IWorker) {
	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
	}
	app.init(workers)
	app.(interface{ MainEntry() }).MainEntry()
}
