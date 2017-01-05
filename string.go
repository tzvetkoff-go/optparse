package optparse

// String value
type StringValue string

type StringValueAssertion interface {
	Value
	IsString() bool
}

func NewStringValue(val string, p *string) *StringValue {
	*p = val
	return (*StringValue)(p)
}

func (s *StringValue) Set(val string) error {
	*s = StringValue(val)
	return nil
}

func (s *StringValue) IsString() bool {
	return true
}

// Defines a string option with a pointer to its value
func (o *OptionParser) StringVar(p *string, long string, short rune, def string) {
	opt := &Option{long, short, NewStringValue(def, p)}
	o.Options = append(o.Options, opt)
}

// Defines a string option with a pointer to its value
func StringVar(p *string, long string, short rune, def string) {
	CommandLine.StringVar(p, long, short, def)
}

// Defines a string option and returns a pointer to its value
func (o *OptionParser) String(long string, short rune, def string) (p *string) {
	p = new(string)
	o.StringVar(p, long, short, def)
	return
}

// Defines a string option and returns a pointer to its value
func String(long string, short rune, def string) (p *string) {
	p = CommandLine.String(long, short, def)
	return
}
