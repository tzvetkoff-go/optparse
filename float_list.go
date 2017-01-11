package optparse

import (
	"strconv"
)

// FloatListValue represents a list of floating point values
type FloatListValue []float64

// FloatListValueAssertion defines an assertion if the Value is a FloatListValue
type FloatListValueAssertion interface {
	Value
	IsFloatList() bool
}

// NewFloatListValue creates a new FloatListValue
func NewFloatListValue(val []float64, p *[]float64) *FloatListValue {
	*p = val
	return (*FloatListValue)(p)
}

// Set appends a value
func (f *FloatListValue) Set(val string) error {
	v, err := strconv.ParseFloat(val, 64)
	*f = append(*f, float64(v))
	return err
}

// IsFloatList returns true if the Value is a FloatListValue
func (f *FloatListValue) IsFloatList() bool {
	return true
}

// FloatListVar defines a 64-bit float list option with a pointer to its value
func (o *OptionParser) FloatListVar(p *[]float64, long string, short rune) {
	opt := &Option{long, short, NewFloatListValue([]float64{}, p)}
	o.Options = append(o.Options, opt)
}

// FloatListVar defines a 64-bit float list option with a pointer to its value
func FloatListVar(p *[]float64, long string, short rune) {
	CommandLine.FloatListVar(p, long, short)
}

// FloatList defines a 64-bit float list option and returns a pointer to its value
func (o *OptionParser) FloatList(long string, short rune) (p *[]float64) {
	p = &[]float64{}
	o.FloatListVar(p, long, short)
	return
}

// FloatList defines a 64-bit float list option and returns a pointer to its value
func FloatList(long string, short rune) (p *[]float64) {
	p = CommandLine.FloatList(long, short)
	return
}
