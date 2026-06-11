package cli

import (
	"flag"
	"fmt"
)

// Int64SliceFlag is a flag with type *Int64Slice
type Int64SliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	Value       *Int64Slice
	DefaultText string
	HasBeenSet  bool
}

// String returns a readable representation of this value (for displays)
func (f *Int64SliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag
func (f *Int64SliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required
func (f *Int64SliceFlag) IsRequired() bool {
	return f.Required
}

// IsSet returns whether the flag has been set
func (f *Int64SliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// Apply populates the flag name and explanation and adds the flag to the flag set
func (f *Int64SliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = &Int64Slice{}
	}

	val := NewInt64Slice(f.Value.Value()...)
	for _, name := range f.Names() {
		set.Var(val, name, f.Usage)
	}
	return nil
}

// Int64Slice looks up the value of a local Int64SliceFlag, returns
// nil if not found
func (c *Context) Int64Slice(name string) []int64 {
	return c.lookupInt64Slice(name)
}

func (c *Context) lookupInt64Slice(name string) []int64 {
	if v, ok := c.Value(name).([]int64); ok {
		return v
	}
	return nil
}