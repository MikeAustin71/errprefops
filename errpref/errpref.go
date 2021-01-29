package errpref

import (
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
// Note that there are no 'pointer' methods provided for this
// type. This is because the type is not designed to store
// information. Its only function is to receive process and
// return strings of error prefix information.
//
//
type ErrPref struct {
	maxErrPrefixTextLineLength uint
	lock                       *sync.Mutex
}

// AddContext - Adds error context string to a pre-existing error
// prefix.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
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
//  errPref             string
//     - A string containing error prefixes in series. This method
//       will add an error context string to the last error prefix
//       string in this series.
//
//
//  errContext          string
//     - This is the error context information that will be added
//       and associated with the last error prefix in the error
//       prefix string, 'errPref'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  string
//     - Returns a string consisting of a revised series of error
//       prefixes. This string is identical to input parameter,
//       'errPref' except that the error context information
//       provided by input parameter 'newErrContext' has been
//       attached to the last error prefix in 'errPref'.
//
//
func (ePref ErrPref) AddContext(
	errPref string,
	newErrContext string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	if len(newErrContext) == 0 {
		return errPref
	}

	ePrefAtom := errPrefAtom{}

	oldErrPref,
		newErrPref,
		lastErrContext :=
		ePrefAtom.extractLastErrPrefInSeries(errPref)

	if len(lastErrContext) > 0 {
		newErrContext += " " + lastErrContext
	}

	ePrefMech := errPrefMechanics{}

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		newErrContext,
		ePref.maxErrPrefixTextLineLength)
}

// ConvertNonPrintableChars - Receives a string containing
// non-printable characters and converts them to 'printable'
// characters returned in a string.
//
// Examples of non-printable characters are '\n', '\t' or 0x06
// (Acknowledge). These example characters would be translated into
// printable string characters as: "\\n", "\\t" and "[ACK]".
//
// Space characters are typically translated as " ". However, if
// the input parameter 'convertSpace' is set to 'true' then all
// spaces are converted to "[SPACE]" in the returned string.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
// This method is useful for verifying error prefix strings which
// are routinely populated with non-printable characters.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters will be converted to
//       printable characters.
//
//  convertSpace        bool
//     - Space or white space characters (0x20) are by default
//       translated as " ". However, if this parameter is set to
//       'true', space characters will be converted to "[SPACE]".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  printableChars      string
//     - This returned string is identical to input parameter
//       'nonPrintableChars' with the exception that non-printable
//       characters are translated into printable characters.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "Hello world!\n"
//  testRunes := []rune(testStr)
//  ePrefix := "theCallingFunction()"
//
//  ePrefQuark := errPrefQuark{}
//
//  actualStr :=
//    ePrefQuark.
//      convertNonPrintableChars(
//           testRunes,
//           true,
//           ePrefix)
//
//  ----------------------------------------------------
//  'actualStr' is now equal to:
//     "Hello[SPACE]world!\\n"
//
//
func (ePref ErrPref) ConvertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool) (
	printableChars string) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	printableChars,
		_ = ePrefQuark.convertNonPrintableChars(
		nonPrintableChars,
		convertSpace,
		"")

	return printableChars
}

// FmtString - Returns a formatted text representation of all error
// prefix and error context information contained in the input
// parameter string, 'errPref'.
//
// Error prefix text is designed to be configured at the beginning
// of error messages and is most often used to document the thread
// of code execution by listing the calling sequence for a specific
// list of functions and methods.
//
// The error context descriptive text provides additional
// information about the function or method identified in the
// associated error prefix string. Typical context information
// might include variable names, variable values and further
// details on function execution.
//
func (ePref ErrPref) FmtString(errPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	ePrefNanobot := errPrefNanobot{}

	return ePrefNanobot.formatErrPrefix(
		ePref.maxErrPrefixTextLineLength,
		errPref)
}

// GetMaxErrPrefTextLineLength - Returns the current maximum number
// of characters allowed in an error prefix text line output
// display.
//
// To change or reset this maximum limit value, see method:
//     ErrPref.SetMaxErrPrefTextLineLength().
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  maxErrPrefixStringLength      uint
//     - This method will return an unsigned integer value
//       specifying the maximum number of characters allowed
//       in an error prefix text display line.
//
func (ePref ErrPref) GetMaxErrPrefTextLineLength() (
	maxErrPrefixStringLength uint) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	maxErrPrefixStringLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	return maxErrPrefixStringLength
}

// NewErrPref - Returns a string concatenating the old error prefix the
// new custom, user-defined error prefix. The new error prefix is
// typically used to document method or function chains in error
// messages.
//
// The old error prefix contains the function chain or series which
// led to the function next in line for execution.
//
// The error prefix text is designed to be configured at the
// beginning of error messages and is most often used to document
// the thread of code execution by listing the calling sequence for
// specific functions and methods.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref          string
//     - This includes the previous error prefix string. This string
//       will be formatted and concatenated with the new error prefix
//       information provided by input parameter, 'newErrPref',
//       below.
//
//
//  newErrPref          string
//     - The new error prefix represents the error prefix string
//       identifies with the function or method which is currently
//       executing for purposes of documenting execution flow in
//       error messages. This parameter is optional and will accept
//       an empty string, but there isn't much point in calling
//       this method without a substantive value for 'newErrPref'.
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
func (ePref ErrPref) NewErrPref(
	oldErrPref string,
	newErrPref string) string {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	ePrefMech := errPrefMechanics{}

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		"",
		ePref.maxErrPrefixTextLineLength)
}

// NewContext - Receives an old error prefix, new error prefix and
// a new context string which are concatenated and returned as a
// combined string.
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

	ePrefQuark := errPrefQuark{}

	ePref.maxErrPrefixTextLineLength =
		ePrefQuark.getErrPrefDisplayLineLength()

	ePrefMech := errPrefMechanics{}

	return ePrefMech.assembleErrPrefix(
		oldErrPref,
		newErrPref,
		newContext,
		ePref.maxErrPrefixTextLineLength)
}

// SetMaxErrPrefTextLineLength - Sets the maximum limit on the
// number of characters allowed in an error prefix text line output
// display.
//
// -IMPORTANT -
// Setting this value will control the maximum character limit not
// only for this ErrPref instance, but will also control all that
// limit for all other instances of ErrPref created in this session.
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
func (ePref ErrPref) SetMaxErrPrefTextLineLength(
	maxErrPrefixTextLineLength uint) {

	if ePref.lock == nil {
		ePref.lock = new(sync.Mutex)
	}

	ePref.lock.Lock()

	defer ePref.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	ePrefQuark.setErrPrefDisplayLineLength(
		maxErrPrefixTextLineLength)

	return
}
