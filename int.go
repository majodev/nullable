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
	_ runtime.Validatable        = &Int{}
	_ runtime.ContextValidatable = &Int{}
	_ runtime.Validatable        = &IntSlice{}
	_ runtime.ContextValidatable = &IntSlice{}
)

// Int represents a int that may be null or not
// present in json at all.
type Int struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid int
	Value   int
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int) Ptr() *int {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int) UnmarshalJSON(data []byte) error {
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

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *Int) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *Int) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// IntSlice represents a []int that may be null or not
// present in json at all.
type IntSlice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []int
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *IntSlice) UnmarshalJSON(data []byte) error {
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

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *IntSlice) Ptr() *[]int {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// Validate implements runtime.Validateable interface from github.com/go-openapi/runtime
func (i *IntSlice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate implements runtime.ContextValidatable from github.com/go-openapi/runtime
func (i *IntSlice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
