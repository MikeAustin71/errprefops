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
	prefixContextCol []ErrorPrefixDto) {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	prefixContextCol = make([]ErrorPrefixDto, 0, 150)

	if len(errPrefix) == 0 {
		return prefixContextCol
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	var lenCleanStr int

	errPrefix,
		lenCleanStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanStr == 0 {
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

	lastIdx := lCollection - 1

	isLastIdx := false

	//inLinePrefixDelimiter,
	//newLinePrefixDelimiter,
	//inLineContextDelimiter,
	//newLineContextDelimiter

	for i := 0; i < lCollection; i++ {

		if i == lastIdx {
			isLastIdx = true
		}

		s := errPrefixContextCollection[i]

		s,
			lenCleanStr = ePrefElectron.cleanErrorPrefixStr(s)

		if lenCleanStr == 0 {
			continue
		}

		contextIdx = strings.Index(s,
			delimiters.GetInLineContextDelimiter())

		element := ErrorPrefixDto{}.New()

		element.isValid = false

		if isLastIdx {
			element.isLastIdx = true
		} else {
			element.isLastIdx = false
		}

		if contextIdx == -1 {
			element.errorPrefixStr = s
			element.lenErrorPrefixStr = uint(len(s))
			element.errPrefixHasContextStr = false
			element.errorContextStr = ""
			element.lenErrorContextStr = 0

			prefixContextCol = append(prefixContextCol, element)

		} else {

			element.errorPrefixStr = s[0:contextIdx]
			element.lenErrorPrefixStr = uint(len(element.errorPrefixStr))
			element.errPrefixHasContextStr = false
			element.errorContextStr = s[contextIdx+
				idxLenInLineContextDelimiter:]
			element.lenErrorContextStr = uint(len(element.errorContextStr))

			prefixContextCol = append(prefixContextCol, element)
		}
	}

	return prefixContextCol
}
