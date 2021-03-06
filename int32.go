// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable

import (
	"bytes"
	"encoding/json"
	"github.com/go-openapi/strfmt"
)

// Int32 represents a int32 that may be null or not
// present in json at all.
type Int32 struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid int32
	Value   int32
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int32) Ptr() *int32 {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int32) UnmarshalJSON(data []byte) error {
	i.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface for go-swagger generation.
func (i *Int32) Validate(formats strfmt.Registry) error {
	return nil
}

// Int32Slice represents a []int32 that may be null or not
// present in json at all.
type Int32Slice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []int32
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int32Slice) UnmarshalJSON(data []byte) error {
	i.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}
