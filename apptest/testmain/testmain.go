package testmain

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
	"strings"
)

type TestMain struct {
	testStr01 string
}

func (tMain TestMain) TestMain026() {

	var twoDSlice01 [][2]string

	twoDSlice01 = make([][2]string, 14)

	twoDSlice01[0][0] = "Tx1.Something()"
	twoDSlice01[0][1] = ""

	twoDSlice01[1][0] = "Tx2.SomethingElse()"
	twoDSlice01[1][1] = ""

	twoDSlice01[2][0] = "Tx3.DoSomething()"
	twoDSlice01[2][1] = ""

	twoDSlice01[3][0] = "Tx4()"
	twoDSlice01[3][1] = ""

	twoDSlice01[4][0] = "Tx5()"
	twoDSlice01[4][1] = ""

	twoDSlice01[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice01[5][1] = ""

	twoDSlice01[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice01[6][1] = "something->newSomething"

	twoDSlice01[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice01[7][1] = ""

	twoDSlice01[8][0] = "Tx9.TryAHammer()"
	twoDSlice01[8][1] = "x->y"

	twoDSlice01[9][0] = "Tx10.X()"
	twoDSlice01[9][1] = ""

	twoDSlice01[10][0] = "Tx11.TryAnything()"
	twoDSlice01[10][1] = ""

	twoDSlice01[11][0] = "Tx12.TryASalad()"
	twoDSlice01[11][1] = ""

	twoDSlice01[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice01[12][1] = ""

	twoDSlice01[13][0] = "Tx14.MoreAwesomeGoodness"
	twoDSlice01[13][1] = "A=7 B=8 C=9"

	var twoDSlice02 [][2]string

	twoDSlice02 = make([][2]string, 14)

	twoDSlice02[0][0] = "Tx1.TryBeer()"
	twoDSlice02[0][1] = ""

	twoDSlice02[1][0] = "Tx2.TryMescal()"
	twoDSlice02[1][1] = ""

	twoDSlice02[2][0] = "Tx3.TryWhiskey()"
	twoDSlice02[2][1] = ""

	twoDSlice02[3][0] = "Tx4.TryMoonShine()"
	twoDSlice02[3][1] = ""

	twoDSlice02[4][0] = "Tx5.TryScotch()"
	twoDSlice02[4][1] = ""

	twoDSlice02[5][0] = "Tx6.TryWine()"
	twoDSlice02[5][1] = ""

	twoDSlice02[6][0] = "Tx7.TryBrandy()"
	twoDSlice02[6][1] = "something->newSomething"

	twoDSlice02[7][0] = "Tx8.TryCognac()"
	twoDSlice02[7][1] = ""

	twoDSlice02[8][0] = "Tx9.TryAle()"
	twoDSlice02[8][1] = "x->y"

	twoDSlice02[9][0] = "Tx10.Vodka()"
	twoDSlice02[9][1] = ""

	twoDSlice02[10][0] = "Tx11.TryRum()"
	twoDSlice02[10][1] = ""

	twoDSlice02[11][0] = "Tx12.TryBourbon()"
	twoDSlice02[11][1] = ""

	twoDSlice02[12][0] = "Tx13.TryTequila()"
	twoDSlice02[12][1] = ""

	twoDSlice02[13][0] = "Tx14.TryWater"
	twoDSlice02[13][1] = "A=7 B=8 C=9"

	var ePDto1, ePDto3 *errpref.ErrPrefixDto
	var err error

	ePDto1,
		err = errpref.ErrPrefixDto{}.NewIEmpty(
		twoDSlice01,
		"",
		"")

	if err != nil {
		fmt.Printf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(twoDSlice01)\n"+
			"%v\n", err.Error())

		return
	}

	ePDto1.AddEPrefStrings(twoDSlice02)

	ePDto1.SetMaxTextLineLen(40)

	var twoDSlice03 [][2]string

	twoDSlice03 = make([][2]string,
		len(twoDSlice01)+len(twoDSlice02))

	copy(twoDSlice03[:len(twoDSlice01)],
		twoDSlice01[:])

	copy(twoDSlice03[len(twoDSlice01):],
		twoDSlice02[:])

	ePDto3,
		err = errpref.ErrPrefixDto{}.NewIEmpty(
		twoDSlice03,
		"",
		"")

	if err != nil {
		fmt.Printf("Unexpected error return from\n"+
			"ErrPrefixDto{}.NewIEmpty(twoDSlice03)\n"+
			"%v\n", err.Error())

		return
	}

	ePDto3.SetMaxTextLineLen(40)

	outputOne := ePDto3.String()

	fmt.Println(outputOne)

	expectedStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto3.String()),
		true)

	actualStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto1.String()),
		true)

	if expectedStr != actualStr {
		fmt.Printf("Error: Expected expectedStr == actualStr\n"+
			"Instead:\n"+
			"expectedStr=\n%v\n\n"+
			"actualStr=\n%v\n\n",
			expectedStr,
			actualStr)
		return
	}

	if !ePDto1.Equal(ePDto3) {
		fmt.Printf("Error: Expected ePDto1 == ePDto3\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"ePDto1.String()=\n%v\n\n"+
			"ePDto3.String()=\n%v\n\n",
			ePDto1.String(),
			ePDto3.String())
	}

}

