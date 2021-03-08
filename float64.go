// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// Compile time validation that our types implement the expected interfaces
var (
	_ runtime.Validatable        = &Float64{}
	_ runtime.ContextValidatable = &Float64{}
	_ runtime.Validatable        = &Float64Slice{}
	_ runtime.ContextValidatable = &Float64Slice{}
)

// Float64 represents a float64 that may be null or not
// present in json at all.
type Float64 struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid float64
	Value   float64
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (f *Float64) Ptr() *float64 {
	if f.Present && f.Valid {
		return &f.Value
	}

	return nil
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float64) UnmarshalJSON(data []byte) error {
	f.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	f.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (f *Float64) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (f *Float64) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Float64Slice represents a []float64 that may be null or not
// present in json at all.
type Float64Slice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []float64
}

// UnmarshalJSON implements json.Marshaler interface.
func (f *Float64Slice) UnmarshalJSON(data []byte) error {
	f.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &f.Value); err != nil {
		return err
	}

	f.Valid = true
	return nil
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (f *Float64Slice) Ptr() *[]float64 {
	if f.Present && f.Valid {
		return &f.Value
	}

	return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (f *Float64Slice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (f *Float64Slice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
