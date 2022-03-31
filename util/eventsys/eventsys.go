package eventsys

import (
	"context"
	"github.com/brucewang585/eventsys/util/worker"
	"time"
)

type EventSys interface {
	Stop()

	Context() context.Context
}

type baseEventSys struct {
	w worker.Worker

	//
	cap      int       //同时执行任务的能力
	interval int       //调度Source的间隔时间
	src      Source    //任务源
	proc     Processor //任务执行器

	//
	evtSrc  chan *event //source 接受请求消息
	evtProc chan *event //processor 接受请求消息
	evtResp chan *event //调度 接受返回消息
}

func NewEventSys(cap int, interval int, s Source, p Processor) EventSys {
	o := &baseEventSys{
		cap:      cap,
		interval: interval,
		src:      s,
		proc:     p,

		//
		evtSrc:  make(chan *event, 1000),
		evtProc: make(chan *event, 1000),
		evtResp: make(chan *event, 1000),
	}

	//根据cap启动服务
	fs := []worker.WorkerFunc{o.do_schedule, o.do_source}
	for i := 0; i < cap; i++ {
		fs = append(fs, o.do_processor)
	}
	o.w = worker.NewWorker(fs...)
	o.w.Start()
	return o
}

func (e *baseEventSys) Stop() {
	e.w.Stop()

	close(e.evtSrc)
	close(e.evtProc)
	close(e.evtResp)
}

func (e *baseEventSys) Context() context.Context {
	return e.w.Context()
}

//调度控制,调度source,调度Processor
func (e *baseEventSys) do_schedule() {
	var (
		ev *event
		ok bool

		proc_remain = e.cap
		tasks       []interface{}
		task        interface{}
	)

	tm := time.NewTimer(time.Millisecond * 100)
	defer tm.Stop()

	for {
		select {
		case <-e.Context().Done():
			return

		case <-tm.C:
			//通知source
			e.evtSrc <- newEvent(eventTypeSourceReq, proc_remain)

		case ev, ok = <-e.evtResp:
			if !ok || ev == nil {
				return
			}
			switch ev.Type {
			case eventTypeProcessorRsp:
				//代表有空闲processor
				proc_remain += 1

			case eventTypeSourceRsp:
				//通知source collect
				tm.Reset(time.Duration(e.interval) * time.Second)

				//通知processor exec
				tasks, ok = ev.Data.([]interface{})
				if len(tasks) > 0 {
					if len(tasks) > proc_remain {
						tasks = tasks[:proc_remain]
						proc_remain = 0
					} else {
						proc_remain -= len(tasks)
					}

					//投递任务
					for _, task = range tasks {
						e.evtProc <- newEvent(eventTypeProcessorReq, task)
					}
				}
			}
		}
	}
}

func (e *baseEventSys) do_source() {
	var (
		ev *event
		ok bool

		//
		num int
		evs []interface{}
	)

	for {
		select {
		case <-e.Context().Done():
			return

		case ev, ok = <-e.evtSrc:
			if !ok || ev == nil {
				return
			}
			evs = nil

			num, ok = ev.Data.(int)
			if num == 0 {
				//如果没有空闲processor,就不需要调用source就返回
				e.evtResp <- newEvent(eventTypeSourceRsp, evs)
			} else {
				evs, _ = e.src.Collect(e.Context(), num)
				e.evtResp <- newEvent(eventTypeSourceRsp, evs)
			}

		}
	}
}

func (e *baseEventSys) do_processor() {
	var (
		ev *event
		ok bool
	)

	for {
		select {
		case <-e.Context().Done():
			return

		case ev, ok = <-e.evtProc:
			if !ok || ev == nil {
				return
			}
			e.proc.Exec(e.Context(), ev.Data)
			e.evtResp <- newEvent(eventTypeProcessorRsp, nil)
		}
	}
}
