// Copyright@daidai53 2023
package interfaces

import (
	"github.com/daidai53/freqdata_cache_demo/public"
)

type FreqDataItf interface {
	RecoverData() any
	UpdateData(any)
	DataType() public.FreqDataType
}
