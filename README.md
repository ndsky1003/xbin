# xbin

专用 go 的序列化,非 proto 的大包大揽,想弄个这个缘故是 proto 不支持泛型

#### 写数据支持

```go
type (
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
```

#### 读数据支持

```go
type (
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
			*int | *[]int | *[]*int |
			*uint | *[]uint | *[]*uint |
			*string | *[]string | *[]*string
	}
)

```

#### 初版完成

##### 待支持自动生成序列化与反序列化
