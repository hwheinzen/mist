// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Package mist provides an extended error type,
// a function for creation, methods for retrieving content,
// ... and two convenience functions.
package mist

// xerror is an extended error interface.
type XError interface {
	error                // Error() returns an error string
	Details() string     // returns a string (e.g. an error trace)
	Vars() []interface{} // returns variables (e.g. for substitution)
}

// Mistake stores extended error information.
type mistake struct {
	error   string
	details string
	vars    []interface{}
}

// New returns an extended error using the given text strings.
func New(txt, det string, vars ...interface{}) error {
	if txt == "" {
		return nil
	}
	return &mistake{error: txt, details: det, vars: vars}
}

// Error returns the error string,
// thus implementing the built-in error interface.
func (m *mistake) Error() string {
	return m.error
}

// Vars returns a slice of empty interfaces (containing variables),
// presumably for variable substitution.
func (m *mistake) Vars() []interface{} {
	return m.vars
}

// Details returns the details string.
func (m *mistake) Details() string {
	return m.details
}

// ----------

// Prepend adds a prefix to the details of the extended error,
// and returns true.
func Prepend(pre string, errp *error) bool {
	if *errp == nil {
		return false
	}
	xerr, ok := (*errp).(XError)
	if ok {
		*errp = New(xerr.Error(), pre+xerr.Details(), xerr.Vars())
	} else {
		*errp = New((*errp).Error(), pre)
	}
	return true
}

// Append adds a suffix to the details of the extended error,
// and returns true.
func Append(suf string, errp *error) bool {
	if *errp == nil {
		return false
	}
	xerr, ok := (*errp).(XError)
	if ok {
		*errp = New(xerr.Error(), xerr.Details()+suf, xerr.Vars())
	} else {
		*errp = New((*errp).Error(), suf)
	}
	return true
}
