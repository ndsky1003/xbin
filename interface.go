package xbin

type Marshaler interface {
	MarshalXBIN() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalXBIN([]byte) error
}
