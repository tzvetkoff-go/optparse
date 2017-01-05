package optparse

import (
	"strconv"
)

// List of boolean values
type BoolListValue []bool

type BoolListValueAssertion interface {
	Value
	IsBoolList() bool
}

func NewBoolListValue(val []bool, p *[]bool) *BoolListValue {
	*p = val
	return (*BoolListValue)(p)
}

func (b *BoolListValue) Set(val string) error {
	v, err := strconv.ParseBool(val)
	*b = append(*b, v)
	return err
}

func (b *BoolListValue) IsBoolList() bool {
	return true
}

// Defines a boolean list option with a pointer to its value
func (o *OptionParser) BoolListVar(p *[]bool, long string, short rune) {
	opt := &Option{long, short, NewBoolListValue([]bool{}, p)}
	o.Options = append(o.Options, opt)
}

// Defines a boolean list option with a pointer to its value
func BoolListVar(p *[]bool, long string, short rune) {
	CommandLine.BoolListVar(p, long, short)
}

// Defines a boolean list option and returns a pointer to its value
func (o *OptionParser) BoolList(long string, short rune) (p *[]bool) {
	p = &[]bool{}
	o.BoolListVar(p, long, short)
	return
}

// Defines a boolean list option and returns a pointer to its value
func BoolList(long string, short rune) (p *[]bool) {
	p = CommandLine.BoolList(long, short)
	return
}
