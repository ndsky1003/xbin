package buf

import (
	"encoding/binary"
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

// bool ,int, uint
func ReadT[T Constraint](r *ReadBuffer) (T, error) {
	var v T
	var d any = v
	var err error
	switch d.(type) {
	case bool:
		d, err = r.BitReadBuffer.Read()
		return d.(T), err
	case int:
		d = r.ReadInt()
		return d.(T), nil
	case uint:
		d = r.ReadUint()
		return d.(T), nil
	case string:
		length := r.ReadLength()
		datas := r.normal_bytes[r.normal_bytes_offset : r.normal_bytes_offset+uint(length)]
		d = string(datas)
		r.normal_bytes_offset += uint(length)
		return d.(T), nil
	case int8:
		currentbyte := r.normal_bytes[r.normal_bytes_offset]
		r.normal_bytes_offset += 1
		currentint8 := int8(currentbyte)
		d = currentint8
		return d.(T), nil
	case uint8:
		d = r.normal_bytes[r.normal_bytes_offset]
		r.normal_bytes_offset += 1
		return d.(T), nil
		// TODO

	}
	return v, nil
}

// T
func ReadPtrT[T Constraint](r *ReadBuffer) (*T, error) {
	b, err := r.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := ReadT[T](r)
		return &b, err
	}
}

// []bool
func ReadSliceT[T Constraint](r *ReadBuffer) ([]T, error) {
	length := r.ReadLength()
	bs := make([]T, length)
	for i := range bs {
		b, err := ReadT[T](r)
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]bool
func ReadPtrSliceT[T Constraint](r *ReadBuffer) (*[]T, error) {
	b, err := r.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := ReadSliceT[T](r)
		return &b, err
	}
}

// []*bool
func ReadSlicePtrT[T Constraint](r *ReadBuffer) ([]*T, error) {
	length := r.ReadLength()
	bs := make([]*T, length)
	for i := range bs {
		b, err := ReadPtrT[T](r)
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]*bool
func ReadPtrSlicePtrT[T Constraint](r *ReadBuffer) (*[]*T, error) {
	b, err := r.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := ReadSlicePtrT[T](r)
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
