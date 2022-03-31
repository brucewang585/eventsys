package aflag

import (
	"sync/atomic"
)

type AUSize uint64

func (i *AUSize) Add(d uint64) {
	atomic.AddUint64((*uint64)(i), d)
}

func (i *AUSize) Get() uint64 {
	return atomic.LoadUint64((*uint64)(i))
}

func (i *AUSize) Set(d uint64) {
	atomic.StoreUint64((*uint64)(i), d)
}

type AISize int64

func (i *AISize) Add(d int64) {
	atomic.AddInt64((*int64)(i), d)
}

func (i *AISize) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

func (i *AISize) Set(d int64) {
	atomic.StoreInt64((*int64)(i), d)
}
