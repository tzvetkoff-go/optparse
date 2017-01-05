package optparse

import (
	"strconv"
)

// List of floating point value
type FloatListValue []float64

type FloatListValueAssertion interface {
	Value
	IsFloatList() bool
}

func NewFloatListValue(val []float64, p *[]float64) *FloatListValue {
	*p = val
	return (*FloatListValue)(p)
}

func (f *FloatListValue) Set(val string) error {
	v, err := strconv.ParseFloat(val, 64)
	*f = append(*f, float64(v))
	return err
}

func (f *FloatListValue) IsFloatList() bool {
	return true
}

// Defines a 64-bit float list option with a pointer to its value
func (o *OptionParser) FloatListVar(p *[]float64, long string, short rune) {
	opt := &Option{long, short, NewFloatListValue([]float64{}, p)}
	o.Options = append(o.Options, opt)
}

// Defines a 64-bit float list option with a pointer to its value
func FloatListVar(p *[]float64, long string, short rune) {
	CommandLine.FloatListVar(p, long, short)
}

// Defines a 64-bit float list option and returns a pointer to its value
func (o *OptionParser) FloatList(long string, short rune) (p *[]float64) {
	p = &[]float64{}
	o.FloatListVar(p, long, short)
	return
}

// Defines a 64-bit float list option and returns a pointer to its value
func FloatList(long string, short rune) (p *[]float64) {
	p = CommandLine.FloatList(long, short)
	return
}
