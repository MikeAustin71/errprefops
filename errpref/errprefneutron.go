package errpref

import (
	"strings"
	"sync"
)

type errPrefNeutron struct {
	lock *sync.Mutex
}

// errPrefix - Receives an error prefix string containing error
// prefix strings and error context strings. This method separates
// these two elements and returns them in a two-dimensional string
// array as a collection of error prefixes and error contexts.
func (ePrefNeutron *errPrefNeutron) getEPrefContextArray(
	errPrefix string) (
	prefixContextCol []ErrorPrefixInfo) {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	prefixContextCol = make([]ErrorPrefixInfo, 0, 150)

	if len(errPrefix) == 0 {
		return prefixContextCol
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	var (
		lenCleanPrefixStr,
		lenCleanContextStr int
	)

	errPrefix,
		lenCleanPrefixStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanPrefixStr == 0 {
		return prefixContextCol
	}

	errPrefix = strings.ReplaceAll(
		errPrefix,
		delimiters.GetNewLineContextDelimiter(),
		delimiters.GetInLineContextDelimiter())

	errPrefix = strings.ReplaceAll(
		errPrefix,
		delimiters.GetInLinePrefixDelimiter(),
		delimiters.GetNewLinePrefixDelimiter())

	errPrefix += delimiters.GetNewLinePrefixDelimiter()

	errPrefixContextCollection := strings.Split(
		errPrefix,
		delimiters.GetNewLinePrefixDelimiter())

	lCollection := len(errPrefixContextCollection)

	if lCollection == 0 {
		return prefixContextCol
	}

	var contextIdx int
	var idxLenInLineContextDelimiter = int(delimiters.GetLengthInLineContextDelimiter() - 1)
	var errCtxStr string

	lastIdx := lCollection - 1

	for i := 0; i < lCollection; i++ {

		s := errPrefixContextCollection[i]

		s,
			lenCleanPrefixStr = ePrefElectron.cleanErrorPrefixStr(s)

		if lenCleanPrefixStr == 0 {
			continue
		}

		contextIdx = strings.Index(s,
			delimiters.GetInLineContextDelimiter())

		element := ErrorPrefixInfo{}.New()

		if i == 0 {
			element.SetIsFirstIndex(true)
		} else {
			element.SetIsFirstIndex(false)
		}

		if i == lastIdx {
			element.SetIsLastIndex(true)
		} else {
			element.SetIsLastIndex(false)
		}

		if contextIdx == -1 {

			element.SetErrPrefixStr(s)

			element.SetErrPrefixHasContext(false)

			prefixContextCol = append(prefixContextCol, element)

		} else {

			element.SetErrPrefixStr(s[0:contextIdx])

			errCtxStr = s[contextIdx+
				idxLenInLineContextDelimiter:]

			errCtxStr,
				lenCleanContextStr = ePrefElectron.cleanErrorContextStr(
				errCtxStr)

			if lenCleanContextStr == 0 {
				element.SetErrPrefixHasContext(false)
				prefixContextCol = append(prefixContextCol, element)
				continue
			}

			element.SetErrPrefixHasContext(true)

			element.SetErrContextStr(errCtxStr)

			prefixContextCol = append(prefixContextCol, element)
		}

	}

	return prefixContextCol
}

// ptr() - Returns a pointer to a new instance of
// errPrefNeutron.
//
func (ePrefNeutron errPrefNeutron) ptr() *errPrefNeutron {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	return &errPrefNeutron{}
}

// writeCurrentLineStr - Writes the contents of a current
// line of error prefix and error context characters to
// a sting builder for text display output.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder          *strings.Builder
//     - A pointer to a string builder. The contents of the current
//       line of text will be written to this string builder.
//
//
//  ePrefLineLenCalc    *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc, the Error
//       Prefix Line Length Calculator. This types encapsulates all
//       the data necessary to perform line length calculations and
//       format the error prefix and error context strings for
//       text display output.
//
//       type EPrefixLineLenCalc struct {
//         ePrefDelimiters    ErrPrefixDelimiters
//         errorPrefixInfo    *ErrorPrefixInfo
//         currentLineStr     string
//         maxErrStringLength uint
//       }
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefNeutron *errPrefNeutron) writeCurrentLineStr(
	strBuilder *strings.Builder,
	ePrefLineLenCalc *EPrefixLineLenCalc) {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	if strBuilder == nil ||
		ePrefLineLenCalc == nil ||
		!ePrefLineLenCalc.IsValidInstance() {
		return
	}

	strBuilder.WriteString(
		ePrefLineLenCalc.GetCurrLineStr())

	if !ePrefLineLenCalc.IsErrPrefixLastIndex() {
		strBuilder.WriteString(
			ePrefLineLenCalc.GetDelimiterNewLineErrPrefix())
	}

	ePrefLineLenCalc.SetCurrentLineStr("")

	return
}
