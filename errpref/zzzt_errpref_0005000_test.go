package errpref

import (
	"fmt"
	"testing"
)

type testFuncAlpha01 struct {
	input string
}

func (tFuncAlpha01 *testFuncAlpha01) Tx1DoSomething(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha01." +
			"Tx1DoSomething()")

	tFuncAlpha2 := testFuncAlpha02{}

	return tFuncAlpha2.Tx2DoSomethingElse(
		ePrefix)
}

type testFuncAlpha02 struct {
	input string
}

func (tFuncAlpha02 *testFuncAlpha02) Tx2DoSomethingElse(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha02." +
			"Tx2DoSomethingElse()")

	tFuncAlpha03 := testFuncAlpha03{}

	err := tFuncAlpha03.Tx3DoAnything(
		ePrefix)

	return err
}

type testFuncAlpha03 struct {
	input string
}

func (tFuncAlpha03 *testFuncAlpha03) Tx3DoAnything(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha03." +
			"Tx3DoAnything()")

	tFuncAlpha04 := testFuncAlpha04{}

	err := tFuncAlpha04.Tx4DoNothing(
		ePrefix)

	return err
}

type testFuncAlpha04 struct {
	input string
}

func (tFuncAlpha04 *testFuncAlpha04) Tx4DoNothing(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha04." +
			"Tx4DoNothing()")

	tFuncAlpha05 := testFuncAlpha05{}

	err := tFuncAlpha05.Tx5DoSomethingBig(
		ePrefix.XEPrefCtx(
			"Tx4DoNothing()",
			"A/B==4"))

	return err
}

type testFuncAlpha05 struct {
	input string
}

func (tFuncAlpha05 *testFuncAlpha05) Tx5DoSomethingBig(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha05." +
			"Tx5DoSomethingBig()")

	tFuncAlpha06 := testFuncAlpha06{}

	err := tFuncAlpha06.Tx6GiveUp(
		ePrefix.XCtx(
			"A->B"))

	return err
}

type testFuncAlpha06 struct {
	input string
}

func (tFuncAlpha06 *testFuncAlpha06) Tx6GiveUp(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha06." +
			"Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	return nil
}

type testFuncBravo01 struct {
	input string
}

func (tFuncBravo01 *testFuncBravo01) Tx1DoSomethingSpecial(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo01." +
			"Tx1DoSomethingSpecial()")

	tFuncBravo02 := testFuncBravo02{}

	return tFuncBravo02.Tx2DoSomethingElse(ePrefix)
}

type testFuncBravo02 struct {
	input string
}

func (tFuncBravo02 *testFuncBravo02) Tx2DoSomethingElse(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo02." +
			"Tx2DoSomethingElse()")

	tFuncBravo03 := testFuncBravo03{}

	err := tFuncBravo03.Tx3DoAnything(ePrefix)

	return err
}

type testFuncBravo03 struct {
	input string
}

func (tFuncBravo03 *testFuncBravo03) Tx3DoAnything(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo03." +
			"Tx3DoAnything()")

	tFuncBravo04 := testFuncBravo04{}

	err := tFuncBravo04.Tx4DoNothing(ePrefix)

	return err
}

type testFuncBravo04 struct {
	input string
}

func (tFuncBravo04 *testFuncBravo04) Tx4DoNothing(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo04." +
			"Tx4DoNothing()")

	tFuncBravo05 := testFuncBravo05{}

	ePrefix.SetCtx("A/B==4")

	err := tFuncBravo05.Tx5DoSomethingBig(ePrefix)

	return err
}

type testFuncBravo05 struct {
	input string
}

func (tFuncBravo05 *testFuncBravo05) Tx5DoSomethingBig(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo05." +
			"Tx5DoSomethingBig()")

	tFuncBravo06 := testFuncBravo06{}

	ePrefix.SetCtx("A->B")

	err := tFuncBravo06.Tx6GiveUp(ePrefix)

	return err
}

type testFuncBravo06 struct {
	input string
}

func (tFuncY6 *testFuncBravo06) Tx6GiveUp(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	return nil
}

type testFuncCharlie01 struct {
	input string
}

