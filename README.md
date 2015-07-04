# mist
Extended error package for Go. 

This package provides an extended error type,
methods for adding variables and retrieving content, 
... and two convenience functions.

The extended error satifies the standard Go error interface.

The extended error can also carry (variabel) information via its details part.
You can add prefixes or sufixes to the details string at any time.
(Thus the details part becomes an error trace if you prefix it
with the function name every time a function returns an error.)

If the error message contains text/template expressions
(like " ... {{printf \"%f\" .FloatVariable}} ... )
which are to be resolved later in the program the extended error
can carry the associated variables.

### Motivation
For translating error messages to any users language - even in a 
multi-user/multi-language environment - it is necessary, either
to transport the language code to every place an error message
is created, or to transport the variable parts of an error message
seperated from the static error string and do translation and
variable substitution at one place high up in the logic hierarchie.

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/mist`

### Usage
The TestCascade function in `mist_test.go` shows a possible usage.
(Run with `go test -v`.)
