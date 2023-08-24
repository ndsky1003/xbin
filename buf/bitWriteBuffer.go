/*
作用，存储一个bool值会浪费7个位
存储是否是指针需要一个信号
综上会涉及到很多标志信号，综合考虑定义了这个bit_buffer
0 false,1 true
*/
package buf

type BitWriteBuffer struct {
	bytes    []byte
	byte     *byte
	bitIndex uint
}

func NewBitWriteBuffer() *BitWriteBuffer {
	return &BitWriteBuffer{
		bytes: make([]byte, 0, 2),
	}
}

func (this *BitWriteBuffer) write_flag(f bool) {
	index := this.bitIndex / 8
	bit_index := this.bitIndex % 8
	if bit_index == 0 {
		var b byte
		this.bytes = append(this.bytes, b)
		this.byte = &this.bytes[index]
	}
	if f {
		*this.byte |= 1 << bit_index
	} else {
		*this.byte = (*this.byte) &^ (1 << bit_index)
	}
	this.bitIndex++
}

// true => 1
// false => 0
func (this *BitWriteBuffer) WriteBool(b bool) {
	this.write_flag(b)
}

func (this *BitWriteBuffer) Bytes() []byte {
	return this.bytes
}
