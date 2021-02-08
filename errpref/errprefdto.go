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

// AddEPrefCollectionStr - Receives a string containing one or more
// error prefix and error context pairs. This error prefix
// information is parsed and added to the internal store of
// error prefix elements.
//
// Upon completion this method returns an integer value identifying
// the number of error prefix elements successfully parsed and
// added to internal error prefix storage.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefixCollectionStr      string
//     - A string consisting of a series of error prefix strings.
//       Error prefixes should be delimited by either a new line
//       character ('\n') or the in-line delimiter string, " - ".
//
//       If the collection string contains error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  numberOfCollectionItemsParsed int
//     - This returned integer value contains the number of error
//       prefix elements parsed from input parameter
//       'errorPrefixCollection' and added to the internal store
//       of error prefix elements.
//
func (ePrefDto *ErrPrefixDto) AddEPrefCollectionStr(
	errorPrefixCollectionStr string) (
	numberOfCollectionItemsParsed int) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)
	}

	if ePrefDto.maxErrPrefixTextLineLength < 10 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	}

	previousCollectionLen := len(ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().getEPrefContextArray(
		errorPrefixCollectionStr,
		ePrefDto.ePrefCol)

	numberOfCollectionItemsParsed =
		len(ePrefDto.ePrefCol) -
			previousCollectionLen

	return numberOfCollectionItemsParsed
}

// EmptyEPrefCollection - All instances of ErrPrefixDto store
// error prefix information internally in an array of
// ErrorPrefixInfo objects. This method will delete or empty the
// current ErrorPrefixInfo array and initialize it to a zero length
// array.
//
// Effectively, this will delete all existing error prefix
// information stored in the current ErrPrefixDto instance.
//
//
func (ePrefDto *ErrPrefixDto) EmptyEPrefCollection() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 256)
}

// GetIsTerminatedWithNewLine - Returns the current setting which
// determines whether error prefix strings returned by this
// ErrPrefixDto instance will be terminated with a new line
// character ('\n').
//
// By default, error prefix strings returned by the method
// ErrPrefixDto.String() do NOT have a new line character as the
// last character in the error prefix string. In other words,
// returned error prefix strings are NOT terminated with new
// line characters ('\n') by default.
//
// If error prefix strings terminated with a new line character is
// a preferred option, use the method 'SetIsTerminatedWithNewLine()'
// to configure this feature.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
// -----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If this return parameter is set to 'true', all error
//       prefix strings returned by this ErrPrefixDto instance WILL
//       BE terminated with a new line character ('\n').
//
func (ePrefDto *ErrPrefixDto) GetIsTerminatedWithNewLine() bool {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	return ePrefDto.isTerminatedWithNewLine
}

// GetErrorPrefixCollection - Returns a deep copy of the current
// error prefix collection maintained by this ErrPrefixDto
// instance.
//
func (ePrefDto *ErrPrefixDto) GetErrorPrefixCollection() []ErrorPrefixInfo {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	var newErrorPrefixCollection []ErrorPrefixInfo

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol = make([]ErrorPrefixInfo, 0, 256)
		newErrorPrefixCollection = make([]ErrorPrefixInfo, 0, 256)
		return newErrorPrefixCollection
	}

	lenEPrefCol := len(ePrefDto.ePrefCol)

	if lenEPrefCol == 0 {
		newErrorPrefixCollection = make([]ErrorPrefixInfo, 0, 256)
		return newErrorPrefixCollection
	}

	newErrorPrefixCollection =
		make([]ErrorPrefixInfo,
			lenEPrefCol,
			lenEPrefCol+256)

	_ = copy(newErrorPrefixCollection, ePrefDto.ePrefCol)

	return newErrorPrefixCollection
}

// GetMaxErrPrefTextLineLength - Returns the maximum limit on the
// number of characters allowed in an error prefix text line output
// for display purposes.
//
// This unsigned integer value controls the maximum character limit
// for this specific ErrPrefixDto instance. This maximum limit will
// remain in force for the life of this ErrPrefixDto instance or
// until a call is made to the method 'SetMaxErrPrefTextLineLength()'
// which is used to change the value.
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
//  uint
//     - This unsigned integer specifies the maximum limit on the
//       number of characters allowed in an error prefix text line
//       output for display purposes.
//
//
func (ePrefDto *ErrPrefixDto) GetMaxErrPrefTextLineLength() uint {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.maxErrPrefixTextLineLength < 10 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	}

	return ePrefDto.maxErrPrefixTextLineLength
}

// MergeErrPrefixDto - Receives a pointer to another ErrPrefixDto
// object and proceeds to add the error prefix collection to that
// of the current ErrPrefixDto instance.
//
// This means that the error prefix collection for
// 'incomingErrPrefixDto' will be added to the end of the
// ePrefDto collection.
//
func (ePrefDto *ErrPrefixDto) MergeErrPrefixDto(
	incomingErrPrefixDto *ErrPrefixDto) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if ePrefDto.ePrefCol == nil {
		ePrefDto.ePrefCol =
			make([]ErrorPrefixInfo, 0, 256)
	}

	if incomingErrPrefixDto.ePrefCol == nil {
		incomingErrPrefixDto.ePrefCol =
			make([]ErrorPrefixInfo, 0, 256)
		return
	}

	lenIncomingDto :=
		len(incomingErrPrefixDto.ePrefCol)

	if lenIncomingDto == 0 {
		return
	}

	ePrefDto.ePrefCol =
		append(ePrefDto.ePrefCol,
			incomingErrPrefixDto.ePrefCol...)

}

