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
	prefixContextCol []createErrPrefixDto) {

	if ePrefNeutron.lock == nil {
		ePrefNeutron.lock = new(sync.Mutex)
	}

	ePrefNeutron.lock.Lock()

	defer ePrefNeutron.lock.Unlock()

	prefixContextCol = make([]createErrPrefixDto, 0, 150)

	if len(errPrefix) == 0 {
		return prefixContextCol
	}

	ePrefQuark := errPrefQuark{}

	if maxErrStringLength == 0 {
		maxErrStringLength = ePrefQuark.getErrPrefDisplayLineLength()
	}

	inLinePrefixDelimiter,
		newLinePrefixDelimiter,
		inLineContextDelimiter,
		newLineContextDelimiter := ePrefQuark.getDelimiters()

	ePrefElectron := errPrefElectron{}

	var lenCleanStr int

	errPrefix,
		lenCleanStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanStr == 0 {
		return prefixContextCol
	}

	errPrefix = strings.ReplaceAll(
		errPrefix,
		newLineContextDelimiter,
		inLineContextDelimiter)

	errPrefix = strings.ReplaceAll(
		errPrefix,
		inLinePrefixDelimiter, newLinePrefixDelimiter)

	errPrefix += newLinePrefixDelimiter

	errPrefixContextCollection := strings.Split(
		errPrefix,
		newLinePrefixDelimiter)

	lCollection := len(errPrefixContextCollection)

	if lCollection == 0 {
		return prefixContextCol
	}

	var contextIdx int
	var idxLenInLineContextDelimiter = len(inLineContextDelimiter) - 1

	lenInLinePrefixDelimiter :=
		uint(len(inLinePrefixDelimiter))

	lenNewLinePrefixDelimiter :=
		uint(len(newLinePrefixDelimiter))

	lenInLineContextDelimiter :=
		uint(len(inLineContextDelimiter))

	lenNewLineContextDelimiter :=
		uint(len(newLineContextDelimiter))

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

		contextIdx = strings.Index(s, inLineContextDelimiter)

		element := createErrPrefixDto{}.New()

		element.errPrefIsEmpty = false

		if isLastIdx {
			element.isLastIdx = true
		} else {
			element.isLastIdx = false
		}

		element.inLinePrefixDelimiter = inLinePrefixDelimiter
		element.lenInLinePrefixDelimiter = lenInLinePrefixDelimiter
		element.newLinePrefixDelimiter = newLinePrefixDelimiter
		element.lenNewLinePrefixDelimiter = lenNewLinePrefixDelimiter
		element.inLineContextDelimiter = inLineContextDelimiter
		element.lenInLineContextDelimiter = lenInLineContextDelimiter
		element.newLineContextDelimiter = newLineContextDelimiter
		element.lenNewLineContextDelimiter = lenNewLineContextDelimiter
		element.totLenInLineEPrefWDelim = 0
		element.totLenInLineEPrefCtxWDelim = 0
		element.maxErrStringLength = maxErrStringLength

		if contextIdx == -1 {
			element.newErrPrefStr = s
			element.lenNewErrPrefStr = uint(len(s))
			element.funcHasContext = false
			element.newErrContextStr = ""
			element.lenNewErrContextStr = 0

			prefixContextCol = append(prefixContextCol, element)

		} else {

			element.newErrPrefStr = s[0:contextIdx]
			element.lenNewErrPrefStr = uint(len(element.newErrPrefStr))
			element.funcHasContext = false
			element.newErrContextStr = s[contextIdx+
				idxLenInLineContextDelimiter:]
			element.lenNewErrContextStr = uint(len(element.newErrContextStr))

			prefixContextCol = append(prefixContextCol, element)
		}
	}

	return prefixContextCol
}
