package xbin

import (
	"bytes"
	"encoding/binary"

	"github.com/ndsky1003/xbin/options"
)

func Write(w *bytes.Buffer, data any, opts ...*options.Option) error {
	if w == nil || data == nil {
		return nil
	}
	opt := options.New().Merge(DefaultOption).Merge(opts...)
	switch v := data.(type) {
	case string:
		return write_str(w, v)
	case *string:
		return write_str(w, *v)
	case []string:
		return write_strs(w, v)
	case []*string:
		return write_strs_ptr(w, v)
	case int:
		return write_int(w, v)
	case *int:
		return write_int(w, *v)
	case uint:
		return write_uint(w, v)
	case *uint:
		return write_uint(w, *v)
	default:
		return binary.Write(w, opt.Order, data)
	}
}

func write_strs(w *bytes.Buffer, s []string) (err error) {
	length := len(s)
	if err = write_int(w, length); err != nil {
		return
	}
	for _, str := range s {
		if err = write_str(w, str); err != nil {
			return
		}
	}
	return
}

func write_strs_ptr(w *bytes.Buffer, s []*string) (err error) {
	length := len(s)
	if err = write_int(w, length); err != nil {
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

func write_str(w *bytes.Buffer, s string) (err error) {
	length := len(s)
	if err = write_int(w, length); err != nil {
		return
	}
	if s != "" {
		_, err = w.WriteString(s)
	}
	return
}

func write_int(w *bytes.Buffer, i int) (err error) {
	t := [4]byte{}
	n := binary.PutVarint(t[:], int64(i))
	_, err = w.Write(t[:n])
	return
}

func write_uint(w *bytes.Buffer, i uint) (err error) {
	t := [4]byte{}
	n := binary.PutUvarint(t[:], uint64(i))
	_, err = w.Write(t[:n])
	return
}
