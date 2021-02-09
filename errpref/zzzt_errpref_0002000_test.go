package errpref

import "testing"

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
