package testFunctions

import (
	"fmt"
	exDto "github.com/MikeAustin71/errorPrefixExamples/examplesDto"
	erPref "github.com/MikeAustin71/errpref"
)

// TestFuncsStrings - This type includes methods designed to test
// the string formatting capabilities of type, 'ErrPref'.
//
// 'ErrPref' is a simple, light-weight type used to receive error prefix
// strings, format them and return the formatted error prefix string immediately
// to the caller. No error prefix information is retained or maintained in the
// the 'ErrPref' object. It is simply used to format strings.
//
type TestFuncsStrings struct{}

func (tFuncStrs TestFuncsStrings) TestMethodSeries001(callingMethodName string) {

	errPref := erPref.ErrPref{}

	errPrefStr := errPref.EPref(
		callingMethodName,
		"TestFuncsStrings.TestMethodSeries001()")

	tAlpha := exDto.TestFuncDtoAlpha01{}

	err := tAlpha.Tx1DoSomething(
		false, // Do NOT return error
		errPrefStr)

	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	tBravo := exDto.TestFuncDtoBravo01{}

	err = tBravo.Tx1DoSomethingSpecial(
		false, // Do NOT return an error
		errPrefStr)

	tCharlie := exDto.TestFuncDtoCharlie01{}

	// This returns an error
	err = tCharlie.Tx1DoStuff(
		errPrefStr)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	} else {
		fmt.Printf("Expected an error return from tCharlie.Tx1DoStuff().\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!!\n" +
			"THAT IS A REAL ERROR CONDITION!!!!\n")
	}

}
