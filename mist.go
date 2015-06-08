// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

/*
 mist.go is home to an extended error type.
*/

// Package mist provides an extended error type,
// methods for creation, and for retrieving content,
// ... and two convenience functions.
package mist

// XError is an extended error interface.
type XError interface {
	error                   // Error() returns an error string
	Details() string        // returns a string (e.g. an error trace)
	Vars()    []interface{} // returns variables (e.g. for substitution)
}

// Mistake stores extended error information.
type mistake struct {
	error   string
	details string
	vars    []interface{}
}

// New returns an extended error using the given text strings.
func New(txt, det string, v ...interface{}) (xerr XError) {
	if txt == "" {
		return nil
	}
	return &mistake{error: txt, details: det, vars: v}
}

// FromError returns an extended error using the given error.
//func FromError(err error, det string, v ...interface{}) (xerr XError) {
func FromError(err error) (xerr XError) {
	if err == nil {
		return nil
	}
	return &mistake{error: err.Error()}
}

// Error returns the error string,
// thus implementing the built-in error interface.
func (m *mistake) Error() string {
	return m.error
}

// Vars returns a slice of empty interfaces containing variables.
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
func Prepend(pre string, xerr *XError) bool {
	if *xerr == nil {
		return false
	}
	*xerr = New((*xerr).Error(), pre+(*xerr).Details())
	return true
}

// Append adds a suffix to the details of the extended error,
// and returns true.
func Append(suf string, xerr *XError) bool {
	if *xerr == nil {
		return false
	}
	*xerr = New((*xerr).Error(), (*xerr).Details()+suf)
	return true
}
