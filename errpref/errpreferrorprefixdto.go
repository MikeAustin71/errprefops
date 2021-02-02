package errpref

import (
	"sync"
)

type ErrorPrefixDto struct {
	isLastIdx              bool // Signals the last index in the array
	errorPrefixStr         string
	lenErrorPrefixStr      uint
	errPrefixHasContextStr bool
	errorContextStr        string
	lenErrorContextStr     uint
	lock                   *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// ErrorPrefixDto ('inComingErrPrefixDto') to the data fields of
// the current ErrorPrefixDto instance ('errorPrefixDto').
//
// If 'inComingErrPrefixDto' is judged to be invalid, this method
// will return an error.
//
// All of the data fields in current ErrorPrefixDto instance
// ('errorPrefixDto') will be modified and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//
//  inComingErrPrefixDto       *ErrorPrefixDto
//     - A pointer to an instance of ErrorPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrorPrefixDto instance will be
//       copied to current ErrorPrefixDto instance
//       ('errorPrefixDto').
//
//       If this ErrorPrefixDto instance proves to be invalid, an
//       error will be returned.
//
//
//  ePrefix             string
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
//  err                 error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
func (errorPrefixDto *ErrorPrefixDto) CopyIn(
	inComingErrPrefixDto *ErrorPrefixDto,
	ePrefix string) error {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()
	ePrefix += "ErrorPrefixDto.CopyIn()\n"

	ePrefDtoElectron := errorPrefixDtoElectron{}

	return ePrefDtoElectron.copyIn(
		errorPrefixDto,
		inComingErrPrefixDto,
		ePrefix)
}

// CopyOut - Creates a deep copy of the data fields contained in
// the current ErrorPrefixDto instance, and returns that data as a
// new instance of ErrorPrefixDto.
//
//
// ------------------------------------------------------------------------
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  ErrorPrefixDto
//     - If this method completes successfully, a deep copy of the
//       current ErrorPrefixDto instance will be returned through
//       this parameter as a completely new instance of
//       ErrorPrefixDto.
//
//
//  error
//     - If this method completes successfully, the returned error Type
//       is set to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'. The
//       'ePrefix' text will be prefixed to the beginning of the returned
//       error message.
//
func (errorPrefixDto *ErrorPrefixDto) CopyOut(
	ePrefix string) (
	ErrorPrefixDto,
	error) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	ePrefix += "ErrorPrefixDto.CopyOut() "

	ePrefDtoElectron := errorPrefixDtoElectron{}

	return ePrefDtoElectron.copyOut(
		errorPrefixDto,
		ePrefix)
}

// GetErrContextStr - Returns the Error Context String.
func (errorPrefixDto *ErrorPrefixDto) GetErrContextStr() string {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return errorPrefixDto.errorContextStr
}

// GetErrPrefixStr - Returns the Error Prefix String.
//
func (errorPrefixDto *ErrorPrefixDto) GetErrPrefixStr() string {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return errorPrefixDto.errorPrefixStr
}

// GetErrPrefixHasContextStr - Returns a boolean flag signaling
// whether the Error Prefix String defined in this ErrorPrefixDto
// instance has an associated Error Context string.
//
// If the returned boolean value is 'true' is signals that this
// Error Prefix has an associated Error Context String.
//
func (errorPrefixDto *ErrorPrefixDto) GetErrPrefixHasContextStr() bool {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return errorPrefixDto.errPrefixHasContextStr
}

// GetIsLastIndex - Returns a boolean flag signaling whether the
// current ErrorPrefixDto instance is the last element in an array.
//
// If the returned boolean value is 'true', it signals that this
// ErrorPrefixDto object is the last element in an array.
//
func (errorPrefixDto *ErrorPrefixDto) GetIsLastIndex() bool {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return errorPrefixDto.isLastIdx
}

// IsValidInstance - Returns a boolean flag signaling whether the
// current ErrorPrefixDto instance is invalid.
//
// If the returned boolean value is 'true', it signals that the
// current ErrorPrefixDto instance is valid and populated with data.
//
// If the boolean value is 'false', it signals that the current
// ErrorPrefixDto instance is invalid.
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
//  isValid             bool
//     - If the current ErrorPrefixDto instance, 'errorPrefixDto',
//       contains invalid data, this boolean value is set to 'false'.
//
//       If the current ErrorPrefixDto instance, 'errorPrefixDto',
//       is valid, this boolean value is set to 'true'.
//
func (errorPrefixDto *ErrorPrefixDto) IsValidInstance() (
	isValid bool) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	ePrefDtoQuark := errorPrefixDtoQuark{}

	isValid,
		_ = ePrefDtoQuark.testValidityOfErrorPrefixDto(
		errorPrefixDto,
		"")

	return isValid
}

