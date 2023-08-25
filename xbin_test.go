package xbin

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ndsky1003/xbin/buf"
	"github.com/ndsky1003/xbin/options"
)

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
	t.Log(datas)

	read_buf := buf.NewReadBuffer(datas)
	var rb bool
	if err := Read(read_buf, false, &rb); err != nil {
		t.Error(err)
	} else {
		t.Log(rb)
	}
	rb = true
	if err := Read(read_buf, true, &rb, options.New().SetClearOldValue(false)); err != nil {
		t.Error(err)
	} else {
		t.Log(rb)
	}
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
