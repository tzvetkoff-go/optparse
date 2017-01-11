package optparse

import (
	"strconv"
)

// BoolValue represents a boolean value
type BoolValue bool

// BoolValueAssertion defines an assertion if the Value is a BoolValue
type BoolValueAssertion interface {
	Value
	IsBool() bool
}

// NewBoolValue creates a new BoolValue
func NewBoolValue(val bool, p *bool) *BoolValue {
	*p = val
	return (*BoolValue)(p)
}

// Set sets the value
func (b *BoolValue) Set(val string) error {
	v, err := strconv.ParseBool(val)
	*b = BoolValue(v)
	return err
}

// IsBool returns true if the Value is a BoolValue
func (b *BoolValue) IsBool() bool {
	return true
}

// BoolVar defines a boolean option with a pointer to its value
func (o *OptionParser) BoolVar(p *bool, long string, short rune, def bool) {
	opt := &Option{long, short, NewBoolValue(def, p)}
	o.Options = append(o.Options, opt)
}

// BoolVar defines a boolean option with a pointer to its value
func BoolVar(p *bool, long string, short rune, def bool) {
	CommandLine.BoolVar(p, long, short, def)
}

// Bool defines a boolean option and returns a pointer to its value
func (o *OptionParser) Bool(long string, short rune, def bool) (p *bool) {
	p = new(bool)
	o.BoolVar(p, long, short, def)
	return
}

// Bool defines a boolean option and returns a pointer to its value
func Bool(long string, short rune, def bool) (p *bool) {
	p = CommandLine.Bool(long, short, def)
	return
}
