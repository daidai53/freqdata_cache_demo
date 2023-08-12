// Copyright@daidai53 2023
package structures

type freqUserHash map[interface{}]*FreqNode

func (f *freqUserHash) Init() {
	if f == nil {
		*f = make(map[interface{}]*FreqNode, 10)
	}
}
