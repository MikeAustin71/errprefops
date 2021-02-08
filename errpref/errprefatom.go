package errpref

import (
	"strings"
	"sync"
)

type errPrefAtom struct {
	lock *sync.Mutex
}

// getEPrefContextArray - Receives an error prefix string
// containing a series of error prefix strings and error context
// strings. This method then proceeds to separate the error prefix
// and error context elements returning the results in an array of
// ErrorPrefixInfo objects passed as input parameter,
// 'prefixContextCol' .
//
func (ePrefAtom *errPrefAtom) getEPrefContextArray(
	errPrefix string,
	prefixContextCol []ErrorPrefixInfo) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	if prefixContextCol == nil {
		prefixContextCol =
			make([]ErrorPrefixInfo, 0, 150)
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

			prefixContextCol = append(prefixContextCol, element)

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

			prefixContextCol = append(prefixContextCol, element)

		}
	}

	lCollection = len(prefixContextCol)

	if lCollection > 0 {
		prefixContextCol[0].SetIsFirstIndex(true)

		prefixContextCol[lCollection-1].SetIsLastIndex(true)
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
