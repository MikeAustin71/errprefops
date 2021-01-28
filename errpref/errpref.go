package errpref

import (
	"strings"
	"sync"
)

// ErrPref - This type is provides methods useful in formatting
// error prefix and error context strings.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to
// document the thread of code execution by listing the calling
// sequence for specific functions and methods.
//
// The error context string is designed to provide additional
// error context information associated with the currently
// executing function or method. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
type ErrPref struct {
	maxErrStringLength uint
	lock               *sync.Mutex
}

// New - Returns a string concatenating the old error prefix the
// new custom, user-defined error prefix. The new error prefix is
// typically used to document method or function chains in error
// messages.
//
// The old error prefix contains the function chain which led to
// the function next in line for execution.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to
// document the thread of code execution by listing the calling
// sequence for specific functions and methods.
//
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

	ePrefMech := errPrefMechanics{}

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		"",
		ePref.maxErrStringLength)
}

// NewContext - Receives an old error prefix, new error prefix and
// a new context string which are concatenated and returned as a
// combined string.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to
// document the thread of code execution by listing the calling
// sequence for specific functions and methods.
//
// The error context string is designed to provide additional
// error context information associated with the currently
// executing function or method. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref          string
//     - This includes the previous error prefix string. This string
//       will be formatted and concatenated with the new error prefix
//       and the associated error prefix.
//
//
//  newErrPref          string
//     - The new error prefix represents the error prefix string
//       associated with the function or method which is currently
//       executing. This parameter is optional and will accept an
//       empty string, but there isn't much point in calling this
//       method without a substantive value for 'newErrPref'.
//
//
//  newContext          string
//     - This is the error context information associated with the
//       new error prefix ('newErrPref'). This parameter is
//       optional and will accept an empty string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method will return the consolidated error prefix text.
//
//       The error prefix text is designed to be configured at the
//       beginning of error messages and is most often used to
//       document the thread of code execution by listing the calling
//       sequence for specific functions and methods.
//
//
//
// -----------------------------------------------------------------
//
// Usage Examples
//
//  errorPrefix = ErrPref{}.NewContext(
//                           errorPrefix, // Assuming this is the old
//                                        // error prefix
//                           newErrPref,
//                           newContext)
//
//
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

	ePrefMech := errPrefMechanics{}

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		newContext,
		ePref.maxErrStringLength)
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

	ePrefNanobot := errPrefNanobot{}

	return ePrefNanobot.formatErrPrefix(
		ePref.maxErrStringLength,
		errPref)
}
