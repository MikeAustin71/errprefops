package errpref

import (
	"fmt"
	"sync"
)

type ErrorPrefixDto struct {
	isValid                bool
	isLastIdx              bool // Signals the last index in the array
	errorPrefixStr         string
	lenErrorPrefixStr      uint
	errPrefixHasContextStr bool
	errorContextStr        string
	lenErrorContextStr     uint
	lock                   *sync.Mutex
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

// IsValidInstanceError - Returns the 'isValid' flag and an error object
// signaling whether the current ErrorPrefixDto instance is
// invalid.
//
// If the returned boolean value is 'true', it signals that the
// current ErrorPrefixDto is valid and populated with data.
//
// If the returned boolean value is 'false', it signals that the
// current ErrorPrefixDto is empty and NOT populated with valid
// data.
//
//
//
func (errorPrefixDto *ErrorPrefixDto) IsValidInstanceError(
	ePrefix string) (
	bool,
	error) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	ePrefix += "ErrorPrefixDto.IsValidInstanceError() "

	errorPrefixDto.lenErrorPrefixStr =
		uint(len(errorPrefixDto.errorPrefixStr))

	if errorPrefixDto.lenErrorPrefixStr == 0 {

		errorPrefixDto.isValid = false

		return errorPrefixDto.isValid,
			fmt.Errorf("%v\n"+
				"Error: Error Prefix is an empty string!\n",
				ePrefix)

	}

	errorPrefixDto.lenErrorContextStr =
		uint(len(errorPrefixDto.errorContextStr))

	if errorPrefixDto.lenErrorPrefixStr == 0 {
		errorPrefixDto.errPrefixHasContextStr = false
	} else {
		errorPrefixDto.errPrefixHasContextStr = true
	}

	errorPrefixDto.isValid = true

	return errorPrefixDto.isValid, nil
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

// SetIsEmpty - Sets the 'isValid' flag signaling whether the
// current ErrorPrefixDto instance is empty, or populated with
// data.
//
// If the 'isValid' flag is set to 'true', it signals that the
// current ErrorPrefixDto is valid and populated with data.
//
// If the 'isValid' flag is set to  'false', it signals that the
// current ErrorPrefixDto is empty and NOT populated with valid
// data.
//
//
func (errorPrefixDto *ErrorPrefixDto) SetIsEmpty(isValid bool) {

	if errorPrefixDto.lock == nil {
		errorPrefixDto.lock = new(sync.Mutex)
	}

	errorPrefixDto.lock.Lock()

	defer errorPrefixDto.lock.Unlock()

	errorPrefixDto.isValid = isValid
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
