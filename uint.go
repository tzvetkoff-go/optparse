package optparse

import (
	"strconv"
)

// Unsigned uinteger value
type UintValue uint

type UintValueAssertion interface {
	Value
	IsUint() bool
}

func NewUintValue(val uint, p *uint) *UintValue {
	*p = val
	return (*UintValue)(p)
}

func (u *UintValue) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 64)
	*u = UintValue(v)
	return err
}

func (u *UintValue) IsUint() bool {
	return true
}

// Defines an unsigned integer option with a pointer to its value
func (o *OptionParser) UintVar(p *uint, long string, short rune, def uint) {
	opt := &Option{long, short, NewUintValue(def, p)}
	o.Options = append(o.Options, opt)
}

// Defines an unsigned integer option with a pointer to its value
func UintVar(p *uint, long string, short rune, def uint) {
	CommandLine.UintVar(p, long, short, def)
}

// Defines an unsigned integer option and returns a pointer to its value
func (o *OptionParser) Uint(long string, short rune, def uint) (p *uint) {
	p = new(uint)
	o.UintVar(p, long, short, def)
	return
}

// Defines an unsigned integer option and returns a pointer to its value
func Uint(long string, short rune, def uint) (p *uint) {
	p = CommandLine.Uint(long, short, def)
	return
}
