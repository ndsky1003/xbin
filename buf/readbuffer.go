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

// encode slice length
// bool
func (this *ReadBuffer) ReadBool() (bool, error) {
	return this.BitReadBuffer.Read()
}

// *bool
func (this *ReadBuffer) ReadPtrBool() (*bool, error) {
	b, err := this.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := this.ReadBool()
		return &b, err
	}
}

// []bool
func (this *ReadBuffer) ReadSliceBool() ([]bool, error) {
	length := this.ReadLength()
	bs := make([]bool, length)
	for i := range bs {
		b, err := this.ReadBool()
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]bool
func (this *ReadBuffer) ReadPtrSliceBool() (*[]bool, error) {
	b, err := this.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := this.ReadSliceBool()
		return &b, err
	}
}

// []*bool
func (this *ReadBuffer) ReadSlicePtrBool() ([]*bool, error) {
	length := this.ReadLength()
	bs := make([]*bool, length)
	for i := range bs {
		b, err := this.ReadPtrBool()
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]*bool
func (this *ReadBuffer) ReadPtrSlicePtrBool() (*[]*bool, error) {
	b, err := this.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := this.ReadSlicePtrBool()
		return &b, err
	}
}

// 10个字节，与定长还是有区别的
func (this *ReadBuffer) ReadInt() int {
	v, length := binary.Varint(this.normal_bytes[this.normal_bytes_offset:])
	this.normal_bytes_offset += uint(length)
	return int(v)
}

// *int
func (this *ReadBuffer) ReadPtrInt() (*int, error) {
	b, err := this.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b := this.ReadInt()
		return &b, nil
	}
}

// []int
func (this *ReadBuffer) ReadSliceInt() ([]int, error) {
	length := this.ReadLength()
	bs := make([]int, length)
	for i := range bs {
		bs[i] = this.ReadInt()
	}
	return bs, nil
}

// *[]int
func (this *ReadBuffer) ReadPtrSliceInt() (*[]int, error) {
	b, err := this.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := this.ReadSliceInt()
		return &b, err
	}
}

// []*int
func (this *ReadBuffer) ReadSlicePtrInt() ([]*int, error) {
	length := this.ReadLength()
	bs := make([]*int, length)
	for i := range bs {
		b, err := this.ReadPtrInt()
		if err != nil {
			return nil, err
		}
		bs[i] = b
	}
	return bs, nil
}

// *[]*int
func (this *ReadBuffer) ReadPtrSlicePtrInt() (*[]*int, error) {
	b, err := this.ReadIsNotNil()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, nil
	} else {
		b, err := this.ReadSlicePtrInt()
		return &b, err
	}
}

func (this *ReadBuffer) ReadUint() uint {
	v, length := binary.Uvarint(this.normal_bytes[this.normal_bytes_offset:])
	this.normal_bytes_offset += uint(length)
	return uint(v)
}
