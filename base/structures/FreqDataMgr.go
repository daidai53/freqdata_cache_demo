// Copyright@daidai53 2023
package structures

import "sync"

type FreqDataMgr struct {
	userNodeMap freqUserHash
	freqLink    freqLink

	Mutex sync.RWMutex
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
}
