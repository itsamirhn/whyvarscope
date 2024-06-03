# WhyVarScope

A go linter to check if you can avoid using variable scope inside if statements.

## Example

```go
package main

func zero() int {
    return 0
}

func main() {
    if o := zero(); o > 0 { // Why to use o here? 
        println("o is greater than 0")
    }
}
```

## Install

    go install github.com/itsamirhn/whyvarscope/cmd/whyvarscope@latest

# Usage
whyvarscope uses [`singlechecker`](https://pkg.go.dev/golang.org/x/tools/go/analysis/singlechecker) package to run:

```
whyvarscope: Check for unnecessary variable scope

Usage: whyvarscope [-flag] [package]


Flags:
  -V    print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -source
        no effect (deprecated)
  -tags string
        no effect (deprecated)
  -test
        indicates whether test files should be analyzed, too (default true)
  -trace string
        write trace log to this file
  -v    no effect (deprecated)
```