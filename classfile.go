package cdor

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

const GopPackage = true

type IApp interface {
	init()
	setWorkers([]IWorker)
}

type IWorker interface {
	init()
	Gen() ([]byte, error)
	getNodes() []*node
	getCons() []*connection
}

var _ IApp = (*App)(nil)
var _ IWorker = (*Cdor)(nil)

func Gopt_App_Main(app IApp, workers ...IWorker) {
	app.init()
	app.setWorkers(workers)
	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
	}
	app.(interface{ MainEntry() }).MainEntry()
}

// --- App, project class ---

type App struct {
	Cdor
	Workers []IWorker
}

func (a *App) RangeDiagrams(action func(string, []byte, error) error) {
	for _, w := range a.Workers {
		data, err := w.Gen()
		if action(className(w), data, err) != nil {
			return
		}
	}
}

func (a *App) SaveFiles(dir ...string) (err error) {
	directory := ""
	if len(dir) == 0 || dir[0] == "" {
		if directory, err = os.Getwd(); err != nil {
			return
		}
	} else {
		directory = dir[0]
	}
	var data []byte
	for _, w := range a.Workers {
		if data, err = w.Gen(); err != nil {
			return err
		}
		file := filepath.Join(directory, fmt.Sprintf("%s.svg", className(w)))
		if err = os.WriteFile(file, data, 0600); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Merge() {
	for _, worker := range a.Workers {
		a.Cdor.nodes = append(a.Cdor.nodes, worker.getNodes()...)
		a.Cdor.connections = append(a.Cdor.connections, worker.getCons()...)
	}
}

func (a *App) init() {
	a.Cdor.init()
}

func (a *App) setWorkers(workers []IWorker) {
	a.Workers = workers
}

func className(worker IWorker) string {
	return reflect.TypeOf(worker).Elem().Name()
}
