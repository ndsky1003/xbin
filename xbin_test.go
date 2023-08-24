package xbin

import (
	"testing"

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
	if err := Write(write_buf, []bool{true, false, true, false}); err != nil {
		t.Error(err)
	}
	True := true
	False := false
	if err := Write(write_buf, []*bool{&False, &True, nil, &False, &True}); err != nil {
		t.Error(err)
	}
	datas := write_buf.Bytes()
	t.Log(datas)

	read_buf := buf.NewReadBuffer(datas)
	if b, err := read_buf.ReadBool(); err != nil {
		t.Error(err)
	} else {
		t.Log(b)
	}

	if b, err := read_buf.ReadPtrBool(); err != nil {
		t.Error(err)
	} else {
		t.Log(b)
	}
}
