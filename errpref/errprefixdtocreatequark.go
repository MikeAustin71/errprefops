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
	crEPrefDto *createErrPrefixDto) (
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
		newLenLastStr == 0 {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	strBuilder.WriteString(lastStr)

	if !crEPrefDto.isLastIdx {
		strBuilder.WriteString(
			crEPrefDto.newLinePrefixDelimiter)
	}

	newLastStr = ""
	newLenLastStr = 0
	newRemainingLineLen =
		crEPrefDto.maxErrStringLength

	return newLastStr, newLenLastStr, newRemainingLineLen
}
