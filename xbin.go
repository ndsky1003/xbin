package xbin

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/ndsky1003/xbin/buf"
	"github.com/ndsky1003/xbin/options"
)

func Write[T write_type](
	w *buf.WriteBuffer,
	data T,
	opts ...*options.Option,
) error {
	if w == nil {
		return errors.New("buffer is nil")
	}
	var d any = data
	opt := options.New().Merge(DefaultOption).Merge(opts...)
	switch v := d.(type) {
	case bool:
		return w.WriteBool(v)
	case *bool:
		return w.WritePtrBool(v)
	case []bool:
		return w.WriteSliceBool(v)
	case *[]bool:
		return w.WritePtrSliceBool(v)
	case []*bool:
		return w.WriteSlicePtrBool(v)
	case *[]*bool:
		return w.WritePtrSlicePtrBool(v)
	default:
		rv := reflect.Indirect(reflect.ValueOf(data))
		switch rv.Kind() {
		case reflect.Array:
		case reflect.Map:

		}

		return binary.Write(w, opt.Order, data)
	}
}

/*
saveIsPtr 判定数据存入的时候是否是指针，用于决定是否进行指针的nil判断
T 一定是一个指针类型
*/

func Read[T read_type](r *buf.ReadBuffer, saveIsPtr bool, data T, opts ...*options.Option) error {
	if r == nil {
		return errors.New("read buffer is nil")
	}
	var d any = data
	opt := options.New().Merge(DefaultOption).Merge(opts...)
	switch v := d.(type) {
	case *bool:
		if !saveIsPtr {
			if b, err := r.ReadBool(); err != nil {
				return err
			} else {
				if data != nil {
					*v = b
				}
			}
		} else {
			if b, err := r.ReadPtrBool(); err != nil {
				return err
			} else {
				if data != nil {
					if b != nil {
						*v = *b
					} else if opt.ClearOldValue != nil && *opt.ClearOldValue {
						*v = false // 用零值抹掉原来的值
					}
				}
			}
		}
	case *[]bool:
		if !saveIsPtr {
			if b, err := r.ReadSliceBool(); err != nil {
				return err
			} else {
				if data != nil {
					*v = b
				}
			}
		} else {
			if b, err := r.ReadPtrSliceBool(); err != nil {
				return err
			} else {
				if data != nil {
					if b != nil {
						*v = *b
					} else if opt.ClearOldValue != nil && *opt.ClearOldValue {
						*v = []bool{}
					}
				}
			}
		}
	case *[]*bool:
		if !saveIsPtr {
			if b, err := r.ReadSlicePtrBool(); err != nil {
				return err
			} else {
				if data != nil {
					*v = b
				}
			}
		} else {
			if b, err := r.ReadPtrSlicePtrBool(); err != nil {
				return err
			} else {
				if data != nil {
					if b != nil {
						*v = *b
					} else if opt.ClearOldValue != nil && *opt.ClearOldValue {
						*v = []*bool{}
					}
				}
			}
		}
	default:
		fmt.Println(opt)
		return nil
	}
	return nil
}