func (tMain TestMain) TestMain025() {

	ePDto := errpref.ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	twoDSlice := make([][2]string, 2)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	outputStr := fmt.Sprintf("%v"+
		"Error: Divide by Zero!",
		ePDto.String())

	fmt.Printf(outputStr)

	expectedStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(outputStr),
		true)

	fmt.Println()
	fmt.Printf(expectedStr + "\n")

}

func (tMain TestMain) TestMain024() {

	ePDto := errpref.ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

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

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness()"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	outputStr := fmt.Sprintf("%v"+
		"Error: Divide by Zero!",
		ePDto.String())

	fmt.Printf(outputStr)

}

func (tMain TestMain) TestMain023() {

	funcName := "TestErrPrefixDto_NewIEmptyWithDelimiters_000900()"

	ePDto := errpref.ErrPrefixDto{}.New()

	maxLineLen := 60

	ePDto.SetMaxTextLineLen(maxLineLen)

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

	var inputDelimiters errpref.ErrPrefixDelimiters

	inputDelimiters,
		err = errpref.ErrPrefixDelimiters{}.New(
		"\n  #",
		" ### ",
		"\n      -",
		" --- ",
		funcName)

	if err != nil {
		fmt.Printf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetInputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		fmt.Printf("Error from ePDto.SetInputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	err =
		ePDto.SetOutputStringDelimiters(
			inputDelimiters,
			funcName)

	if err != nil {
		fmt.Printf("Error from ePDto.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	initialStr = ePDto.String()

	var outputDelimiters errpref.ErrPrefixDelimiters

	outputDelimiters,
		err = errpref.ErrPrefixDelimiters{}.New(
		"\n  &",
		" *&* ",
		"\n      %",
		" %%% ",
		funcName)

	if err != nil {
		fmt.Printf("Error from ErrPrefixDelimiters{}.New()\n"+
			"%v\n", err.Error())
		return
	}

	iStringerEPref := testIStringerErrPref{}

	iStringerEPref.Set(initialStr)

	err =
		ePDto.SetOutputStringDelimiters(
			outputDelimiters,
			funcName)

	if err != nil {
		fmt.Printf("Error from ePDto.SetOutputStringDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	var ePDto2 *errpref.ErrPrefixDto

	ePDto2,
		err = errpref.ErrPrefixDto{}.NewIEmptyWithDelimiters(
		&iStringerEPref,
		"",
		"",
		inputDelimiters,
		outputDelimiters,
		funcName)

	if err != nil {
		fmt.Printf("Error from  ErrPrefixDto{}.NewIEmptyWithDelimiters()\n"+
			"%v\n", err.Error())
		return
	}

	ePDto2.SetMaxTextLineLen(maxLineLen)

	ePDtoColLen := ePDto.GetEPrefCollectionLen()

	ePDto2ColLen := ePDto2.GetEPrefCollectionLen()

	if ePDtoColLen != ePDto2ColLen {
		fmt.Printf("ERROR:\n"+
			"Expected ePDtoColLen == ePDto2ColLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePDtoColLen='%v'\n"+
			"ePDto2ColLen='%v'\n",
			ePDtoColLen,
			ePDto2ColLen)

		return
	}

	ePDtoStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto.String()),
		false)

	ePDto2Str := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(ePDto2.String()),
		false)

	if ePDtoStr != ePDto2Str {
		fmt.Printf("Error: Expected ePDtoStr==ePDto2Str.\n"+
			"However, THEY ARE NOT EQUAL!\n"+
			"ePDto=\n%v\n\nePDto2=\n%v\n\n",
			ePDtoStr,
			ePDto2Str)
		return
	}

	if !ePDto.Equal(ePDto2) {
		fmt.Printf("ERROR:\n" +
			"Expected that ePDto and ePDto2 would be equal\n" +
			"in all respects.\n" +
			"HOWEVER, THEY ARE NOT EQUAL\n")
	}

}

func (tMain TestMain) TestMain022() {

	ePDto := errpref.ErrPrefixDto{}.New()

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

	ePDto2 := errpref.ErrPrefixDto{}.New()

	_ = ePDto2.AddEPrefCollectionStr(
		initialStr)

	ePDto2.SetMaxTextLineLen(40)

	twoAreEqual :=
		ePDto.Equal(&ePDto2)

	if !twoAreEqual {
		fmt.Printf("FAILURE: ePDto and ePDto2 ARE NOT EQUAL!\n\n")
	} else {
		fmt.Printf("SUCCESS: ePDto and ePDto2 ARE EQUAL!\n\n")

	}

}

func (tMain TestMain) TestMain021() {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness : A=7 B=8 C=9"

	ePDto := errpref.ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(3)

	ePDto.SetLeftMarginChar('*')

	outputStr := ePDto.String()

	fmt.Printf("Left Margin Actual String:\n\n"+
		"%v\n\n\n",
		outputStr)

	expectedStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(outputStr),
		true)

	fmt.Printf("Left Margin Converted String:\n\n"+
		"%v",
		expectedStr)

}

func (tMain TestMain) TestMain020() {

	initialStr :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\n" +
			"Tx4() - Tx5() - Tx6.DoSomethingElse()\n" +
			"Tx7.TrySomethingNew() : something->newSomething\n" +
			"Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y - Tx10.X()\n" +
			"Tx11.TryAnything() - Tx12.TryASalad()\n" +
			"Tx13.SomeFabulousAndComplexStuff()\n" +
			"Tx14.MoreAwesomeGoodness() : A=7 B=8 C=9"

	ePDto := errpref.ErrPrefixDto{}.NewEPrefOld(initialStr)

	ePDto.SetMaxTextLineLen(40)

	ePDto.SetLeftMarginLength(3)

	ePDto.SetLeftMarginChar('*')

	outputStr := ePDto.String()

	fmt.Printf("Left Margin Length: %v\n",
		ePDto.GetLeftMarginLength())
	fmt.Printf("Left Margin Character: %v\n",
		string(ePDto.GetLeftMarginChar()))
	fmt.Printf("Left Margin Actual String:\n\n"+
		"%v\n\n\n",
		outputStr)

	expectedStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(outputStr),
		true)

	fmt.Printf("Left Margin Converted String:\n\n"+
		"%v",
		expectedStr)

}

func (tMain TestMain) TestMain019() {

	var err error
	//var ePDto2 *errpref.ErrPrefixDto

	_,
		err = errpref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"",
		"")

	if err == nil {
		fmt.Printf("Expected an error return from ErrPrefixDto{}.NewIEmpty(sbPtr)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func (tMain TestMain) TestMain018() {

	var sbPtr *strings.Builder

	var err error
	_,
		err = errpref.ErrPrefixDto{}.NewIEmpty(
		sbPtr,
		"",
		"")

	if err == nil {
		fmt.Printf("Expected an error return from ErrPrefixDto{}.NewIEmpty(sbPtr)\n" +
			"HOWEVER, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func (tMain TestMain) TestMain017() {

	var actualStr string
	//	ePDto := new(errpref.ErrPrefixDto)

	ePDto := errpref.ErrPrefixDto{}.Ptr()
	actualStr = ePDto.String()

	ePDto.SetEPref("Tx1.Something()")
	ePDto.SetEPref("Tx2.SomethingElse()")
	ePDto.SetEPrefCtx("Tx3.DoEverything()", "A==b^c")
	ePDto.SetEPrefCtx("Tx4()", "A/10==4")
	ePDto.SetEPref("Tx5.BrandNewMethod()")

	actualStr = ePDto.String()

	fmt.Printf("Actual String= '%v'\n",
		actualStr)

	fmt.Println("----------------------")

}

func (tMain *TestMain) TestMain016() {

	tFuncB1 := TestFuncB1{}

	err := tFuncB1.Tx1DoSomething()

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
	} else {
		fmt.Printf("Error: Expected an error return from\n" +
			"tFuncB1.Tx1DoSomething(). However, no error was returned!\n")
		fmt.Println()
		return
	}

	fmt.Println()
	fmt.Println("Error String Showing Non-Printable Characters")
	fmt.Println("---------------------------------------------")
	fmt.Println()

	expectedStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(err.Error()),
		true)

	fmt.Printf("%v\n",
		expectedStr)

	return
}

func (tMain *TestMain) TestMain015() {

	tFunc1 := testFunc1{}

	err := tFunc1.Tx1DoSomething()

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
	} else {
		return
	}

	expectedStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(err.Error()),
		true)

	fmt.Printf("%v\n",
		expectedStr)

	return
}

