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
	ePrefDelimiters     ErrPrefixDelimiters
	errorPrefixInfo     *ErrorPrefixInfo
	currentLineStr      string
	lenCurrentLineStr   uint
	remainingLineLength uint
	maxErrStringLength  uint
	lock                *sync.Mutex
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
func (ePrefLineLenCalc *EPrefixLineLenCalc) IsValidInstance(
	ePrefix string) bool {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.IsValidInstance() "

	ePrefLineLenCalcQuark := ePrefixLineLenCalcQuark{}

	isValid,
		_ := ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		ePrefLineLenCalc,
		ePrefix+
			"ePrefLineLenCalc\n")

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
func (ePrefLineLenCalc *EPrefixLineLenCalc) IsValidInstanceError(
	ePrefix string) error {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.IsValidInstanceError() "

	ePrefLineLenCalcQuark := ePrefixLineLenCalcQuark{}

	_,
		err := ePrefLineLenCalcQuark.testValidityOfEPrefixLineLenCalc(
		ePrefLineLenCalc,
		ePrefix+
			"ePrefLineLenCalc\n")

	return err
}

// GetCurrLineStr - Return the current line string. This string
// includes the characters which have been formatted and included
// in a single text line but which have not yet been written out
// text display. As soon as the current line string fills up with
// characters to the maximum allowed line length, the text line
// will be written out to the display device.
//
func (ePrefLineLenCalc *EPrefixLineLenCalc) GetCurrLineStr() string {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

	return ePrefLineLenCalc.currentLineStr
}

// GetMaxErrStringLength - Returns the current the value for
// maximum error string length. This limit controls the line length
// for text displays of error prefix strings.
//
// The value of maximum error string length is returned as an
// unsigned integer.
//
func (ePrefLineLenCalc *EPrefixLineLenCalc) GetMaxErrStringLength() uint {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

	return ePrefLineLenCalc.maxErrStringLength
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
func (ePrefLineLenCalc *EPrefixLineLenCalc) SetEPrefDelimiters(
	ePrefDelimiters *ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

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

	return ePrefLineLenCalc.ePrefDelimiters.CopyIn(
		ePrefDelimiters,
		ePrefix)

}

// SetEPrefDto - Sets the Error Prefix Data Transfer Object member
// variable.
//
// The Error Prefix Data Transfer Object stores information on the
// error prefix and error context strings.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefixInfo      *ErrorPrefixInfo
//     - This Error Prefix Data Transfer Object stores information
//       on the error prefix and error context strings.
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
func (ePrefLineLenCalc *EPrefixLineLenCalc) SetEPrefDto(
	errorPrefixDto *ErrorPrefixInfo,
	ePrefix string) error {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

	ePrefix += "EPrefixLineLenCalc.SetEPrefDto() "

	if errorPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errorPrefixInfo' is a "+
			"'nil' pointer\n",
			ePrefix)
	}

	err := errorPrefixDto.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'errorPrefixInfo' validity check\n"+
			"returned an error message!\n"+
			"%v\n",
			ePrefix,
			err.Error())
	}

	ePrefLineLenCalc.errorPrefixInfo = errorPrefixDto

	return nil
}

// SetMaxErrStringLength - Sets EPrefixLineLenCalc.maxErrStringLength
//
// This method sets the value for maximum error string length. This
// limit controls the line length for text displays of error prefix
// strings.
//
func (ePrefLineLenCalc *EPrefixLineLenCalc) SetMaxErrStringLength(
	maxErrStringLength uint) {

	if ePrefLineLenCalc.lock == nil {
		ePrefLineLenCalc.lock = new(sync.Mutex)
	}

	ePrefLineLenCalc.lock.Lock()

	defer ePrefLineLenCalc.lock.Unlock()

	ePrefLineLenCalc.maxErrStringLength = maxErrStringLength

}
