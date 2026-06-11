package cli

import (
	"fmt"
)

// StringSlice wraps a []string to satisfy flag.Value
type StringSlice struct {
	slice      []string
	hasBeenSet bool
}

// NewStringSlice creates a *StringSlice with default values
func NewStringSlice(defaults ...string) *StringSlice {
	return &StringSlice{slice: append([]string{}, defaults...)}
}

// Set appends the value to the slice or overrides the default if it's the first time Set is called
func (s *StringSlice) Set(value string) error {
	if !s.hasBeenSet {
		s.slice = []string{value}
		s.hasBeenSet = true
	} else {
		s.slice = append(s.slice, value)
	}
	return nil
}

// Get returns the slice of strings
func (s *StringSlice) Get() interface{} {
	return s.slice
}

// String returns a readable representation of this value (for displays)
func (s *StringSlice) String() string {
	return fmt.Sprintf("%s", s.slice)
}

// Value returns the slice of strings
func (s *StringSlice) Value() []string {
	return s.slice
}