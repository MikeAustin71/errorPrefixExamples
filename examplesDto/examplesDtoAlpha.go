package examplesDto

import (
	"fmt"

	erPref "github.com/MikeAustin71/errpref"
)

// TestFuncDtoAlpha01 - This type is designed to call a series
// of methods using Error Prefix Data Transfer Objects otherwise
// known as 'ErrPrefixDto'.  For source code documentation on this
// type, reference:
//  https://pkg.go.dev/github.com/MikeAustin71/errpref#ErrPrefixDto
//
// This method chain will return an error on the last function
// call, if, and only if, parameter, 'returnError' is set to 'true'.
//
// The entry point for this method chain is:
//    TestFuncAlpha01.Tx1DoSomething()
//
// These examples showcase the 'ErrPrefixDto' type and the use of
// an empty interface to initialize a new instance of
// 'ErrPrefixDto'.
//
// For source code documentation on the 'ErrPrefixDto' type,
// reference:
//  https://pkg.go.dev/github.com/MikeAustin71/errpref#ErrPrefixDto
//
type TestFuncDtoAlpha01 struct {
	input string
}

func (tFuncAlpha01 *TestFuncDtoAlpha01) Tx1DoSomething(
	returnError bool,
	errorPrefix interface{}) error {

	var ePrefix *erPref.ErrPrefixDto
	var err error

	ePrefix,
		err = erPref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestFuncDtoAlpha01.Tx1DoSomething()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha2 := testFuncDtoAlpha02{}

	return tFuncAlpha2.tx2DoSomethingElse(
		returnError,
		ePrefix)
}

type testFuncDtoAlpha02 struct {
	input string
}

func (tFuncAlpha02 *testFuncDtoAlpha02) tx2DoSomethingElse(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoAlpha02.tx2DoSomethingElse()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha03 := testFuncDtoAlpha03{}

	err = tFuncAlpha03.tx3DoAnything(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoAlpha03 struct {
	input string
}

func (tFuncAlpha03 *testFuncDtoAlpha03) tx3DoAnything(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoAlpha03.tx3DoAnything()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha04 := testFuncDtoAlpha04{}

	err = tFuncAlpha04.tx4DoNothing(
		returnError,
		ePrefix)

	return err
}

type testFuncDtoAlpha04 struct {
	input string
}

func (tFuncAlpha04 *testFuncDtoAlpha04) tx4DoNothing(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoAlpha04.tx4DoNothing()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha05 := testFuncDtoAlpha05{}

	err = tFuncAlpha05.tx5DoSomethingBig(
		returnError,
		ePrefix.XCtx("A/B==4"))

	return err
}

type testFuncDtoAlpha05 struct {
	input string
}

func (tFuncAlpha05 *testFuncDtoAlpha05) tx5DoSomethingBig(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoAlpha05.tx5DoSomethingBig()",
		"")

	if err != nil {
		return err
	}

	tFuncAlpha06 := testFuncDtoAlpha06{}

	err = tFuncAlpha06.tx6GiveUp(
		returnError,
		ePrefix.XCtx(
			"A->B"))

	return err
}

type testFuncDtoAlpha06 struct {
	input string
}

func (tFuncAlpha06 *testFuncDtoAlpha06) tx6GiveUp(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoAlpha06.tx6GiveUp()",
		"A/B = C B==0")

	if err != nil {
		return err
	}

	if returnError {
		err = fmt.Errorf(
			"%v\n"+
				"Example Error: An Error Ocurred! Something bad.\n"+
				"Like Divide by Zero!\n",
			ePrefix.String())
	}

	return err
}
