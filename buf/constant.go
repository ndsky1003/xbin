package buf

const INT_DATA_MAX_LENGTH = 10

// 就是所有基本类型
type (
	Constraint                     = SupportPrimitiveTypeConstraint
	SupportPrimitiveTypeConstraint interface {
		bool | int | uint | string |
			// 固定长度
			int8 |
			uint8 |
			int16 |
			uint16 |
			int32 |
			uint32 |
			int64 |
			uint64 |
			float32 |
			float64
	}
)

// 下面是基本类型的变体
type (
	WConstraint         = WriteTypeConstraint
	WriteTypeConstraint interface {
		bool | *bool | []bool | *[]bool | []*bool | *[]*bool |
			int8 | *int8 | []int8 | *[]int8 | []*int8 | *[]*int8 |
			uint8 | *uint8 | []uint8 | *[]uint8 | []*uint8 | *[]*uint8 |
			int16 | *int16 | []int16 | *[]int16 | []*int16 | *[]*int16 |
			uint16 | *uint16 | []uint16 | *[]uint16 | []*uint16 | *[]*uint16 |
			int32 | *int32 | []int32 | *[]int32 | []*int32 | *[]*int32 |
			uint32 | *uint32 | []uint32 | *[]uint32 | []*uint32 | *[]*uint32 |
			int64 | *int64 | []int64 | *[]int64 | []*int64 | *[]*int64 |
			uint64 | *uint64 | []uint64 | *[]uint64 | []*uint64 | *[]*uint64 |
			float32 | *float32 | []float32 | *[]float32 | []*float32 | *[]*float32 |
			float64 | *float64 | []float64 | *[]float64 | []*float64 | *[]*float64 |
			// 以上 binary 支持
			int | *int | []int | *[]int | []*int | *[]*int |
			uint | *uint | []uint | *[]uint | []*uint | *[]*uint |
			string | *string | []string | *[]string | []*string | *[]*string
	}
)

type (
	RConstraint        = ReadTypeConstraint
	ReadTypeConstraint interface {
		*bool | *[]bool | *[]*bool |
			*int8 | *[]int8 | *[]*int8 |
			*uint8 | *[]uint8 | *[]*uint8 |
			*int16 | *[]int16 | *[]*int16 |
			*uint16 | *[]uint16 | *[]*uint16 |
			*int32 | *[]int32 | *[]*int32 |
			*uint32 | *[]uint32 | *[]*uint32 |
			*int64 | *[]int64 | *[]*int64 |
			*uint64 | *[]uint64 | *[]*uint64 |
			*float32 | *[]float32 | *[]*float32 |
			*float64 | *[]float64 | *[]*float64 |
			// 以上 binary 支持
			*int | *[]int | *[]*int |
			*uint | *[]uint | *[]*uint |
			*string | *[]string | *[]*string
	}
)
