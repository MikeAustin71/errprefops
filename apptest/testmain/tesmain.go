package testmain

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
	"strings"
)

type TestMain struct {
	testStr01 string
}

func (tMain *TestMain) TestMain003() (
	testResult string) {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	errpref.ErrPref{}.SetMaxErrPrefTextLineLength(60)

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx8.TryAnyCombination()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx9.TryAHammer()")
	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "x->y")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx10.X()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx11.TryAnything()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx12.TryASalad()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx13.SomeFabulousAndComplexStuff()")
	ePrefix = errpref.ErrPref{}.NewContext(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = errpref.ErrPref{}.FmtString(ePrefix)

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

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx8.TryAnyCombination()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx9.TryAHammer()")
	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "x->y")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx10.X()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx11.TryAnything()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx12.TryASalad()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx13.SomeFabulousAndComplexStuff()")
	ePrefix = errpref.ErrPref{}.NewContext(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = errpref.ErrPref{}.FmtString(ePrefix)

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
