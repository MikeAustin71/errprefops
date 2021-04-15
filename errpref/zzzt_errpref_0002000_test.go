package errpref

import "testing"

func TestErrPrefixDto_AddEPrefCollectionStr_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto.SetEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto2 := ErrPrefixDto{}.New()

	numOfItemsAdded := ePDto2.AddEPrefCollectionStr(
		initialStr)

	if numOfItemsAdded != ePDto.GetEPrefCollectionLen() {
		t.Errorf("Error: Number of Items Added NOT equal to Collection Items!\n"+
			"numOfItemsAdded=='%v'\n"+
			"ePDto.GetEPrefCollectionLen()=='%v'\n",
			numOfItemsAdded,
			ePDto.GetEPrefCollectionLen())

		return
	}

	ePDto2.SetMaxTextLineLen(40)

	if ePDto.GetEPrefCollectionLen() != ePDto2.GetEPrefCollectionLen() {
		t.Errorf("Expected ePDto.GetEPrefCollectionLen() == ePDto2.GetEPrefCollectionLen()\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto.GetEPrefCollectionLen()=='%v'\n"+
			"ePDto2.GetEPrefCollectionLen()=='%v'\n",
			ePDto.GetEPrefCollectionLen(),
			ePDto2.GetEPrefCollectionLen())
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Errorf("Error: Expected ePDto==ePDto2.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDto.String(),
			ePDto2.String())
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if expectedConvertedStr != actualConvertedStr {
		t.Errorf("Error: Expected expectedConvertedStr==actualConvertedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedConvertedStr=\n%v\n\n"+
			"actualConvertedStr=\n%v\n\n",
			expectedConvertedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_CopyPtr_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

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

	ePDto2 := ePDto.CopyPtr()

	if !ePDto.Equal(ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
	}

}

func TestErrPrefixDto_CopyIn_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

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

	ePDto2 := ErrPrefixDto{}.Ptr()

	err := ePDto2.CopyIn(
		&ePDto,
		"TestErrPrefixDto_CopyIn_000100")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if !ePDto.Equal(ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
	}
}

func TestErrPrefixDto_CopyIn_000200(t *testing.T) {
	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(5)

	ePDto.SetLeftMarginChar('*')

	ePDto.SetIsLastLineTermWithNewLine(true)

	ePDto2 := ErrPrefixDto{}.Ptr()

	err := ePDto2.CopyIn(
		&ePDto,
		"TestErrPrefixDto_CopyIn_000200")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if ePDto2.GetLeftMarginLength() != 5 {
		t.Errorf("Error: After Copy expected Left Margin Length==5.\n"+
			"Instead, Left Margin Length = %v\n",
			ePDto2.GetLeftMarginLength())
		return
	}

	if ePDto2.GetLeftMarginChar() != '*' {
		t.Errorf("Error: After Copy expected Left Margin Char=='*'.\n"+
			"Instead, Left Margin Char == %v\n",
			string(ePDto2.GetLeftMarginChar()))
		return
	}

	if !ePDto2.GetIsLastLineTerminatedWithNewLine() {
		t.Error("Error: After Copy expected GetIsLastLineTerminatedWithNewLine()=='true'.\n" +
			"Instead, GetIsLastLineTerminatedWithNewLine()=='false'.\n")
		return
	}

	if ePDto2.GetMaxErrPrefTextLineLength() != 40 {
		t.Errorf("Error: After Copy expected GetMaxErrPrefTextLineLength=='40'.\n"+
			"Instead, GetMaxErrPrefTextLineLength == %v\n",
			ePDto2.GetMaxErrPrefTextLineLength())
		return
	}

	if !ePDto.Equal(ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if expectedConvertedStr != actualConvertedStr {
		t.Errorf("Error: Expected expectedConvertedStr==actualConvertedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedConvertedStr=\n%v\n\n"+
			"actualConvertedStr=\n%v\n\n",
			expectedConvertedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_CopyOut_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

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

	ePDto2,
		err := ePDto.CopyOut("TestErrPrefixDto_CopyOut_000100")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if expectedConvertedStr != actualConvertedStr {
		t.Errorf("Error: Expected expectedConvertedStr==actualConvertedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedConvertedStr=\n%v\n\n"+
			"actualConvertedStr=\n%v\n\n",
			expectedConvertedStr,
			actualConvertedStr)
	}

}

func TestErrPrefixDto_CopyOut_000200(t *testing.T) {
	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(5)

	ePDto.SetLeftMarginChar('*')

	ePDto.SetIsLastLineTermWithNewLine(true)

	ePDto2,
		err := ePDto.CopyOut("TestErrPrefixDto_CopyOut_000200")

	if err != nil {
		t.Errorf("Error:\n"+
			"%v", err.Error())
		return
	}

	if ePDto2.GetLeftMarginLength() != 5 {
		t.Errorf("Error: After Copy expected Left Margin Length==5.\n"+
			"Instead, Left Margin Length = %v\n",
			ePDto2.GetLeftMarginLength())
		return
	}

	if ePDto2.GetLeftMarginChar() != '*' {
		t.Errorf("Error: After Copy expected Left Margin Char=='*'.\n"+
			"Instead, Left Margin Char == %v\n",
			string(ePDto2.GetLeftMarginChar()))
		return
	}

	if !ePDto2.GetIsLastLineTerminatedWithNewLine() {
		t.Error("Error: After Copy expected GetIsLastLineTerminatedWithNewLine()=='true'.\n" +
			"Instead, GetIsLastLineTerminatedWithNewLine()=='false'.\n")
		return
	}

	if ePDto2.GetMaxErrPrefTextLineLength() != 40 {
		t.Errorf("Error: After Copy expected GetMaxErrPrefTextLineLength=='40'.\n"+
			"Instead, GetMaxErrPrefTextLineLength == %v\n",
			ePDto2.GetMaxErrPrefTextLineLength())
		return
	}

	if !ePDto.Equal(&ePDto2) {
		t.Error("Expected ePDto to Equal ePDto2\n" +
			"However, THEY ARE NOT EQUAL!\n")
		return
	}

	expectedConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		true)

	actualConvertedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		true)

	if expectedConvertedStr != actualConvertedStr {
		t.Errorf("Error: Expected expectedConvertedStr==actualConvertedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedConvertedStr=\n%v\n\n"+
			"actualConvertedStr=\n%v\n\n",
			expectedConvertedStr,
			actualConvertedStr)
	}
}

func TestErrPrefixDto_Empty_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetLeftMarginLength(3)
	ePDto.SetLeftMarginChar('*')
	ePDto.SetMaxTextLineLen(40)
	ePDto.SetIsLastLineTermWithNewLine(true)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen == 0 {
		t.Error("Error: After Initialization with data,\n" +
			"Error Prefix Collection Length==0\n")
		return
	}

	ePDto.Empty()

	if ePDto.GetLeftMarginLength() != 0 {
		t.Errorf("Error: After Empty() expected Left Margin\n"+
			"Length == 0.\n"+
			"Instead, ePDto.GetLeftMarginLength()=='%v'\n",
			ePDto.GetLeftMarginLength())
		return
	}

	if ePDto.GetEPrefCollectionLen() != 0 {
		t.Errorf("Error: After Empty() expected Collection\n"+
			"Length == 0.\n"+
			"Instead, ePDto.GetEPrefCollectionLen()=='%v'\n",
			ePDto.GetEPrefCollectionLen())
		return
	}

	if ePDto.GetLeftMarginChar() != 0 {
		t.Errorf("Error: After Empty() expected Left Margin\n"+
			"Character == 0.\n"+
			"Instead, ePDto.GetLeftMarginChar()=='%v'\n",
			string(ePDto.GetLeftMarginChar()))
		return
	}

	if ePDto.GetIsLastLineTerminatedWithNewLine() {
		t.Error("Error: After Empty() expected isLastLineTerminatedWithNewLine==false\n" +
			"Instead, isLastLineTerminatedWithNewLine==true")
	}

}

func TestErrPrefixDto_EmptyEPrefCollection_000100(t *testing.T) {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	collectionLen := ePDto.GetEPrefCollectionLen()

	if collectionLen == 0 {
		t.Error("Error: After Initialization with data,\n" +
			"Error Prefix Collection Length==0\n")
		return
	}

	ePDto.EmptyEPrefCollection()

	collectionLen = ePDto.GetEPrefCollectionLen()

	if collectionLen != 0 {
		t.Errorf("Error: After EmptyEPrefCollection(),\n"+
			"Error Prefix Collection Length!=0\n"+
			"Error Prefix Collection Length=='%v'\n",
			ePDto.GetEPrefCollectionLen())

	}

}

func TestErrPrefixDto_Equal_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

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

	ePDto2 := ePDto.Copy()

	if !ePDto.Equal(&ePDto2) {

		t.Error("Error:\n" +
			"Expected ePDto to Equal ePDto2.\n" +
			"However, THEY ARE NOT EQUAL!\n")

		return
	}

	ePDto2.SetMaxTextLineLen(60)

	if ePDto.Equal(&ePDto2) {

		t.Error("Error:\n" +
			"Expected ePDto to be UNEqual to ePDto2.\n" +
			"However, THEY ARE EQUAL!\n")

		return
	}
}

