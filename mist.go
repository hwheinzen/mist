// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// OBSOLETE. THIS PACKAGE IS NO LONGER MAINTAINED;
// IT HAS BEEN SUPERSEDED BY stringl10n/message.
//
// Package mist (short for mistake) provides an extended error type,
// methods for adding variables and retrieving content,
// ... and two convenience functions.
package mist

import "strings"

// XError is an extended error interface.
type XError interface {
	error                           // Error() returns an error string
	Details() string                // returns details in a string (e.g. an error trace)
	AddVar(n string, v interface{}) // attaches a variable to the extended error
	Vars() []struct {               // returns the attached variables
		Name  string
		Value interface{}
	}
}

// Mistake stores extended error information.
type mistake struct {
	error   string
	details string
	vars    []struct {
		Name  string
		Value interface{}
	}
}

// New returns an extended error using the given text strings.
// First string is error message, others are details.
func New(err string, dets ...string) error {
	if err == "" {
		return nil
	}
	if len(dets) == 0 {
		return &mistake{error: err}
	}
	return &mistake{error: err, details: strings.Join(dets, "")}
}

// AddVar adds a Variable to the extended error.
func (m *mistake) AddVar(n string, v interface{}) {
	m.vars = append(
		m.vars,
		struct {
			Name  string
			Value interface{}
		}{
			Name: n,
			Value: v,
		},
	)
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

// Vars returns the Variables attached to the extended error.
func (m *mistake) Vars() []struct {
	Name  string
	Value interface{}
} {
	return m.vars
}

// ----------

// Prepend adds a prefix and ":" to the details of the extended error,
// and returns true.
func Prepend(pre string, errp *error) bool {
	if *errp == nil { // no error occured
		return false
	}
	xerr, ok := (*errp).(XError)
	if ok {
		*errp = XError(&mistake{
			error:   xerr.Error(),
			details: pre + ":" + xerr.Details(),
			vars:    xerr.Vars(),
		})
	} else {
		*errp = XError(&mistake{
			error:   (*errp).Error(),
			details: pre,
		})
	}
	return true
}

// Append adds ";" and a suffix to the details of the extended error,
// and returns true.
func Append(suf string, errp *error) bool {
	if *errp == nil { // no error occured
		return false
	}
	xerr, ok := (*errp).(XError)
	if ok {
		*errp = XError(&mistake{
			error:   xerr.Error(),
			details: xerr.Details() + ";" + suf,
			vars:    xerr.Vars(),
		})
	} else {
		*errp = XError(&mistake{
			error:   (*errp).Error(),
			details: suf,
		})
	}
	return true
}
