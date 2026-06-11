package cli

import (
	"flag"
	"fmt"
)

// StringSliceFlag is a flag with type *StringSlice
type StringSliceFlag struct {
	Name        string
	Aliases     []string
	Usage       string
	EnvVars     []string
	FilePath    string
	Required    bool
	Hidden      bool
	TakesFile   bool
	Value       *StringSlice
	DefaultText string
	HasBeenSet  bool
}

// String returns a readable representation of this value (for displays)
func (f *StringSliceFlag) String() string {
	return FlagStringer(f)
}

// Names returns the names of the flag
func (f *StringSliceFlag) Names() []string {
	return flagNames(f.Name, f.Aliases)
}

// IsRequired returns whether the flag is required
func (f *StringSliceFlag) IsRequired() bool {
	return f.Required
}

// IsSet returns whether the flag has been set
func (f *StringSliceFlag) IsSet() bool {
	return f.HasBeenSet
}

// Apply populates the flag name and explanation and adds the flag to the flag set
func (f *StringSliceFlag) Apply(set *flag.FlagSet) error {
	if f.Value == nil {
		f.Value = &StringSlice{}
	}

	val := NewStringSlice(f.Value.Value()...)
	for _, name := range f.Names() {
		set.Var(val, name, f.Usage)
	}
	return nil
}

// StringSlice looks up the value of a local StringSliceFlag, returns
// nil if not found
func (c *Context) StringSlice(name string) []string {
	return c.lookupStringSlice(name)
}

func (c *Context) lookupStringSlice(name string) []string {
	if v, ok := c.Value(name).([]string); ok {
		return v
	}
	return nil
}