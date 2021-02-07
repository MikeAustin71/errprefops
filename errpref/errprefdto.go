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
	ePrefCol                []ErrorPrefixInfo
	isTerminatedWithNewLine bool
	lock                    *sync.Mutex
}

// New - Returns a new instance of ErrPrefixDto.
//
func (ePrefDto ErrPrefixDto) New() ErrPrefixDto {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	return newErrPrefixDto
}

// SetEPref - Adds an error context string to the list of
// previous error context strings.
func (ePrefDto *ErrPrefixDto) SetEPref()

// String - Returns a formatted error prefix/context string
// incorporating all error prefixes previously added to this
// instance.
//
// Error prefix information is stored internally in the 'ePrefCol'
// array.
//
func (ePrefDto *ErrPrefixDto) String() string {

}
