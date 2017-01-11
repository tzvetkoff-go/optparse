package optparse

import (
	"strconv"
)

// BoolListValue represents a list of boolean values
type BoolListValue []bool

// BoolListValueAssertion defines an assertion if the Value is a BoolListValue
type BoolListValueAssertion interface {
	Value
	IsBoolList() bool
}

// NewBoolListValue creates a new BoolListValue
func NewBoolListValue(val []bool, p *[]bool) *BoolListValue {
	*p = val
	return (*BoolListValue)(p)
}

// Set appends a value
func (b *BoolListValue) Set(val string) error {
	v, err := strconv.ParseBool(val)
	*b = append(*b, v)
	return err
}

// IsBoolList returns true if the Value is a BoolListValue
func (b *BoolListValue) IsBoolList() bool {
	return true
}

// BoolListVar defines a boolean list option with a pointer to its value
func (o *OptionParser) BoolListVar(p *[]bool, long string, short rune) {
	opt := &Option{long, short, NewBoolListValue([]bool{}, p)}
	o.Options = append(o.Options, opt)
}

// BoolListVar defines a boolean list option with a pointer to its value
func BoolListVar(p *[]bool, long string, short rune) {
	CommandLine.BoolListVar(p, long, short)
}

// BoolList defines a boolean list option and returns a pointer to its value
func (o *OptionParser) BoolList(long string, short rune) (p *[]bool) {
	p = &[]bool{}
	o.BoolListVar(p, long, short)
	return
}

// BoolList defines a boolean list option and returns a pointer to its value
func BoolList(long string, short rune) (p *[]bool) {
	p = CommandLine.BoolList(long, short)
	return
}