func TestErrPrefixDto_GetDelimiters_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_GetDelimiters_000100"

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

	errPrefDelims := ePDto.GetDelimiters()

	errPrefDelimsOne := ErrPrefixDelimiters{}

	err := errPrefDelimsOne.CopyIn(
		&errPrefDelims,
		funcName)

	if err != nil {
		t.Errorf("Error Returned by errPrefDelimsOne.CopyIn()\n"+
			"%v\n",
			err.Error())

		return
	}

	ep := ErrPref{}

	epDelims := ep.GetDelimiters()

	var errPrefDelimsTwo ErrPrefixDelimiters

	errPrefDelimsTwo,
		err = epDelims.CopyOut(
		funcName)

	if err != nil {
		t.Errorf("Error Returned by epDelims.CopyOut()\n"+
			"%v\n",
			err.Error())

		return
	}

	errPrefDelimsOneStr := errPrefDelimsOne.String()

	errPrefDelimsTwoStr := errPrefDelimsTwo.String()

	if !errPrefDelimsOne.Equal(&errPrefDelimsTwo) {
		t.Errorf("Error: Expected errPrefDelimsOne==errPrefDelimTwo\n"+
			"HOWEVER, THE TWO ARE NOT EQUAL\n"+
			"errPrefDelimsOne=\n%v\n"+
			"errPrefDelimsTwo=\n%v\n",
			errPrefDelimsOneStr,
			errPrefDelimsTwoStr)

		return
	}

	if errPrefDelimsOneStr != errPrefDelimsTwoStr {
		t.Errorf("Error: Expected errPrefDelimsOneStr==errPrefDelimsTwoStr\n"+
			"HOWEVER, THE TWO ARE NOT EQUAL\n"+
			"errPrefDelimsOne=\n%v\n"+
			"errPrefDelimsTwo=\n%v\n",
			errPrefDelimsOneStr,
			errPrefDelimsTwoStr)
	}

}

