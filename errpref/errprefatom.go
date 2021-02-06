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
// and error context elements returning the results in an array as
// a collection of ErrorPrefixInfo objects.
//
func (ePrefAtom *errPrefAtom) getEPrefContextArray(
	errPrefix string) (
	prefixContextCol []ErrorPrefixInfo) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

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
