package cdor

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type Mgr struct {
	Cdor
	cdors     map[string]*Cdor
	cdorNames []string
}

func (m *Mgr) RangeCdors(action func(string, *Cdor, error) error, diagrams ...string) {
	m.rangeAction(func(name string, c *Cdor) {
		if action(name, c, c.err) != nil {
			return
		}
	}, diagrams)
}

func (a *Mgr) ApplyConfig(cfg *config, diagrams ...string) {
	a.rangeAction(func(name string, c *Cdor) {
		c.cdor().ApplyConfig(cfg)
	}, diagrams)
}

func (a *Mgr) ApplyOption(opt *option, diagrams ...string) {
	a.rangeAction(func(name string, c *Cdor) {
		c.cdor().ApplyOption(opt)
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
	a.rangeAction(func(name string, c *Cdor) {
		res.ApplyConfig(c.cdor().config)
		res.ApplyOption(c.cdor().cdor().globalOption)
		res.nodes = append(res.nodes, c.cdor().nodes...)
		res.connections = append(res.connections, c.cdor().connections...)
	}, diagrams)

	return res
}

func (a *Mgr) SaveFiles(dir string, diagrams ...string) (err error) {
	if dir == "" {
		if dir, err = os.Getwd(); err != nil {
			return
		}
	}
	a.rangeAction(func(name string, c *Cdor) {
		var data []byte
		if data, err = c.cdor().Gen(); err != nil {
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
	a.cdors = make(map[string]*Cdor, len(workers))
	a.cdorNames = make([]string, len(workers))
	for i, worker := range workers {
		name := reflect.TypeOf(worker).Elem().Name()
		a.cdorNames[i] = name
		a.cdors[name] = worker.cdor()
	}
	a.Cdor.init()
}

func (a *Mgr) rangeAction(action func(string, *Cdor), diagrams []string) {
	if len(diagrams) == 0 {
		diagrams = a.cdorNames // for all cdors
	}
	for _, name := range diagrams {
		cdor, ok := a.cdors[name]
		if !ok {
			continue
		}
		action(name, cdor)
	}
}
