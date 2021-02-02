package errpref

import (
	"fmt"
	"sync"
)

// EPrefixLineLenCalc - Error Prefix Line Length Calculator
// This type is used to perform calculations on the line length of
// error prefix text output strings. Among the calculations
// performed, these associated methods determine how many error
// prefix and error context strings can be accommodated on the same
// line of text give a maximum line length limit.
//
type EPrefixLineLenCalc struct {
	ePrefDelimiters    ErrPrefixDelimiters
	errorPrefixInfo    *ErrorPrefixInfo
	currentLineStr     string
	maxErrStringLength uint
	lock               *sync.Mutex
}

// CopyIn - Receives an instance of type EPrefixLineLenCalc and
// proceeds to copy the internal member data variable values to the
// current EPrefixLineLenCalc instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  incomingLineLenCalc        *EPrefixLineLenCalc
//     - A pointer to an instance of EPrefixLineLenCalc. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this EPrefixLineLenCalc instance will
//       be copied to current EPrefixLineLenCalc instance
//       ('ePrefixLineLenCalc').
//
//       If this EPrefixLineLenCalc instance proves to be invalid,
//       an error will be returned.
//
//
//  ePrefix                    string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) CopyIn(
	incomingLineLenCalc *EPrefixLineLenCalc,
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.CopyIn() "

	ePrefLineLenCalcElectron := ePrefixLineLenCalcElectron{}

	return ePrefLineLenCalcElectron.copyIn(
		ePrefixLineLenCalc,
		incomingLineLenCalc,
		ePrefix)
}

// IsValidInstance - Returns a boolean flag signalling whether the
// current EPrefixLineLenCalc instance is valid, or not.
//
// If this method returns a boolean value of 'false', it signals
// that the current EPrefixLineLenCalc instance is invalid.
//
// If this method returns a boolean value of 'true', it signals
// that the current EPrefixLineLenCalc instance is valid in all
// respects.
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
//  bool
//     - This boolean flag signals whether the current
//       EPrefixLineLenCalc instance is valid.
//
//       If this method returns a value of 'false', it signals that
//       the current EPrefixLineLenCalc instance is invalid.
//
//       If this method returns a value of 'true', it signals that
//       the current EPrefixLineLenCalc instance is valid in all
//       respects.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) IsValidInstance(
	ePrefix string) bool {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.IsValidInstance() "

	ePrefLineLenCalcQuark := ePrefixLineLenCalcQuark{}

	isValid,
		_ := ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		ePrefixLineLenCalc,
		ePrefix+
			"ePrefixLineLenCalc\n")

	return isValid
}

// IsValidInstanceError - Returns an error type signalling whether
// the current EPrefixLineLenCalc instance is valid, or not.
//
// If this method returns an error value NOT equal to 'nil', it
// signals that the current EPrefixLineLenCalc instance is
// invalid.
//
// If this method returns an error value which IS equal to 'nil',
// it signals that the current EPrefixLineLenCalc instance is
// valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this returned error type is set equal to 'nil', it
//       signals that the current EPrefixLineLenCalc is valid in
//       all respects.
//
//       If this returned error type is NOT equal to 'nil', it
//       signals that the current EPrefixLineLenCalc is invalid.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) IsValidInstanceError(
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.IsValidInstanceError() "

	ePrefLineLenCalcQuark := ePrefixLineLenCalcQuark{}

	_,
		err := ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		ePrefixLineLenCalc,
		ePrefix+
			"ePrefixLineLenCalc\n")

	return err
}

// GetRemainingLineLength - Returns the remaining line length. This is
// the difference between current line length and Maximum Error
// String Length.
//
//   remainingLineLen = maxErrStringLen - currentLineStringLen
//
// If currentLineStringLen is greater than Maximum Error String
// Length, the Remaining String Length is zero.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetRemainingLineLength() uint {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	var (
		lenCurrentLineStr,
		remainingLineLength uint
	)

	lenCurrentLineStr =
		uint(len(ePrefixLineLenCalc.currentLineStr))

	if lenCurrentLineStr >
		ePrefixLineLenCalc.maxErrStringLength {

		remainingLineLength = 0

	} else {

		remainingLineLength =
			ePrefixLineLenCalc.maxErrStringLength -
				lenCurrentLineStr

	}

	return remainingLineLength
}

// GetCurrLineStr - Return the current line string. This string
// includes the characters which have been formatted and included
// in a single text line but which have not yet been written out
// text display. As soon as the current line string fills up with
// characters to the maximum allowed line length, the text line
// will be written out to the display device.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetCurrLineStr() string {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.currentLineStr
}

