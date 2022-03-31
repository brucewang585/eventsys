package aflag

import (
	"sync/atomic"
)

type AFlag int32

func (i *AFlag) Set(flag bool) {
	if flag {
		atomic.StoreInt32((*int32)(i), 1)
	} else {
		atomic.StoreInt32((*int32)(i), 0)
	}
}

func (i *AFlag) Get() bool {
	n := int(atomic.LoadInt32((*int32)(i)))
	if n > 0 {
		return true
	} else {
		return false
	}
}
