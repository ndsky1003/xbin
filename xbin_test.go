package xbin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndsky1003/xbin/buf"
	"github.com/ndsky1003/xbin/options"
)

// bool
func TestWriteReadBool(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	if err := Write(write_buf, true); err != nil {
		t.Error(err)
	}
	var b *bool
	if err := Write(write_buf, b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb bool
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, true)
	rb = true
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, false)
}

func TestWriteReadSliceBool(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writeslicebools := []bool{true, false, true, false}
	if err := Write(write_buf, writeslicebools); err != nil {
		t.Error(err)
	}
	True := true
	False := false
	ptrbools := []bool{true, false, false, false, True, False, True, True, false, false}
	if err := Write(write_buf, &ptrbools); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var bools []bool
	if err := Read(read_buf, false, &bools); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeslicebools, bools)
	var boolsptr []bool
	if err := Read(read_buf, true, &boolsptr); err != nil {
		t.Error(err)
	}
	assert.Equal(t, boolsptr, ptrbools)
}

func TestWriteReadSlicePtrBool(t *testing.T) {
	True1 := true
	False1 := false
	True := &True1
	False := &False1
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrBools := []*bool{True, False, True, False}
	if err := Write(write_buf, writeSlicePtrBools); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrBools := []*bool{
		True,
		False,
		False,
		False,
		True,
		False,
		True,
		True,
		False,
		False,
	}
	if err := Write(write_buf, &writePtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrBools []*bool
	if err := Read(read_buf, false, &readSlicePtrBools); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrBools, readSlicePtrBools)
	var readPtrSlicePtrBools []*bool
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writePtrSlicePtrBools, readPtrSlicePtrBools)
}

// int
func TestWriteReadInt(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	w := 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b int = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb int
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 131
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, 18)
}

