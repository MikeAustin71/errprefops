package errpref

import (
	"strings"
	"sync"
)

type errorPrefixDtoQuark struct {
	lock *sync.Mutex
}

func (errPrefixDtoQuark *errorPrefixDtoQuark) writeLastStr(
	strBuilder *strings.Builder,
	lastStr string,
	remainingLineLen uint,
	crEPrefDto *ErrorPrefixDto,
	delimiters *EPrefixDelimiters) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if errPrefixDtoQuark.lock == nil {
		errPrefixDtoQuark.lock = new(sync.Mutex)
	}

	errPrefixDtoQuark.lock.Lock()

	defer errPrefixDtoQuark.lock.Unlock()

	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))
	newRemainingLineLen = remainingLineLen

	if strBuilder == nil ||
		crEPrefDto == nil ||
		delimiters == nil ||
		newLenLastStr == 0 {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	strBuilder.WriteString(lastStr)

	if !crEPrefDto.isLastIdx {
		strBuilder.WriteString(
			delimiters.GetNewLinePrefixDelimiter())
	}

	newLastStr = ""
	newLenLastStr = 0
	newRemainingLineLen =
		delimiters.GetMaxErrStringLength()

	return newLastStr, newLenLastStr, newRemainingLineLen
}
