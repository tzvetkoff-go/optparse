# Command line option parsing for Go

[![GoDoc](https://godoc.org/github.com/go2c/optparse?status.svg)](http://godoc.org/github.com/go2c/optparse)
[![Build Status](https://travis-ci.org/go2c/optparse.svg?branch=master)](https://travis-ci.org/go2c/optparse)

A slightly better CLI option parsing for Go.
Because [flag](https://golang.org/pkg/flag/) is somewhat crippled.

## Example

``` go
package main

import (
    "os"
    "io"
    "fmt"
    "github.com/go2c/optparse"
)

func usage(f io.Writer, name string) {
    fmt.Fprintln(f, "Usage:")
    fmt.Fprintf(f, "  %s [options] [arguments]\n", name)
    fmt.Fprintln(f)

    fmt.Fprintln(f, "Common options:")
    fmt.Fprintln(f, "  -h, --help                Print help and exit")
    fmt.Fprintln(f, "  -v, --version             Print version and exit")
    fmt.Fprintln(f)

    fmt.Fprintln(f, "Example single-value options:")
    fmt.Fprintln(f, "  -i I, --int=I             Set a signed integer value (default: 0)")
    fmt.Fprintln(f, "  -u U, --uint=U            Set an unsigned integer value (default: 0)")
    fmt.Fprintln(f, "  -f F, --float=F           Set a floating point value (default: 0.0)")
    fmt.Fprintln(f, "  -s S, --string=S          Set a string value (default: \"\")")
    fmt.Fprintln(f)

    fmt.Fprintln(f, "Example list-value options:")
    fmt.Fprintln(f, "  -I I, --int-list=I        Append a signed integer to a list")
    fmt.Fprintln(f, "  -U U, --uint-list=I       Append an unsigned integer to a list")
    fmt.Fprintln(f, "  -F F, --float-list=I      Append a float to a list")
    fmt.Fprintln(f, "  -S S, --string-list=I     Append a string to a list")

    if f == os.Stderr {
        os.Exit(1)
    } else {
        os.Exit(0)
    }
}

func main() {
    pHelp := optparse.Bool("help", 'h', false)
    pVersion := optparse.Bool("version", 'v', false)

    pInt := optparse.Int("int", 'i', 0)
    pUint := optparse.Uint("uint", 'u', 0)
    pFloat := optparse.Float("float", 'f', 0.0)
    pString := optparse.String("string", 's', "")

    pIntList := optparse.IntList("int-list", 'I')
    pUintList := optparse.UintList("uint-list", 'U')
    pFloatList := optparse.FloatList("float-list", 'F')
    pStringList := optparse.StringList("string-list", 'S')

    args, err := optparse.Parse(os.Args[1:])

    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %s\n\n", os.Args[0], err.Error())
        usage(os.Stderr, os.Args[0])
    }

    if *pHelp {
        usage(os.Stdout, os.Args[0])
    }

    if *pVersion {
        fmt.Println("0.1.0")
        os.Exit(0)
    }

    fmt.Printf("int         : %T(%v)\n", *pInt, *pInt)
    fmt.Printf("uint        : %T(%v)\n", *pUint, *pUint)
    fmt.Printf("float       : %T(%f)\n", *pFloat, *pFloat)
    fmt.Printf("string      : %T(%q)\n", *pString, *pString)
    fmt.Println()

    fmt.Printf("int-list    : %T{%v}\n", *pIntList, *pIntList)
    fmt.Printf("uint-list   : %T{%v}\n", *pUintList, *pUintList)
    fmt.Printf("float-list  : %T{%f}\n", *pFloatList, *pFloatList)
    fmt.Printf("string-list : %T{%q}\n", *pStringList, *pStringList)
    fmt.Println()

    fmt.Printf("args        : %T{%q}\n", args, args)
    fmt.Println()
}
```

If you need to parse arguments more than once, you can allocate a `OptionParser` handle instead of using `optparse.CommandLine`:

``` go
o := optparse.New()
pHelp := o.Bool("help", 'h', false)
// ...
args, err := o.Parse([]string{"-h"})
```

## License

The code is subject to the [MIT license](https://opensource.org/licenses/MIT).
