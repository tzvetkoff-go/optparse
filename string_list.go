package optparse

// List of string value
type StringListValue []string

type StringListValueAssertion interface {
	Value
	IsStringList() bool
}

func NewStringListValue(val []string, p *[]string) *StringListValue {
	*p = val
	return (*StringListValue)(p)
}

func (s *StringListValue) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func (s *StringListValue) IsStringList() bool {
	return true
}

// Defines a string list option with a pointer to its value
func (o *OptionParser) StringListVar(p *[]string, long string, short rune) {
	opt := &Option{long, short, NewStringListValue([]string{}, p)}
	o.Options = append(o.Options, opt)
}

// Defines a string list option with a pointer to its value
func StringListVar(p *[]string, long string, short rune) {
	CommandLine.StringListVar(p, long, short)
}

// Defines a string list option and returns a pointer to its value
func (o *OptionParser) StringList(long string, short rune) (p *[]string) {
	p = &[]string{}
	o.StringListVar(p, long, short)
	return
}

// Defines a string list option and returns a pointer to its value
func StringList(long string, short rune) (p *[]string) {
	p = CommandLine.StringList(long, short)
	return
}
