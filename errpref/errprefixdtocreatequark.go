package errpref

import (
	"strings"
	"sync"
)

type createErrPrefixDtoQuark struct {
	lock *sync.Mutex
}

func (createEPrefDtoQuark *createErrPrefixDtoQuark) writeLastStr(
	strBuilder *strings.Builder,
	lastStr string,
	remainingLineLen uint,
	crEPrefDto *createErrPrefixDto,
	delimiters *EPrefixDelimiters) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if createEPrefDtoQuark.lock == nil {
		createEPrefDtoQuark.lock = new(sync.Mutex)
	}

	createEPrefDtoQuark.lock.Lock()

	defer createEPrefDtoQuark.lock.Unlock()

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
