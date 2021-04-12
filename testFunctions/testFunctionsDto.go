package testFunctions

import (
	"fmt"
	exDto "github.com/MikeAustin71/errorPrefixExamples/examplesDto"
	erPref "github.com/MikeAustin71/errpref"
)

type TestFuncsDto struct{}

// TestAlphaDto001
// This method calls the Alpha Series method chain with
// a two-dimensional array of strings.
func (tFuncDto TestFuncsDto) TestAlphaDto001(
	callingMethodName string,
	returnError bool) {

	tAlpha := exDto.TestFuncDtoAlpha01{}

	twoDStrAry := make([][2]string, 2)

	twoDStrAry[0][0] = callingMethodName
	twoDStrAry[0][1] = ""

	twoDStrAry[1][0] = "TestFuncsDto.TestAlphaDto001() "
	twoDStrAry[1][1] = "A->B" // Error Context Information

	err := tAlpha.Tx1DoSomething(
		returnError,
		twoDStrAry)

	if err != nil {
		fmt.Printf("%v", err.Error())
	}

}

// TestAlphaDto002
// This method calls the Alpha Series method chain with
// an ErrPrefixDto object.
func (tFuncDto TestFuncsDto) TestAlphaDto002(
	callingMethodName string,
	returnError bool) {

	ePrefDto := erPref.ErrPrefixDto{}

	ePref := ePrefDto.NewEPrefOld(callingMethodName)

	ePref.SetEPref("TestFuncsDto.TestAlphaDto002()")

	tAlpha := exDto.TestFuncDtoAlpha01{}

	err := tAlpha.Tx1DoSomething(
		returnError,
		ePref)

	if err != nil {
		fmt.Printf("%v", err.Error())
	}

}

// TestMethodSeries001
// This method calls the Alpha, Bravo and Charlie method chains
// with an ErrPrefixDto object. An error is only triggered on the
// last method in the 'Charlie' method chain.
func (tFuncDto TestFuncsDto) TestMethodSeries001(
	callingMethodName string) {

	ePrefDto := erPref.ErrPrefixDto{}

	ePref,
		err := ePrefDto.NewIEmpty(
		callingMethodName,
		"TestMethodSeries001",
		"Alpha-Bravo-Charlie Test")

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	tAlpha := exDto.TestFuncDtoAlpha01{}

	err = tAlpha.Tx1DoSomething(
		false, // Do NOT return error
		ePref)

	if err != nil {
		fmt.Printf("%v", err.Error())
	}

	tBravo := exDto.TestFuncDtoBravo01{}

	err = tBravo.Tx1DoSomethingSpecial(
		false, // Do NOT return an error
		ePref)

	tCharlie := exDto.TestFuncDtoCharlie01{}

	// This returns an error
	err = tCharlie.Tx1DoStuff(
		ePref)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	} else {
		fmt.Printf("Expected an error return from tCharlie.Tx1DoStuff().\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!!\n" +
			"THAT IS A REAL ERROR CONDITION!!!!\n")
	}

}