func (tFuncCharlie01 *testFuncCharlie01) Tx1DoStuff(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPrefCtx(
		"testFuncCharlie01."+
			"Tx1DoStuff()",
		"X->Y")

	tFuncCharlie02 := testFuncCharlie02{}

	return tFuncCharlie02.Tx2DoMoreStuff(
		ePrefix)
}

type testFuncCharlie02 struct {
	input string
}

func (tFuncCharlie02 *testFuncCharlie02) Tx2DoMoreStuff(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncCharlie02." +
			"Tx2DoMoreStuff()")

	tFuncCharlie03 := testFuncCharlie03{}

	return tFuncCharlie03.Tx3DoLessStuff(
		ePrefix.XCtx(
			"B->C"))
}

type testFuncCharlie03 struct {
	input string
}

func (tFuncCharlie03 *testFuncCharlie03) Tx3DoLessStuff(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncCharlie03." +
			"Tx3DoLessStuff()")

	tFuncCharlie04 := testFuncCharlie04{}

	return tFuncCharlie04.Tx4DoFunStuff(
		ePrefix)
}

type testFuncCharlie04 struct {
	input string
}

func (tFuncCharlie04 *testFuncCharlie04) Tx4DoFunStuff(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncCharlie04." +
			"Tx4DoFunStuff()")

	tFuncCharlie05 := testFuncCharlie05{}

	return tFuncCharlie05.Tx5DoExcitingStuff(
		ePrefix)

}

type testFuncCharlie05 struct {
	input string
}

func (tFuncCharlie05 *testFuncCharlie05) Tx5DoExcitingStuff(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncCharlie05." +
			"Tx5DoExcitingStuff()")

	tFuncCharlie06 := testFuncCharlie06{}

	return tFuncCharlie06.Tx6DoSpaceStuff(
		ePrefix.XCtx(
			"X*Y"))

}

type testFuncCharlie06 struct {
	input string
}

func (tFuncCharlie06 *testFuncCharlie06) Tx6DoSpaceStuff(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPrefCtx(
		"testFuncCharlie06."+
			"Tx6DoSpaceStuff()",
		"Asteroid Collision!")

	return fmt.Errorf("%v\n",
		ePrefix.String())
}

func TestErrPrefixDto_CallSeries01_000100(t *testing.T) {

	funcName := "StartingMethod()"
	twoDAry := make([][2]string, 7)

	twoDAry[0][0] = funcName
	twoDAry[0][1] = ""

	twoDAry[1][0] = "testFuncCharlie01.Tx1DoStuff()"
	twoDAry[1][1] = "X->Y"

	twoDAry[2][0] = "testFuncCharlie02.Tx2DoMoreStuff()"
	twoDAry[2][1] = "B->C"

	twoDAry[3][0] = "testFuncCharlie03.Tx3DoLessStuff()"
	twoDAry[3][1] = ""

	twoDAry[4][0] = "testFuncCharlie04.Tx4DoFunStuff()"
	twoDAry[4][1] = ""

	twoDAry[5][0] = "testFuncCharlie05.Tx5DoExcitingStuff()"
	twoDAry[5][1] = "X*Y"

	twoDAry[6][0] = "testFuncCharlie06.Tx6DoSpaceStuff()"
	twoDAry[6][1] = "Asteroid Collision!"

	ePDtoBaseLine,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDAry,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoBaseLine.SetMaxTextLineLen(40)

	ePDto1 := new(ErrPrefixDto)

	ePDto1.SetEPref(funcName)

	ePDto1.SetMaxTextLineLen(40)

	tFuncAlpha01 := testFuncAlpha01{}

	err =
		tFuncAlpha01.Tx1DoSomething(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncAlpha01.Tx1DoSomething(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncBravo01 := testFuncBravo01{}

	err =
		tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncCharlie01 := testFuncCharlie01{}

	err2 :=
		tFuncCharlie01.Tx1DoStuff(ePDto1)

	if err2 == nil {
		t.Error("Expected error return from\n" +
			"tFuncCharlie01.Tx1DoStuff(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")

		return
	}

	if ePDto1.String() != funcName {
		t.Errorf("Error: Expected ePDto1.String() != funcName\n"+
			"Instead:\n"+
			"ePDto1.String()='%v'\n"+
			"funcName='%v'\n\n",
			ePDto1.String(),
			funcName)
		return
	}

	var ePDtoFinal *ErrPrefixDto

	ePDtoFinal,
		err = ErrPrefixDto{}.NewIEmpty(
		err2.Error(),
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(err.Error())\n"+
			"%v\n", err.Error())

		return
	}

	ePDtoFinal.SetMaxTextLineLen(40)

	ePDtoBaseLineStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoBaseLine.String()),
		false)

	ePDtoFinalStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoFinal.String()),
		false)

	if ePDtoBaseLineStr != ePDtoFinalStr {
		t.Errorf("Error: ePDtoBaseLineStr != ePDtoFinalStr\n"+
			"ePDtoBaseLineStr=\n%v\n\n"+
			"ePDtoFinalStr=\n%v\n\n",
			ePDtoBaseLineStr,
			ePDtoFinalStr)
	}

}

