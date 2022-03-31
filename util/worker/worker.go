package worker

import (
	"context"
	"sync"
)

/*
worker 对象
*/

type WorkerFunc func()

type Worker interface {
	Start()
	Stop()
	Context() context.Context
}

type baseWorker struct {
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	//
	fns []WorkerFunc
}

func NewWorker(fns ...WorkerFunc) Worker {
	ctx, cancel := context.WithCancel(context.Background())

	o := &baseWorker{
		wg:     &sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
		fns:    fns,
	}
	return o
}

func (this *baseWorker) Start() {
	this.wg.Add(len(this.fns))
	for _, fn := range this.fns {
		this.run(fn)
	}
}

func (this *baseWorker) Stop() {
	this.cancel()
	this.wg.Wait()
}

func (this *baseWorker) Context() context.Context {
	return this.ctx
}

func (this *baseWorker) run(fn func()) {
	go func() {
		defer func() {
			recover()
			this.wg.Done()
		}()

		//执行
		fn()
	}()
}
