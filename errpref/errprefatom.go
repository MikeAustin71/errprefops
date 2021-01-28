package errpref

import (
	"strings"
	"sync"
)

type errPrefAtom struct {
	lock *sync.Mutex
}

func (ePrefAtom *errPrefAtom) getLastErrPrefInSeriesInfo(
	oldErrPref string) (
	LastPrefixIdx int,
	LastPrefixSeparator string,
	LastContextIdx int,
	LastContextSeparator string) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	LastPrefixIdx = -1
	LastContextIdx = -1

	var lastIdxDashPref, lastIdxNewLinePref int

	lastIdxDashPref =
		strings.LastIndex(
			oldErrPref,
			" - ")

	lastIdxNewLinePref =
		strings.LastIndex(
			oldErrPref,
			"\n")

	if lastIdxDashPref > lastIdxNewLinePref {

		LastPrefixIdx = lastIdxDashPref

		LastPrefixSeparator =
			oldErrPref[LastPrefixIdx : LastPrefixIdx+3]

	} else if lastIdxNewLinePref > lastIdxDashPref {
		LastPrefixIdx = lastIdxNewLinePref
		LastPrefixSeparator =
			oldErrPref[LastPrefixIdx : LastPrefixIdx+1]

	} else {
		// lastIdxNewLinePref == lastIdxDashPref
		// This signals Last Error Prefix Does Not
		//   Exist!
		return LastPrefixIdx,
			LastPrefixSeparator,
			LastContextIdx,
			LastContextSeparator
	}

	return LastPrefixIdx,
		LastPrefixSeparator,
		LastContextIdx,
		LastContextSeparator

}