// New - Returns a new and properly initialized instance of
// ErrPrefixDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  ErrPrefixDto
//     - This method will return a new, properly initialized
//       instance of ErrPrefixDto.
//
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

// NewEPrefCollection - Returns a new and properly initialized
// instance of ErrPrefixDto. This method receives a string
// containing one or more error prefix and error context pairs.
// This error prefix information is parsed and stored internally
// as individual error prefix elements.
//
// Upon completion this method returns an integer value identifying
// the number of error prefix elements successfully parsed and
// stored for future use.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefixCollection         string
//     - A string consisting of a series of error prefix strings.
//       Error prefixes should be delimited by either a new line
//       character ('\n') or the in-line delimiter string, " - ".
//
//       If the collection string contains error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  numberOfCollectionItemsParsed int
//     - This returned integer value contains the number of error
//       prefix elements parsed from input parameter
//       'errorPrefixCollection' and stored internally for future
//       use.
//
//
//  newErrPrefixDto               ErrPrefixDto
//     - This method will return a newly created instance of
//       ErrPrefixDto configured with the error prefix information
//       parsed from input parameter, 'errorPrefixCollection'.
//
//
func (ePrefDto ErrPrefixDto) NewEPrefCollection(
	errorPrefixCollection string) (
	numberOfCollectionItemsParsed int,
	newErrPrefixDto ErrPrefixDto) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	newErrPrefixDto.ePrefCol = make([]ErrorPrefixInfo, 0, 100)

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	previousCollectionLen := len(ePrefDto.ePrefCol)

	errPrefAtom{}.ptr().getEPrefContextArray(
		errorPrefixCollection,
		ePrefDto.ePrefCol)

	numberOfCollectionItemsParsed =
		len(ePrefDto.ePrefCol) -
			previousCollectionLen

	return numberOfCollectionItemsParsed, newErrPrefixDto
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
// previous error prefix strings maintained by this instance of
// ErrPrefixDto.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to document
// the thread of code execution by listing the calling sequence for
// specific functions and methods.
//
// This method is designed to process a single error prefix string
// passed in input parameter 'ErrPrefixDto'. If this string
// contains multiple error prefixes, use method
// 'ErrPrefixDto.SetEPrefCollection()'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This new error prefix will be added to the internal list
//       of error prefixes maintained by this ErrPrefixDto
//       instance.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
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
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context string is designed to provide additional
// information about the function or method identified by the
// associated error prefix string. Typical context information
// might include variable names, variable values and additional
// details on function execution.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPrefix        string
//     - The new error prefix represents typically identifies
//       the function or method which is currently executing. This
//       information is used to document source code execution flow
//       in error messages.
//
//       This method is designed to process a single new error prefix
//       string. To process a collection of error prefix strings, see
//       method 'ErrPrefixDto.SetEPrefCollection()'.
//
//
//  newErrContext       string
//     - This is the error context information associated with the
//       new error prefix ('newErrPrefix'). This parameter is
//       optional and will accept an empty string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
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

// SetIsTerminatedWithNewLine - By default, error prefix strings
// returned by the method ErrPrefixDto.String() do NOT have a new
// line character ('\n') as the last character in the error prefix
// string. In other words, returned error prefix strings are NOT
// terminated with new line characters ('\n') by default.
//
// If the user prefers to terminate error prefix strings returned
// by this instance of ErrPrefixDto with a new line character, the
// input parameter, 'isTerminatedWithNewLine', should be set to
// 'true'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  isTerminatedWithNewLine    bool
//     - If this parameter is set to 'true', all error prefix
//       strings returned by this ErrPrefixDto instance WILL BE
//       terminated with a new line character.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefDto *ErrPrefixDto) SetIsTerminatedWithNewLine(
	isTerminatedWithNewLine bool) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.isTerminatedWithNewLine =
		isTerminatedWithNewLine
}

// SetMaxErrPrefTextLineLength - Sets the maximum limit on the
// number of characters allowed in an error prefix text line output
// for display purposes.
//
// Setting this value will control the maximum character limit for
// this specific ErrPrefixDto instance. This maximum limit will
// remain in force for the life of this ErrPrefixDto instance or
// until another call is made to this method changing the value.
//
// When this instance of ErrPrefixDto was initially created, a
// default value of 40-characters was applied.
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
//       If 'maxErrPrefixTextLineLength' is set to a value less
//       than ten (10), this method will take no action and return.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
//
func (ePrefDto *ErrPrefixDto) SetMaxErrPrefTextLineLength(
	maxErrPrefixTextLineLength uint) {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	if maxErrPrefixTextLineLength < 10 {
		return
	}

	ePrefDto.maxErrPrefixTextLineLength =
		maxErrPrefixTextLineLength
}

// SetMaxErrPrefTextLineLengthToDefault - Maximum Error Prefix Line
// Length is the maximum limit on the number of characters allowed
// in a single error prefix text line.
//
// This method resets that maximum limit to its default value of
// 40-characters.
//
func (ePrefDto *ErrPrefixDto) SetMaxErrPrefTextLineLengthToDefault() {

	if ePrefDto.lock == nil {
		ePrefDto.lock = new(sync.Mutex)
	}

	ePrefDto.lock.Lock()

	defer ePrefDto.lock.Unlock()

	ePrefDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()
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

	if ePrefDto.maxErrPrefixTextLineLength == 0 {

		ePrefDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().
				getMasterErrPrefDisplayLineLength()
	}

	return ""
}
