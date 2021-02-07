package errpref

import "sync"

// ErrPrefixDto - Error Prefix Data Transfer Object. This type
// encapsulates and formats error prefix and error context strings.
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
// ErrPrefixDto is similar in function to type 'ErrPref'. However,
// unlike 'ErrPref', ErrPrefixDto is configured encapsulate and
// transmit error prefix information between methods. 'ErrPref'
// only returns strings. This type, ErrPrefixDto, returns an object
// which may be exchanged between methods.
//
type ErrPrefixDto struct {
	ePrefCol                   []ErrorPrefixInfo
	isTerminatedWithNewLine    bool
	maxErrPrefixTextLineLength uint
	lock                       *sync.Mutex
}

// New - Returns a new and properly initialized instance of
// ErrPrefixDto.
//
func (ePrefDto ErrPrefixDto) New() ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	return newErrPrefixDto
}

// SetCtx - Sets or resets the error context for the last error
// prefix. This operation either adds, or replaces, the error
// context string associated with the last error prefix the
// current list of error prefixes maintained by this instance.
//
// If the last error prefix already has an error context string, it
// will be replaced by input parameter, 'newErrContext'.
//
// If the last error prefix does NOT have an associated error
// context, this new error context string will be associated
// with that error prefix.
//
// If the internal list of error prefixes is currently empty, this
// method will take no action and exit.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrContext       string
//     - This string holds the new error context information. If
//       the last error prefix in internal storage already has an
//       associated error context, that context will be deleted and
//       replaced by 'newErrContext'. If, however, the last error
//       prefix does NOT have an associated error context, this
//       'newErrContext' string will be added and associated with
//       that last error prefix.
//
//       If this string is 'empty', this method will take no action
//       and exit.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetCtx(
	newErrContext string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
		return
	}

	if len(ePrefDto.ePrefCol) == 0 {
		return
	}

	errPrefNanobot{}.ptr().setLastCtx(
		newErrContext,
		ePrefDto.ePrefCol)
}

// SetEPref - Adds an error prefix string to the list of
// previous error prefix strings.
func (ePrefDto *ErrPrefixDto) SetEPref(
	newErrPrefix string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		"",
		ePrefDto.ePrefCol)

	return
}

// SetEPrefCtx - Adds an error prefix and an error context string
// to the list of previous error prefix/context information.
//
func (ePrefDto *ErrPrefixDto) SetEPrefCtx(
	newErrPrefix string,
	newErrContext string) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	errPrefNanobot{}.ptr().addEPrefInfo(
		newErrPrefix,
		newErrContext,
		ePrefDto.ePrefCol)

	return
}

// String - Returns a formatted error prefix/context string
// incorporating all error prefixes previously added to this
// instance.
//
// Error prefix information is stored internally in the 'ePrefCol'
// array.
//
func (ePrefDto *ErrPrefixDto) String() string {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	return ""
}
