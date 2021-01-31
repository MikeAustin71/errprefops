package errpref

import (
	"sync"
)

type createErrPrefixDto struct {
	errPrefIsEmpty             bool
	isLastIdx                  bool // Signals the last index in the array
	newErrPrefStr              string
	lenNewErrPrefStr           uint
	funcHasContext             bool
	newErrContextStr           string
	lenNewErrContextStr        uint
	inLinePrefixDelimiter      string
	lenInLinePrefixDelimiter   uint
	newLinePrefixDelimiter     string
	lenNewLinePrefixDelimiter  uint
	inLineContextDelimiter     string
	lenInLineContextDelimiter  uint
	newLineContextDelimiter    string
	lenNewLineContextDelimiter uint
	totLenInLineEPrefWDelim    uint
	totLenInLineEPrefCtxWDelim uint
	maxErrStringLength         uint
	lock                       *sync.Mutex
}

// New - Creates and returns a new instance of type,
// 'createErrPrefixDto'.
//
func (crErrPrefDto createErrPrefixDto) New() createErrPrefixDto {

	if crErrPrefDto.lock == nil {
		crErrPrefDto.lock = new(sync.Mutex)
	}

	crErrPrefDto.lock.Lock()

	defer crErrPrefDto.lock.Unlock()

	return createErrPrefixDto{}
}
