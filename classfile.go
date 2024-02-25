package cdor

const GopPackage = true

type App struct {
	Cdor
	Workers []IWorker
}

func (a *App) RangeDiagrams(action func(data []byte, err error) error) {
	for _, w := range a.Workers {
		data, err := w.Gen()
		if action(data, err) != nil {
			return
		}
	}
}

func (a *App) init() {
	a.Cdor.init()
}

func (a *App) addWorker(w IWorker) {
	a.Workers = append(a.Workers, w)
}

type IApp interface {
	init()
	addWorker(IWorker)
}

type IWorker interface {
	init()
	Gen() ([]byte, error)
}

func Gopt_App_Main(app IApp, workers ...IWorker) {
	app.init()

	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
		app.addWorker(worker)
	}

	if me, ok := app.(interface{ MainEntry() }); ok {
		me.MainEntry()
	}
}
