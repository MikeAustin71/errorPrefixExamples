package main

import (
	"fmt"
	"github.com/MikeAustin71/errorPrefixExamples/testFunctions"
	"sync"
)

// main
// Call the various test methods in:
//  testFunctions.TestFuncsDto
//            and
//  testFunctions.TestFuncsStrings{}
//
func main() {

	mt := MainTest{}

	mt.mainTest007()
}

type MainTest struct {
}

// Runs a series of methods and prints the
// example error message returned by the last
// method in the series.
//
func (mt MainTest) mainTest001() {

	tFuncDto := testFunctions.TestFuncsDto{}

	err := tFuncDto.TestMethodSeries001(
		true, // returnExampleError
		true, // Print Example Error Message
		"main() - mainTest001")

	if err != nil {
		fmt.Printf("A REAL SYSTEM ERROR WAS RETURNED!\n"+
			"Error=\n%v\n\n",
			err.Error())
	} else {
		fmt.Printf("\n\n\nSUCCESSFUL COMPLETION!\n\n")
	}

}

// Runs a series of methods and prints the
// example error message returned by the last
// method in the series.
//
// In addition, this test specifies a maximum
// error prefix line length of 70
func (mt MainTest) mainTest002() {

	tFuncDto := testFunctions.TestFuncsDto{}

	err := tFuncDto.TestMethodSeries002(
		true, // returnExampleError
		true, // Print Example Error Message
		"main() - mainTest002")

	if err != nil {
		fmt.Printf("A REAL SYSTEM ERROR WAS RETURNED!\n"+
			"Error=\n%v\n\n",
			err.Error())
	} else {
		fmt.Printf("\n\n\nSUCCESSFUL COMPLETION!\n\n")
	}

}

// mainTest003 - This is designed as a "Concurrency" example.
//
func (mt MainTest) mainTest003() {

	tFuncDto := testFunctions.TestFuncsDto{}.Ptr()

	const NUMOfGOROUTINES = 200

	var c = make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(NUMOfGOROUTINES)

	for i := 1; i <= NUMOfGOROUTINES; i++ {

		go tFuncDto.TestMethodSeries003(c, &wg)
	}

	// Wait For It!
	go func() {

		wg.Wait()
		close(c)
	}()

	fmt.Printf("\n\nmainTest003()\n"+
		"Results Of %v x TestMethodSeries003\n\n",
		NUMOfGOROUTINES)

	var resultIdx = 1

	var resultStatus string
	var runSuccessful = true

	for n := range c {

		if n == 1 {
			resultStatus = "SUCCESS"
		} else {
			resultStatus = "FAILURE"
			runSuccessful = false
		}

		fmt.Printf("Run No: %3d   Result: %v   Status: %s\n",
			resultIdx,
			n,
			resultStatus)

		resultIdx++
	}

	if !runSuccessful {
		fmt.Printf("\n\n*** FAILURE ***\n" +
			"One or more methods returned a real error!\n")
	} else {
		fmt.Printf("\n\nSUCCESSFUL COMPLETION!\n")
	}

}

// mainTest004 - Executes Test TestFuncsDto.TestAlphaDto001()
// with a returned error message printed to the terminal.
//
// The printed example error messages demonstrates error prefix
// information displayed with the default Maximum Text Line
// Length of 40-characters.
//
func (mt MainTest) mainTest004() {

	tFuncDto := testFunctions.TestFuncsDto{}

	err := tFuncDto.TestAlphaDto001(
		true,
		true,
		"main()-mainTest004()")

	if err != nil {
		fmt.Printf("A REAL SYSTEM ERROR WAS RETURNED!\n"+
			"Error=\n%v\n\n",
			err.Error())
	} else {
		fmt.Printf("\n\n\nSUCCESSFUL COMPLETION!\n\n")
	}

}

// mainTest005 - Executes Test TestFuncsDto.TestAlphaDto001()
// with returned error message printed to the terminal.
//
// The maximum text line length is set to 70-characters.
//
func (mt MainTest) mainTest005() {

	tFuncDto := testFunctions.TestFuncsDto{}

	err := tFuncDto.TestBravoDto002(
		true,
		true,
		70,
		"main()-mainTest005()")

	if err != nil {
		fmt.Printf("A REAL SYSTEM ERROR WAS RETURNED!\n"+
			"Error=\n%v\n\n",
			err.Error())
	} else {
		fmt.Printf("\n\n\nSUCCESSFUL COMPLETION!\n\n")
	}

}

// This method calls TestFuncsDto.TestIBuilder004() which provides
// example usage of the IBuilderErrorPrefix interface.
//
// This method will display an example error message at the
// terminal.
//
func (mt MainTest) mainTest006() {

	tFuncDto := testFunctions.TestFuncsDto{}

	err := tFuncDto.TestIBuilder004(
		true,
		true)

	if err != nil {
		fmt.Printf("A REAL SYSTEM ERROR WAS RETURNED!\n"+
			"Error=\n%v\n\n",
			err.Error())
	} else {
		fmt.Printf("\n\n\nSUCCESSFUL COMPLETION!\n\n")
	}

}

// mainTest007() - Calls TestFuncsDto.TestDelimiters005() which
// demonstrates the variable input/output string delimiters
// feature.
//
// This method will display an example error message at the
// terminal.
//
func (mt MainTest) mainTest007() {

	tFuncDto := testFunctions.TestFuncsDto{}

	err := tFuncDto.TestDelimiters005(
		true,
		true)

	if err != nil {
		fmt.Printf("A REAL SYSTEM ERROR WAS RETURNED!\n"+
			"Error=\n%v\n\n",
			err.Error())
	} else {
		fmt.Printf("\n\n\nSUCCESSFUL COMPLETION!\n\n")
	}

}
