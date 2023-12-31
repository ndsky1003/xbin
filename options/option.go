package options

import "encoding/binary"

// 保证么个字段都有值，至少有个默认值
type Option struct {
	Order         binary.ByteOrder
	ClearOldValue *bool // 意义就是写入的是个nil，读的时候是否将nil当做一个零值覆盖到占坑指针的那个非默认值
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

func (this *Option) SetClearOldValue(b bool) *Option {
	if this == nil {
		return nil
	}
	this.ClearOldValue = &b
	return this
}

func (this *Option) merge(delta *Option) {
	if this == nil || delta == nil {
		return
	}
	if delta.Order != nil {
		this.Order = delta.Order
	}
	if delta.ClearOldValue != nil {
		this.ClearOldValue = delta.ClearOldValue
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
