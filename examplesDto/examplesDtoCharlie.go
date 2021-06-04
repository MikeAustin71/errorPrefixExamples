package examplesDto

import (
	"fmt"
	erPref "github.com/MikeAustin71/errpref"
)

// TestFuncDtoCharlie01 - This type is designed to call a series
// of methods using Error Prefix Data Transfer Objects otherwise
// known as 'ErrPrefixDto'. For source code documentation on this
// type, reference:
//  https://pkg.go.dev/github.com/MikeAustin71/errpref#ErrPrefixDto
//
// This method chain will always return an error on the last function
// call.
//
// The entry point for this method chain is:
//    TestFuncDtoCharlie01.Tx1DoStuff()
//
// These examples showcase the 'ErrPrefixDto' type and the use of
// an empty interface to initialize a new instance of
// 'ErrPrefixDto'.
//
type TestFuncDtoCharlie01 struct {
	input string
}

func (tFuncDtoCharlie01 *TestFuncDtoCharlie01) Tx1DoStuff(
	returnError bool,
	errorPrefix interface{}) error {

	var ePrefix *erPref.ErrPrefixDto
	var err error

	ePrefix,
		err = erPref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestFuncDtoAlpha01.Tx1DoStuff()",
		"X->Y")

	if err != nil {
		return err
	}

	tFuncCharlie02 := testFuncDtoCharlie02{}

	return tFuncCharlie02.Tx2DoMoreStuff(
		returnError,
		ePrefix)
}

type testFuncDtoCharlie02 struct {
	input string
}

func (tFuncCharlie02 *testFuncDtoCharlie02) Tx2DoMoreStuff(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoCharlie02.Tx2DoMoreStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie03 := testFuncDtoCharlie03{}

	return tFuncCharlie03.Tx3DoLessStuff(
		returnError,
		ePrefix.XCtx(
			"B->C"))
}

type testFuncDtoCharlie03 struct {
	input string
}

func (tFuncCharlie03 *testFuncDtoCharlie03) Tx3DoLessStuff(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoCharlie03.Tx3DoLessStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie04 := testFuncDtoCharlie04{}

	return tFuncCharlie04.Tx4DoFunStuff(
		returnError,
		ePrefix)
}

type testFuncDtoCharlie04 struct {
	input string
}

func (tFuncCharlie04 *testFuncDtoCharlie04) Tx4DoFunStuff(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoCharlie04.Tx4DoFunStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie05 := testFuncDtoCharlie05{}

	return tFuncCharlie05.Tx5DoExcitingStuff(
		returnError,
		ePrefix)

}

type testFuncDtoCharlie05 struct {
	input string
}

func (tFuncCharlie05 *testFuncDtoCharlie05) Tx5DoExcitingStuff(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoCharlie05.Tx5DoExcitingStuff()",
		"")

	if err != nil {
		return err
	}

	tFuncCharlie06 := testFuncDtoCharlie06{}

	return tFuncCharlie06.Tx6DoSpaceStuff(
		returnError,
		ePrefix.XCtx(
			"X*Y"))

}

type testFuncDtoCharlie06 struct {
	input string
}

func (tFuncCharlie06 *testFuncDtoCharlie06) Tx6DoSpaceStuff(
	returnError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoCharlie06.Tx6DoSpaceStuff()",
		"Error Context: Asteroid Collision!")

	if err != nil {
		return err
	}

	if returnError {
		err = fmt.Errorf("%v\n"+
			"Example Error: %v",
			ePrefix.String(),
			"Real bad error! Something hit the space ship!\n")

	}

	return err
}
