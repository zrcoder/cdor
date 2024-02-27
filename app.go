package cdor

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type App struct {
	Cdor
	workers     map[string]IWorker
	workerNames []string
}

func (a *App) rangeAction(action func(string, IWorker), diagrams []string) {
	if len(diagrams) == 0 {
		diagrams = a.workerNames // for all workers
	}
	for _, name := range diagrams {
		worker, ok := a.workers[name]
		if !ok {
			continue
		}
		action(name, worker)
	}
}

func (a *App) RangeDiagrams(action func(string, []byte, error) error, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		data, err := worker.Gen()
		if action(name, data, err) != nil {
			return
		}
	}, diagrams)
}

func (a *App) ApplyGlobalConfig(cfg *config, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.ApplyConfig(cfg)
	}, diagrams)
}

func (a *App) BaseGlobalConfig(cfg *config, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.BaseConfig(cfg)
	}, diagrams)
}

func (a *App) BaseOptionAll(opt *option, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.BaseOption(opt)
	}, diagrams)
}

func (a *App) ApplyOptionAll(opt *option, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.ApplyOption(opt)
	}, diagrams)
}

func (a *App) BaseConOptionAll(opt *conOption, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.BaseConOption(opt)
	}, diagrams)
}

func (a *App) ApployConOptionAll(opt *conOption, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.ApplyConOption(opt)
	}, diagrams)
}

func (a *App) Merge(diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		a.Cdor.nodes = append(a.Cdor.nodes, worker.getNodes()...)
		a.Cdor.connections = append(a.Cdor.connections, worker.getCons()...)
	}, diagrams)
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
	for _, worker := range a.workers {
		if data, err = worker.Gen(); err != nil {
			return err
		}
		file := filepath.Join(directory, fmt.Sprintf("%s.svg", className(worker)))
		if err = os.WriteFile(file, data, 0600); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) init() {
	a.workers = map[string]IWorker{}
	a.Cdor.init()
}

func (a *App) addWorker(worker IWorker) {
	name := className(worker)
	a.workerNames = append(a.workerNames, name)
	a.workers[name] = worker
}

func className(worker IWorker) string {
	return reflect.TypeOf(worker).Elem().Name()
}
