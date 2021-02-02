package errpref

import (
	"fmt"
	"sync"
)

type errorPrefixDtoQuark struct {
	lock *sync.Mutex
}

// testValidityOfErrorPrefixDto - Performs a diagnostic review of
// the input parameter 'errPrefixDto', an instance of
// ErrorPrefixDto. The purpose of this diagnostic review is to
// determine whether this ErrorPrefixDto instance is valid in all
// respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefixDto        *ErrorPrefixDto
//     - A pointer to an instance of ErrorPrefixDto. This object
//       will be evaluated to determine whether or not it is a
//       valid instance.
//
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
//  isValid             bool
//     - This returned boolean value will signal whether the input
//       parameter, 'errPrefixDto', is valid, or not. If the
//       'errPrefixDto' object contains valid data,  this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
//  err                 error
//     - If the input parameter object, 'errPrefixDto', contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the input parameter object, 'errPrefixDto', is valid,
//       this error parameter will be set to 'nil'.
//
func (ePrefDtoQuark *errorPrefixDtoQuark) testValidityOfErrorPrefixDto(
	errPrefixDto *ErrorPrefixDto,
	ePrefix string) (
	isValid bool,
	err error) {

	if ePrefDtoQuark.lock == nil {
		ePrefDtoQuark.lock = new(sync.Mutex)
	}

	ePrefDtoQuark.lock.Lock()

	defer ePrefDtoQuark.lock.Unlock()

	ePrefix += "errorPrefixDtoQuark.testValidityOfErrorPrefixDto() "
	isValid = false

	if errPrefixDto == nil {
		err = fmt.Errorf("%v\n"+
			"\nInput parameter 'errPrefixDto' is INVALID!\n"+
			"'errPrefixDto' is a nil pointer!\n",
			ePrefix)

		return isValid, err
	}

	errPrefixDto.lenErrorPrefixStr =
		uint(len(errPrefixDto.errorPrefixStr))

	if errPrefixDto.lenErrorPrefixStr == 0 {

		err =
			fmt.Errorf("%v\n"+
				"Error: Error Prefix is an empty string!\n",
				ePrefix)
		return isValid, err
	}

	errPrefixDto.lenErrorContextStr =
		uint(len(errPrefixDto.errorContextStr))

	if errPrefixDto.lenErrorPrefixStr == 0 {
		errPrefixDto.errPrefixHasContextStr = false
	} else {
		errPrefixDto.errPrefixHasContextStr = true
	}

	isValid = true

	return isValid, err
}
