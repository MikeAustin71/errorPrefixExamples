package examplesDto

import (
	"fmt"
	erPref "github.com/MikeAustin71/errpref"
)

// TestFuncDtoDelta01 - This type is designed to call a series
// of methods using Error Prefix Data Transfer Objects otherwise
// known as 'ErrPrefixDto'. For source code documentation on this
// type, reference:
//  https://pkg.go.dev/github.com/MikeAustin71/errpref#ErrPrefixDto
//
// This method chain will always return an error on the last function
// call.
//
// The entry point for this method chain is:
//    TestFuncDtoDelta01.Tx1DoGreatThings()
//
// These examples showcase the 'ErrPrefixDto' type and the use of
// an empty interface to initialize a new instance of
// 'ErrPrefixDto'.
//
type TestFuncDtoDelta01 struct {
	input string
}

func (tFuncDtoDelta01 *TestFuncDtoDelta01) Tx1DoGreatThings(
	returnError bool,
	errorPrefix interface{}) error {

	var ePrefix *erPref.ErrPrefixDto
	var err error

	ePrefix,
		err = erPref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestFuncDtoDelta01.Tx1DoGreatThings()",
		"X->Y")

	if err != nil {
		return err
	}

	tFuncDelta02 := testFuncDtoDelta02{}

	return tFuncDelta02.Tx2DoSomeThings(
		returnError,
		ePrefix)
}

type testFuncDtoDelta02 struct {
	input string
}

func (tFuncDelta02 *testFuncDtoDelta02) Tx2DoSomeThings(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoDelta02.Tx2DoSomeThings()",
		"")

	if err != nil {
		return err
	}

	tFuncDelta03 := testFuncDtoDelta03{}

	return tFuncDelta03.Tx3DoFewerThings(
		returnError,
		ePrefix.XCtx(
			"B->C"))
}

type testFuncDtoDelta03 struct {
	input string
}

func (tFuncDelta03 *testFuncDtoDelta03) Tx3DoFewerThings(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoDelta03.Tx3DoFewerThings()",
		"")

	if err != nil {
		return err
	}

	tFuncDelta04 := testFuncDtoDelta04{}

	return tFuncDelta04.Tx4DoFunThings(
		returnError,
		ePrefix)
}

type testFuncDtoDelta04 struct {
	input string
}

func (tFuncDelta04 *testFuncDtoDelta04) Tx4DoFunThings(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoDelta04.Tx4DoFunThings()",
		"")

	if err != nil {
		return err
	}

	tFuncDelta05 := testFuncDtoDelta05{}

	return tFuncDelta05.Tx5DoExcitingThings(
		returnError,
		ePrefix)

}

type testFuncDtoDelta05 struct {
	input string
}

func (tFuncDelta05 *testFuncDtoDelta05) Tx5DoExcitingThings(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoDelta05.Tx5DoExcitingThings()",
		"")

	if err != nil {
		return err
	}

	tFuncDelta06 := testFuncDtoDelta06{}

	return tFuncDelta06.Tx6DoUnbelievableThings(
		returnError,
		ePrefix.XCtx(
			"X*Y"))

}

type testFuncDtoDelta06 struct {
	input string
}

func (tFuncDelta06 *testFuncDtoDelta06) Tx6DoUnbelievableThings(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoDelta06.Tx6DoUnbelievableThings()",
		"Error Context: Looking For Higgs Boson!")

	if err != nil {
		return err
	}

	if returnError {
		err = fmt.Errorf("%v\n"+
			"Example Error: %v",
			ePrefix.String(),
			"Real bad error! Yikes! We created a local black hole by mistake!\n")

	}

	return err
}
