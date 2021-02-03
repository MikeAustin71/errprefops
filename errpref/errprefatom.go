package errpref

import (
	"strings"
	"sync"
)

type errPrefAtom struct {
	lock *sync.Mutex
}

// extractLastErrPrefInSeries - Receives a string containing error
// prefixes and proceeds to extract and return the last error
// prefix in the series along with the last error context and the
// modified error prefix series.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPref             string
//     - A string containing error prefixes in series. This method
//       will extract the last error prefix and error context from
//       this string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  oldErrPref          string
//     - This string holds the modified series of error prefixes.
//       It is identical to input parameter, 'errPref', except that
//       the last error prefix and error context have been removed.
//
//
//  newErrPref          string
//     - A string containing the last error prefix extracted from
//       input parameter, 'errPref'.
//
//
//  newErrContext       string
//     - A string containing the last error context description
//       extracted from input parameter, 'errPref'.
//
func (ePrefAtom *errPrefAtom) extractLastErrPrefInSeries(
	errPref string) (
	oldErrPref string,
	newErrPref string,
	newErrContext string) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	ePrefElectron := errPrefElectron{}
	var lenCleanStr int

	errPref,
		lenCleanStr =
		ePrefElectron.cleanErrorPrefixStr(errPref)

	if lenCleanStr == 0 {
		return oldErrPref, newErrPref, newErrContext
	}

	var lastIdxDashPref, lastIdxNewLinePref,
		lastPrefixIdx int

	lastIdxDashPref =
		strings.LastIndex(
			errPref,
			" - ")

	lastIdxNewLinePref =
		strings.LastIndex(
			errPref,
			"\n")

	if lastIdxDashPref > lastIdxNewLinePref {

		lastPrefixIdx = lastIdxDashPref

	} else if lastIdxNewLinePref > lastIdxDashPref {

		lastPrefixIdx = lastIdxNewLinePref

	} else {
		// lastIdxNewLinePref == lastIdxDashPref
		// This signals that there is only one
		// error prefix in the series.

		lastPrefixIdx = 0
	}

	if lastPrefixIdx > 0 {
		oldErrPref = errPref[0:lastPrefixIdx]

		oldErrPref,
			lenCleanStr =
			ePrefElectron.cleanErrorPrefixStr(oldErrPref)

	}

	tempNewErrPref := errPref[lastPrefixIdx:]

	tempNewErrPref,
		lenCleanStr =
		ePrefElectron.cleanErrorPrefixStr(tempNewErrPref)

	if lenCleanStr == 0 {
		return oldErrPref, newErrPref, newErrContext
	}

	var (
		lastIdxColonContext, lastIdxNewLineContext,
		lastContextIdx int
	)

	lastIdxNewLineContext =
		strings.LastIndex(
			tempNewErrPref,
			"\n : ")

	lastIdxColonContext =
		strings.LastIndex(
			tempNewErrPref,
			" : ")

	if lastIdxNewLineContext > lastIdxColonContext {

		lastContextIdx = lastIdxNewLineContext

	} else if lastIdxNewLineContext < lastIdxColonContext {

		lastContextIdx = lastIdxColonContext

	} else {
		//  lastIdxNewLineContext == lastIdxColonContext
		// This signals that there is no context
		// present in tempNewErrPref
		lastContextIdx = -1
	}

	if lastContextIdx > -1 {
		newErrPref = tempNewErrPref[0:lastContextIdx]
		newErrContext = tempNewErrPref[lastContextIdx:]

		newErrPref,
			_ =
			ePrefElectron.cleanErrorPrefixStr(newErrPref)

		newErrContext,
			_ =
			ePrefElectron.cleanErrorContextStr(newErrContext)

	} else {
		// lastContextIdx = -1
		// This signals that there is no context
		// present in tempNewErrPref
		newErrPref = tempNewErrPref
	}

	return oldErrPref, newErrPref, newErrContext
}

// ptr() - Returns a pointer to a new instance of errPrefAtom.
//
func (ePrefAtom errPrefAtom) ptr() *errPrefAtom {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	newErrPrefAtom := errPrefAtom{}

	return &newErrPrefAtom
}
