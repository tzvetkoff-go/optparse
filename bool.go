package optparse

import (
	"strconv"
)

// Boolean value
type BoolValue bool

type BoolValueAssertion interface {
	Value
	IsBool() bool
}

func NewBoolValue(val bool, p *bool) *BoolValue {
	*p = val
	return (*BoolValue)(p)
}

func (b *BoolValue) Set(val string) error {
	v, err := strconv.ParseBool(val)
	*b = BoolValue(v)
	return err
}

func (b *BoolValue) IsBool() bool {
	return true
}

// Defines a boolean option with a pointer to its value
func (o *OptionParser) BoolVar(p *bool, long string, short rune, def bool) {
	opt := &Option{long, short, NewBoolValue(def, p)}
	o.Options = append(o.Options, opt)
}

// Defines a boolean option with a pointer to its value
func BoolVar(p *bool, long string, short rune, def bool) {
	CommandLine.BoolVar(p, long, short, def)
}

// Defines a boolean option and returns a pointer to its value
func (o *OptionParser) Bool(long string, short rune, def bool) (p *bool) {
	p = new(bool)
	o.BoolVar(p, long, short, def)
	return
}

// Defines a boolean option and returns a pointer to its value
func Bool(long string, short rune, def bool) (p *bool) {
	p = CommandLine.Bool(long, short, def)
	return
}
