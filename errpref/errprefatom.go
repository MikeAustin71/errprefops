package errpref

import (
	"strings"
	"sync"
)

type errPrefAtom struct {
	lock *sync.Mutex
}

// getEPrefContextArray - Receives a string containing a series of
// error prefix strings and error context strings. This method then
// proceeds to parse the error prefix and error context elements
// returning the results in an array of ErrorPrefixInfo objects
// passed as input parameter, 'prefixContextCol' .
//
// All error prefix elements parsed from the 'errPrefix' string
// will be appended to 'prefixContextCol'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefix           string
//     - This string contains a series of error prefix and error
//       context strings. Input parameter 'errPrefix' will be
//       parsed into error prefix and error context components
//       before being appended to array of ErrorPrefixInfo objects
//       passed in input parameter, 'prefixContextCol'
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited by either a
//       new line character ('\n') or the in-line delimiter string,
//       " - ".
//
//       If this string contains associated error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
//  prefixContextCol    *[]ErrorPrefixInfo
//     - A pointer to an array of ErrorPrefixInfo objects. The
//       error prefix and error context data elements parsed from
//       input parameter, 'errPrefix', will be appended to this
//       array.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefAtom *errPrefAtom) getEPrefContextArray(
	errPrefix string,
	prefixContextCol *[]ErrorPrefixInfo) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	if *prefixContextCol == nil {
		*prefixContextCol = make([]ErrorPrefixInfo, 0, 256)
	}

	if len(errPrefix) == 0 {
		return
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	var lenCleanPrefixStr int

	errPrefix,
		lenCleanPrefixStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanPrefixStr == 0 {
		return
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
		return
	}

	var contextIdx int
	var idxLenInLineContextDelimiter = int(delimiters.GetLengthInLineContextDelimiter() - 1)
	var (
		errPrefixStr,
		errCtxStr string
	)

	ePrefNeutron := errPrefNeutron{}

	for i := 0; i < lCollection; i++ {

		s := errPrefixContextCollection[i]

		contextIdx = strings.Index(s,
			delimiters.GetInLineContextDelimiter())

		if contextIdx == -1 {

			element,
				err := ePrefNeutron.createNewEPrefInfo(
				s,
				"",
				"")

			if err != nil {
				continue
			}

			*prefixContextCol = append(*prefixContextCol, element)

		} else {

			errPrefixStr =
				s[0:contextIdx]

			errCtxStr = s[contextIdx+
				idxLenInLineContextDelimiter:]

			element,
				err := ePrefNeutron.createNewEPrefInfo(
				errPrefixStr,
				errCtxStr,
				"")

			if err != nil {
				continue
			}

			*prefixContextCol = append(*prefixContextCol, element)

		}
	}

	lCollection = len(*prefixContextCol)

	return
}

// setFlagsErrorPrefixInfoArray - Sets internal flags in an array
// of ErrorPrefixInfo objects
//
func (ePrefAtom *errPrefAtom) setFlagsErrorPrefixInfoArray(
	prefixContextCol []ErrorPrefixInfo) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	if prefixContextCol == nil {
		prefixContextCol = make([]ErrorPrefixInfo, 0, 256)
	}

	lenCollection := len(prefixContextCol)

	if lenCollection == 0 {
		return
	}

	lastIdx := lenCollection - 1

	for i := 0; i < lenCollection; i++ {

		_ = prefixContextCol[i].GetIsPopulated()

		switch i {
		case 0:
			prefixContextCol[i].SetIsFirstIndex(true)
			prefixContextCol[i].SetIsLastIndex(false)
		case lastIdx:
			prefixContextCol[i].SetIsFirstIndex(false)
			prefixContextCol[i].SetIsLastIndex(true)
		default:
			prefixContextCol[i].SetIsFirstIndex(false)
			prefixContextCol[i].SetIsLastIndex(false)
		}
	}

	return
}

// ptr() - Returns a pointer to a new instance of
// errPrefAtom.
//
func (ePrefAtom errPrefAtom) ptr() *errPrefAtom {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	return &errPrefAtom{}
}