func (tMain *TestMain) TestMain014() {

	ePDto := new(errpref.ErrPrefixDto)

	ePDto.SetEPref("Tx1.Something()")
	ePDto.SetEPref("Tx2.SomethingElse()")
	ePDto.SetEPrefCtx("Tx4()", "A/10==4")
	ePDto.SetEPref("Tx5.BrandNewMethod()")

	err := fmt.Errorf("%v\n"+
		"Test Error Message!\n",
		ePDto)

	fmt.Println("TestMain014()")
	fmt.Println("-------------------------------------------")
	fmt.Println("Test Error Message #1 - Using v")
	fmt.Println("-------------------------------------------")

	fmt.Printf("%v",
		err.Error())

	fmt.Println()
	fmt.Println("-------------------------------------------")
	fmt.Println("Test Error Message #2 - Using s")
	fmt.Println("-------------------------------------------")

	fmt.Printf("%s",
		err.Error())

	fmt.Println()
	fmt.Println("-------------------------------------------")
	fmt.Println("Test Error Message #3 - Using String()")
	fmt.Println("-------------------------------------------")

	fmt.Printf("%s",
		err.Error())

}

func (tMain *TestMain) TestMain012() {

	ePDto := errpref.ErrPrefixDto{}.New()

	ePDto.SetMaxTextLineLen(40)

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

	expectedStr = errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(expectedStr),
		true)

	fmtActualStr := errpref.ErrPref{}.ConvertNonPrintableChars(
		[]rune(actualStr),
		true)

	var msg string
	if expectedStr != fmtActualStr {
		msg =
			fmt.Sprintf("Error:\n"+
				"Expected actualStr= '%v'\n"+
				"Instead, actualStr= '%v'\n",
				expectedStr,
				fmtActualStr)

		fmt.Printf(msg + "\n")

	} else {
		msg = fmt.Sprintf("**** Success ****\n"+
			"Expected String == Actual String!\n"+
			"Actual String=\n"+
			"'%v'\n", actualStr)
	}

}