func TestErrPrefixDto_CallSeries01_000200(t *testing.T) {

	funcName := "StartingMethod()"
	twoDAry := make([][2]string, 7)

	twoDAry[0][0] = funcName
	twoDAry[0][1] = ""

	twoDAry[1][0] = "testFuncCharlie01.Tx1DoStuff()"
	twoDAry[1][1] = "X->Y"

	twoDAry[2][0] = "testFuncCharlie02.Tx2DoMoreStuff()"
	twoDAry[2][1] = "B->C"

	twoDAry[3][0] = "testFuncCharlie03.Tx3DoLessStuff()"
	twoDAry[3][1] = ""

	twoDAry[4][0] = "testFuncCharlie04.Tx4DoFunStuff()"
	twoDAry[4][1] = ""

	twoDAry[5][0] = "testFuncCharlie05.Tx5DoExcitingStuff()"
	twoDAry[5][1] = "X*Y"

	twoDAry[6][0] = "testFuncCharlie06.Tx6DoSpaceStuff()"
	twoDAry[6][1] = "Asteroid Collision!"

	ePDtoBaseLine,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDAry,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDtoBaseLine.SetMaxTextLineLen(40)

	ePDto1 := new(ErrPrefixDto)

	ePDto1.SetEPref(funcName)

	ePDto1.SetMaxTextLineLen(40)

	tFuncAlpha01 := testFuncAlpha01{}

	err =
		tFuncAlpha01.Tx1DoSomething(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncAlpha01.Tx1DoSomething(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncBravo01 := testFuncBravo01{}

	err =
		tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncBravo01.Tx1DoSomethingSpecial(ePDto1)\n"+
			"%v\n", err.Error())

		return
	}

	tFuncCharlie01 := testFuncCharlie01{}

	err2 :=
		tFuncCharlie01.Tx1DoStuff(ePDto1)

	if err2 == nil {
		t.Error("Expected error return from\n" +
			"tFuncCharlie01.Tx1DoStuff(ePDto1)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")

		return
	}

	if ePDto1.String() != funcName {
		t.Errorf("Error: Expected ePDto1.String() != funcName\n"+
			"Instead:\n"+
			"ePDto1.String()='%v'\n"+
			"funcName='%v'\n\n",
			ePDto1.String(),
			funcName)
		return
	}

	var ePDtoFinal *ErrPrefixDto

	ePDtoFinal,
		err = ErrPrefixDto{}.NewIEmpty(
		err2.Error(),
		"",
		"")

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(err.Error())\n"+
			"%v\n", err.Error())

		return
	}

	ePDtoFinal.SetMaxTextLineLen(40)

	ePDtoBaseLineStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoBaseLine.String()),
		false)

	ePDtoFinalStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDtoFinal.String()),
		false)

	if ePDtoBaseLineStr != ePDtoFinalStr {
		t.Errorf("Error: ePDtoBaseLineStr != ePDtoFinalStr\n"+
			"ePDtoBaseLineStr=\n%v\n\n"+
			"ePDtoFinalStr=\n%v\n\n",
			ePDtoBaseLineStr,
			ePDtoFinalStr)
	}

}