// IsValidInstanceError - Returns an error object signaling whether
// the current ErrorPrefixDto instance is invalid.
//
// If the returned error value is 'nil', it signals that the
// current ErrorPrefixDto instance is valid and populated with
// data.
//
// If the returned error value is NOT 'nil', it signals that the
// current ErrorPrefixDto is invalid. In this case, the error
// object will contain an appropriate error message.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the current ErrorPrefixDto instance, 'errorPrefixDto',
//       contains invalid data, a detailed error message will be
//       returned identifying the invalid data item.
//
//       If the current ErrorPrefixDto instance, 'errorPrefixDto',
//       is valid, this error parameter will be set to 'nil'.
//
func (errorPrefixDto *ErrorPrefixDto) IsValidInstanceError(
	ePrefix string) error {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	ePrefix += "ErrorPrefixDto.IsValidInstanceError() "

	ePrefDtoQuark := errorPrefixDtoQuark{}

	_,
		err := ePrefDtoQuark.testValidityOfErrorPrefixDto(
		errorPrefixDto,
		ePrefix)

	return err
}

// GetLengthErrContextStr - Returns the number of characters
// in the Error Context String. This is also known as the string
// length and it is returned as a type unsigned integer.
//
func (errorPrefixDto *ErrorPrefixDto) GetLengthErrContextStr() uint {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return errorPrefixDto.lenErrorContextStr
}

// GetLengthErrPrefixStr - Returns the number of characters in
// the Error Prefix String. This also known as the string length
// and it is returned as a type unsigned integer.
//
func (errorPrefixDto *ErrorPrefixDto) GetLengthErrPrefixStr() uint {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return errorPrefixDto.lenErrorPrefixStr
}

// New - Creates and returns a new instance of type,
// 'ErrorPrefixDto'.
//
func (errorPrefixDto ErrorPrefixDto) New() ErrorPrefixDto {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	return ErrorPrefixDto{}
}

// SetErrContextStr - Sets the Error Context String value.
//
// In addition, this method also calculates and set the value of
// length or number of characters contained in the Error Context
// String.
//
func (errorPrefixDto *ErrorPrefixDto) SetErrContextStr(
	errContext string) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	errorPrefixDto.errorContextStr = errContext

	errorPrefixDto.lenErrorContextStr =
		uint(len(errorPrefixDto.errorContextStr))

	if errorPrefixDto.lenErrorContextStr > 0 {
		errorPrefixDto.errPrefixHasContextStr = true
	} else {
		errorPrefixDto.errPrefixHasContextStr = false
	}

}

// SetErrPrefixHasContext - Sets an internal boolean flag
// specifying whether the Error Prefix String for this
// ErrorPrefixDto instance has an associated Error Context String.
//
// If this boolean flag is set to 'true' is signals that this
// Error Prefix String has an associated Error Context String.
//
func (errorPrefixDto *ErrorPrefixDto) SetErrPrefixHasContext(
	errPrefixHasContextStr bool) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	errorPrefixDto.errPrefixHasContextStr = errPrefixHasContextStr
}

// SetErrPrefixStr - Sets the value of the error prefix string.
//
// In addition, this method also calculates and set the value of
// length or number of characters contained in the error prefix
// string.
//
func (errorPrefixDto *ErrorPrefixDto) SetErrPrefixStr(
	errPrefix string) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	errorPrefixDto.errorPrefixStr = errPrefix

	errorPrefixDto.lenErrorPrefixStr =
		uint(len(errorPrefixDto.errorPrefixStr))
}

// SetIsLastIndex - Sets an internal flag designating whether the
// current ErrorPrefixDto instance is the last element in an array
// of ErrorPrefixDto objects.
//
func (errorPrefixDto *ErrorPrefixDto) SetIsLastIndex(isLastIdx bool) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	errorPrefixDto.isLastIdx = isLastIdx
}
