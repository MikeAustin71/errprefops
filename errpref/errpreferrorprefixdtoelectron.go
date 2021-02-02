package errpref

import (
	"fmt"
	"sync"
)

type errorPrefixDtoElectron struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from an incoming instance of
// ErrorPrefixDto ('inComingErrPrefixDto') to the data fields of
// the target ErrorPrefixDto instance ('targetErrPrefixDto')
//
// If 'inComingErrPrefixDto' is judged to be invalid, this method
// will return an error.
//
// All of the data fields in the 'targetErrPrefixDto' instance will
// be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetErrPrefixDto         *ErrorPrefixDto
//     - A pointer to an instance of ErrorPrefixDto. This method
//       WILL CHANGE AND OVERWRITE all values of internal member
//       variables encapsulated in this object to achieve the
//       method's objectives.
//
//       This ErrorPrefixDto instance will receive all the data
//       values contained from input parameter
//       'inComingErrPrefixDto'. When the copy operation is
//       completed, the data values in 'targetErrPrefixDto' will be
//       identical to those contained in input parameter,
//       'inComingErrPrefixDto'.
//
//
//  inComingErrPrefixDto       *ErrorPrefixDto
//     - A pointer to an instance of ErrorPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrorPrefixDto instance will be
//       copied to input parameter 'targetErrPrefixDto'.
//
//       If this ErrorPrefixDto instance proves to be invalid, an
//       error will be returned.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'ePrefix'.
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
func (ePrefDtoElectron *errorPrefixDtoElectron) copyIn(
	targetErrPrefixDto *ErrorPrefixDto,
	inComingErrPrefixDto *ErrorPrefixDto,
	ePrefix string) error {

	if ePrefDtoElectron.lock == nil {
		ePrefDtoElectron.lock = new(sync.Mutex)
	}

	ePrefDtoElectron.lock.Lock()

	defer ePrefDtoElectron.lock.Unlock()

	ePrefix += "errorPrefixDtoElectron.copyIn() "

	if targetErrPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetErrPrefixDto' is INVALID!\n"+
			"'targetErrPrefixDto' is a nil pointer!\n",
			ePrefix)

	}

	if inComingErrPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'inComingErrPrefixDto' is INVALID!\n"+
			"'inComingErrPrefixDto' is a nil pointer!\n",
			ePrefix)
	}

	ePrefDtoQuark := errorPrefixDtoQuark{}

	_,
		err := ePrefDtoQuark.testValidityOfErrorPrefixDto(
		inComingErrPrefixDto,
		ePrefix+
			"inComingErrPrefixDto\n")

	if err != nil {
		return err
	}

	targetErrPrefixDto.isLastIdx =
		inComingErrPrefixDto.isLastIdx

	targetErrPrefixDto.errorPrefixStr =
		inComingErrPrefixDto.errorPrefixStr

	targetErrPrefixDto.lenErrorPrefixStr =
		uint(len(targetErrPrefixDto.errorPrefixStr))

	if targetErrPrefixDto.lenErrorPrefixStr !=
		inComingErrPrefixDto.lenErrorPrefixStr {

		inComingErrPrefixDto.lenErrorPrefixStr =
			targetErrPrefixDto.lenErrorPrefixStr
	}

	targetErrPrefixDto.errorContextStr =
		inComingErrPrefixDto.errorContextStr

	targetErrPrefixDto.lenErrorContextStr =
		uint(len(targetErrPrefixDto.errorContextStr))

	if targetErrPrefixDto.lenErrorContextStr !=
		inComingErrPrefixDto.lenErrorContextStr {

		inComingErrPrefixDto.lenErrorContextStr =
			targetErrPrefixDto.lenErrorContextStr
	}

	if targetErrPrefixDto.lenErrorContextStr > 0 {
		targetErrPrefixDto.errPrefixHasContextStr = true
	} else {
		targetErrPrefixDto.errPrefixHasContextStr = false
	}

	if targetErrPrefixDto.errPrefixHasContextStr !=
		inComingErrPrefixDto.errPrefixHasContextStr {

		inComingErrPrefixDto.errPrefixHasContextStr =
			targetErrPrefixDto.errPrefixHasContextStr
	}

	return nil
}

// copyOut - Creates a deep copy of the data fields contained in
// input parameter 'errPrefixDto' and returns that data as a new
// instance ErrorPrefixDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errPrefixDto        *ErrorPrefixDto
//     - A pointer to an instance of ErrorPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this object.
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
//  ErrorPrefixDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'errPrefixDto' will be returned.
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
func (ePrefDtoElectron *errorPrefixDtoElectron) copyOut(
	errPrefixDto *ErrorPrefixDto,
	ePrefix string) (
	ErrorPrefixDto,
	error) {

	if ePrefDtoElectron.lock == nil {
		ePrefDtoElectron.lock = new(sync.Mutex)
	}

	ePrefDtoElectron.lock.Lock()

	defer ePrefDtoElectron.lock.Unlock()

	ePrefix += "errorPrefixDtoElectron.copyOut() "

	newErrPrefixDto := ErrorPrefixDto{}

	if errPrefixDto == nil {
		return newErrPrefixDto,
			fmt.Errorf("%v\n"+
				"\nInput parameter 'errPrefixDto' is INVALID!\n"+
				"'errPrefixDto' is a nil pointer!\n",
				ePrefix)

	}

	ePrefDtoQuark := errorPrefixDtoQuark{}

	_,
		err := ePrefDtoQuark.testValidityOfErrorPrefixDto(
		errPrefixDto,
		ePrefix+
			"errPrefixDto\n")

	if err != nil {
		return newErrPrefixDto, err
	}

	newErrPrefixDto.isLastIdx =
		errPrefixDto.isLastIdx

	newErrPrefixDto.errorPrefixStr =
		errPrefixDto.errorPrefixStr

	newErrPrefixDto.lenErrorPrefixStr =
		uint(len(newErrPrefixDto.errorPrefixStr))

	if newErrPrefixDto.lenErrorPrefixStr !=
		errPrefixDto.lenErrorPrefixStr {

		errPrefixDto.lenErrorPrefixStr =
			newErrPrefixDto.lenErrorPrefixStr
	}

	newErrPrefixDto.errorContextStr =
		errPrefixDto.errorContextStr

	newErrPrefixDto.lenErrorContextStr =
		uint(len(newErrPrefixDto.errorContextStr))

	if newErrPrefixDto.lenErrorContextStr !=
		errPrefixDto.lenErrorContextStr {

		errPrefixDto.lenErrorContextStr =
			newErrPrefixDto.lenErrorContextStr
	}

	if newErrPrefixDto.lenErrorContextStr > 0 {
		newErrPrefixDto.errPrefixHasContextStr = true
	} else {
		newErrPrefixDto.errPrefixHasContextStr = false
	}

	if newErrPrefixDto.errPrefixHasContextStr !=
		errPrefixDto.errPrefixHasContextStr {

		errPrefixDto.errPrefixHasContextStr =
			newErrPrefixDto.errPrefixHasContextStr
	}

	return newErrPrefixDto, nil
}
