package errpref

import (
	"strings"
	"sync"
)

type ErrPref struct {
	maxErrStringLength uint
	lock               *sync.Mutex
}

// New - Returns a string concatenating the old error prefix the
// new custom, user-defined error prefix. This is typically used
// to create method or function chains. The old error prefix
// contains the function chain which led to a a new function
// and the new error prefix contains the name of the currently
// executing method or function.
//
func (ePref ErrPref) New(
	oldErrPref string,
	newErrPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	if ePref.maxErrStringLength == 0 {
		ePref.maxErrStringLength = 40
	}

	newErrPref = strings.TrimLeft(strings.TrimRight(newErrPref, " "), " ")
	newErrPref = strings.TrimRight(newErrPref, "\n")

	if len(newErrPref) == 0 {
		return oldErrPref
	}

	oldErrPref = strings.TrimLeft(strings.TrimRight(oldErrPref, " "), " ")
	oldErrPref = strings.TrimRight(oldErrPref, "\n")

	if len(oldErrPref) == 0 {
		return newErrPref
	}

	if uint(len(oldErrPref)+len(newErrPref)+3) > ePref.maxErrStringLength {
		return oldErrPref + "\n" + newErrPref
	}

	return oldErrPref + " - " + newErrPref
}

func (ePref ErrPref) NewContext(
	oldErrPref string,
	newErrPref string,
	newContext string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	if ePref.maxErrStringLength == 0 {
		ePref.maxErrStringLength = 40
	}

	if len(newErrPref) == 0 {
		return oldErrPref
	}

	newErrPref = strings.TrimLeft(strings.TrimRight(newErrPref, " "), " ")

	newErrPref = strings.TrimRight(newErrPref, "\n")

	if len(newErrPref) == 0 {
		return oldErrPref
	}

	var newEPrefContext string

	if len(newContext) == 0 {
		newEPrefContext = newErrPref
	} else {

		newContext = strings.TrimLeft(strings.TrimRight(newContext, " "), " ")

		newContext = strings.TrimRight(newContext, "\n")

		if len(newContext) == 0 {
			newEPrefContext = newErrPref
		} else {

			if uint(len(newErrPref)+len(newContext)+3) > ePref.maxErrStringLength {
				newEPrefContext = newErrPref + "\n : " + newContext
			} else {
				newEPrefContext = newErrPref + " : " + newContext
			}
		}
	}

	return oldErrPref + "\n" + newEPrefContext
}

func (ePref ErrPref) AddContext(
	errPref string,
	errContext string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	if ePref.maxErrStringLength == 0 {
		ePref.maxErrStringLength = 40
	}

	if len(errPref) == 0 {
		return errContext
	}

	if len(errContext) == 0 {
		return errPref
	}

	errPref = strings.TrimLeft(strings.TrimRight(errPref, " "), " ")

	errPref = strings.TrimRight(errPref, "\n")

	if len(errPref) == 0 {
		return errContext
	}

	errContext = strings.TrimLeft(strings.TrimRight(errContext, " "), " ")

	errContext = strings.TrimRight(errContext, "\n")

	if len(errContext) == 0 {
		return errPref
	}

	if uint(len(errPref)+len(errContext)+3) > ePref.maxErrStringLength {
		return errPref + "\n : " + errContext
	}

	return errPref + " : " + errContext
}

func (ePref ErrPref) FmtString(errPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	if ePref.maxErrStringLength == 0 {
		ePref.maxErrStringLength = 40
	}

	ePrefQuark := errPrefQuark{}

	return ePrefQuark.formatErrPrefix(
		ePref.maxErrStringLength,
		errPref)
}
