package optparse

import (
	"strconv"
)

// UintValue represents an unsigned uinteger value
type UintValue uint

// UintValueAssertion defines an assertion if the Value is a UintValue
type UintValueAssertion interface {
	Value
	IsUint() bool
}

// NewUintValue creates a new UintValue
func NewUintValue(val uint, p *uint) *UintValue {
	*p = val
	return (*UintValue)(p)
}

// Set sets the value
func (u *UintValue) Set(val string) error {
	v, err := strconv.ParseUint(val, 0, 64)
	*u = UintValue(v)
	return err
}

// IsUint returns true if the Value is a UintValue
func (u *UintValue) IsUint() bool {
	return true
}

// UintVar defines an unsigned integer option with a pointer to its value
func (o *OptionParser) UintVar(p *uint, long string, short rune, def uint) {
	opt := &Option{long, short, NewUintValue(def, p)}
	o.Options = append(o.Options, opt)
}

// UintVar defines an unsigned integer option with a pointer to its value
func UintVar(p *uint, long string, short rune, def uint) {
	CommandLine.UintVar(p, long, short, def)
}

// Uint defines an unsigned integer option and returns a pointer to its value
func (o *OptionParser) Uint(long string, short rune, def uint) (p *uint) {
	p = new(uint)
	o.UintVar(p, long, short, def)
	return
}

// Uint defines an unsigned integer option and returns a pointer to its value
func Uint(long string, short rune, def uint) (p *uint) {
	p = CommandLine.Uint(long, short, def)
	return
}
