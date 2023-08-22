package options

import "encoding/binary"

// 保证么个字段都有值，至少有个默认值
type Option struct {
	Order binary.ByteOrder
}

func New() *Option {
	return new(Option)
}

func (this *Option) SetOrder(order binary.ByteOrder) *Option {
	if this == nil {
		return nil
	}
	this.Order = order
	return this
}

func (this *Option) merge(delta *Option) {
	if this == nil || delta == nil {
		return
	}
	if delta.Order != nil {
		this.Order = delta.Order
	}
}

func (this *Option) Merge(deltas ...*Option) *Option {
	if this == nil {
		return nil
	}

	for _, delta := range deltas {
		this.merge(delta)
	}
	return this
}
