package errpref

import "testing"

func TestErrPrefixDto_Multiple_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxErrPrefTextLineLength(60)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew() : something->newSomething\\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\\n" +
			"Tx13.SomeFabulousAndComplexStuff()\\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	actualStr := ePDto.String()

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		false)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

	ePDto.SetMaxErrPrefTextLineLengthToDefault()
}

func TestErrPrefixDto_Multiple_000200(t *testing.T) {

	initialStr :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 40-Characters
	ePDto.SetMaxErrPrefTextLineLength(40)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPref("Tx9.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx10.X()")

	ePDto.SetEPrefCtx(
		"Tx11.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx12.TryASalad()",
		"")

	ePDto.SetEPref("Tx13.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto.SetEPrefCtx(
		"Tx15.AVeryVeryLongMethodNameCalledSomething()",
		"A=7 B=8 C=9")

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse()\\n" +
			"Tx3.DoSomething() - Tx4() - Tx5()\\n" +
			"Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew()\\n" +
			" :  something->newSomething\\n" +
			"Tx8.TryAnyCombination()\\n" +
			"Tx9.TryAHammer() : x->y - Tx10.X()\\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\\n" +
			"Tx13.SomeFabulousAndComplexStuff()\\n" +
			"Tx14.MoreAwesomeGoodness\\n" +
			" :  A=7 B=8 C=9\\n" +
			"Tx15.AVeryVeryLongMethodNameCalledSomething()\\n" +
			" :  A=7 B=8 C=9"

	actualStr := ePDto.String()

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		false)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

	ePDto.SetMaxErrPrefTextLineLengthToDefault()
}

func TestErrPrefixDto_Multiple_000300(t *testing.T) {

	initialStr :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 40-Characters
	ePDto.SetMaxErrPrefTextLineLength(40)

	ePDto.SetEPrefCtx(
		"Tx7.TrySomethingNew()",
		"")

	ePDto.SetCtx("something->newSomething")

	ePDto.SetEPrefCtx(
		"Tx8.TryAnyCombination()",
		"")

	ePDto.SetEPrefCtx(
		"Tx9.AVeryVeryLongMethodNameCalledSomething()",
		"A=7 B=8 C=9")

	ePDto.SetEPref("Tx10.TryAHammer()")

	ePDto.SetCtx("x->y")

	ePDto.SetEPref("Tx11.X()")

	ePDto.SetEPrefCtx(
		"Tx12.TryAnything()",
		"")

	ePDto.SetEPrefCtx(
		"Tx13.TryASalad()",
		"")

	ePDto.SetEPref("Tx14.SomeFabulousAndComplexStuff()")

	ePDto.SetEPrefCtx(
		"Tx15.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse()\\n" +
			"Tx3.DoSomething() - Tx4() - Tx5()\\n" +
			"Tx6.DoSomethingElse()\\n" +
			"Tx7.TrySomethingNew()\\n" +
			" :  something->newSomething\\n" +
			"Tx8.TryAnyCombination()\\n" +
			"Tx9.AVeryVeryLongMethodNameCalledSomething()\\n" +
			" :  A=7 B=8 C=9\\n" +
			"Tx10.TryAHammer() : x->y - Tx11.X()\\n" +
			"Tx12.TryAnything() - Tx13.TryASalad()\\n" +
			"Tx14.SomeFabulousAndComplexStuff()\\n" +
			"Tx15.MoreAwesomeGoodness\\n" +
			" :  A=7 B=8 C=9"

	actualStr := ePDto.String()

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		false)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

	ePDto.SetMaxErrPrefTextLineLengthToDefault()
}

func TestErrPrefixDto_NewEPrefOld_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\nTx3.DoSomething() - Tx4() - Tx5()\nTx6.DoSomethingElse()"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxErrPrefTextLineLength(40)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_NewEPrefOld_000200(t *testing.T) {

	initialStr := "Tx1.StartThis()"

	expectedStr := "Tx1.StartThis()"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxErrPrefTextLineLength(40)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetCtx_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetCtx("A!=B")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetCtx_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]A!=B"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetCtx("A!=B")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetCtx_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	expectedStr := ""

	ePDto.SetCtx("A!=B")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetEPref_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPref("Tx5.BrandNewMethod()")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetEPref_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	expectedStr := "Tx5.BrandNewMethod()"

	ePDto.SetEPref("Tx5.BrandNewMethod()")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetEPref_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPref("")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetEPref_000400(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr := "Tx1.StartThis()"

	expectedStr := "Tx1.StartThis()"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPref("")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetEPrefCtx_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()[SPACE]:[SPACE]X->G"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPrefCtx(
		"Tx5.BrandNewMethod()",
		"X->G")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPrefixDto_SetEPrefCtx_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999\\n" +
		"Tx5.BrandNewMethod()"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPrefCtx(
		"Tx5.BrandNewMethod()",
		"")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPrefixDto_SetEPrefCtx_000300(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPrefCtx(
		"",
		"G->X")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPrefixDto_SetEPrefCtx_000400(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
		"[SPACE]:[SPACE][SPACE]A->B\\n" +
		"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
		"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
		"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPrefCtx(
		"",
		"")

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:\n"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}
}

func TestErrPrefixDto_SetIsLastLineTermWithNewLine_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.Something()\n" +
			"Tx2.SomethingElse()\n" +
			"Tx3.DoSomething()\n" +
			"Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()"

	expectedStr :=
		"Tx1.Something() - Tx2.SomethingElse()\n" +
			"Tx3.DoSomething() - Tx4() - Tx5()\n" +
			"Tx6.DoSomethingElse()\n"

	ePDto.SetEPrefOld(initialStr)

	ePDto.
		SetIsLastLineTermWithNewLine(true)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_SetMaxErrPrefTextLineLength_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(60)

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\nTx4() - Tx5() - Tx6.DoSomethingElse()"

	ePDto.SetEPrefOld(initialStr)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_String_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\nTx3.DoSomething() - Tx4() - Tx5()\nTx6.DoSomethingElse()"

	ePDto.SetEPrefOld(initialStr)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}

func TestErrPrefixDto_String_000200(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxErrPrefTextLineLength(40)

	initialStr := "Tx1.StartThis()"

	expectedStr := "Tx1.StartThis()"

	ePDto.SetEPrefOld(initialStr)

	actualStr := ePDto.String()

	expectedStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	actualStr = ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	if expectedStr != actualStr {

		t.Errorf("Error:"+
			"Expected actualStr= '%v'\n"+
			"Instead, actualStr= '%v'\n",
			expectedStr,
			actualStr)
	}

}