func (tMain *TestMain) TestMain011() {

	funcName := "TestMain011()"

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	actualStr := errpref.ErrPref{}.FmtStr(
		initialStr)

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)
	fmt.Println(initialStr)
	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)
	fmt.Println()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()

}

func (tMain *TestMain) TestMain010() {

	funcName := "TestMain010()"

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5() : B==999"

	actualStr := errpref.ErrPref{}.EPref(
		initialStr,
		"")

	//expectedStr := "Tx1.AVeryVeryLongMethodNameCalledSomething()\\n" +
	//	"[SPACE]:[SPACE][SPACE]A->B\\n" +
	//	"Tx2.SomethingElse()[SPACE]:[SPACE]A==B\\n" +
	//	"Tx3.DoSomething()[SPACE]:[SPACE]A==10\\n" +
	//	"Tx4()[SPACE]:[SPACE]A/10==4[SPACE]-[SPACE]Tx5()[SPACE]:[SPACE]B==999"

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)
	fmt.Println(initialStr)
	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)
	fmt.Println()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()

}

func (tMain *TestMain) TestMain009() {

	funcName := "TestMain009()"

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	actualStr := errpref.ErrPref{}.SetCtxt(
		initialStr,
		"A!=B")

	initialStr += " : A!=B"

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)
	fmt.Println(initialStr)
	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)
	fmt.Println()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()

}

func (tMain *TestMain) TestMain008() {

	funcName := "TestMain008()"

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()"

	actualStr := errpref.ErrPref{}.GetLastEPref(initialStr)

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)
	fmt.Println(initialStr)
	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)
	fmt.Println()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()

}

