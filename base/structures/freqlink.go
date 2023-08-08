// Copyright@daidai53 2023
package structures

type freqLink struct {
	head FreqNode
	tail FreqNode
}

func (f *freqLink) EmptyFreqLink() bool {
	if f.head.Prev() == nil &&
		f.tail.Prev() == nil {
		return true
	}
	return false
}

// RangeAllNodes
// Range all freq nodes in freq list with WLock. User cannot query data temporarily.
func (f *freqLink) RangeAllNodes(fun func(node *FreqNode)) {
	if f.EmptyFreqLink() {
		return
	}
	tmp := f.head.Next()
	for tmp.Next() != nil {
		tmp.Prev().Mutex.Lock()
		tmp.Mutex.Lock()
		tmp.Next().Mutex.Lock()
		fun(tmp)
		tmp.Prev().Mutex.Unlock()
		tmp.Mutex.Unlock()
		tmp.Next().Mutex.Unlock()
		tmp = tmp.Next()
	}
}

func (f *freqLink) addNode(newNode *FreqNode) {
	f.tail.Mutex.Lock()
	f.tail.prev.Mutex.Lock()
	newNode.prev = f.tail.prev
	f.tail.prev.next = newNode
	newNode.next = &f.tail
	f.tail.prev = newNode
	if newNode.period < 4*DefaultGuardPeriod {
		newNode.period = 4 * DefaultGuardPeriod
	}
	f.tail.Mutex.Unlock()
	newNode.prev.Mutex.Unlock()
}
