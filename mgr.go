package cdor

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type Mgr struct {
	Cdor
	workers     map[string]IWorker
	workerNames []string
}

func (a *Mgr) RangeDiagrams(action func(string, []byte, error) error, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		data, err := worker.Gen()
		if action(name, data, err) != nil {
			return
		}
	}, diagrams)
}

func (a *Mgr) ApplyConfig(cfg *config, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.ApplyConfig(cfg)
	}, diagrams)
}

func (a *Mgr) ApplyOptionAll(opt *option, diagrams ...string) {
	a.rangeAction(func(name string, worker IWorker) {
		worker.ApplyOption(opt)
	}, diagrams)
}

func (a *Mgr) Merge(diagrams ...string) *Cdor {
	res := Ctx()
	if len(diagrams) == 0 { // merge all diagrams
		res.ApplyConfig(a.config)
		res.ApplyOption(a.globalOption)
		res.nodes = append(res.nodes, a.nodes...)
		res.connections = append(res.connections, a.connections...)
	}
	a.rangeAction(func(name string, worker IWorker) {
		res.ApplyConfig(worker.getConfig())
		res.ApplyOption(worker.getBaseOption())
		res.nodes = append(res.nodes, worker.getNodes()...)
		res.connections = append(res.connections, worker.getCons()...)
	}, diagrams)

	return res
}

func (a *Mgr) SaveFiles(dir string, diagrams ...string) (err error) {
	if dir == "" {
		if dir, err = os.Getwd(); err != nil {
			return
		}
	}
	a.rangeAction(func(name string, worker IWorker) {
		var data []byte
		if data, err = worker.Gen(); err != nil {
			return
		}
		file := filepath.Join(dir, fmt.Sprintf("%s.svg", name))
		if err = os.WriteFile(file, data, 0600); err != nil {
			return
		}
	}, diagrams)

	return nil
}

func (a *Mgr) init(workers []IWorker) {
	a.workers = make(map[string]IWorker, len(workers))
	a.workerNames = make([]string, len(workers))
	for i, worker := range workers {
		name := reflect.TypeOf(worker).Elem().Name()
		a.workerNames[i] = name
		a.workers[name] = worker
	}
	a.Cdor.init()
}

func (a *Mgr) rangeAction(action func(string, IWorker), diagrams []string) {
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
