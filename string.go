package optparse

// StringValue represents a string value
type StringValue string

// StringValueAssertion defines an assertion if the Value is a StringValue
type StringValueAssertion interface {
	Value
	IsString() bool
}

// NewStringValue creates a new StringValue
func NewStringValue(val string, p *string) *StringValue {
	*p = val
	return (*StringValue)(p)
}

// Set sets the value
func (s *StringValue) Set(val string) error {
	*s = StringValue(val)
	return nil
}

// IsString returns true if the Value is a StringValue
func (s *StringValue) IsString() bool {
	return true
}

// StringVar defines a string option with a pointer to its value
func (o *OptionParser) StringVar(p *string, long string, short rune, def string) {
	opt := &Option{long, short, NewStringValue(def, p)}
	o.Options = append(o.Options, opt)
}

// StringVar defines a string option with a pointer to its value
func StringVar(p *string, long string, short rune, def string) {
	CommandLine.StringVar(p, long, short, def)
}

// String defines a string option and returns a pointer to its value
func (o *OptionParser) String(long string, short rune, def string) (p *string) {
	p = new(string)
	o.StringVar(p, long, short, def)
	return
}

// String defines a string option and returns a pointer to its value
func String(long string, short rune, def string) (p *string) {
	p = CommandLine.String(long, short, def)
	return
}
