package optparse

import (
	"os"
	"fmt"
	"strings"
)

// OptionParser structure structure contains options information
type OptionParser struct {
	Options			[]*Option
}

// Option structure contains option information
type Option struct {
	Long			string
	Short			rune
	Value			Value
}

// Value interface represents an abstract value type
type Value interface {
	Set(string)		error
}

// New creates a new OptionParser structure
func New() *OptionParser {
	return &OptionParser{}
}

// Reset resets the list of options
func (o *OptionParser) Reset() {
	o.Options = nil
}

// Parse parses a list of command line arguments and returns positional arguments
func (o *OptionParser) Parse(args []string) (a []string, e error) {
	var i int

ArgumentLoop:
	for i = 0; i < len(args); i++ {
		arg := args[i]

		// Stop parsing?
		if arg == "--" {
			break
		}

		// Long options
		if len(arg) > 2 && arg[0] == '-' && arg[1] == '-' {
			// `--key="value"` style
			if idx := strings.IndexRune(arg, '='); idx != -1 {
				key, val := arg[2:idx], arg[idx + 1:]

				for j := 0; j < len(o.Options); j++ {
					opt := o.Options[j]

					if key == opt.Long {
						e = opt.Value.Set(val)

						if e != nil {
							e = fmt.Errorf("invalid value `%s' for option `--%s'", val, key)
							return
						}

						continue ArgumentLoop
					}
				}

				e = fmt.Errorf("unrecognized option `--%s'", key)
				return
			}

			// `--key "value"` style
			key := arg[2:]

			for j := 0; j < len(o.Options); j++ {
				opt := o.Options[j]

				if key == opt.Long {
					// Boolean values have no value, they're always true
					if _, ok := opt.Value.(BoolValueAssertion); ok {
						opt.Value.Set("true")
						continue ArgumentLoop
					}
					if _, ok := opt.Value.(BoolListValueAssertion); ok {
						opt.Value.Set("true")
						continue ArgumentLoop
					}

					// Everything else requires a value
					if i < len(args) - 1 {
						val := args[i + 1]
						i++

						e = opt.Value.Set(val)
						if e != nil {
							e = fmt.Errorf("invalid value `%s' for option `--%s'", val, key)
							return
						}

						continue ArgumentLoop
					}

					e = fmt.Errorf("option `--%s' requires a value", key)
					return
				}
			}

			e = fmt.Errorf("unrecognized option `--%s'", key)
			return
		}

		// Short options
		if arg[0] == '-' {
			runes := []rune(arg[1:])

ShortLoop:
			for ri := 0; ri < len(runes); ri++ {
				r := runes[ri]

				for j := 0; j < len(o.Options); j++ {
					opt := o.Options[j]

					if r == opt.Short {
						// Boolean values have no value
						if _, ok := opt.Value.(BoolValueAssertion); ok {
							opt.Value.Set("true")
							continue ShortLoop
						}
						if _, ok := opt.Value.(BoolListValueAssertion); ok {
							opt.Value.Set("true")
							continue ShortLoop
						}

						// Everything else requires a value
						if ri == len(runes) - 1 {
							if i >= len(args) - 1 {
								e = fmt.Errorf("option `-%c' requires a value", r)
								return
							}

							e = opt.Value.Set(args[i + 1])
							if e != nil {
								e = fmt.Errorf("invalid value `%s' for option `-%c'", args[i + 1], r)
								return
							}

							i++
						} else {
							e = opt.Value.Set(string(runes[ri + 1:]))
							if e != nil {
								e = fmt.Errorf("invalid value `%s' for option `-%c'", string(runes[ri + 1:]), r)
								return
							}
						}

						continue ArgumentLoop
					}
				}

				e = fmt.Errorf("unrecognized option `-%c'", r)
				return
			}

			continue
		}

		a = append(a, arg)
	}

	// Append rest of the arguments
	if i < len(args) {
		a = append(a, args[i + 1:]...)
	}

	return
}

// Parse parses the list of command line arguments (`os.Args[1:]``) and returns positional arguments
func Parse() (a []string, e error) {
	a, e = CommandLine.Parse(os.Args[1:])
	return
}

// CommandLine is the default OptionParser, working with `os.Args[1:]`
var CommandLine = New()
