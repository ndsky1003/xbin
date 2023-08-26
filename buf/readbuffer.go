package buf

import (
	"encoding/binary"
	"math"

	"github.com/ndsky1003/xbin/options"
)

type ReadBuffer struct {
	normal_bytes        []byte
	normal_bytes_offset uint
	*BitReadBuffer
}

func NewReadBuffer(data []byte) *ReadBuffer {
	flag_bytes_length, n := binary.Uvarint(data)
	header_length := uint64(n)
	split_dot := header_length + flag_bytes_length
	flag_bytes := data[n:split_dot]
	return &ReadBuffer{
		BitReadBuffer: NewBitReadBuffer(flag_bytes),
		normal_bytes:  data[split_dot:],
	}
}

// not nil=>true,nil => false
func (this *ReadBuffer) ReadIsNotNil() (bool, error) {
	return this.BitReadBuffer.Read()
}

// int 不符合语义，长度不可能为负数，但是len这个内置函数的返回值是int，避免无效思想负担，直接这么用
func (this *ReadBuffer) ReadLength() int {
	return this.ReadInt()
}

func (this *ReadBuffer) read_bytes(length uint) []byte {
	datas := this.normal_bytes[this.normal_bytes_offset : this.normal_bytes_offset+length]
	this.normal_bytes_offset += length
	return datas
}

// bool ,int, uint
func ReadT[T Constraint](r *ReadBuffer, opt *options.Option) (T, error) {
	var v T
	var d any = v
	var err error
	switch d.(type) {
	case bool:
		d, err = r.BitReadBuffer.Read()
	case int:
		d = r.ReadInt()
	case uint:
		d = r.ReadUint()
	case string:
		length := r.ReadLength()
		datas := r.read_bytes(uint(length))
		d = string(datas)
	case int8:
		d = int8(r.read_bytes(1)[0])
	case uint8:
		d = r.read_bytes(1)[0]
	case int16:
		d = int16(opt.Order.Uint16(r.read_bytes(2)))
	case uint16:
		d = opt.Order.Uint16(r.read_bytes(2))
	case int32:
		d = int32(opt.Order.Uint32(r.read_bytes(4)))
	case uint32:
		d = opt.Order.Uint32(r.read_bytes(4))
	case int64:
		d = int64(opt.Order.Uint64(r.read_bytes(8)))
	case uint64:
		d = opt.Order.Uint64(r.read_bytes(8))
	case float32:
		d = math.Float32frombits(opt.Order.Uint32(r.read_bytes(4)))
	case float64:
		d = math.Float64frombits(opt.Order.Uint64(r.read_bytes(8)))
	}
	return d.(T), err
}

// T
func ReadPtrT[T Constraint](r *ReadBuffer, opt *options.Option) (*T, error) {
	b, err := r.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := ReadT[T](r, opt)
		return &b, err
	}
}

// []bool
func ReadSliceT[T Constraint](r *ReadBuffer, opt *options.Option) ([]T, error) {
	length := r.ReadLength()
	bs := make([]T, length)
	for i := range bs {
		b, err := ReadT[T](r, opt)
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]bool
func ReadPtrSliceT[T Constraint](r *ReadBuffer, opt *options.Option) (*[]T, error) {
	b, err := r.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := ReadSliceT[T](r, opt)
		return &b, err
	}
}

// []*bool
func ReadSlicePtrT[T Constraint](r *ReadBuffer, opt *options.Option) ([]*T, error) {
	length := r.ReadLength()
	bs := make([]*T, length)
	for i := range bs {
		b, err := ReadPtrT[T](r, opt)
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]*bool
func ReadPtrSlicePtrT[T Constraint](r *ReadBuffer, opt *options.Option) (*[]*T, error) {
	b, err := r.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := ReadSlicePtrT[T](r, opt)
		return &b, err
	}
}

// 10个字节，与定长还是有区别的
func (this *ReadBuffer) ReadInt() int {
	v, length := binary.Varint(this.normal_bytes[this.normal_bytes_offset:])
	this.normal_bytes_offset += uint(length)
	return int(v)
}

func (this *ReadBuffer) ReadUint() uint {
	v, length := binary.Uvarint(this.normal_bytes[this.normal_bytes_offset:])
	this.normal_bytes_offset += uint(length)
	return uint(v)
}
