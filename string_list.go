package optparse

// StringListValue represents a list of string values
type StringListValue []string

// StringListValueAssertion defines an assertion if the Value is a StringListValue
type StringListValueAssertion interface {
	Value
	IsStringList() bool
}

// NewStringListValue creates a new StringListValue
func NewStringListValue(val []string, p *[]string) *StringListValue {
	*p = val
	return (*StringListValue)(p)
}

// Set appends a value
func (s *StringListValue) Set(val string) error {
	*s = append(*s, val)
	return nil
}

// IsStringList returns true if the Value is a StringListValue
func (s *StringListValue) IsStringList() bool {
	return true
}

// StringListVar defines a string list option with a pointer to its value
func (o *OptionParser) StringListVar(p *[]string, long string, short rune) {
	opt := &Option{long, short, NewStringListValue([]string{}, p)}
	o.Options = append(o.Options, opt)
}

// StringListVar defines a string list option with a pointer to its value
func StringListVar(p *[]string, long string, short rune) {
	CommandLine.StringListVar(p, long, short)
}

// StringList defines a string list option and returns a pointer to its value
func (o *OptionParser) StringList(long string, short rune) (p *[]string) {
	p = &[]string{}
	o.StringListVar(p, long, short)
	return
}

// StringList defines a string list option and returns a pointer to its value
func StringList(long string, short rune) (p *[]string) {
	p = CommandLine.StringList(long, short)
	return
}
