package optparse

import (
	"strconv"
)

// List of signed integer values
type IntListValue []int

type IntListValueAssertion interface {
	Value
	IsIntList() bool
}

func NewIntListValue(val []int, p *[]int) *IntListValue {
	*p = val
	return (*IntListValue)(p)
}

func (i *IntListValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)
	*i = append(*i, int(v))
	return err
}

func (i *IntListValue) IsIntList() bool {
	return true
}

// Defines a signed integer list option with a pointer to its value
func (o *OptionParser) IntListVar(p *[]int, long string, short rune) {
	opt := &Option{long, short, NewIntListValue([]int{}, p)}
	o.Options = append(o.Options, opt)
}

// Defines a signed integer list option with a pointer to its value
func IntListVar(p *[]int, long string, short rune) {
	CommandLine.IntListVar(p, long, short)
}

// Defines a signed integer list option and returns a pointer to its value
func (o *OptionParser) IntList(long string, short rune) (p *[]int) {
	p = &[]int{}
	o.IntListVar(p, long, short)
	return
}

// Defines a signed integer list option and returns a pointer to its value
func IntList(long string, short rune) (p *[]int) {
	p = CommandLine.IntList(long, short)
	return
}
