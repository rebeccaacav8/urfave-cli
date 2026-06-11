package cli

import (
	"flag"
	"fmt"
)

// Float64SliceFlag is a flag with type *Float64Slice
type Float64SliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	Value       *Float64Slice
	DefaultText string
	HasBeenSet  bool
}

// String returns a readable representation of this value (for displays)
func (f *Float64SliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag
func (f *Float64SliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required
func (f *Float64SliceFlag) IsRequired() bool {
	return f.Required
}

// IsSet returns whether the flag has been set
func (f *Float64SliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// Apply populates the flag name and explanation and adds the flag to the flag set
func (f *Float64SliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = &Float64Slice{}
	}

	val := NewFloat64Slice(f.Value.Value()...)
	for _, name := range f.Names() {
		set.Var(val, name, f.Usage)
	}
	return nil
}

// Float64Slice looks up the value of a local Float64SliceFlag, returns
// nil if not found
func (c *Context) Float64Slice(name string) []float64 {
	return c.lookupFloat64Slice(name)
}

func (c *Context) lookupFloat64Slice(name string) []float64 {
	if v, ok := c.Value(name).([]float64); ok {
		return v
	}
	return nil
}