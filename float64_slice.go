package cli

import (
	"fmt"
	"strconv"
)

// Float64Slice wraps a []float64 to satisfy flag.Value
type Float64Slice struct {
	slice      []float64
	hasBeenSet bool
}

// NewFloat64Slice creates a *Float64Slice with default values
func NewFloat64Slice(defaults ...float64) *Float64Slice {
	return &Float64Slice{slice: append([]float64{}, defaults...)}
}

// Set parses the value and appends it to the slice
func (f *Float64Slice) Set(value string) error {
	tmp, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	if !f.hasBeenSet {
		f.slice = []float64{tmp}
		f.hasBeenSet = true
	} else {
		f.slice = append(f.slice, tmp)
	}
	return nil
}

// Get returns the slice of float64s
func (f *Float64Slice) Get() interface{} {
	return f.slice
}

// String returns a readable representation of this value (for displays)
func (f *Float64Slice) String() string {
	return fmt.Sprintf("%s", f.slice)
}

// Value returns the slice of float64s
func (f *Float64Slice) Value() []float64 {
	return f.slice
}