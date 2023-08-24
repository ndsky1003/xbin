package buf

import (
	"bytes"
	"encoding/binary"
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

	t := [10]byte{}
	n := binary.PutUvarint(t[:], uint64(flag_bytes_length))
	flag_bytes_length_bytes := t[:n]
	result_data := make([]byte, 0, len(flag_bytes_length_bytes)+flag_bytes_length+len(normal_bytes))
	result_data = append(result_data, flag_bytes_length_bytes...)
	result_data = append(result_data, flag_bytes...)
	result_data = append(result_data, normal_bytes...)
	return result_data
}

// not nil=>true,nil => false
func (this *WriteBuffer) WriteIsNotNil(b bool) {
	_ = this.WriteBool(b)
}

// int 不符合语义，长度不可能为负数，但是len这个内置函数的返回值是int，避免无效思想负担，直接这么用
func (this *WriteBuffer) WriteLength(l int) error {
	return this.WriteInt(l)
}

// encode slice length

// bool
func (this *WriteBuffer) WriteBool(b bool) error {
	this.BitWriteBuffer.WriteBool(b)
	return nil
}

// *bool
func (this *WriteBuffer) WritePtrBool(b *bool) error {
	this.WriteIsNotNil(b != nil)
	if b != nil {
		if err := this.WriteBool(*b); err != nil {
			return err
		}
	}
	return nil
}

// []bool
func (this *WriteBuffer) WriteSliceBool(bs []bool) error {
	length := len(bs)
	if err := this.WriteLength(length); err != nil {
		return err
	}
	for _, b := range bs {
		_ = this.WriteBool(b)
	}
	return nil
}

// *[]bool
func (this *WriteBuffer) WritePtrSliceBool(bs *[]bool) error {
	this.WriteIsNotNil(bs != nil)
	if bs != nil {
		if err := this.WriteSliceBool(*bs); err != nil {
			return err
		}
	}
	return nil
}

// []*bool
func (this *WriteBuffer) WriteSlicePtrBool(bs []*bool) error {
	length := len(bs)
	if err := this.WriteLength(length); err != nil {
		return err
	}
	for _, b := range bs {
		_ = this.WritePtrBool(b)
	}
	return nil
}

// *[]*bool
func (this *WriteBuffer) WritePtrSlicePtrBool(bs *[]*bool) error {
	this.WriteIsNotNil(bs != nil)
	if bs != nil {
		return this.WriteSlicePtrBool(*bs)
	}
	return nil
}

// 10个字节，与定长还是有区别的
func (this *WriteBuffer) WriteInt(i int) (err error) {
	t := [10]byte{}
	n := binary.PutVarint(t[:], int64(i))
	_, err = this.Write(t[:n])
	return
}

func (this *WriteBuffer) WriteUint(i uint) (err error) {
	t := [10]byte{}
	n := binary.PutUvarint(t[:], uint64(i))
	_, err = this.Write(t[:n])
	return
}

// func write_strs(w *buf.WriteBuffer, s []string) (err error) {
// 	length := len(s)
// 	if err = encode_int(w, length); err != nil {
// 		return
// 	}
// 	for _, str := range s {
// 		if err = write_str(w, str); err != nil {
// 			return
// 		}
// 	}
// 	return
// }
//
// func write_strs_ptr(w *buf.WriteBuffer, s []*string) (err error) {
// 	length := len(s)
// 	if err = encode_int(w, length); err != nil {
// 		return
// 	}
// 	for _, str := range s {
// 		var s string
// 		if str != nil {
// 			s = *str
// 		}
// 		if err = write_str(w, s); err != nil {
// 			return
// 		}
// 	}
// 	return
// }
//
// func write_str(w *buf.WriteBuffer, s string) (err error) {
// 	length := len(s)
// 	if err = encode_int(w, length); err != nil {
// 		return
// 	}
// 	if s != "" {
// 		_, err = w.WriteString(s)
// 	}
// 	return
// }
