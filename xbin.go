package xbin

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ndsky1003/xbin/buf"
	"github.com/ndsky1003/xbin/options"
)

func WriteForm(w *buf.WriteBuffer, s Marshaler, opts ...*options.Option) error {
	data, err := s.MarshalXBIN()
	if err != nil {
		return err
	}
	return Write(w, data, opts...)
}

func Write[T buf.WConstraint](
	w *buf.WriteBuffer,
	data T,
	opts ...*options.Option,
) error {
	if w == nil {
		return errors.New("buffer is nil")
	}

	opt := options.New().Merge(DefaultOption).Merge(opts...)
	return write(w, data, opt)
}

func write(
	w *buf.WriteBuffer,
	data any,
	opt *options.Option,
) error {
	switch v := data.(type) {
	case bool:
		return buf.WriteT(w, v, opt)
	case *bool:
		return buf.WritePtrT(w, v, opt)
	case []bool:
		return buf.WriteSliceT(w, v, opt)
	case *[]bool:
		return buf.WritePtrSliceT(w, v, opt)
	case []*bool:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*bool:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case int:
		return buf.WriteT(w, v, opt)
	case *int:
		return buf.WritePtrT(w, v, opt)
	case []int:
		return buf.WriteSliceT(w, v, opt)
	case *[]int:
		return buf.WritePtrSliceT(w, v, opt)
	case []*int:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*int:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case uint:
		return buf.WriteT(w, v, opt)
	case *uint:
		return buf.WritePtrT(w, v, opt)
	case []uint:
		return buf.WriteSliceT(w, v, opt)
	case *[]uint:
		return buf.WritePtrSliceT(w, v, opt)
	case []*uint:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*uint:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case string:
		return buf.WriteT(w, v, opt)
	case *string:
		return buf.WritePtrT(w, v, opt)
	case []string:
		return buf.WriteSliceT(w, v, opt)
	case *[]string:
		return buf.WritePtrSliceT(w, v, opt)
	case []*string:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*string:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case int8:
		return buf.WriteT(w, v, opt)
	case *int8:
		return buf.WritePtrT(w, v, opt)
	case []int8:
		return buf.WriteSliceT(w, v, opt)
	case *[]int8:
		return buf.WritePtrSliceT(w, v, opt)
	case []*int8:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*int8:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case uint8:
		return buf.WriteT(w, v, opt)
	case *uint8:
		return buf.WritePtrT(w, v, opt)
	case []uint8:
		return buf.WriteSliceT(w, v, opt)
	case *[]uint8:
		return buf.WritePtrSliceT(w, v, opt)
	case []*uint8:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*uint8:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case int16:
		return buf.WriteT(w, v, opt)
	case *int16:
		return buf.WritePtrT(w, v, opt)
	case []int16:
		return buf.WriteSliceT(w, v, opt)
	case *[]int16:
		return buf.WritePtrSliceT(w, v, opt)
	case []*int16:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*int16:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case uint16:
		return buf.WriteT(w, v, opt)
	case *uint16:
		return buf.WritePtrT(w, v, opt)
	case []uint16:
		return buf.WriteSliceT(w, v, opt)
	case *[]uint16:
		return buf.WritePtrSliceT(w, v, opt)
	case []*uint16:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*uint16:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case int32:
		return buf.WriteT(w, v, opt)
	case *int32:
		return buf.WritePtrT(w, v, opt)
	case []int32:
		return buf.WriteSliceT(w, v, opt)
	case *[]int32:
		return buf.WritePtrSliceT(w, v, opt)
	case []*int32:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*int32:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case uint32:
		return buf.WriteT(w, v, opt)
	case *uint32:
		return buf.WritePtrT(w, v, opt)
	case []uint32:
		return buf.WriteSliceT(w, v, opt)
	case *[]uint32:
		return buf.WritePtrSliceT(w, v, opt)
	case []*uint32:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*uint32:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case int64:
		return buf.WriteT(w, v, opt)
	case *int64:
		return buf.WritePtrT(w, v, opt)
	case []int64:
		return buf.WriteSliceT(w, v, opt)
	case *[]int64:
		return buf.WritePtrSliceT(w, v, opt)
	case []*int64:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*int64:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case uint64:
		return buf.WriteT(w, v, opt)
	case *uint64:
		return buf.WritePtrT(w, v, opt)
	case []uint64:
		return buf.WriteSliceT(w, v, opt)
	case *[]uint64:
		return buf.WritePtrSliceT(w, v, opt)
	case []*uint64:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*uint64:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case float32:
		return buf.WriteT(w, v, opt)
	case *float32:
		return buf.WritePtrT(w, v, opt)
	case []float32:
		return buf.WriteSliceT(w, v, opt)
	case *[]float32:
		return buf.WritePtrSliceT(w, v, opt)
	case []*float32:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*float32:
		return buf.WritePtrSlicePtrT(w, v, opt)

	case float64:
		return buf.WriteT(w, v, opt)
	case *float64:
		return buf.WritePtrT(w, v, opt)
	case []float64:
		return buf.WriteSliceT(w, v, opt)
	case *[]float64:
		return buf.WritePtrSliceT(w, v, opt)
	case []*float64:
		return buf.WriteSlicePtrT(w, v, opt)
	case *[]*float64:
		return buf.WritePtrSlicePtrT(w, v, opt)
	default:
		return fmt.Errorf("write not support:%v", reflect.TypeOf(data).Name())
	}
}

