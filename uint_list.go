package optparse

import (
	"strconv"
)

// List of unsigned integer values
type UintListValue []uint

type UintListValueAssertion interface {
	Value
	IsUintList() bool
}

func NewUintListValue(val []uint, p *[]uint) *UintListValue {
	*p = val
	return (*UintListValue)(p)
}

func (u *UintListValue) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 64)
	*u = append(*u, uint(v))
	return err
}

func (u *UintListValue) IsUintList() bool {
	return true
}

// Defines an unsigned integer list option with a pointer to its value
func (o *OptionParser) UintListVar(p *[]uint, long string, short rune) {
	opt := &Option{long, short, NewUintListValue([]uint{}, p)}
	o.Options = append(o.Options, opt)
}

// Defines an unsigned integer list option with a pointer to its value
func UintListVar(p *[]uint, long string, short rune) {
	CommandLine.UintListVar(p, long, short)
}

// Defines an unsigned integer list option and returns a pointer to its value
func (o *OptionParser) UintList(long string, short rune) (p *[]uint) {
	p = &[]uint{}
	o.UintListVar(p, long, short)
	return
}

// Defines an unsigned integer list option and returns a pointer to its value
func UintList(long string, short rune) (p *[]uint) {
	p = CommandLine.UintList(long, short)
	return
}
