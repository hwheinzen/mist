// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package mist

import (
	//"errors"
	//"fmt"
	"testing"
)

func TestNewErrorDetails(t *testing.T) {
	err := New("")
	if err != nil {
		t.Error("err should be nil after New(\"\")")
	}
	if Prepend("prefix", &err) {
		t.Error("Prepend nil should be false")
	}
	err = New("error", "details")
	xerr, ok := err.(XError)
	if !ok {
		t.Error("xerr should be of interface type XError")
	}
	if xerr.Error() != "error" {
		t.Error("\nExpected: " + "error" + "\ngot:      " + xerr.Error())
	}
	if xerr.Details() != "details" {
		t.Error("\nExpected: " + "details" + "\ngot:      " + xerr.Details())
	}
}

type prependTest struct {
	det string
	in  string
	out string
}

var prependTests = []prependTest{
	{det: "details", in: "", out: ":details"},
	{det: "details", in: "prefix", out: "prefix:details"},
}

func TestPrepend(t *testing.T) {
	for _, v := range prependTests {
		err := New("error", v.det)
		ok := Prepend(v.in, &err)
		if !ok {
			t.Error("err should be an error")
		}
		xerr, ok := err.(XError)
		if !ok {
			t.Error("xerr should be of interface type XError")
		} else {
			if xerr.Details() != v.out {
				t.Error("\nExpected: " + v.out + "\ngot:      " + xerr.Details())
			}
		}
	}
}

type appendTest struct {
	det string
	in  string
	out string
}

var appendTests = []appendTest{
	{det: "details", in: "", out: "details;"},
	{det: "details", in: "suffix", out: "details;suffix"},
}

func TestAppend(t *testing.T) {
	for _, v := range appendTests {
		err := New("error", v.det)
		ok := Append(v.in, &err)
		if !ok {
			t.Error("err should be an error")
		}
		xerr, ok := err.(XError)
		if !ok {
			t.Error("xerr should be of interface type XError")
		} else {
			if xerr.Details() != v.out {
				t.Error("\nExpected: " + v.out + "\ngot:      " + xerr.Details())
			}
		}
	}
}

func TestCascade(t *testing.T) {
	fncname := "TestCascade"

	err := cascade1(t)
	if Prepend(fncname, &err) {
		t.Log("Error:   " + err.Error())
		t.Log("Details: " + err.(XError).Details())
	}

	if err.Error() != "intentional error occured" {
		t.Error("\nExpected: " + "intentional error occured" + "\ngot:      " + err.Error())
	}
	xerr, ok := err.(XError)
	if !ok {
		t.Error("xerr should be of interface type XError")
	}
	if xerr.Details() != "TestCascade:cascade1:cascade2:details" {
		t.Error("\nExpected: " + "TestCascade:cascade1:cascade2:details" + "\ngot:      " + xerr.Details())
	}
}

func cascade1(t *testing.T) (err error) {
	fncname := "cascade1"

	err = cascade2(t)
	if Prepend(fncname, &err) {
		return
	}
	return
}

func cascade2(t *testing.T) (err error) {
	fncname := "cascade2"

	err = New("intentional error occured", "details")
	if Prepend(fncname, &err) {
		return
	}
	return
}
