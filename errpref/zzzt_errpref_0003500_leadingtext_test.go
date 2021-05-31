package errpref

import (
	"fmt"
	"strings"
	"testing"
)

func TestLeadingText_000100(t *testing.T) {

	ePDto := ErrPrefixDto{}.New()

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

	expectedStr :=
		"\\n--------------------------------------------------\\n" +
			"Tx1.Something()[SPACE]-[SPACE]Tx2.SomethingElse()" +
			"\\n--------------------------------------------------\\n" +
			"Error:[SPACE]Divide[SPACE]by[SPACE]Zero!"

	actualStr := ErrPref{}.ConvertNonPrintableChars(
		[]rune(outputStr),
		true)

	if expectedStr != actualStr {
		t.Errorf("ERROR Expected String DID NOT MATCH Actual String!\n"+
			"Expected Str: %v\n"+
			"  Actual Str: %v\n",
			expectedStr, actualStr)
	}

}
