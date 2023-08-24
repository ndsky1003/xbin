package xbin

const (
	True  byte = 0b00000001
	False byte = 0b11111111
)

func bool2byte(b bool) (byte byte) {
	if b {
		byte = True
	} else {
		byte = False
	}
	return
}

type write_type interface {
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

type read_type interface {
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
