package eventsys

import (
	"time"
	"github.com/brucewang585/eventsys/util/aflag"
)

var evtId aflag.AInt64

func init() {
	evtId = aflag.AInt64(0)
}

const (
	eventTypeProcessorReq = iota + 1
	eventTypeProcessorRsp
	eventTypeSourceReq
	eventTypeSourceRsp
)

type event struct {
	//基本信息
	Type  int
	Data  interface{}

	//扩展信息
	Uuid  int64
	Begin time.Time
}

func newEvent(etype int,data interface{}) *event {
	o := &event{
		Type: etype,
		Data: data,
		Uuid: evtId.Add(),
		Begin: time.Now(),
	}
	return o
}
