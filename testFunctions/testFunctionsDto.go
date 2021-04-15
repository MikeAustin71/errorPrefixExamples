package testFunctions

import (
	"fmt"
	exDto "github.com/MikeAustin71/errorPrefixExamples/examplesDto"
	erPref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type TestFuncsDto struct{}

// Ptr - Returns a pointer to a new instance of
// TestFuncsDto
func (tFuncDto TestFuncsDto) Ptr() *TestFuncsDto {

	return &TestFuncsDto{}
}

// TestAlphaDto001
// This method calls the Alpha Series method chain with
// a two-dimensional array of strings.
//
// An example error is only triggered on the last method in the
// chain if input parameter, 'returnExampleError' is set to 'true'.
//
// If the input parameter 'returnExampleError' is set to 'true',
// this method will print the example error message to the
// terminal.
//
func (tFuncDto TestFuncsDto) TestAlphaDto001(
	returnExampleError bool,
	callingMethodName string) error {

	tAlpha := exDto.TestFuncDtoAlpha01{}

	twoDStrAry := make([][2]string, 2)

	twoDStrAry[0][0] = callingMethodName
	twoDStrAry[0][1] = ""

	twoDStrAry[1][0] = "TestFuncsDto.TestAlphaDto001() "
	twoDStrAry[1][1] = "A->B" // Error Context Information

	var err error
	var testMsg string

	err = tAlpha.Tx1DoSomething(
		returnExampleError,
		twoDStrAry)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

	}

	if returnExampleError == true &&
		err != nil {

		fmt.Printf("\nPRINTING EXAMPLE ERROR MESSAGE\n\n")

		fmt.Printf("%v",
			err.Error())
	}

	return nil
}

// TestBravoDto002
// This method calls the Bravo Series method chain with
// an ErrPrefixDto object.
//
// An example error is only triggered on the last method in the
// chain if input parameter, 'returnExampleError' is set to 'true'.
//
// If the input parameter 'returnExampleError' is set to 'true',
// this method will print the example error message to the
// terminal.
//
func (tFuncDto TestFuncsDto) TestBravoDto002(
	returnExampleError bool,
	callingMethodName string) error {

	ePrefDto := erPref.ErrPrefixDto{}

	ePref := ePrefDto.NewEPrefOld(callingMethodName)

	ePref.SetEPref("TestFuncsDto.TestAlphaDto002()")

	var testMsg string
	var err error

	tAlpha := exDto.TestFuncDtoAlpha01{}

	err = tAlpha.Tx1DoSomething(
		returnExampleError,
		ePref)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

	}

	if returnExampleError == true &&
		err != nil {

		fmt.Printf("\nPRINTING EXAMPLE ERROR MESSAGE\n\n")

		fmt.Printf("%v",
			err.Error())
	}

	return nil
}

// TestMethodSeries001 - This method calls the Alpha, Bravo and
// Charlie method chains with an ErrPrefixDto object.
//
// An example error is only triggered on the last method in the
// chain if input parameter, 'returnExampleError' is set to 'true'.
//
// If the input parameter 'returnExampleError' is set to 'true' and
// the input parameter 'showExampleErrorMsg' is set to true, this
// method will print the example error message to the terminal.
//
// This method relies on default maximum line length for error prefix
// strings and does NOT implement a left margin.
//
func (tFuncDto TestFuncsDto) TestMethodSeries001(
	returnExampleError bool,
	showExampleErrorMsg bool,
	callingMethodName string) error {

	ePrefDto := erPref.ErrPrefixDto{}

	ePrefDto.SetEPrefOld(callingMethodName)

	ePrefDto.SetEPrefCtx("TestMethodSeries001",
		"Alpha-Bravo-Charlie-Delta Test")

	tAlpha := exDto.TestFuncDtoAlpha01{}
	var err error

	err = tAlpha.Tx1DoSomething(
		false, // Do NOT return error
		ePrefDto)

	var testMsg string

	if err != nil {
		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

		err = nil
	}

	tBravo := exDto.TestFuncDtoBravo01{}

	err = tBravo.Tx1TrySomethingSpecial(
		false, // Do NOT return an error
		ePrefDto)

	if err != nil {
		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

		err = nil
	}

	tCharlie := exDto.TestFuncDtoCharlie01{}

	err = tCharlie.Tx1DoStuff(
		false, // Do NOT return an error
		ePrefDto)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

		err = nil
	}

	tDelta := exDto.TestFuncDtoDelta01{}

	// This may return an example error depending
	// on the value of 'returnExampleError'.
	err = tDelta.Tx1DoGreatThings(
		returnExampleError,
		ePrefDto)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

	}

	if returnExampleError == true &&
		showExampleErrorMsg == true &&
		err != nil {

		fmt.Printf("\nPRINTING EXAMPLE ERROR MESSAGE\n\n")

		fmt.Printf("%v",
			err.Error())
	}

	return nil
}

