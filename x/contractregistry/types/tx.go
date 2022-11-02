package types

import (
	"encoding/json"
	"errors"
)

type RawSigningData []byte

func (r RawSigningData) MarshalJSON() ([]byte, error) {
	return json.RawMessage(r).MarshalJSON()
}

func (r *RawSigningData) UnmarshalJSON(b []byte) error {
	if r == nil {
		return errors.New("unmarshalJSON on nil pointer")
	}
	*r = append((*r)[0:0], b...)
	return nil
}

func (r *RawSigningData) ValidateBasic() error {
	if r == nil {
		return ErrEmpty
	}
	if !json.Valid(*r) {
		return ErrInvalid
	}
	return nil
}

// Bytes returns raw bytes type
func (r RawSigningData) Bytes() []byte {
	return r
}
