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
	mainTest003()
}

// Runs a series of methods and prints the
// example error message returned by the last
// method in the series.
//
func mainTest001() {

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
func mainTest002() {

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
func mainTest003() {

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
