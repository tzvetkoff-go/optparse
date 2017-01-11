package optparse

import (
	"strconv"
)

// UintListValue represents a list of unsigned integer values
type UintListValue []uint

// UintListValueAssertion defines an assertion if the Value is a UintListValue
type UintListValueAssertion interface {
	Value
	IsUintList() bool
}

// NewUintListValue creates a new UintListValue
func NewUintListValue(val []uint, p *[]uint) *UintListValue {
	*p = val
	return (*UintListValue)(p)
}

// Set appends a value
func (u *UintListValue) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 64)
	*u = append(*u, uint(v))
	return err
}

// IsUintList returns true if the Value is a UintListValue
func (u *UintListValue) IsUintList() bool {
	return true
}

// UintListVar defines an unsigned integer list option with a pointer to its value
func (o *OptionParser) UintListVar(p *[]uint, long string, short rune) {
	opt := &Option{long, short, NewUintListValue([]uint{}, p)}
	o.Options = append(o.Options, opt)
}

// UintListVar defines an unsigned integer list option with a pointer to its value
func UintListVar(p *[]uint, long string, short rune) {
	CommandLine.UintListVar(p, long, short)
}

// UintList defines an unsigned integer list option and returns a pointer to its value
func (o *OptionParser) UintList(long string, short rune) (p *[]uint) {
	p = &[]uint{}
	o.UintListVar(p, long, short)
	return
}

// UintList defines an unsigned integer list option and returns a pointer to its value
func UintList(long string, short rune) (p *[]uint) {
	p = CommandLine.UintList(long, short)
	return
}
