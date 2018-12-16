package optparse

import (
	"fmt"
	"path"
	"runtime"
	"errors"
	"testing"
)

type test struct {
	where		string
	setup		func(o *OptionParser, t *test)
	args		[]string
	expected	interface {}
	err			error
	result		interface {}
}

var tests = []test{
	// Boolean
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Bool("bool", 'b', false) },
		args: 		[]string{"-b"},
		expected:	true,
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Bool("bool", 'b', false) },
		args: 		[]string{"--bool"},
		expected:	true,
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Bool("bool", 'b', false) },
		args: 		[]string{"--bool=false"},
		expected:	false,
	},

	// Signed integer
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Int("int", 'i', 0) },
		args: 		[]string{"-i", "-31337"},
		expected:	-31337,
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Int("int", 'i', 0) },
		args: 		[]string{"-i-031337"},
		expected:	-031337,
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Int("int", 'i', 0) },
		args: 		[]string{"--int", "-31337"},
		expected:	-31337,
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Int("int", 'i', 0) },
		args: 		[]string{"--int=-0x31337"},
		expected:	-0x31337,
	},

	// Unsigned integer
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Uint("uint", 'u', 0) },
		args: 		[]string{"-u", "31337"},
		expected:	uint(31337),
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Uint("uint", 'u', 0) },
		args: 		[]string{"-u031337"},
		expected:	uint(031337),
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Uint("uint", 'u', 0) },
		args: 		[]string{"--uint", "31337"},
		expected:	uint(31337),
	},
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.Uint("uint", 'u', 0) },
		args: 		[]string{"--uint=0x31337"},
		expected:	uint(0x31337),
	},

	// Float
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Float("float", 'f', 0.0) },
		args:		[]string{"-f3.1337"},
		expected:	3.1337,
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Float("float", 'f', 0.0) },
		args:		[]string{"-f", "31.337"},
		expected:	31.337,
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Float("float", 'f', 0.0) },
		args:		[]string{"--float", "313.37"},
		expected:	313.37,
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Float("float", 'f', 0.0) },
		args:		[]string{"--float=31337e-1"},
		expected:	3133.7,
	},

	// String
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.String("string", 's', "") },
		args:		[]string{"-sfoo"},
		expected:	"foo",
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.String("string", 's', "") },
		args:		[]string{"-s", "foo"},
		expected:	"foo",
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.String("string", 's', "") },
		args:		[]string{"--string", "foo"},
		expected:	"foo",
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.String("string", 's', "") },
		args:		[]string{"--string=foo"},
		expected:	"foo",
	},

	// Boolean list
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.BoolList("bool", 'b') },
		args: 		[]string{"-b", "--bool=false", "--bool"},
		expected:	[]bool{true, false, true},
	},

	// Signed integer list
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.IntList("int", 'i') },
		args: 		[]string{"-i-1", "-i", "-2", "--int=-3"},
		expected:	[]int{-1, -2, -3},
	},

	// Unsigned integer list
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.UintList("uint", 'u') },
		args: 		[]string{"-u1", "-u", "2", "--uint=3"},
		expected:	[]uint{1, 2, 3},
	},

	// Float list
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.FloatList("float", 'f') },
		args: 		[]string{"-f1", "-f", "2.2", "--float=333e-2"},
		expected:	[]float64{1, 2.2, 333e-2},
	},

	// String list
	{
		where: 		here(),
		setup: 		func(o *OptionParser, tt *test) { tt.result = o.StringList("string", 's') },
		args: 		[]string{"-sfoo", "-s", "bar", "--string=baz"},
		expected:	[]string{"foo", "bar", "baz"},
	},

	// Unrecognized option
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Bool("bool", 'b', false) },
		args:		[]string{"-x"},
		err:		errors.New("unrecognized option `-x'"),
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Bool("bool", 'b', false) },
		args:		[]string{"--unrecognized"},
		err:		errors.New("unrecognized option `--unrecognized'"),
	},

	// Invalid value
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Bool("bool", 'b', false) },
		args:		[]string{"--bool=not-bool"},
		err:		errors.New("invalid value `not-bool' for option `--bool'"),
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Bool("int", 'i', false) },
		args:		[]string{"--int=not-int"},
		err:		errors.New("invalid value `not-int' for option `--int'"),
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Bool("uint", 'f', false) },
		args:		[]string{"--uint=not-uint"},
		err:		errors.New("invalid value `not-uint' for option `--uint'"),
	},
	{
		where:		here(),
		setup:		func(o *OptionParser, tt *test) { tt.result = o.Bool("float", 'f', false) },
		args:		[]string{"--float=not-float"},
		err:		errors.New("invalid value `not-float' for option `--float'"),
	},
}

func TestOptionParser(t *testing.T) {
	for _, tt := range tests {
		o := New()
		tt.setup(o, &tt)
		_, err := o.Parse(tt.args)

		if err != nil {
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("%s: got error(%s), expected error(%s)", tt.where, err, tt.err)
				}
			} else {
				t.Errorf("%s: %s", tt.where, err.Error())
			}

			continue
		}

		if r, ok := tt.result.(*bool); ok {
			if *r != tt.expected {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			}
		} else if r, ok := tt.result.(*int); ok {
			if *r != tt.expected {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			}
		} else if r, ok := tt.result.(*uint); ok {
			if *r != tt.expected {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			}
		} else if r, ok := tt.result.(*float64); ok {
			if *r != tt.expected {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			}
		} else if r, ok := tt.result.(*string); ok {
			if *r != tt.expected {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			}
		} else if r, ok := tt.result.(*[]bool); ok {
			ev := tt.expected.([]bool)
			rv := *r
			if len(rv) != len(ev) {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			} else {
				for idx := range *r {
					if rv[idx] != ev[idx] {
						t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
						break
					}
				}
			}
		} else if r, ok := tt.result.(*[]int); ok {
			ev := tt.expected.([]int)
			rv := *r
			if len(rv) != len(ev) {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			} else {
				for idx := range *r {
					if rv[idx] != ev[idx] {
						t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
						break
					}
				}
			}
		} else if r, ok := tt.result.(*[]uint); ok {
			ev := tt.expected.([]uint)
			rv := *r
			if len(rv) != len(ev) {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			} else {
				for idx := range *r {
					if rv[idx] != ev[idx] {
						t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
						break
					}
				}
			}
		} else if r, ok := tt.result.(*[]float64); ok {
			ev := tt.expected.([]float64)
			rv := *r
			if len(rv) != len(ev) {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			} else {
				for idx := range *r {
					if rv[idx] != ev[idx] {
						t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
						break
					}
				}
			}
		} else if r, ok := tt.result.(*[]string); ok {
			ev := tt.expected.([]string)
			rv := *r
			if len(rv) != len(ev) {
				t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
			} else {
				for idx := range *r {
					if rv[idx] != ev[idx] {
						t.Errorf("%s: got %#v, expected %#v", tt.where, *r, tt.expected)
						break
					}
				}
			}
		} else {
			t.Errorf("%s: got unexpected result %#v, expected %#v", tt.where, *r, tt.expected)
		}
	}
}

func here() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", path.Base(file), line)
}