func TestWriteReadSliceInt(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []int{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []int{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, i)
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	t.Log(len(datas))
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []int
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []int
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrInt(t *testing.T) {
	One := 1
	Two := 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*int{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*int{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*int
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*int
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// uint
func TestWriteReadUint(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w uint = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b uint = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb uint
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	var rb1 uint = 131
	if err := Read(read_buf, true, &rb1); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb1, b)
}

func TestWriteReadSliceUint(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []uint{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []uint{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, uint(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []uint
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []uint
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrUint(t *testing.T) {
	var One uint = 1
	var Two uint = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*uint{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*uint{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*uint
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*uint
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// string
func TestWriteReadString(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w string = "I am a gool boy"
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	// var w2 string = "I am not bad girl"
	var www2 *string = nil
	if err := Write(write_buf, www2); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rs1 string
	if err := Read(read_buf, false, &rs1); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rs1, w)

	var rs2 string = "ddd"
	if err := Read(read_buf, true, &rs2, options.New().SetClearOldValue(true)); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rs2, "")
}

func TestWriteReadSliceString(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []string{"one", "two", "three", "lllllllllllllllllllllllllllllllllllllll"}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []string{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, fmt.Sprintf("---------%v", i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readslicestrings []string
	if err := Read(read_buf, false, &readslicestrings); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readslicestrings)
	var ptrreadslicestrings []string
	if err := Read(read_buf, true, &ptrreadslicestrings); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadslicestrings)
}

func TestWriteReadSlicePtrString(t *testing.T) {
	var One string = "oooooo"
	var Two string = "gggppppppppppppppp"
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtr := []*string{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtr); err != nil {
		t.Error(err)
	}
	writePtrSlicePtr := []*string{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtr); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtr []*string
	if err := Read(read_buf, false, &readSlicePtr); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtr, readSlicePtr)
	var readPtrSlicePtr []*string
	if err := Read(read_buf, true, &readPtrSlicePtr); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writePtrSlicePtr, readPtrSlicePtr)
}

// int8
func TestWriteReadInt8(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w int8 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b int8 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb int8
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	t.Log(rb)
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, int8(18))
}

func TestWriteReadSliceInt8(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []int8{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []int8{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, int8(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	t.Log(len(datas))
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []int8
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []int8
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrInt8(t *testing.T) {
	var One int8 = 1
	var Two int8 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*int8{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*int8{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*int8
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*int8
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// uint8
func TestWriteReadUint8(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w byte = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b byte = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb byte
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, uint8(18))
}

func TestWriteReadSliceUint8(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []uint8{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []uint8{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, uint8(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	t.Log(len(datas))
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []uint8
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []uint8
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrUint8(t *testing.T) {
	var One uint8 = 1
	var Two uint8 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*uint8{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*uint8{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*uint8
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*uint8
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// int16
func TestWriteReadInt16(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w int16 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b int16 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb int16
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, int16(18))
}

func TestWriteReadSliceInt16(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []int16{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []int16{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, int16(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []int16
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []int16
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrInt16(t *testing.T) {
	var One int16 = 1
	var Two int16 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*int16{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*int16{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*int16
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*int16
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// uint16
func TestWriteReadUint16(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w uint16 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b uint16 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb uint16
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, uint16(18))
}

func TestWriteReadSliceUint16(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []uint16{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []uint16{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, uint16(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []uint16
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []uint16
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrUInt16(t *testing.T) {
	var One uint16 = 1
	var Two uint16 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*uint16{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*uint16{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*uint16
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*uint16
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// int32
func TestWriteReadInt32(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w int32 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b int32 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb int32
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, int32(18))
}

func TestWriteReadSliceInt32(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []int32{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []int32{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, int32(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []int32
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []int32
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrInt32(t *testing.T) {
	var One int32 = 1
	var Two int32 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*int32{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*int32{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*int32
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*int32
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// uint32
func TestWriteReadUint32(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w uint32 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b uint32 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb uint32
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, uint32(18))
}

func TestWriteReadSliceUint32(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []uint32{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []uint32{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, uint32(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []uint32
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []uint32
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrUInt32(t *testing.T) {
	var One uint32 = 1
	var Two uint32 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*uint32{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*uint32{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*uint32
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*uint32
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// int64
func TestWriteReadInt64(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w int64 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b int64 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb int64
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, int64(18))
}

func TestWriteReadSliceInt64(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []int64{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []int64{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, int64(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []int64
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []int64
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrInt64(t *testing.T) {
	var One int64 = 1
	var Two int64 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*int64{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*int64{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*int64
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*int64
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// uint64
func TestWriteReadUint64(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w uint64 = 12
	if err := Write(write_buf, w); err != nil {
		t.Error(err)
	}

	var b uint64 = 18
	if err := Write(write_buf, &b); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var rb uint64
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, w)

	rb = 11
	if err := Read(read_buf, true, &rb); err != nil {
		t.Error(err)
	}
	assert.Equal(t, rb, uint64(18))
}

func TestWriteReadSliceUint64(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []uint64{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []uint64{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, uint64(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []uint64
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []uint64
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrUInt64(t *testing.T) {
	var One uint64 = 1
	var Two uint64 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*uint64{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*uint64{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*uint64
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*uint64
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// float32
func TestWriteReadFloat32(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w1 float32 = 12.1233456
	t.Log("w1:", w1)
	if err := Write(write_buf, w1); err != nil {
		t.Error(err)
	}

	var w2 float32 = 12.12334567
	fmt.Println("w2:", w2)
	if err := Write(write_buf, &w2); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var r1 float32
	if err := Read(read_buf, false, &r1); err != nil {
		t.Error(err)
	}
	assert.Equal(t, r1, w1)

	var r2 float32
	if err := Read(read_buf, true, &r2); err != nil {
		t.Error(err)
	}
	t.Log(r2, w2)
	assert.Equal(t, r2, w2)
}

func TestWriteReadSliceFloat32(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []float32{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []float32{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, float32(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []float32
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []float32
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrFloat32(t *testing.T) {
	var One float32 = 1
	var Two float32 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*float32{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*float32{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*float32
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*float32
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}

// float64
func TestWriteReadfloat64(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	var w1 float64 = 12.12334562890932566
	t.Log("w1:", w1)
	if err := Write(write_buf, w1); err != nil {
		t.Error(err)
	}

	var w2 float64 = 12.12334567
	t.Log("w2:", w2)
	if err := Write(write_buf, &w2); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var r1 float64
	if err := Read(read_buf, false, &r1); err != nil {
		t.Error(err)
	}
	t.Log("r1:", r1)
	assert.Equal(t, r1, w1)

	var r2 float64
	if err := Read(read_buf, true, &r2); err != nil {
		t.Error(err)
	}
	t.Log("r2:", r2)
	assert.Equal(t, r2, w2)
}

func TestWriteReadSlicefloat64(t *testing.T) {
	write_buf := buf.NewWriteBuffer()
	writesliceints := []float64{1, 2, 3, 4}
	if err := Write(write_buf, writesliceints); err != nil {
		t.Error(err)
	}
	ptrwritesliceints := []float64{}
	for i := 0; i < 10; i++ {
		ptrwritesliceints = append(ptrwritesliceints, float64(i))
	}
	if err := Write(write_buf, &ptrwritesliceints); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readsliceints []float64
	if err := Read(read_buf, false, &readsliceints); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writesliceints, readsliceints)
	var ptrreadsliceints []float64
	if err := Read(read_buf, true, &ptrreadsliceints); err != nil {
		t.Error(err)
	}

	assert.Equal(t, ptrwritesliceints, ptrreadsliceints)
}

func TestWriteReadSlicePtrfloat64(t *testing.T) {
	var One float64 = 1
	var Two float64 = 2
	POne := &One
	PTwo := &Two
	write_buf := buf.NewWriteBuffer()
	writeSlicePtrInts := []*float64{POne, nil, PTwo, POne}
	if err := Write(write_buf, writeSlicePtrInts); err != nil {
		t.Error(err)
	}
	writePtrSlicePtrInts := []*float64{
		POne,
		PTwo,
		nil,
		nil,
		POne,
		PTwo,
		POne,
		PTwo,
		POne,
		nil,
	}
	if err := Write(write_buf, &writePtrSlicePtrInts); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	read_buf := buf.NewReadBuffer(datas)
	var readSlicePtrInts []*float64
	if err := Read(read_buf, false, &readSlicePtrInts); err != nil {
		t.Error(err)
	}
	assert.Equal(t, writeSlicePtrInts, readSlicePtrInts)
	var readPtrSlicePtrBools []*float64
	if err := Read(read_buf, true, &readPtrSlicePtrBools); err != nil {
		t.Error(err)
	}
	for _, v := range readPtrSlicePtrBools {
		if v != nil {
			t.Log(*v)
		} else {
			t.Log(v)
		}
	}
	assert.Equal(t, writePtrSlicePtrInts, readPtrSlicePtrBools)
}
