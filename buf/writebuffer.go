package buf

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/ndsky1003/xbin/options"
)

type WriteBuffer struct {
	bytes.Buffer
	BitWriteBuffer
}

func NewWriteBuffer() *WriteBuffer {
	return new(WriteBuffer)
}

func (this *WriteBuffer) Bytes() []byte {
	flag_bytes := this.BitWriteBuffer.Bytes()
	flag_bytes_length := len(flag_bytes)
	normal_bytes := this.Buffer.Bytes()

	t := [INT_DATA_MAX_LENGTH]byte{}
	n := binary.PutUvarint(t[:], uint64(flag_bytes_length))
	flag_bytes_length_bytes := t[:n]
	result_data := make([]byte, 0, len(flag_bytes_length_bytes)+flag_bytes_length+len(normal_bytes))
	result_data = append(result_data, flag_bytes_length_bytes...)
	result_data = append(result_data, flag_bytes...)
	result_data = append(result_data, normal_bytes...)
	return result_data
}

// not nil=>true,nil => false
func (this *WriteBuffer) WriteIsNotNil(b bool, opt *options.Option) {
	_ = WriteT(this, b, opt)
}

// int 不符合语义，长度不可能为负数，但是len这个内置函数的返回值是int，避免无效思想负担，直接这么用
func (this *WriteBuffer) WriteLength(l int, opt *options.Option) error {
	return WriteT(this, l, opt)
}

// bool,int,uint
func WriteT[T Constraint](w *WriteBuffer, v T, opt *options.Option) (err error) {
	var d any = v
	switch vv := d.(type) {
	case bool:
		w.BitWriteBuffer.WriteBool(vv)
	case int:
		t := [INT_DATA_MAX_LENGTH]byte{}
		n := binary.PutVarint(t[:], int64(vv))
		_, err = w.Write(t[:n])
	case uint:
		t := [INT_DATA_MAX_LENGTH]byte{}
		n := binary.PutUvarint(t[:], uint64(vv))
		_, err = w.Write(t[:n])
	case string:
		length := len(vv)
		if err = w.WriteLength(length, opt); err != nil {
			return
		}
		if vv != "" {
			_, err = w.WriteString(vv)
		}
	case int8:
		err = w.WriteByte(byte(vv))
	case uint8:
		err = w.WriteByte(vv)

	case int16:
		bs := make([]byte, 2)
		opt.Order.PutUint16(bs, uint16(vv))
		_, err = w.Write(bs)
	case uint16:
		bs := make([]byte, 2)
		opt.Order.PutUint16(bs, vv)
		_, err = w.Write(bs)
	case int32:
		bs := make([]byte, 4)
		opt.Order.PutUint32(bs, uint32(vv))
		_, err = w.Write(bs)
	case uint32:
		bs := make([]byte, 4)
		opt.Order.PutUint32(bs, vv)
		_, err = w.Write(bs)
	case int64:
		bs := make([]byte, 8)
		opt.Order.PutUint64(bs, uint64(vv))
		_, err = w.Write(bs)
	case uint64:
		bs := make([]byte, 8)
		opt.Order.PutUint64(bs, vv)
		_, err = w.Write(bs)
	case float32:
		bs := make([]byte, 4)
		opt.Order.PutUint32(bs, math.Float32bits(vv))
		_, err = w.Write(bs)
	case float64:
		bs := make([]byte, 8)
		opt.Order.PutUint64(bs, math.Float64bits(vv))
		_, err = w.Write(bs)
	}
	return
}

// *bool,*int,*uint
func WritePtrT[T Constraint](w *WriteBuffer, v *T, opt *options.Option) error {
	w.WriteIsNotNil(v != nil, opt)
	if v != nil {
		if err := WriteT(w, *v, opt); err != nil {
			return err
		}
	}
	return nil
}

// []bool,[]int,[]uint
func WriteSliceT[T Constraint](w *WriteBuffer, vs []T, opt *options.Option) error {
	length := len(vs)
	if err := w.WriteLength(length, opt); err != nil {
		return err
	}
	for _, b := range vs {
		_ = WriteT(w, b, opt)
	}
	return nil
}

// *[]bool,*[]int,*[]uint
func WritePtrSliceT[T Constraint](w *WriteBuffer, bs *[]T, opt *options.Option) error {
	w.WriteIsNotNil(bs != nil, opt)
	if bs != nil {
		if err := WriteSliceT(w, *bs, opt); err != nil {
			return err
		}
	}
	return nil
}

// []*bool,[]*int,[]*uint
func WriteSlicePtrT[T Constraint](w *WriteBuffer, bs []*T, opt *options.Option) error {
	length := len(bs)
	if err := w.WriteLength(length, opt); err != nil {
		return err
	}
	for _, b := range bs {
		_ = WritePtrT(w, b, opt)
	}
	return nil
}

// *[]*bool,*[]*int,*[]*uint
func WritePtrSlicePtrT[T Constraint](w *WriteBuffer, bs *[]*T, opt *options.Option) error {
	w.WriteIsNotNil(bs != nil, opt)
	if bs != nil {
		return WriteSlicePtrT(w, *bs, opt)
	}
	return nil
}
