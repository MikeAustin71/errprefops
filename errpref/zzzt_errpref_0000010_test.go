package errpref

import (
	"testing"
)

func TestErrPref_FmtStr_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\\nTx3.DoSomething() - Tx4() - Tx5()\\nTx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPref_FmtStr_000200(t *testing.T) {

	initialStr :=
		"Tx1.Something() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.Something()[SPACE]:[SPACE]A->B\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()\n" +
		"Tx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPref_FmtStr_000300(t *testing.T) {

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\n" +
		"[SPACE]:[SPACE][SPACE]A->B\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()\n" +
		"Tx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPref_FmtStr_000400(t *testing.T) {
	initialStr :=
		"Tx1.Something() : A->B\nTx2.AVeryVeryLongMethodNameCalledSomething() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	expectedStr := "Tx1.Something() : A->B\n" +
		"Tx2.AVeryVeryLongMethodNameCalledSomething()\n" +
		"[SPACE]:[SPACE][SPACE]A==B\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()\n" +
		"Tx6.DoSomethingElse()"

	actualStr := ErrPref{}.FmtStr(
		initialStr)

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error: Expected actualStr= '%v'\n"+
			"Instead, actualStr='%v'\n",
			expectedStr,
			actualStr)
	}

}

/*
func TestNewErrPref_0010(t *testing.T) {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ErrPref{}.SetMaxErrPrefTextLineLength(60)

	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = ErrPref{}.SetCtxt(ePrefix, "something->newSomething")

	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx8.TryAnyCombination()")
	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx9.TryAHammer()")
	ePrefix = ErrPref{}.SetCtxt(ePrefix, "x->y")
	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx10.X()")
	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx11.TryAnything()")
	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx12.TryASalad()")
	ePrefix = ErrPref{}.EPrefCtx(ePrefix, "Tx13.SomeFabulousAndComplexStuff()")
	ePrefix = ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedResult :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\nTx4() - Tx5() - Tx6.DoSomethingElse()\\nTx7.TrySomethingNew() :  something->newSomething\\nTx8.TryAnyCombination() - Tx9.TryAHammer() :  x->y\\nTx10.X() - Tx11.TryAnything() - Tx12.TryASalad()\\nTx13.SomeFabulousAndComplexStuff()\\nTx14.MoreAwesomeGoodness :  A=7 B=8 C=9"

	ePrefix = ErrPref{}.FmtStr(ePrefix)

	actualResult := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePrefix),
		false)

	if expectedResult != actualResult {
		t.Errorf("Error: Expected Result=='%v'\n"+
			"Instead, Actual Result=='%v'\n",
			expectedResult,
			actualResult)
	}

}
*/