// TestMethodSeries002 - This method calls the Alpha, Bravo and
// Charlie method chains with an ErrPrefixDto object.
//
// An example error is only triggered on the last method in the
// chain if input parameter, 'returnExampleError' is set to 'true'.
//
// If the input parameter 'returnExampleError' is set to 'true' and
// the input parameter 'showExampleErrorMsg' is set to true, this
// method will print the example error message to the terminal.
//
// This method relies on default maximum line length for error prefix
// strings and does NOT implement a left margin.
//
func (tFuncDto TestFuncsDto) TestMethodSeries002(
	returnExampleError bool,
	showExampleErrorMsg bool,
	callingMethodName string) error {

	ePrefDto := erPref.ErrPrefixDto{}

	ePrefDto.SetEPrefOld(callingMethodName)

	ePrefDto.SetEPrefCtx("TestMethodSeries002",
		"Alpha-Bravo-Charlie-Delta Test - Max Line Len=70")

	ePrefDto.SetMaxTextLineLen(70)

	tAlpha := exDto.TestFuncDtoAlpha01{}
	var err error

	err = tAlpha.Tx1DoSomething(
		false, // Do NOT return error
		ePrefDto)

	var testMsg string

	if err != nil {
		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

		err = nil
	}

	tBravo := exDto.TestFuncDtoBravo01{}

	err = tBravo.Tx1TrySomethingSpecial(
		false, // Do NOT return an error
		ePrefDto)

	if err != nil {
		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

		err = nil
	}

	tCharlie := exDto.TestFuncDtoCharlie01{}

	err = tCharlie.Tx1DoStuff(
		false, // Do NOT return an error
		ePrefDto)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

		err = nil
	}

	tDelta := exDto.TestFuncDtoDelta01{}

	// This may return an example error depending
	// on the value of 'returnExampleError'.
	err = tDelta.Tx1DoGreatThings(
		returnExampleError,
		ePrefDto)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

	}

	if returnExampleError == true &&
		showExampleErrorMsg == true &&
		err != nil {

		fmt.Printf("\nPRINTING EXAMPLE ERROR MESSAGE\n\n")

		fmt.Printf("%v",
			err.Error())
	}

	return nil
}

// TestMethodSeries003 - This method calls the Alpha, Bravo and
// Charlie method chains with an ErrPrefixDto object.
//
// An example error is only triggered on the last method in the
// chain if input parameter, 'returnExampleError' is set to 'true'.
//
// This method will NOT print the example error message to the
// terminal.
//
// THIS METHOD IS DESIGNED TO BE USED IN A CONCURRENCY EXAMPLE!
//
// See mainTest003() in app/main.go
//
func (tFuncDto *TestFuncsDto) TestMethodSeries003(c chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	ePrefDto := erPref.ErrPrefixDto{}

	ePrefDto.SetEPrefCtx("TestMethodSeries003",
		"Alpha-Bravo-Charlie-Delta Test")

	tAlpha := exDto.TestFuncDtoAlpha01{}
	var err error

	err = tAlpha.Tx1DoSomething(
		false, // Do NOT return error
		ePrefDto)

	var testMsg string

	if err != nil {
		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {
			fmt.Printf("REAL SYSTEM ERROR\n"+
				"%v\n",
				err.Error())

			c <- -101
			return
		}

		err = nil
	}

	tBravo := exDto.TestFuncDtoBravo01{}

	err = tBravo.Tx1TrySomethingSpecial(
		false, // Do NOT return an error
		ePrefDto)

	if err != nil {
		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			c <- -102
			return
		}

		err = nil
	}

	tCharlie := exDto.TestFuncDtoCharlie01{}

	err = tCharlie.Tx1DoStuff(
		false, // Do NOT return an error
		ePrefDto)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			c <- -103
			return
		}

		err = nil
	}

	tDelta := exDto.TestFuncDtoDelta01{}

	// This may return an example error depending
	// on the value of 'returnExampleError'.
	err = tDelta.Tx1DoGreatThings(
		true,
		ePrefDto)

	if err != nil {

		testMsg = fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			c <- -104
			return
		}

	}

	c <- 1
	return
}
