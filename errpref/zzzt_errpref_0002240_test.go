package errpref

import "testing"

func TestErrPrefixDto_CopyInFromIBuilder_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_CopyInFromIBuilder"
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

	iBuilder := testIBuilderErrPref{}

	iBuilder.SetEPrefStrings(twoDSlice)

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefStrings(twoDSlice)

	ePDto1.SetMaxTextLineLen(40)

	bogusStr :=
		"Xt1.Something() - Xt2.SomethingElse() - Xt3.DoSomething()\n" +
			"Xt4() - Xt5() - Xt6.DoSomethingElse()\n" +
			"Xt7.TrySomethingNew() : something->newSomething\n" +
			"Xt8.TryAnyCombination() - Xt9.TryAHammer() : x->y - Xt10.X()\n" +
			"Xt11.TryAnything() - Xt12.TryASalad()\n" +
			"Xt13.SomeFabulousAndComplexStuff()\n" +
			"Xt14.MoreAwesomeGoodness\n" +
			"Xt15.MustardSandwiches()\n" +
			"Xt16.TomatoSandwiches()\n" +
			"Xt17.Benitos() : Z=4 Y=3 X=9"

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetEPrefOld(bogusStr)

	ePDto2.SetMaxTextLineLen(40)

	err :=
		ePDto2.CopyInFromIBuilder(
			&iBuilder,
			funcName)

	if err != nil {
		t.Errorf("Error from ErrPrefixDto{}.CopyInFromIBuilder()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto1.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto1==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1.String(),
			ePDto2.String())
		return
	}

	if ePDto1Str != ePDto2Str {
		t.Errorf("Error: Expected ePDto1Str==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1Str,
			ePDto2Str)
	}

}

func TestErrPrefixDto_CopyOutToIBuilder_000100(t *testing.T) {

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

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefStrings(twoDSlice)

	ePDto1.SetMaxTextLineLen(40)

	iBuilder := testIBuilderErrPref{}

	ePDto1.SetEPrefStrings(twoDSlice)

	ePDto1.SetMaxTextLineLen(40)

	ePDto1.CopyOutToIBuilder(&iBuilder)

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefStrings(
		iBuilder.GetEPrefStrings())

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto1.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto1==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1.String(),
			ePDto2.String())
		return
	}

	if ePDto1Str != ePDto2Str {
		t.Errorf("Error: Expected ePDto1Str==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1Str,
			ePDto2Str)
	}

}

func TestErrPrefixDto_SetIBuilder(t *testing.T) {

	funcName := "TestErrPrefixDto_SetIBuilder"

	bogusStr :=
		"Xt1.Something() - Xt2.SomethingElse() - Xt3.DoSomething()\n" +
			"Xt4() - Xt5() - Xt6.DoSomethingElse()\n" +
			"Xt7.TrySomethingNew() : something->newSomething\n" +
			"Xt8.TryAnyCombination() - Xt9.TryAHammer() : x->y - Xt10.X()\n" +
			"Xt11.TryAnything() - Xt12.TryASalad()\n" +
			"Xt13.SomeFabulousAndComplexStuff()\n" +
			"Xt14.MoreAwesomeGoodness\n" +
			"Xt15.MustardSandwiches()\n" +
			"Xt16.TomatoSandwiches()\n" +
			"Xt17.Benitos() : Z=4 Y=3 X=9"

	ePDto1 := ErrPrefixDto{}.New()

	ePDto1.SetEPrefOld(bogusStr)

	ePDto1.SetMaxTextLineLen(40)

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

	iBuilder := testIBuilderErrPref{}

	iBuilder.SetEPrefStrings(twoDSlice)

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

	ePDto2.SetEPrefStrings(twoDSlice)

	err := ePDto1.SetIBuilder(
		&iBuilder,
		funcName)

	if err != nil {
		t.Errorf("Error from ePDto1.SetIBuilder()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto1Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		false)

	ePDto2Str := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if !ePDto1.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto1==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1.String(),
			ePDto2.String())
		return
	}

	if ePDto1Str != ePDto2Str {
		t.Errorf("Error: Expected ePDto1Str==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto1=\n%v\n\nePDto2=\n%v\n\n",
			ePDto1Str,
			ePDto2Str)
	}

}
