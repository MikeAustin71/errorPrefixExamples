package examplesDto

import (
	"fmt"

	erPref "github.com/MikeAustin71/errprefops/errpref"
)

// TestFuncDtoBravo01 - This type is designed to call a series
// of methods using Error Prefix Data Transfer Objects otherwise
// known as 'ErrPrefixDto'. For source code documentation on this
// type, reference:
//  https://pkg.go.dev/github.com/MikeAustin71/errprefops/errpref#ErrPrefixDto
//
// This method chain will return an error on the last function
// call, if, and only if, parameter, 'returnError' is set to 'true'.
//
// The entry point for this method chain is:
//    TestFuncBravo01.Tx1DoSomethingSpecial()
//
// These examples showcase the 'ErrPrefixDto' type and the use of
// an empty interface to initialize a new instance of
// 'ErrPrefixDto'.
//
type TestFuncDtoBravo01 struct {
	input string
}

func (tFuncDtoBravo01 *TestFuncDtoBravo01) Tx1DoSomethingSpecial(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"TestFuncDtoBravo01." +
			"Tx1DoSomethingSpecial()")

	tFuncBravo02 := testFuncDtoBravo02{}

	return tFuncBravo02.tx2DoSomethingElse(
		returnError,
		ePrefix)
}

type testFuncDtoBravo02 struct {
	input string
}

func (tFuncBravo02 *testFuncDtoBravo02) tx2DoSomethingElse(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo02." +
			"tx2DoSomethingElse()")

	tFuncBravo03 := testFuncDtoBravo03{}

	err := tFuncBravo03.tx3DoAnything(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo03 struct {
	input string
}

func (tFuncBravo03 *testFuncDtoBravo03) tx3DoAnything(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo03." +
			"tx3DoAnything()")

	tFuncBravo04 := testFuncDtoBravo04{}

	err := tFuncBravo04.tx4DoNothing(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo04 struct {
	input string
}

func (tFuncBravo04 *testFuncDtoBravo04) tx4DoNothing(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo04." +
			"tx4DoNothing()")

	tFuncBravo05 := testFuncDtoBravo05{}

	ePrefix.SetCtx("A/B==4")

	err := tFuncBravo05.tx5DoSomethingBig(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo05 struct {
	input string
}

func (tFuncBravo05 *testFuncDtoBravo05) tx5DoSomethingBig(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo05." +
			"tx5DoSomethingBig()")

	tFuncBravo06 := testFuncDtoBravo06{}

	ePrefix.SetCtx("A->B")

	err := tFuncBravo06.tx6GiveUp(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo06 struct {
	input string
}

func (tFuncY6 *testFuncDtoBravo06) tx6GiveUp(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	var err error

	if returnError {
		err = fmt.Errorf("%v\n"+
			"An Error Ocurred! Something Bad...\n"+
			"Maybe its Divide By Zero!\n",
			ePrefix.String())
	}

	return err
}