func (tMain *TestMain) TestMain007() {

	funcName := "TestMain007()"

	initialStr :=
		"Tx1.AVeryVeryLongMethodNameCalledSomething() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	actualStr := errpref.ErrPref{}.FmtStr(
		initialStr)

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)
	fmt.Println(initialStr)
	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)
	fmt.Println()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()

}

func (tMain *TestMain) TestMain006() {

	funcName := "TestMain006()"

	initialStr :=
		"Tx1.Something() : A->B\nTx2.SomethingElse() : A==B\n" +
			"Tx3.DoSomething() : A==10\nTx4() : A/10==4 - Tx5()\nTx6.DoSomethingElse()\n"

	actualStr := errpref.ErrPref{}.FmtStr(
		initialStr)

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)
	fmt.Println(initialStr)
	fmt.Println()

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)
	fmt.Println()

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()

}
func (tMain *TestMain) TestMain005() {

	funcName := "TestMain005()"

	initialStr :=
		"Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	actualStr := errpref.ErrPref{}.FmtStr(
		initialStr)

	fmt.Println()
	fmt.Println(funcName)

	fmt.Println("--------------------------------------------------")

	fmt.Println("Initial String With Non-Printable Characters")

	fmt.Printf(initialStr)

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Non-Printable Characters")
	fmt.Println(actualStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Initial String With Printable Characters")

	tMain2 := TestMain{}

	initialStr = tMain2.ConvertNonPrintableChars(
		[]rune(initialStr), true, funcName)

	fmt.Println(initialStr)

	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Formatted String With Printable Characters")

	actualStr = tMain2.ConvertNonPrintableChars(
		[]rune(actualStr), true, funcName)
	fmt.Println(actualStr)

	fmt.Println()
	fmt.Println("--------------------------------------------------")
	fmt.Println()
	fmt.Println("Comparison: 'actualStr' vs expectedStr")
	fmt.Println(actualStr)
	fmt.Println(initialStr)
	fmt.Println()
}

func (tMain *TestMain) TestMain004() {

	xPref := errpref.ErrPref{}.EPrefCtx(
		"",
		"TestMain004()",
		"")

	ePrefix := errpref.ErrPref{}.EPrefCtx("",
		"Tx1.Something()",
		"A->B")

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx2.SomethingElse()",
		"(A+B)")

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx3.DoSomething()",
		"(A+B) + C = 9")

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx4()",
		"CopyPtr (A+B) -> C")

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx5.MoreAwesomeGoodness()",
		"(A+B) x C = D")

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx6.SomeFabulousAndComplexStuff()",
		"(A+B)^C = C x E")

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx8.TryAHammer()",
		"ErrNo: 5007-6004-9175")
	fmt.Println()
	fmt.Println(xPref)
	fmt.Println()

	tm2 := TestMain{}

	printableStr :=
		tm2.ConvertNonPrintableChars(
			[]rune(ePrefix),
			true, "")

	fmt.Println("Original Non-Printable Characters")

	fmt.Println(printableStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()

	ePrefix = errpref.ErrPref{}.FmtStr(ePrefix)

	printableStr =
		tm2.ConvertNonPrintableChars([]rune(ePrefix), true, "")

	fmt.Println("Formatted Non-Printable Characters")

	fmt.Println(printableStr)
	fmt.Println("--------------------------------------------------")
	fmt.Println()

	/*
		Un-formatted:

		   Formatted:

	*/
}

func (tMain *TestMain) TestMain003() (
	testResult string) {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	errpref.ErrPref{}.SetMaxErrPrefTextLineLength(60)

	ePrefix = errpref.ErrPref{}.EPrefCtx(ePrefix,
		"Tx7.TrySomethingNew()",
		"")

	ePrefix = errpref.ErrPref{}.SetCtxt(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx8.TryAnyCombination()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx9.TryAHammer()",
		"")

	ePrefix = errpref.ErrPref{}.SetCtxt(
		ePrefix,
		"x->y")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx10.X()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx11.TryAnything()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx12.TryASalad()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx13.SomeFabulousAndComplexStuff()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = errpref.ErrPref{}.FmtStr(ePrefix)

	fmt.Printf(ePrefix)

	tm2 := TestMain{}
	fmt.Println()
	fmt.Println()
	printableStr :=
		tm2.ConvertNonPrintableChars([]rune(ePrefix), false, "")

	return printableStr
}

func (tMain *TestMain) TestMain002() {
	tStr := "Test01%Test02%Test03%Test04%Test05%Test06"

	fmt.Printf("Test String: %v\n", tStr)

	tStrArray := strings.Split(tStr, "%")

	for _, s := range tStrArray {
		fmt.Printf("%v\n", s)
	}

}

