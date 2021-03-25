package main

import (
	"errors"
	"log"
)

type taskResult struct {
	result interface{}
	err    error
}

type task struct {
	render Render
	//args interface{}
	//fn func(v interface{}) taskResult

	chResult chan taskResult
}

type worker struct {
	typ RenderType
	p   *Pool

	chTask chan *task
}

func (w *worker) run() {
	go func() {
		defer func() {
			log.Printf("worker exit.")
		}()

		for task := range w.chTask {
			if task == nil {
				return
			}

			var res taskResult
			if err := task.render.Rend(); err != nil {
				res.err = err
			}

			task.chResult <- res
			close(task.chResult)
		}
	}()
}

type Pool struct {
	workersMap map[RenderType]chan *worker
}

func NewPool() *Pool {
	p := &Pool{}
	p.workersMap = make(map[RenderType]chan *worker, RenderTypeCount)

	cores := 4
	for i := 0; i < int(RenderTypeCount); i++ {
		typ := RenderType(i)
		p.workersMap[typ] = make(chan *worker, cores)

		// init N worker for each type
		for j := 0; j < cores; j++ {
			w := &worker{
				typ:    typ,
				p:      p,
				chTask: make(chan *task),
			}
			w.run()

			p.workersMap[typ] <- w
		}
	}

	return p
}

func (p *Pool) retrieveWorker(typ RenderType) *worker {
	chWorkers, ok := p.workersMap[typ]
	if !ok {
		return nil
	}

	w := <-chWorkers
	return w
}

func (p *Pool) Submit(render Render) (result taskResult, err error) {
	w := p.retrieveWorker(render.Type())
	if w == nil {
		err = errors.New("no valid worker in pool")
		return
	}

	task := &task{
		render:   render,
		chResult: make(chan taskResult, 1),
	}
	w.chTask <- task

	result = <-task.chResult // wait result

	p.workersMap[w.typ] <- w // put worker back
	return
}
