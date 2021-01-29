package errpref

import (
	"sync"
)

const constDefaultErrPrefLineLength = uint(40)

var defaultErrorPrefixDisplayLineLength uint

var defaultErrPrefLineLenLock *sync.Mutex

type errPrefQuark struct {
	lock *sync.Mutex
}

// getErrPrefDisplayLineLength - Returns the current value for the
// maximum number of characters allowed in an error prefix text
// display line.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method will return an integer value specifying the
//       maximum number of characters allowed in an error prefix
//       text display line.
//
func (ePrefQuark *errPrefQuark) getErrPrefDisplayLineLength() uint {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if defaultErrPrefLineLenLock == nil {
		defaultErrPrefLineLenLock = new(sync.Mutex)
	}

	var maxErrPrefixStringLength uint

	defaultErrPrefLineLenLock.Lock()

	defer defaultErrPrefLineLenLock.Unlock()

	if defaultErrorPrefixDisplayLineLength == 0 {

		defaultErrorPrefixDisplayLineLength =
			constDefaultErrPrefLineLength

	}

	maxErrPrefixStringLength =
		defaultErrorPrefixDisplayLineLength

	return maxErrPrefixStringLength
}

// isEmptyOrWhiteSpace - Analyzes the incoming string and returns
// 'true' if the strings is empty or consists of all white space.
//
func (ePrefQuark *errPrefQuark) isEmptyOrWhiteSpace(
	targetStr string) bool {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}

// setErrPrefDisplayLineLength - Sets the value of the maximum
// error prefix line length. This maximum limit controls the length
// of text lines produced for display of error prefix information.
//
// This method will set a global variable to the line length
// specified by input parameter 'maxErrPrefixTextLineLength'. Once
// this maximum limit is set it will control the line lengths for
// all instances of the 'ErrPref' type unless or until this limit
// is reset to another value by calling this method.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  maxErrPrefixTextLineLength      uint
//     - This unsigned integer value will be used to set the
//       maximum number of characters allowed in a text display
//       line for error prefix information.
//
//       If 'maxErrPrefixTextLineLength' is set to a value of zero
//       (0), this method will take no action and return.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
func (ePrefQuark *errPrefQuark) setErrPrefDisplayLineLength(
	maxErrPrefixStringLength uint) {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	if defaultErrPrefLineLenLock == nil {
		defaultErrPrefLineLenLock = new(sync.Mutex)
	}

	defaultErrPrefLineLenLock.Lock()

	defer defaultErrPrefLineLenLock.Unlock()

	if maxErrPrefixStringLength == 0 {
		return
	}

	defaultErrorPrefixDisplayLineLength =
		maxErrPrefixStringLength

	return
}