func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetCurrLineStringLength() uint {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	var lenCurrentLineStr uint

	lenCurrentLineStr =
		uint(len(ePrefixLineLenCalc.currentLineStr))

	return lenCurrentLineStr
}

// GetMaxErrStringLength - Returns the current the value for
// maximum error string length. This limit controls the line length
// for text displays of error prefix strings.
//
// The value of maximum error string length is returned as an
// unsigned integer.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) GetMaxErrStringLength() uint {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	return ePrefixLineLenCalc.maxErrStringLength
}

// SetCurrentLineStr - Sets the Current Line String. This string
// represents the characters which have been collected thus far
// from error prefix and error context information. The current
// line string is used to control maximum line length and stores
// the characters which have not yet been written out to the
// text display.
//
// This method calculates current line length and remaining line
// length before storing the data in internal member variables.
//
// Be sure to set Maximum Error String Length first, before
// calling this method.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetCurrentLineStr(
	currentLineStr string) {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefixLineLenCalc.currentLineStr = currentLineStr

	return
}

// SetEPrefDelimiters - Sets the Error Prefix Delimiters member
// variable.
//
// The Error Prefix Delimiters object stores string delimiters used
// to terminate error prefix and error context strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefDelimiters     *ErrPrefixDelimiters
//     - A pointer to an Error Prefix Delimiters object. This
//       delimiters object contains information on the delimiter
//       strings used to terminate error prefix and error context
//       strings.
//
//       type ErrPrefixDelimiters struct {
//         inLinePrefixDelimiter      string
//         lenInLinePrefixDelimiter   uint
//         newLinePrefixDelimiter     string
//         lenNewLinePrefixDelimiter  uint
//         inLineContextDelimiter     string
//         lenInLineContextDelimiter  uint
//         newLineContextDelimiter    string
//         lenNewLineContextDelimiter uint
//         maxErrStringLength         uint
//       }
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will
//       encapsulate an error message. Note that this error message
//       will incorporate the method chain and text passed by input
//       parameter, 'ePrefix'. The 'ePrefix' text will be prefixed
//       to the beginning of the error message.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetEPrefDelimiters(
	ePrefDelimiters *ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.SetEPrefDelimiters() "

	if ePrefDelimiters == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefDelimiters' is a "+
			"'nil' pointer\n",
			ePrefix)
	}

	err := ePrefDelimiters.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefDelimiters' validity check\n"+
			"returned an error message!\n"+
			"%v\n",
			ePrefix,
			err.Error())
	}

	return ePrefixLineLenCalc.ePrefDelimiters.CopyIn(
		ePrefDelimiters,
		ePrefix+
			"ePrefDelimiters\n")

}

// SetErrPrefixInfo - Sets the Error Prefix Information Object member
// variable.
//
// The Error Prefix Information Object stores information on the
// error prefix and error context strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefixInfo       *ErrorPrefixInfo
//     - This Error Prefix Information Object stores information on
//       the error prefix and error context strings.
//
//       type ErrorPrefixInfo struct {
//         isValid                bool
//         isLastIdx              bool // Signals the last index in the array
//         errorPrefixStr         string
//         lenErrorPrefixStr      uint
//         errPrefixHasContextStr bool
//         errorContextStr        string
//         lenErrorContextStr     uint
//       }
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will
//       encapsulate an error message. Note that this error message
//       will incorporate the method chain and text passed by input
//       parameter, 'ePrefix'. The 'ePrefix' text will be prefixed
//       to the beginning of the error message.
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetErrPrefixInfo(
	errPrefixInfo *ErrorPrefixInfo,
	ePrefix string) error {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.SetErrPrefixInfo() "

	if errPrefixInfo == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefixInfo' is a "+
			"'nil' pointer\n",
			ePrefix)
	}

	err := errPrefixInfo.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errPrefixInfo' validity check\n"+
			"returned an error message!\n"+
			"%v\n",
			ePrefix,
			err.Error())
	}

	ePrefixLineLenCalc.errorPrefixInfo = errPrefixInfo

	return nil
}

// SetMaxErrStringLength - Sets EPrefixLineLenCalc.maxErrStringLength
//
// This method sets the value for maximum error string length. This
// limit controls the line length for text displays of error prefix
// strings.
//
// Set this value first, before setting Current Line Length
//
func (ePrefixLineLenCalc *EPrefixLineLenCalc) SetMaxErrStringLength(
	maxErrStringLength uint) {

	if ePrefixLineLenCalc.lock == nil {
		ePrefixLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefixLineLenCalc.lock.Lock()

	defer ePrefixLineLenCalc.lock.Unlock()

	ePrefixLineLenCalc.maxErrStringLength = maxErrStringLength

}
