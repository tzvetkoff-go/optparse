package optparse

import (
	"strconv"
)

// FloatValue represents a floating point value
type FloatValue float64

// FloatValueAssertion defines an assertion if the Value is a FloatValue
type FloatValueAssertion interface {
	Value
	IsFloat() bool
}

// NewFloatValue creates a new FloatValue
func NewFloatValue(val float64, p *float64) *FloatValue {
	*p = val
	return (*FloatValue)(p)
}

// Set sets the value
func (f *FloatValue) Set(val string) error {
	v, err := strconv.ParseFloat(val, 64)
	*f = FloatValue(v)
	return err
}

// IsFloat returns true if the Value is a FloatValue
func (f *FloatValue) IsFloat() bool {
	return true
}

// FloatVar defines a 64-bit float option with a pointer to its value
func (o *OptionParser) FloatVar(p *float64, long string, short rune, def float64) {
	opt := &Option{long, short, NewFloatValue(def, p)}
	o.Options = append(o.Options, opt)
}

// FloatVar defines a 64-bit float option with a pointer to its value
func FloatVar(p *float64, long string, short rune, def float64) {
	CommandLine.FloatVar(p, long, short, def)
}

// Float defines a 64-bit float option and returns a pointer to its value
func (o *OptionParser) Float(long string, short rune, def float64) (p *float64) {
	p = new(float64)
	o.FloatVar(p, long, short, def)
	return
}

// Float defines a 64-bit float option and returns a pointer to its value
func Float(long string, short rune, def float64) (p *float64) {
	p = CommandLine.Float(long, short, def)
	return
}
