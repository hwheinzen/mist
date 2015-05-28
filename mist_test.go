// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package mist

import (
	"errors"
	"testing"
)

func TestNewErrorDetail(t *testing.T) {
	xerr := New("", "xyz")
	if xerr != nil {
		t.Error("xerr should be nil")
	}
	xerr = New("error", "details")
	if xerr.Error() != "error" {
		t.Error("\nExpected: " + "error" + "\ngot:      " + xerr.Error())
	}
	if xerr.Details() != "details" {
		t.Error("\nExpected: " + "details" + "\ngot:      " + xerr.Details())
	}
}

type prependTest struct {
	in  string
	out string
}

var prependTests = []prependTest{
	{in: "", out: "details"},
	{in: "prefix:", out: "prefix:details"},
}

func TestPrepend(t *testing.T) {
	xerr := New("error", "details")
	for _, v := range prependTests {
		ok := Prepend(v.in, &xerr)
		if !ok {
			t.Error("err should be an error")
		}
		if xerr.Details() != v.out {
			t.Error("\nExpected: " + v.out + "\ngot:      " + xerr.Details())
		}
	}
}

type appendTest struct {
	in  string
	out string
}

var appendTests = []appendTest{
	{in: "", out: "details"},
	{in: ":suffix", out: "details:suffix"},
}

func TestAppend(t *testing.T) {
	xerr := New("error", "details")
	for _, v := range appendTests {
		ok := Append(v.in, &xerr)
		if !ok {
			t.Error("err should be an error")
		}
		if xerr.Details() != v.out {
			t.Error("\nExpected: " + v.out + "\ngot:      " + xerr.Details())
		}
	}
}

func TestCascade(t *testing.T) {
	fncname := "TestCascade"

	xerr := cascade1()
	if Prepend(fncname+":", &xerr) {
		t.Log(xerr)
		t.Log("Details: " + xerr.Details())
	}
	if xerr.Error() != "intentional error occured" {
		t.Error("\nExpected: " + "intentional error occured" + "\ngot:      " + xerr.Error())
	}
	if xerr.Details() != "TestCascade:cascade1:cascade2:details" {
		t.Error("\nExpected: " + "TestCascade:cascade1:cascade2:details" + "\ngot:      " + xerr.Details())
	}
}

func cascade1() (xerr XError) {
	fncname := "cascade1"
	
	xerr = cascade2()
	if Prepend(fncname+":", &xerr) {
		return
	}
	return
}

func cascade2() (xerr XError) {
	fncname := "cascade2"
	
	err := errors.New("intentional error occured")
	if err != nil {
		xerr = FromError(err, "details")
		Prepend(fncname+":", &xerr)
		return
	}
	return
}





















