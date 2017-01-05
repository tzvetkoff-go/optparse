package optparse

import (
	"strconv"
)

// Floating point value
type FloatValue float64

type FloatValueAssertion interface {
	Value
	IsFloat() bool
}

func NewFloatValue(val float64, p *float64) *FloatValue {
	*p = val
	return (*FloatValue)(p)
}

func (f *FloatValue) Set(val string) error {
	v, err := strconv.ParseFloat(val, 64)
	*f = FloatValue(v)
	return err
}

func (f *FloatValue) IsFloat() bool {
	return true
}

// Defines a 64-bit float option with a pointer to its value
func (o *OptionParser) FloatVar(p *float64, long string, short rune, def float64) {
	opt := &Option{long, short, NewFloatValue(def, p)}
	o.Options = append(o.Options, opt)
}

// Defines a 64-bit float option with a pointer to its value
func FloatVar(p *float64, long string, short rune, def float64) {
	CommandLine.FloatVar(p, long, short, def)
}

// Defines a 64-bit float option and returns a pointer to its value
func (o *OptionParser) Float(long string, short rune, def float64) (p *float64) {
	p = new(float64)
	o.FloatVar(p, long, short, def)
	return
}

// Defines a 64-bit float option and returns a pointer to its value
func Float(long string, short rune, def float64) (p *float64) {
	p = CommandLine.Float(long, short, def)
	return
}
