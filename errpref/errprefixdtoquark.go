package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDtoQuark struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of errPrefixDtoQuark.
//
func (ePrefDtoQuark errPrefixDtoQuark) ptr() *errPrefixDtoQuark {

	if ePrefDtoQuark.lock == nil {
		ePrefDtoQuark.lock = new(sync.Mutex)
	}

	ePrefDtoQuark.lock.Lock()

	defer ePrefDtoQuark.lock.Unlock()

	return &errPrefixDtoQuark{
		lock: new(sync.Mutex),
	}
}

// emptyEPrefCollection - Receives a pointer to an ErrPrefixDto
// object an proceeds to delete all of the existing error prefix
// and error context information. When completed, the ErrPrefixDto
// object's collection of ErrorPrefixInfo elements will be set to
// 'nil'.
//
func (ePrefDtoQuark *errPrefixDtoQuark) emptyErrPrefInfoCollection(
	ePrefixDto *ErrPrefixDto,
	errPrefStr string) error {

	if ePrefDtoQuark.lock == nil {
		ePrefDtoQuark.lock = new(sync.Mutex)
	}

	ePrefDtoQuark.lock.Lock()

	defer ePrefDtoQuark.lock.Unlock()

	errPrefStr = errPrefStr +
		"\nerrPrefixDtoQuark.emptyErrPrefInfoCollection()"

	if ePrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefixDto' is invalid!\n"+
			"'ePrefixDto' is a 'nil' pointer.\n",
			errPrefStr)
	}

	if ePrefixDto.ePrefCol == nil {
		return nil
	}

	lenEPrefCol := len(ePrefixDto.ePrefCol)

	if lenEPrefCol == 0 {
		return nil
	}

	for i := 0; i < lenEPrefCol; i++ {
		ePrefixDto.ePrefCol[i].Empty()
	}

	ePrefixDto.ePrefCol = nil

	return nil
}

// newZeroErrPrefixDto - Returns a new instance of ErrPrefixDto
// with all internal member variables set to their initial or zero
// values.
//
func (ePrefDtoQuark *errPrefixDtoQuark) newZeroErrPrefixDto() ErrPrefixDto {

	if ePrefDtoQuark.lock == nil {
		ePrefDtoQuark.lock = new(sync.Mutex)
	}

	ePrefDtoQuark.lock.Lock()

	defer ePrefDtoQuark.lock.Unlock()

	newErrPrefixDto := ErrPrefixDto{}

	newErrPrefixDto.lock = new(sync.Mutex)

	newErrPrefixDto.ePrefCol = nil

	newErrPrefixDto.maxErrPrefixTextLineLength =
		errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	newErrPrefixDto.inputStrDelimiters.SetToDefaultIfEmpty()

	newErrPrefixDto.outputStrDelimiters.SetToDefaultIfEmpty()

	return newErrPrefixDto
}

// testValidityOfErrorPrefixInfo - Performs a diagnostic review of
// the input parameter 'ePrefixDto', an instance of ErrPrefixDto.
// The purpose of this diagnostic review is to determine whether
// this ErrPrefixDto instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefixDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This object
//       will be evaluated to determine whether or not it is a
//       valid instance.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the
//       calling method or methods. Be sure to leave a space at the
//       end of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - This returned boolean value will signal whether the input
//       parameter, 'ePrefixDto', is valid, or not. If the
//       'errPrefixDto' object contains valid data, this method
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
func (ePrefDtoQuark *errPrefixDtoQuark) testValidityOfErrPrefixDto(
	ePrefixDto *ErrPrefixDto,
	errPrefStr string) (
	isValid bool,
	err error) {

	if ePrefDtoQuark.lock == nil {
		ePrefDtoQuark.lock = new(sync.Mutex)
	}

	ePrefDtoQuark.lock.Lock()

	defer ePrefDtoQuark.lock.Unlock()

	errPrefStr = errPrefStr +
		"\nerrPrefixDtoQuark.testValidityOfErrPrefixDto()"

	isValid = false

	if ePrefixDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'ePrefixDto' is invalid!\n"+
			"'ePrefixDto' is a 'nil' pointer.\n",
			errPrefStr)

		return isValid, err
	}

	// Minimum length is 10-characters
	if ePrefixDto.maxErrPrefixTextLineLength < 10 {

		err = fmt.Errorf("%v\n"+
			"Error: Maximum Text Line Length is less than 10-Characters!\n"+
			"Maximum Text Line Length='%v'\n",
			errPrefStr,
			ePrefixDto.maxErrPrefixTextLineLength)

		return isValid, err
	}

	err = ePrefixDto.inputStrDelimiters.
		IsValidInstanceError(errPrefStr + " inputStrDelimiters ")

	if err != nil {
		return isValid, err
	}

	err = ePrefixDto.outputStrDelimiters.
		IsValidInstanceError(errPrefStr + " outputStrDelimiters ")

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}
