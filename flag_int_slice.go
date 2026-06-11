package cli

import (
	"flag"
	"fmt"
)

// IntSliceFlag is a flag with type *IntSlice
type IntSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	Value       *IntSlice
	DefaultText string
	HasBeenSet  bool
}

// String returns a readable representation of this value (for displays)
func (f *IntSliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag
func (f *IntSliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required
func (f *IntSliceFlag) IsRequired() bool {
	return f.Required
}

// IsSet returns whether the flag has been set
func (f *IntSliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// Apply populates the flag name and explanation and adds the flag to the flag set
func (f *IntSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = &IntSlice{}
	}

	val := NewIntSlice(f.Value.Value()...)
	for _, name := range f.Names() {
		set.Var(val, name, f.Usage)
	}
	return nil
}

// IntSlice looks up the value of a local IntSliceFlag, returns
// nil if not found
func (c *Context) IntSlice(name string) []int {
	return c.lookupIntSlice(name)
}

func (c *Context) lookupIntSlice(name string) []int {
	if v, ok := c.Value(name).([]int); ok {
		return v
	}
	return nil
}