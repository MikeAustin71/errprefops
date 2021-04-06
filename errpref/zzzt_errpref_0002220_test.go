package errpref

import "testing"

func TestErrPrefixDto_NewEPrefOld_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()"

	expectedStr := "Tx1.Something() - Tx2.SomethingElse()\nTx3.DoSomething() - Tx4() - Tx5()\nTx6.DoSomethingElse()"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

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

	ePDto.SetMaxTextLineLen(40)

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

func TestErrPrefixDto_NewFromIErrorPrefix_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto2 := ErrPrefixDto{}.NewFromIErrorPrefix(
		&ePDto)

	ePDto2.SetMaxTextLineLen(40)

	areEqual := ePDto.Equal(&ePDto2)

	if !areEqual {
		t.Error("Error: Expected ePDto and ePDto2 would be equal.\n" +
			"THEY ARE NOT EQUAL!\n")
	}

}

func TestErrPrefixDto_NewIEmpty_000100(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDSlice,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
	}

}

func TestErrPrefixDto_NewIEmpty_000200(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	funcName := "TestErrPrefixDto_NewIEmpty_000200"

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDSlice,
		funcName,
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	if ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto!=ePDto2.\n"+
			"However, THEY ARE EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
	}

}

func TestErrPrefixDto_NewIEmpty_000300(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetEPrefCtx(
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePDto.SetMaxTextLineLen(40)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 13)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		twoDSlice,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000400(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew()\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness()"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	var oneDSlice []string

	oneDSlice = make([]string, 14)

	oneDSlice[0] = "Tx1.Something()"

	oneDSlice[1] = "Tx2.SomethingElse()"

	oneDSlice[2] = "Tx3.DoSomething()"

	oneDSlice[3] = "Tx4()"

	oneDSlice[4] = "Tx5()"

	oneDSlice[5] = "Tx6.DoSomethingElse()"

	oneDSlice[6] = "Tx7.TrySomethingNew()"

	oneDSlice[7] = "Tx8.TryAnyCombination()"

	oneDSlice[8] = "Tx9.TryAHammer()"

	oneDSlice[9] = "Tx10.X()"

	oneDSlice[10] = "Tx11.TryAnything()"

	oneDSlice[11] = "Tx12.TryASalad()"

	oneDSlice[12] = "Tx13.SomeFabulousAndComplexStuff()"

	oneDSlice[13] = "Tx14.MoreAwesomeGoodness()"

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		oneDSlice,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000500(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		initialStr,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000600(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()" +
			"Tx7.TrySomethingNew() : something->newSomething" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()" +
			"Tx11.TryAnything() - Tx12.TryASalad()" +
			"Tx13.SomeFabulousAndComplexStuff()" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto2,
		err := ErrPrefixDto{}.NewIEmpty(
		initialStr,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000700(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error
	var ePDtoCopy ErrPrefixDto
	var ePDto2 *ErrPrefixDto

	funcName :=
		"TestErrPrefixDto_NewIEmpty_000700()"

	ePDtoCopy,
		err = ePDto.CopyOut(funcName)

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		ePDtoCopy,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000800(t *testing.T) {

	/*
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
		"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
		"Tx7.TrySomethingNew() : something->newSomething\n" +
		"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
		"Tx11.TryAnything() - Tx12.TryASalad()\n" +
		"Tx13.SomeFabulousAndComplexStuff()\n" +
		"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	*/

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	var err error
	var ePDtoCopy *ErrPrefixDto
	var ePDto2 *ErrPrefixDto

	ePDtoCopy = ePDto.CopyPtr()

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		ePDtoCopy,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(40)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto.Equal(ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if ePDtoStr != ePDto2Str {
		t.Errorf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
	}

}

func TestErrPrefixDto_NewIEmpty_000900(t *testing.T) {

	var err error
	var ePDto2 *ErrPrefixDto

	ePDto2,
		err = ErrPrefixDto{}.NewIEmpty(
		nil,
		"",
		"")

	if err != nil {
		t.Errorf("Error from  ErrPrefixDto{}.NewIEmpty()\n"+
			"%v\n", err.Error())
		return
	}

	if ePDto2.GetEPrefCollectionLen() != 0 {
		t.Errorf("Error: Expected Collection Length==0\n"+
			"Instead Collection Length=='%v'\n",
			ePDto2.GetEPrefCollectionLen())
	}
}
