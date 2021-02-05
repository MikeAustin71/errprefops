package errpref

import "testing"

func TestNewErrPref_0010(t *testing.T) {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ErrPref{}.SetMaxErrPrefTextLineLength(60)

	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = ErrPref{}.SetCtxt(ePrefix, "something->newSomething")

	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx8.TryAnyCombination()")
	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx9.TryAHammer()")
	ePrefix = ErrPref{}.SetCtxt(ePrefix, "x->y")
	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx10.X()")
	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx11.TryAnything()")
	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx12.TryASalad()")
	ePrefix = ErrPref{}.NewErrPref(ePrefix, "Tx13.SomeFabulousAndComplexStuff()")
	ePrefix = ErrPref{}.NewCtxt(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedResult :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\nTx4() - Tx5() - Tx6.DoSomethingElse()\\nTx7.TrySomethingNew() :  something->newSomething\\nTx8.TryAnyCombination() - Tx9.TryAHammer() :  x->y\\nTx10.X() - Tx11.TryAnything() - Tx12.TryASalad()\\nTx13.SomeFabulousAndComplexStuff()\\nTx14.MoreAwesomeGoodness :  A=7 B=8 C=9"

	ePrefix = ErrPref{}.FmtString(ePrefix)

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
