// Copyright@daidai53 2023
package structures

import (
	"github.com/daidai53/freqdata_cache_demo/interfaces"
	"sync"
	"time"
)

type FreqDataMgr struct {
	userNodeMap freqUserHash
	freqLink    freqLink
	interfaces.FreqDataItf

	Mutex sync.RWMutex
}

func NewFreqDataMgr(inf interfaces.FreqDataItf) *FreqDataMgr {
	var f = &FreqDataMgr{}
	f.Mutex.Lock()
	f.userNodeMap.Init()
	f.freqLink.Init()
	f.FreqDataItf = inf
	f.Mutex.Unlock()
	return f
}

func (f *FreqDataMgr) AddNode(newNode *FreqNode) {
	if f.userNodeMap == nil {
		//todo:add error log
		return
	}
	if !newNode.Valid() {
		return
	}
	f.Mutex.Lock()
	f.freqLink.addNode(newNode)
	f.userNodeMap[newNode.userId] = newNode
	f.Mutex.Unlock()
	if !GuardAlive() {
		guard(&f.freqLink)
	}
}

// ReadDBData
// 从缓存中读取高频数据，如果没有则从DB中恢复高频数据
func (f *FreqDataMgr) ReadDBData(userId interface{}) any {
	if f == nil || userId == nil {
		return nil
	}

	f.Mutex.RLock()
	dataPtr, ok := f.userNodeMap[userId]
	f.Mutex.RUnlock()
	if ok {
		return dataPtr.ReadData()
	}

	newData := f.RecoverData()
	newNode := &FreqNode{
		userId:       userId,
		dataType:     f.DataType(),
		data:         newData,
		period:       200 * time.Millisecond,
		lastScanTime: time.Duration(time.Now().UnixNano()),
		remain:       10,
	}
	f.AddNode(newNode)
}
