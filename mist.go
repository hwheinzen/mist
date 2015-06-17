// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Package mist provides an extended error type,
// a function for creation, methods for retrieving content,
// ... and two convenience functions.
package mist

// XError is an extended error interface.
type XError interface {
	error                // Error() returns an error string
	Details() string     // returns details in a string (e.g. an error trace)
}

// Mistake stores extended error information.
type mistake struct {
	error   string
	details string
}

// New returns an extended error using the given text strings.
func New(txt, det string) error {
	if txt == "" {
		return nil
	}
	return &mistake{error: txt, details: det}
}

// Error returns the error string,
// thus implementing the built-in error interface.
func (m *mistake) Error() string {
	return m.error
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
		*errp = New(xerr.Error(), pre+xerr.Details())
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
		*errp = New(xerr.Error(), xerr.Details()+suf)
	} else {
		*errp = New((*errp).Error(), suf)
	}
	return true
}
