package xbin

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/ndsky1003/xbin/options"
)

/*
support byte *byte []byte, not support []*byte ,*[]*byte
byte 可以代表 int,uint,string
*/
func Write[T support_type](w *Buffer, data T, opts ...*options.Option) error {
	if w == nil {
		return errors.New("buffer is nil")
	}
	var d any = data
	opt := options.New().Merge(DefaultOption).Merge(opts...)
	switch v := d.(type) {
	case bool:
		return encode_bool(w, v)
	case *bool:
		return encode_ptr_bool(w, v)
	case []bool:
		return encode_slice_bool(w, v)

	case string:
		return write_str(w, v)
	case *string:
		return write_str(w, *v)
	case []string:
		return write_strs(w, v)
	case []*string:
		return write_strs_ptr(w, v)
	case int:
		return encode_int(w, v)
	case *int:
		return encode_int(w, *v)
	case uint:
		return encode_uint(w, v)
	case *uint:
		return encode_uint(w, *v)
	default:
		rv := reflect.Indirect(reflect.ValueOf(data))
		switch rv.Kind() {
		case reflect.Array:
		case reflect.Map:

		}

		return binary.Write(w, opt.Order, data)
	}
}

// NULL
func encode_nil(w *Buffer) error {
	return w.WriteByte(NULL)
}

// int 不符合语义，长度不可能为负数，但是len这个内置函数的返回值是int，避免无效思想负担，直接这么用
func encode_length(w *Buffer, l int) error {
	return encode_int(w, l)
}

// encode slice length

// bool
func encode_bool(w *Buffer, b bool) error {
	return w.WriteByte(bool2byte(b))
}

// *bool
func encode_ptr_bool(w *Buffer, b *bool) error {
	if b == nil {
		return encode_nil(w)
	}
	return encode_bool(w, *b)
}

// []bool
func encode_slice_bool(w *Buffer, bs []bool) error {
	length := len(bs)
	if err := encode_length(w, length); err != nil {
		return err
	}
	byte_slice := make([]byte, length)
	for i, v := range bs {
		byte_slice[i] = bool2byte(v)
	}
	n, err := w.Write(byte_slice)
	if err != nil {
		return err
	}
	if n != length {
		return fmt.Errorf("expect:%d,but use %d", length, n)
	}
	return nil
}

// []*bool
func encode_slice_ptr_bool(w *Buffer, bs []*bool) error {
	length := len(bs)
	if err := encode_length(w, length); err != nil {
		return err
	}
	byte_slice := make([]byte, length)
	for i, v := range bs {
		if v == nil {
			byte_slice[i] = NULL
		} else {
			byte_slice[i] = bool2byte(*v)
		}
	}
	n, err := w.Write(byte_slice)
	if err != nil {
		return err
	}
	if n != length {
		return fmt.Errorf("expect:%d,but use %d", length, n)
	}
	return nil
}

// *[]*bool
func encode_ptr_slice_ptr_bool(w *Buffer, bs *[]*bool) error {
	if bs == nil {
		return encode_nil(w)
	}
	return encode_slice_ptr_bool(w, *bs)
}

// 10个字节，与定长还是有区别的
func encode_int(w *Buffer, i int) (err error) {
	t := [10]byte{}
	n := binary.PutVarint(t[:], int64(i))
	_, err = w.Write(t[:n])
	return
}

func encode_uint(w *Buffer, i uint) (err error) {
	t := [10]byte{}
	n := binary.PutUvarint(t[:], uint64(i))
	_, err = w.Write(t[:n])
	return
}

func write_strs(w *Buffer, s []string) (err error) {
	length := len(s)
	if err = encode_int(w, length); err != nil {
		return
	}
	for _, str := range s {
		if err = write_str(w, str); err != nil {
			return
		}
	}
	return
}

func write_strs_ptr(w *Buffer, s []*string) (err error) {
	length := len(s)
	if err = encode_int(w, length); err != nil {
		return
	}
	for _, str := range s {
		var s string
		if str != nil {
			s = *str
		}
		if err = write_str(w, s); err != nil {
			return
		}
	}
	return
}

func write_str(w *Buffer, s string) (err error) {
	length := len(s)
	if err = encode_int(w, length); err != nil {
		return
	}
	if s != "" {
		_, err = w.WriteString(s)
	}
	return
}
