package xbin

import (
	"encoding/binary"
	"io"
)

func Read(r io.Reader, data any) error {
	if r == nil || data == nil {
		return nil
	}
	switch v := data.(type) {
	case *string:
		return read_str(r, v)
	default:
		return binary.Read(r, big, data)
	}
}

func read_str(r io.Reader, s *string) error {
	return nil
}
