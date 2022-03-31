package aflag

import (
	"sync/atomic"
)

type AInt64 int64

func (i *AInt64) Add() int64 {
	return atomic.AddInt64((*int64)(i), 1)
}

func (i *AInt64) Set(a int64) {
	atomic.StoreInt64((*int64)(i), a)
}

func (i *AInt64) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

type AInt int32

func (i *AInt) Add() int {
	return int(atomic.AddInt32((*int32)(i), 1))
}

func (i *AInt) Set(a int) {
	atomic.StoreInt32((*int32)(i), int32(a))
}

func (i *AInt) Get() int {
	return int(atomic.LoadInt32((*int32)(i)))
}
