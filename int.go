package optparse

import (
	"strconv"
)

// IntValue represents a signed integer value
type IntValue int

// IntValueAssertion defines an assertion if the Value is an IntValue
type IntValueAssertion interface {
	Value
	IsInt() bool
}

// NewIntValue creates a new IntValue
func NewIntValue(val int, p *int) *IntValue {
	*p = val
	return (*IntValue)(p)
}

// Set sets the value
func (i *IntValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)
	*i = IntValue(v)
	return err
}

// IsInt returns true if the Value is an IntValue
func (i *IntValue) IsInt() bool {
	return true
}

// IntVar defines a signed integer option with a pointer to its value
func (o *OptionParser) IntVar(p *int, long string, short rune, def int) {
	opt := &Option{long, short, NewIntValue(def, p)}
	o.Options = append(o.Options, opt)
}

// IntVar defines a signed integer option with a pointer to its value
func IntVar(p *int, long string, short rune, def int) {
	CommandLine.IntVar(p, long, short, def)
}

// Int defines a signed integer option and returns a pointer to its value
func (o *OptionParser) Int(long string, short rune, def int) (p *int) {
	p = new(int)
	o.IntVar(p, long, short, def)
	return
}

// Int defines a signed integer option and returns a pointer to its value
func Int(long string, short rune, def int) (p *int) {
	p = CommandLine.Int(long, short, def)
	return
}
