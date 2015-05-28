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
	error
	Details() string
}

// Mistake stores extended error information.
type mistake struct {
	error   string
	details string
	// contains filtered or unexported fields
}

// New returns an extended error using the given text strings.
func New(txt, det string) (xerr XError) {
	if txt == "" {
		return nil
	}
	return &mistake{error: txt, details: det}
}

// FromError returns an extended error using the given error
// and the details string.
func FromError(err error, det string) (xerr XError) {
	if err == nil {
		return
	}
	return &mistake{error: err.Error(), details: det}
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
func Prepend(pre string, xerr *XError) bool {
	if xerr == nil {
		return false
	}
	*xerr = New((*xerr).Error(), pre+(*xerr).Details())
	return true
}

// Append adds a suffix to the details of the extended error,
// and returns true.
func Append(suf string, xerr *XError) bool {
	if xerr == nil {
		return false
	}
	*xerr = New((*xerr).Error(), (*xerr).Details()+suf)
	return true
}
