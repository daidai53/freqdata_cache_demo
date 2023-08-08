// Copyright@daidai53 2023
package structures

import (
	"time"
)

var guardAlive bool

const DefaultGuardPeriod = 50 * time.Millisecond

var curGuardPeriod = DefaultGuardPeriod

func Alive() bool {
	return guardAlive
}

func GuardFreqList(f *freqLink) {
	if Alive() {
		return
	}
	go guard(f)
}

func guard(f *freqLink) {
	//todo:add info log
	guardAlive = true
	for {
		if f.EmptyFreqLink() {
			break
		}
		time.Sleep(curGuardPeriod)
		f.RangeAllNodes(func(node *FreqNode) {
			node.ScanOnce()
		})
	}
	//todo:add info log
	guardAlive = false
}
