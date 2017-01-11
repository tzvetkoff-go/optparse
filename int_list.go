package optparse

import (
	"strconv"
)

// IntListValue represents a list of signed integer values
type IntListValue []int

// IntListValueAssertion defines an assertion if the Value is an IntListValue
type IntListValueAssertion interface {
	Value
	IsIntList() bool
}

// NewIntListValue creates a new IntListValue
func NewIntListValue(val []int, p *[]int) *IntListValue {
	*p = val
	return (*IntListValue)(p)
}

// Set appends a value
func (i *IntListValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)
	*i = append(*i, int(v))
	return err
}

// IsIntList returns true if the Value is an IntListValue
func (i *IntListValue) IsIntList() bool {
	return true
}

// IntListVar defines a signed integer list option with a pointer to its value
func (o *OptionParser) IntListVar(p *[]int, long string, short rune) {
	opt := &Option{long, short, NewIntListValue([]int{}, p)}
	o.Options = append(o.Options, opt)
}

// IntListVar defines a signed integer list option with a pointer to its value
func IntListVar(p *[]int, long string, short rune) {
	CommandLine.IntListVar(p, long, short)
}

// IntList defines a signed integer list option and returns a pointer to its value
func (o *OptionParser) IntList(long string, short rune) (p *[]int) {
	p = &[]int{}
	o.IntListVar(p, long, short)
	return
}

// IntList defines a signed integer list option and returns a pointer to its value
func IntList(long string, short rune) (p *[]int) {
	p = CommandLine.IntList(long, short)
	return
}
