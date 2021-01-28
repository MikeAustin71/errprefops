package errpref

import (
	"strings"
	"sync"
)

type errPrefElectron struct {
	lock *sync.Mutex
}

// cleanErrorPrefixStr - Receives error prefix strings (not error
// context strings) and proceeds to scrub and remove invalid
// sub-strings and characters. After this process is completed, the
// 'cleaned' string along with the length of the newly cleaned
// string is returned to the calling function.
//
//
func (ePrefElectron *errPrefElectron) cleanErrorPrefixStr(
	errPref string) (
	cleanStr string,
	lenCleanStr int) {

	if ePrefElectron.lock == nil {
		ePrefElectron.lock = new(sync.Mutex)
	}

	ePrefElectron.lock.Lock()

	defer ePrefElectron.lock.Unlock()

	lenCleanStr = 0
	cleanStr = ""

	ePrefQuark := errPrefQuark{}

	if ePrefQuark.isEmptyOrWhiteSpace(errPref) {
		return cleanStr, lenCleanStr
	}

	dirtyChars := []string{
		"\n",
		" ",
		" - ",
	}

	lenDirtyChars := len(dirtyChars)

	for i := 0; i < 3; i++ {
		for j := 0; j < lenDirtyChars; j++ {
			errPref = strings.TrimLeft(strings.TrimRight(errPref, dirtyChars[j]), dirtyChars[j])
		}
	}

	if ePrefQuark.isEmptyOrWhiteSpace(errPref) {
		return cleanStr, lenCleanStr
	}

	cleanStr = errPref
	lenCleanStr = len(errPref)
	return cleanStr, lenCleanStr
}

// cleanErrorContextStr - Receives error context strings (not error
// prefix strings) and proceeds to scrub and remove invalid
// sub-strings and characters. After this process is completed, the
// 'cleaned' string along with the length of the newly cleaned
// string is returned to the calling function.
//
//
func (ePrefElectron *errPrefElectron) cleanErrorContextStr(
	errContext string) (
	cleanStr string,
	lenCleanStr int) {

	if ePrefElectron.lock == nil {
		ePrefElectron.lock = new(sync.Mutex)
	}

	ePrefElectron.lock.Lock()

	defer ePrefElectron.lock.Unlock()

	lenCleanStr = 0
	cleanStr = ""

	ePrefQuark := errPrefQuark{}

	if ePrefQuark.isEmptyOrWhiteSpace(errContext) {
		return cleanStr, lenCleanStr
	}

	dirtyChars := []string{
		"\n",
		" ",
		" : ",
	}

	lenDirtyChars := len(dirtyChars)

	for i := 0; i < 3; i++ {
		for j := 0; j < lenDirtyChars; j++ {
			errContext = strings.TrimLeft(strings.TrimRight(errContext, dirtyChars[j]), dirtyChars[j])
		}
	}

	if ePrefQuark.isEmptyOrWhiteSpace(errContext) {
		return cleanStr, lenCleanStr
	}

	cleanStr = errContext
	lenCleanStr = len(errContext)
	return cleanStr, lenCleanStr
}
