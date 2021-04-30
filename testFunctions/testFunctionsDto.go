package testFunctions

import (
	"fmt"
	exDto "github.com/MikeAustin71/errorPrefixExamples/examplesDto"
	erPref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
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
// If the input parameter 'returnExampleError' is set to 'true' and
// the input parameter 'showExampleErrorMsg' is set to true, this
// method will print the example error message to the terminal.
//
func (tFuncDto TestFuncsDto) TestAlphaDto001(
	returnExampleError bool,
	showExampleErrorMsg bool,
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
		showExampleErrorMsg == true &&
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
// If the input parameter 'returnExampleError' is set to 'true' and
// the input parameter 'showExampleErrorMsg' is set to true, this
// method will print the example error message to the terminal.
//
// Input parameter 'maxLineLength' is used to control the maximum
// line length of the error prefix information printed to the
// terminal. If this parameter is set to a value of minus one (-1),
// the default maximum line length will be applied. The current
// default maximum line length is 40-characters.
//
func (tFuncDto TestFuncsDto) TestBravoDto002(
	returnExampleError bool,
	showExampleErrorMsg bool,
	maxLineLength int,
	callingMethodName string) error {

	ePrefDto := erPref.ErrPrefixDto{}

	ePref := ePrefDto.NewEPrefOld(callingMethodName)

	ePref.SetEPref("TestFuncsDto.TestAlphaDto002()")

	ePref.SetMaxTextLineLen(maxLineLength)

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
		showExampleErrorMsg == true &&
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

	var millisecondTimeDelay time.Duration = 50

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

	time.Sleep(time.Millisecond * millisecondTimeDelay)

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

	time.Sleep(time.Millisecond * millisecondTimeDelay)

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

	time.Sleep(time.Millisecond * millisecondTimeDelay)

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

	time.Sleep(time.Millisecond * millisecondTimeDelay)

	c <- 1
	return
}

// TestIBuilder004 - Demonstrates interoperability of the
// IBuilderErrorPrefix interface declared in the 'errpref'
// software package:
//    "github.com/MikeAustin71/errpref"
//
//       type IBuilderErrorPrefix interface {
//         GetEPrefStrings() [][2]string
//
//         SetEPrefStrings(twoDStrArray [][2]string)
//
//         String() string
//       }
//
// Type TestIBuilderErrPref implements the IBuilderErrorPrefix
// interface.
//
//
// This method calls the Delta method chains with an ErrPrefixDto
// created from an IBuilderErrorPrefix object.
//
// An example error is only triggered on the last method in the
// chain if input parameter, 'returnExampleError' is set to 'true'.
//
// If the input parameter 'returnExampleError' is set to 'true' and
// the input parameter 'showExampleErrorMsg' is set to true, this
// method will print the example error message to the terminal.
//
// This method sets the maximum line length for error prefix
// strings to 60-characters and does NOT implement a left margin.
//
//
func (tFuncDto *TestFuncsDto) TestIBuilder004(
	returnExampleError bool,
	showExampleErrorMsg bool) error {

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	iBuilder := exDto.TestIBuilderErrPref{}

	iBuilder.SetEPrefStrings(twoDSlice)

	ePrefDto := erPref.ErrPrefixDto{}.New()

	ePrefDto.SetMaxTextLineLen(60)

	thisFunc := "TestFuncsDto.TestIBuilder004()"

	err := ePrefDto.SetIBuilder(
		&iBuilder,
		thisFunc)

	if err != nil {
		return err
	}

	ePrefDto.SetEPref(thisFunc)

	tDelta := exDto.TestFuncDtoDelta01{}

	// This may return an example error depending
	// on the value of 'returnExampleError'.
	err = tDelta.Tx1DoGreatThings(
		returnExampleError,
		ePrefDto)

	if err != nil {

		testMsg := fmt.Sprintf("%v",
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

func (tFuncDto *TestFuncsDto) TestDelimiters005(
	returnExampleError bool,
	showExampleErrorMsg bool) error {

	thisFunc := "TestFuncsDto.TestDelimiters005() "

	oldErrPrefix := "Tx1.Something()**" +
		"Tx2.SomethingElse()**" +
		"Tx3.DoSomething()**" +
		"Tx4()**" +
		"Tx5()**" +
		"Tx6.DoSomethingElse()"

	inputDelimiters,
		err :=
		erPref.ErrPrefixDelimiters{}.New(
			"\n",
			"**",
			"\n",
			"--",
			thisFunc)

	if err != nil {
		return err
	}

	var outputDelimiters erPref.ErrPrefixDelimiters

	outputDelimiters,
		err = erPref.ErrPrefixDelimiters{}.New(
		"\n * ",
		" - ",
		"\n  ",
		" $ ",
		thisFunc)

	if err != nil {
		return err
	}

	tfEcho := exDto.TestFuncDtoEcho01{}

	err =
		tfEcho.Tx1CreateSomething(
			returnExampleError,
			oldErrPrefix,
			inputDelimiters,
			outputDelimiters)

	if err != nil {

		testMsg := fmt.Sprintf("%v",
			err.Error())

		if !strings.Contains(testMsg, "Example Error") {

			return err
		}

	}

	if returnExampleError == true &&
		showExampleErrorMsg == true &&
		err != nil {

		fmt.Printf("\nORIGINAL Error Prefix Info in String Format\n")
		fmt.Printf("%v\n\n\n",
			oldErrPrefix)
		fmt.Printf("PRINTING EXAMPLE ERROR MESSAGE\n\n")

		fmt.Printf("%v",
			err.Error())
	}

	return nil
}
