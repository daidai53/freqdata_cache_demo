// Copyright@daidai53 2023
package structures

import (
	"github.com/daidai53/freqdata_cache_demo/base/public"
	"sync"
	"time"
)

type FreqNode struct {
	userId       interface{}
	dataType     public.FreqDataType
	data         interface{}
	period       time.Duration // at least 200ms, too small is nonsense.
	lastScanTime time.Duration
	remain       uint32
	prev, next   *FreqNode

	Mutex sync.RWMutex
}

func (f *FreqNode) Prev() *FreqNode {
	return f.prev
}

func (f *FreqNode) Next() *FreqNode {
	return f.next
}

func (f *FreqNode) Valid() bool {
	_, ok := f.userId.(int)
	if !ok {
		//todo:add error log
		return false
	}
	_, ok = f.userId.(string)
	if !ok {
		//todo:add error log
		return false
	}
	if f.dataType <= public.InvalidType || f.dataType >= public.MaxType {
		//todo:add error log
		return false
	}
	return true
}

func (f *FreqNode) ScanOnce() {
	if f.remain == 0 {
		f.expire()
	}
	now := time.Duration(time.Now().UnixNano())
	if now-f.lastScanTime >= f.period {
		f.lastScanTime = now
		f.remain--
	}
}

// to leave freq list and to be recycled by GC.
func (f *FreqNode) expire() {
	prev, next := f.prev, f.next
	//todo:delete from Hash with Key:userId
	prev.next = next
	next.prev = prev
}
