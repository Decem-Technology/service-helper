package cipher

import "encoding/json"

// BinaryMarshaler binary marshaler
type BinaryMarshaler struct{}

// MarshalBinary marshal binary
func (e *BinaryMarshaler) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary marshal binary
func (e *BinaryMarshaler) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, e)
}
