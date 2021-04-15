package examplesDto

import (
	"fmt"

	erPref "github.com/MikeAustin71/errpref"
)

// TestFuncDtoBravo01 - This type is designed to call a series
// of methods using Error Prefix Data Transfer Objects otherwise
// known as 'ErrPrefixDto'. For source code documentation on this
// type, reference:
//  https://pkg.go.dev/github.com/MikeAustin71/errpref#ErrPrefixDto
//
// This method chain will return an error on the last function
// call, if, and only if, parameter, 'returnError' is set to 'true'.
//
// The entry point for this method chain is:
//    TestFuncBravo01.Tx1TrySomethingSpecial()
//
// These examples showcase the 'ErrPrefixDto' type and the use of
// an empty interface to initialize a new instance of
// 'ErrPrefixDto'.
//
type TestFuncDtoBravo01 struct {
	input string
}

func (tFuncDtoBravo01 *TestFuncDtoBravo01) Tx1TrySomethingSpecial(
	returnError bool,
	errorPrefix interface{}) error {
	var ePrefix *erPref.ErrPrefixDto
	var err error

	ePrefix,
		err = erPref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestFuncDtoAlpha01.Tx1TrySomethingSpecial()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo02 := testFuncDtoBravo02{}

	return tFuncBravo02.tx2TrySomethingElse(
		returnError,
		ePrefix)
}

type testFuncDtoBravo02 struct {
	input string
}

func (tFuncBravo02 *testFuncDtoBravo02) tx2TrySomethingElse(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo02." +
			"tx2TrySomethingElse()")

	tFuncBravo03 := testFuncDtoBravo03{}

	err := tFuncBravo03.tx3TryAnything(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo03 struct {
	input string
}

func (tFuncBravo03 *testFuncDtoBravo03) tx3TryAnything(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo03." +
			"tx3TryAnything()")

	tFuncBravo04 := testFuncDtoBravo04{}

	err := tFuncBravo04.tx4TryDoingNothing(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo04 struct {
	input string
}

func (tFuncBravo04 *testFuncDtoBravo04) tx4TryDoingNothing(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo04." +
			"tx4TryDoingNothing()")

	tFuncBravo05 := testFuncDtoBravo05{}

	ePrefix.SetCtx("A/B==4")

	err := tFuncBravo05.tx5TrySomethingBig(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo05 struct {
	input string
}

func (tFuncBravo05 *testFuncDtoBravo05) tx5TrySomethingBig(
	returnError bool,
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoBravo05." +
			"tx5TrySomethingBig()")

	tFuncBravo06 := testFuncDtoBravo06{}

	ePrefix.SetCtx("A->B")

	err := tFuncBravo06.tx6TryGivingUp(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoBravo06 struct {
	input string
}

func (tFuncY6 *testFuncDtoBravo06) tx6TryGivingUp(
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
			"Error= An Error Ocurred! Something Bad...\n"+
			"Maybe it is Divide By Zero!\n",
			ePrefix.String())
	}

	return err
}
