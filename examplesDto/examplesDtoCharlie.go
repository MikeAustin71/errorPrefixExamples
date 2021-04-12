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

	ePrefix.SetEPrefCtx(
		"testFuncCharlie01."+
			"Tx1DoStuff()",
		"X->Y")

	tFuncCharlie02 := testFuncDtoCharlie02{}

	return tFuncCharlie02.Tx2DoMoreStuff(
		ePrefix)
}

type testFuncDtoCharlie02 struct {
	input string
}

func (tFuncCharlie02 *testFuncDtoCharlie02) Tx2DoMoreStuff(
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoCharlie02." +
			"Tx2DoMoreStuff()")

	tFuncCharlie03 := testFuncDtoCharlie03{}

	return tFuncCharlie03.Tx3DoLessStuff(
		ePrefix.XCtx(
			"B->C"))
}

type testFuncDtoCharlie03 struct {
	input string
}

func (tFuncCharlie03 *testFuncDtoCharlie03) Tx3DoLessStuff(
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoCharlie03." +
			"Tx3DoLessStuff()")

	tFuncCharlie04 := testFuncDtoCharlie04{}

	return tFuncCharlie04.Tx4DoFunStuff(
		ePrefix)
}

type testFuncDtoCharlie04 struct {
	input string
}

func (tFuncCharlie04 *testFuncDtoCharlie04) Tx4DoFunStuff(
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoCharlie04." +
			"Tx4DoFunStuff()")

	tFuncCharlie05 := testFuncDtoCharlie05{}

	return tFuncCharlie05.Tx5DoExcitingStuff(
		ePrefix)

}

type testFuncDtoCharlie05 struct {
	input string
}

func (tFuncCharlie05 *testFuncDtoCharlie05) Tx5DoExcitingStuff(
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncDtoCharlie05." +
			"Tx5DoExcitingStuff()")

	tFuncCharlie06 := testFuncDtoCharlie06{}

	return tFuncCharlie06.Tx6DoSpaceStuff(
		ePrefix.XCtx(
			"X*Y"))

}

type testFuncDtoCharlie06 struct {
	input string
}

func (tFuncCharlie06 *testFuncDtoCharlie06) Tx6DoSpaceStuff(
	ePrefix *erPref.ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = erPref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPrefCtx(
		"testFuncDtoCharlie06."+
			"Tx6DoSpaceStuff()",
		"Error Context: Asteroid Collision!")

	return fmt.Errorf("%v\n"+
		"Error= %v",
		ePrefix.String(),
		"Real bad error! Something hit the space ship!\n")
}