func (tMain *TestMain) TestMain001() {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx7.TrySomethingNew()",
		"")

	ePrefix = errpref.ErrPref{}.SetCtxt(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx8.TryAnyCombination()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx9.TryAHammer()",
		"")

	ePrefix = errpref.ErrPref{}.SetCtxt(ePrefix, "x->y")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx10.X()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx11.TryAnything()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx12.TryASalad()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx13.SomeFabulousAndComplexStuff()",
		"")

	ePrefix = errpref.ErrPref{}.EPrefCtx(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = errpref.ErrPref{}.FmtStr(ePrefix)

	fmt.Printf(ePrefix)

	tm2 := TestMain{}
	fmt.Println()
	fmt.Println()
	printableStr :=
		tm2.ConvertNonPrintableChars([]rune(ePrefix), false, "")

	fmt.Printf(printableStr)
}

// ConvertNonPrintableChars - Receives a string containing
// non-printable characters and converts them to 'printable'
// characters returned in a string.
//
// Examples of non-printable characters are '\n', '\t' or 0x06
// (Acknowledge). These example characters would be translated into
// printable string characters as: "\\n", "\\t" and "[ACK]".
//
// Space characters are typically translated as " ". However, if
// the input parameter 'convertSpace' is set to 'true' then all
// spaces are converted to "[SPACE]" in the returned string.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters will be converted to
//       printable characters.
//
//  convertSpace        bool
//     - Space or white space characters (0x20) are by default
//       translated as " ". However, if this parameter is set to
//       'true', space characters will be converted to "[SPACE]".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  printableChars      string
//     - This returned string is identical to input parameter
//       'nonPrintableChars' with the exception that non-printable
//       characters are translated into printable characters.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "Hello world!\n"
//  testRunes := []rune(testStr)
//
//  actualStr :=
//    StrOps{}.NewPtr().
//      ConvertNonPrintableChars(testRunes, true)
//
//  ----------------------------------------------------
//  'actualStr' is now equal to:
//     "Hello[SPACE]world!\\n"
//
func (tMain *TestMain) ConvertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool,
	ePrefix string) (
	printableChars string) {

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "TestMain.convertNonPrintableChars() "

	lenNonPrintableChars := len(nonPrintableChars)

	if lenNonPrintableChars == 0 {
		return "[EMPTY]"
	}

	var b strings.Builder
	b.Grow(lenNonPrintableChars * 5)

	for i := 0; i < lenNonPrintableChars; i++ {
		cRune := nonPrintableChars[i]

		switch cRune {
		case 0:
			b.WriteString("[NULL]") // 0x00 NULL
		case 1:
			b.WriteString("[SOH]") // 0x01 State of Heading
		case 2:
			b.WriteString("[STX]") // 0x02 State of Text
		case 3:
			b.WriteString("[ETX]") // 0x03 End of Text
		case 4:
			b.WriteString("[EOT]") // 0X04 End of Transmission
		case 5:
			b.WriteString("[ENQ]") // 0x05 Enquiry
		case 6:
			b.WriteString("[ACK]") // 0x06 Acknowledge
		case '\a':
			b.WriteString("\\a") // U+0007 alert or bell
		case '\b':
			b.WriteString("\\b") // U+0008 backspace
		case '\t':
			b.WriteString("\\t") // U+0009 horizontal tab
		case '\n':
			b.WriteString("\\n") // U+000A line feed or newline
		case '\v':
			b.WriteString("\\v") // U+000B vertical tab
		case '\f':
			b.WriteString("\\f") // U+000C form feed
		case '\r':
			b.WriteString("\\r") // U+000D carriage return
		case 14:
			b.WriteString("[SO]") // U+000E Shift Out
		case 15:
			b.WriteString("[SI]") // U+000F Shift In
		case '\\':
			b.WriteString("\\") // U+005c backslash
		case ' ':
			// 0X20 Space character
			if convertSpace {
				b.WriteString("[SPACE]")
			} else {
				b.WriteRune(' ')
			}

		default:
			b.WriteRune(cRune)

		}

	}

	return b.String()
}

type testIStringerErrPref struct {
	locString string
}

func (tIStr *testIStringerErrPref) Set(str string) {
	tIStr.locString = str
}

func (tIStr *testIStringerErrPref) String() string {
	return tIStr.locString
}
