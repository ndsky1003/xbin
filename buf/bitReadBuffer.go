package buf

import "fmt"

type BitReadBuffer struct {
	bytes    []byte
	bitIndex uint
}

func NewBitReadBuffer(bytes []byte) *BitReadBuffer {
	return &BitReadBuffer{
		bytes: bytes,
	}
}

func (this *BitReadBuffer) Read() (bool, error) {
	length := len(this.bytes)
	index := this.bitIndex / 8
	bit_index := this.bitIndex % 8
	if index >= uint(length) {
		return false, fmt.Errorf("bytes length:%d,out of index:%d", length, index)
	}
	byte := this.bytes[index]
	byte >>= bit_index
	byte &= 1
	this.bitIndex++
	return byte == 1, nil
}