/*
saveIsPtr 判定数据存入的时候是否是指针，用于决定是否进行指针的nil判断,这个字段一定保证正确，否则错位，调试还有可能成功
eg:存取*int，你无论读的时候无论传入什么值，都能正确的读取，因为这个nil判断在BitWriteBuffer,这里就会少读一个，就会导致错位
T 一定是一个指针类型
*/

func ReadForm(r *buf.ReadBuffer, s Unmarshaler, opts ...*options.Option) error {
	var bs []byte
	if err := Read(r, false, &bs, opts...); err != nil {
		return err
	}
	return s.UnmarshalXBIN(bs)
}

func Read[T buf.RConstraint](
	r *buf.ReadBuffer,
	saveIsPtr bool,
	data T,
	opts ...*options.Option,
) error {
	if r == nil {
		return errors.New("read buffer is nil")
	}
	var d any = data
	opt := options.New().Merge(DefaultOption).Merge(opts...)
	switch d.(type) {
	case *bool, *[]bool, *[]*bool:
		return read[bool](r, saveIsPtr, data, opt)

	case *int, *[]int, *[]*int:
		return read[int](r, saveIsPtr, data, opt)

	case *uint, *[]uint, *[]*uint:
		return read[uint](r, saveIsPtr, data, opt)

	case *string, *[]string, *[]*string:
		return read[string](r, saveIsPtr, data, opt)

	case *int8, *[]int8, *[]*int8:
		return read[int8](r, saveIsPtr, data, opt)

	case *uint8, *[]uint8, *[]*uint8:
		return read[uint8](r, saveIsPtr, data, opt)

	case *int16, *[]int16, *[]*int16:
		return read[int16](r, saveIsPtr, data, opt)

	case *uint16, *[]uint16, *[]*uint16:
		return read[uint16](r, saveIsPtr, data, opt)

	case *int32, *[]int32, *[]*int32:
		return read[int32](r, saveIsPtr, data, opt)

	case *uint32, *[]uint32, *[]*uint32:
		return read[uint32](r, saveIsPtr, data, opt)

	case *int64, *[]int64, *[]*int64:
		return read[int64](r, saveIsPtr, data, opt)

	case *uint64, *[]uint64, *[]*uint64:
		return read[uint64](r, saveIsPtr, data, opt)

	case *float32, *[]float32, *[]*float32:
		return read[float32](r, saveIsPtr, data, opt)

	case *float64, *[]float64, *[]*float64:
		return read[float64](r, saveIsPtr, data, opt)
	default:
		return fmt.Errorf("not support:%v", reflect.TypeOf(data).Name())
	}
}

func read[T buf.Constraint](
	r *buf.ReadBuffer,
	saveIsPtr bool,
	data any,
	opt *options.Option,
) error {
	switch v := data.(type) {
	case *T:
		if !saveIsPtr {
			if b, err := buf.ReadT[T](r, opt); err != nil {
				return err
			} else {
				if data != nil {
					*v = b
				}
			}
		} else {
			if b, err := buf.ReadPtrT[T](r, opt); err != nil {
				return err
			} else {
				if data != nil {
					if b != nil {
						*v = *b
					} else if opt.ClearOldValue != nil && *opt.ClearOldValue {
						var tmpV T
						*v = tmpV // 用零值抹掉原来的值
					}
				}
			}
		}
	case *[]T:
		if !saveIsPtr {
			if b, err := buf.ReadSliceT[T](r, opt); err != nil {
				return err
			} else {
				if data != nil {
					*v = b
				}
			}
		} else {
			if b, err := buf.ReadPtrSliceT[T](r, opt); err != nil {
				return err
			} else {
				if data != nil {
					if b != nil {
						*v = *b
					} else if opt.ClearOldValue != nil && *opt.ClearOldValue {
						*v = []T{}
					}
				}
			}
		}
	case *[]*T:
		if !saveIsPtr {
			if b, err := buf.ReadSlicePtrT[T](r, opt); err != nil {
				return err
			} else {
				if data != nil {
					*v = b
				}
			}
		} else {
			if b, err := buf.ReadPtrSlicePtrT[T](r, opt); err != nil {
				return err
			} else {
				if data != nil {
					if b != nil {
						*v = *b
					} else if opt.ClearOldValue != nil && *opt.ClearOldValue {
						*v = []*T{}
					}
				}
			}
		}
	default:
		return fmt.Errorf("not support:%v", reflect.TypeOf(data).Name())
	}
	return nil
}
