package main

import "github.com/MikeAustin71/errorPrefixExamples/testFunctions"

// main
// Call the various test methods in:
//  testFunctions.TestFuncsDto
//            and
//  testFunctions.TestFuncsStrings{}
//
func main() {

	tFuncDto := testFunctions.TestFuncsDto{}

	tFuncDto.TestMethodSeries001(
		"main()")

}
