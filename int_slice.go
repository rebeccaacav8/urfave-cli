package cli

import (
	"fmt"
	"strconv"
)

// IntSlice wraps a []int to satisfy flag.Value
type IntSlice struct {
	slice      []int
	hasBeenSet bool
}

// NewIntSlice creates a *IntSlice with default values
func NewIntSlice(defaults ...int) *IntSlice {
	return &IntSlice{slice: append([]int{}, defaults...)}
}

// Set parses the value and appends it to the slice
func (i *IntSlice) Set(value string) error {
	tmp, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if !i.hasBeenSet {
		i.slice = []int{tmp}
		i.hasBeenSet = true
	} else {
		i.slice = append(i.slice, tmp)
	}
	return nil
}

// Get returns the slice of ints
func (i *IntSlice) Get() interface{} {
	return i.slice
}

// String returns a readable representation of this value (for displays)
func (i *IntSlice) String() string {
	return fmt.Sprintf("%s", i.slice)
}

// Value returns the slice of ints
func (i *IntSlice) Value() []int {
	return i.slice
}