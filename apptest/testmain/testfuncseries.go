package testmain

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
)

type testFunc1 struct {
	input string
}

func (tFunc1 *testFunc1) Tx1DoSomething() error {

	tFunc2 := testFunc2{}

	ePrefix := errpref.ErrPrefixDto{}

	ePrefix.SetEPref("Tx1DoSomething()")

	err := tFunc2.Tx2DoSomethingElse(ePrefix)

	return err
}

type testFunc2 struct {
	input string
}

func (tFunc2 *testFunc2) Tx2DoSomethingElse(
	ePrefix errpref.ErrPrefixDto) error {

	ePrefix.SetEPref("Tx2DoSomethingElse()")

	tFunc3 := testFunc3{}

	err := tFunc3.Tx3DoAnything(ePrefix)

	return err
}

type testFunc3 struct {
	input string
}

func (tFunc3 *testFunc3) Tx3DoAnything(
	ePrefix errpref.ErrPrefixDto) error {

	ePrefix.SetEPref("Tx3DoAnything()")

	tFunc4 := testFunc4{}

	err := tFunc4.Tx4DoNothing(ePrefix)

	return err
}

type testFunc4 struct {
	input string
}

func (tFunc4 *testFunc4) Tx4DoNothing(
	ePrefix errpref.ErrPrefixDto) error {

	ePrefix.SetEPref("Tx4DoNothing()")

	tFunc5 := testFunc5{}

	ePrefix.SetCtx("A/B==4")

	err := tFunc5.Tx5DoSomethingBig(ePrefix)

	return err
}

type testFunc5 struct {
	input string
}

func (tFunc5 *testFunc5) Tx5DoSomethingBig(
	ePrefix errpref.ErrPrefixDto) error {

	ePrefix.SetEPref("Tx5DoSomethingBig()")

	tFunc6 := testFunc6{}

	ePrefix.SetCtx("A->B")

	err := tFunc6.Tx6GiveUp(ePrefix)

	return err
}

type testFunc6 struct {
	input string
}

func (tFunc6 *testFunc6) Tx6GiveUp(
	ePrefix errpref.ErrPrefixDto) error {

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	var err = fmt.Errorf("%s\n"+
		"Error: Divide By Zero!\n",
		ePrefix)

	return err
}
