package aflag

import (
	"sync/atomic"
)

type ARef int32

func (i *ARef) Add() {
	atomic.AddInt32((*int32)(i), 1)
}

func (i *ARef) Release() {
	atomic.AddInt32((*int32)(i), -1)
}

func (i *ARef) Reset() {
	atomic.StoreInt32((*int32)(i), 0)
}

func (i *ARef) Set(n int) {
	atomic.StoreInt32((*int32)(i), int32(n))
}

func (i *ARef) Get() int {
	return int(atomic.LoadInt32((*int32)(i)))
}

func (i *ARef) IsEmpty() bool {
	n := int(atomic.LoadInt32((*int32)(i)))
	if n <= 0 {
		return true
	} else {
		return false
	}
}
