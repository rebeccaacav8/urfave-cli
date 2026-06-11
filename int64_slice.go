package cli

import (
	"fmt"
	"strconv"
)

// Int64Slice wraps a []int64 to satisfy flag.Value
type Int64Slice struct {
	slice      []int64
	hasBeenSet bool
}

// NewInt64Slice creates a *Int64Slice with default values
func NewInt64Slice(defaults ...int64) *Int64Slice {
	return &Int64Slice{slice: append([]int64{}, defaults...)}
}

// Set parses the value and appends it to the slice
func (i *Int64Slice) Set(value string) error {
	tmp, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	if !i.hasBeenSet {
		i.slice = []int64{tmp}
		i.hasBeenSet = true
	} else {
		i.slice = append(i.slice, tmp)
	}
	return nil
}

// Get returns the slice of int64s
func (i *Int64Slice) Get() interface{} {
	return i.slice
}

// String returns a readable representation of this value (for displays)
func (i *Int64Slice) String() string {
	return fmt.Sprintf("%s", i.slice)
}

// Value returns the slice of int64s
func (i *Int64Slice) Value() []int64 {
	return i.slice
}