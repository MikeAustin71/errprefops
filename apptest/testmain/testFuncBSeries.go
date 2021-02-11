package testmain

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
)

type TestFuncB1 struct {
	input string
}

func (tFuncB1 *TestFuncB1) Tx1DoSomething() error {

	tFuncB2 := testFuncB2{}

	ePrefix := errpref.ErrPrefixDto{}

	err := tFuncB2.Tx2DoSomethingElse(
		ePrefix.XEPref("Tx1DoSomething()"))

	return err
}

type testFuncB2 struct {
	input string
}

func (tFuncB2 *testFuncB2) Tx2DoSomethingElse(
	ePrefix *errpref.ErrPrefixDto) error {

	tFuncB3 := testFuncB3{}

	err := tFuncB3.Tx3DoAnything(
		ePrefix.XEPref("Tx2DoSomethingElse()"))

	return err
}

type testFuncB3 struct {
	input string
}

func (tFuncB3 *testFuncB3) Tx3DoAnything(
	ePrefix *errpref.ErrPrefixDto) error {

	tFuncB4 := testFuncB4{}

	err := tFuncB4.Tx4DoNothing(
		ePrefix.XEPref("Tx3DoAnything()"))

	return err
}

type testFuncB4 struct {
	input string
}

func (tFuncB4 *testFuncB4) Tx4DoNothing(
	ePrefix *errpref.ErrPrefixDto) error {

	tFuncB5 := testFuncB5{}

	err := tFuncB5.Tx5DoSomethingBig(
		ePrefix.XEPrefCtx(
			"Tx4DoNothing()",
			"A/B==4"))

	return err
}

type testFuncB5 struct {
	input string
}

func (tFuncB5 *testFuncB5) Tx5DoSomethingBig(
	ePrefix *errpref.ErrPrefixDto) error {

	tFuncB6 := testFuncB6{}

	err := tFuncB6.Tx6GiveUp(
		ePrefix.XEPrefCtx(
			"Tx5DoSomethingBig()",
			"A->B"))

	return err
}

type testFuncB6 struct {
	input string
}

func (tFuncB6 *testFuncB6) Tx6GiveUp(
	ePrefix *errpref.ErrPrefixDto) error {

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	var err = fmt.Errorf("%v\n"+
		"Error: Divide By Zero!\n",
		ePrefix.String())

	return err
}
