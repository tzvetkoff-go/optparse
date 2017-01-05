package optparse

import (
	"strconv"
)

// Signed integer value
type IntValue int

type IntValueAssertion interface {
	Value
	IsInt() bool
}

func NewIntValue(val int, p *int) *IntValue {
	*p = val
	return (*IntValue)(p)
}

func (i *IntValue) Set(val string) error {
	v, err := strconv.ParseInt(val, 0, 64)
	*i = IntValue(v)
	return err
}

func (i *IntValue) IsInt() bool {
	return true
}

// Defines a signed integer option with a pointer to its value
func (o *OptionParser) IntVar(p *int, long string, short rune, def int) {
	opt := &Option{long, short, NewIntValue(def, p)}
	o.Options = append(o.Options, opt)
}

// Defines a signed integer option with a pointer to its value
func IntVar(p *int, long string, short rune, def int) {
	CommandLine.IntVar(p, long, short, def)
}

// Defines a signed integer option and returns a pointer to its value
func (o *OptionParser) Int(long string, short rune, def int) (p *int) {
	p = new(int)
	o.IntVar(p, long, short, def)
	return
}

// Defines a signed integer option and returns a pointer to its value
func Int(long string, short rune, def int) (p *int) {
	p = CommandLine.Int(long, short, def)
	return
}
