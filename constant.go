package xbin

const (
	NULL  byte = 0x0
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

// nil 用上面的NULL表示
type support_type interface {
	bool | *bool | []bool | []*bool | *[]*bool |
		int8 | *int8 | []int8 | []*int8 | *[]*int8 |
		uint8 | *uint8 | []uint8 | []*uint8 | *[]*uint8 |
		int16 | *int16 | []int16 | []*int16 | *[]*int16 |
		uint16 | *uint16 | []uint16 | []*uint16 | *[]*uint16 |
		int32 | *int32 | []int32 | []*int32 | *[]*int32 |
		uint32 | *uint32 | []uint32 | []*uint32 | *[]*uint32 |
		int64 | *int64 | []int64 | []*int64 | *[]*int64 |
		uint64 | *uint64 | []uint64 | []*uint64 | *[]*uint64 |
		float32 | *float32 | []float32 | []*float32 | *[]*float32 |
		float64 | *float64 | []float64 | []*float64 | *[]*float64 |
		// binary 支持
		int | *int | []int | []*int | *[]*int |
		uint | *uint | []uint | []*uint | *[]*uint |
		string | *string | []string | []*string | *[]*string
}