func TestErrPrefixDto_MergeErrPrefixDto_000100(t *testing.T) {

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

	ePDto2 := ErrPrefixDto{}.New()

	ePDto2.SetMaxTextLineLen(40)

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

	ePDto2.SetEPrefOld(bogusStr)

	expectedCompositeStr := initialStr +
		"\n" +
		bogusStr

	expectedEPDto := ErrPrefixDto{}.New()

	expectedEPDto.SetMaxTextLineLen(40)

	expectedEPDto.SetEPrefOld(expectedCompositeStr)

	ePDto.MergeErrPrefixDto(
		&ePDto2)

	expectedStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedEPDto.String()),
		false)

	ePDtoStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	if !expectedEPDto.Equal(&ePDto) {
		t.Errorf("Error: Expected expectedEPDto==ePDto.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			expectedStr,
			ePDtoStr)
		return
	}

	if expectedStr != ePDtoStr {
		t.Errorf("Error: Expected expectedStr==ePDtoStr.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"expectedStr=\n%v\n\nePDtoStr=\n%v\n\n",
			expectedStr,
			ePDtoStr)
	}

}

func TestErrPrefixDto_Multiple_000100(t *testing.T) {

	initialStr := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePDto := ErrPrefixDto{}.NewEPrefOld(initialStr)

	// Setting Line Length to 60-Characters
	ePDto.SetMaxTextLineLen(60)

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

	ePDto.SetMaxTextLineLenToDefault()
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
	ePDto.SetMaxTextLineLen(40)

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

	ePDto.SetMaxTextLineLenToDefault()
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
	ePDto.SetMaxTextLineLen(40)

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

	ePDto.SetMaxTextLineLenToDefault()
}
