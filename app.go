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

func (a *App) RangeDiagrams(action func(string, []byte, error) error, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		data, err := worker.Gen()
		if action(name, data, err) != nil {
			return
		}
	}, diagrams)
}

func (a *App) ApplyConfig(cfg *config, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.ApplyConfig(cfg)
	}, diagrams)
}

func (a *App) BaseConfig(cfg *config, diagrams ...string) {
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

func (a *App) Merge(diagrams ...string) *Cdor {
	res := Ctx()
	if len(diagrams) == 0 { // merge all diagrams
		res.ApplyConfig(a.config)
		res.ApplyOption(a.baseOption)
		res.ApplyConOption(a.baseConOption)
		res.nodes = append(res.nodes, a.nodes...)
		res.connections = append(res.connections, a.connections...)
	}
	a.rangeAction(func(name string, worker IWorker) {
		res.ApplyConfig(worker.getConfig())
		res.ApplyOption(worker.getBaseOption())
		res.ApplyConOption(worker.getBaseConOption())
		res.nodes = append(res.nodes, worker.getNodes()...)
		res.connections = append(res.connections, worker.getCons()...)
	}, diagrams)

	return res
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
	for name, worker := range a.workers {
		if data, err = worker.Gen(); err != nil {
			return err
		}
		file := filepath.Join(directory, fmt.Sprintf("%s.svg", name))
		if err = os.WriteFile(file, data, 0600); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) init(workers []IWorker) {
	a.workers = make(map[string]IWorker, len(workers))
	a.workerNames = make([]string, len(workers))
	for i, worker := range workers {
		name := reflect.TypeOf(worker).Elem().Name()
		a.workerNames[i] = name
		a.workers[name] = worker
	}
	a.Cdor.init()
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
