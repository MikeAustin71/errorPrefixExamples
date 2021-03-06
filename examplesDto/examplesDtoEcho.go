package examplesDto

import (
	"fmt"
	erPref "github.com/MikeAustin71/errpref"
)

type TestFuncDtoEcho01 struct {
	input string
}

func (tFuncDtoEcho01 *TestFuncDtoEcho01) Tx1CreateSomething(
	returnExampleError bool,
	oldErrPrefix string,
	inputDelimiters erPref.ErrPrefixDelimiters,
	outputDelimiters erPref.ErrPrefixDelimiters) error {

	var ePrefix erPref.ErrPrefixDto
	var err error
	thisFunc := "TestFuncDtoEcho01.Tx1CreateSomething()"

	ePrefix,
		err = erPref.ErrPrefixDto{}.NewFromStrings(
		oldErrPrefix,
		thisFunc,
		"Entry Point",
		inputDelimiters,
		outputDelimiters,
		"")

	if err != nil {
		return err
	}

	tFuncEcho02 := testFuncDtoEcho02{}

	return tFuncEcho02.Tx2DoResearch(
		returnExampleError,
		&ePrefix)
}

type testFuncDtoEcho02 struct {
	input string
}

func (tFuncEcho02 *testFuncDtoEcho02) Tx2DoResearch(
	returnExampleError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoEcho02.Tx2DoResearch()",
		"")

	if err != nil {
		return err
	}

	tFuncDelta03 := testFuncDtoEcho03{}

	return tFuncDelta03.Tx3DoPlanning(
		returnExampleError,
		ePrefix)
}

type testFuncDtoEcho03 struct {
	input string
}

func (tFuncEcho03 *testFuncDtoEcho03) Tx3DoPlanning(
	returnExampleError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoEcho03.Tx3DoPlanning()",
		"")

	if err != nil {
		return err
	}

	tFuncEcho04 := testFuncDtoEcho04{}

	return tFuncEcho04.Tx4GetResources(
		returnExampleError,
		ePrefix)
}

type testFuncDtoEcho04 struct {
	input string
}

func (tFuncEcho04 *testFuncDtoEcho04) Tx4GetResources(
	returnExampleError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoEcho04.Tx4GetResources()",
		"")

	if err != nil {
		return err
	}

	tFuncEcho05 := testFuncDtoEcho05{}

	return tFuncEcho05.Tx5AssembleTeam(
		returnExampleError,
		ePrefix)

}

type testFuncDtoEcho05 struct {
	input string
}

func (tFuncEcho05 *testFuncDtoEcho05) Tx5AssembleTeam(
	returnExampleError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoEcho05.Tx5AssembleTeam()",
		"")

	if err != nil {
		return err
	}

	tFuncEcho06 := testFuncDtoEcho06{}

	return tFuncEcho06.Tx6DoTheWork(
		returnExampleError,
		ePrefix.XCpy(
			"A+B+C+D"))

}

type testFuncDtoEcho06 struct {
	input string
}

func (tFuncEcho06 *testFuncDtoEcho06) Tx6DoTheWork(
	returnExampleError bool,
	ePrefDto *erPref.ErrPrefixDto) error {

	ePrefix,
		err := erPref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"testFuncDtoEcho06.Tx6DoTheWork()",
		"Must Finish in 3-Days!")

	if err != nil {
		return err
	}

	if returnExampleError {
		err = fmt.Errorf("%v\n"+
			"Example Error: %v",
			ePrefix.String(),
			"Can't finish by deadline. Need more time!\n")
	}

	return err
}
