package xbin

import (
	"encoding/binary"

	"github.com/ndsky1003/xbin/options"
)

var DefaultOption = options.New().SetOrder(binary.BigEndian)

func InitOptions(deltas ...*options.Option) {
	DefaultOption.Merge(deltas...)
}
