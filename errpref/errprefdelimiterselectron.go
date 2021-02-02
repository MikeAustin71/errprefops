package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDelimitersElectron struct {
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
func (ePrefDelimsElectron *errPrefixDelimitersElectron) copyIn(
	targetDelimiters *ErrPrefixDelimiters,
	incomingDelimiters *ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefDelimsElectron.lock == nil {
		ePrefDelimsElectron.lock = new(sync.Mutex)
	}

	ePrefDelimsElectron.lock.Lock()

	defer ePrefDelimsElectron.lock.Unlock()

	ePrefix += "errorPrefixDtoElectron.copyIn() "

	if targetDelimiters == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetDelimiters' is INVALID!\n"+
			"'targetDelimiters' is a nil pointer!\n",
			ePrefix)

	}

	if incomingDelimiters == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'incomingDelimiters' is INVALID!\n"+
			"'incomingDelimiters' is a nil pointer!\n",
			ePrefix)
	}

	ePrefDelimsQuark := errPrefixDelimitersQuark{}

	_,
		err := ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		incomingDelimiters,
		ePrefix+
			"incomingDelimiters\n")

	if err != nil {
		return err
	}

	targetDelimiters.inLinePrefixDelimiter =
		incomingDelimiters.inLinePrefixDelimiter

	targetDelimiters.lenInLinePrefixDelimiter =
		uint(len(targetDelimiters.inLinePrefixDelimiter))

	if incomingDelimiters.lenInLinePrefixDelimiter !=
		targetDelimiters.lenInLinePrefixDelimiter {
		incomingDelimiters.lenInLinePrefixDelimiter =
			targetDelimiters.lenInLinePrefixDelimiter
	}

	targetDelimiters.newLinePrefixDelimiter =
		incomingDelimiters.newLinePrefixDelimiter

	targetDelimiters.lenNewLinePrefixDelimiter =
		uint(len(targetDelimiters.newLinePrefixDelimiter))

	if incomingDelimiters.lenNewLinePrefixDelimiter !=
		targetDelimiters.lenNewLinePrefixDelimiter {
		incomingDelimiters.lenNewLinePrefixDelimiter =
			targetDelimiters.lenNewLinePrefixDelimiter
	}

	targetDelimiters.inLineContextDelimiter =
		incomingDelimiters.inLineContextDelimiter

	targetDelimiters.lenInLineContextDelimiter =
		uint(len(targetDelimiters.inLineContextDelimiter))

	if incomingDelimiters.lenInLineContextDelimiter !=
		targetDelimiters.lenInLineContextDelimiter {
		incomingDelimiters.lenInLineContextDelimiter =
			targetDelimiters.lenInLineContextDelimiter
	}

	targetDelimiters.newLineContextDelimiter =
		incomingDelimiters.newLineContextDelimiter

	targetDelimiters.lenNewLineContextDelimiter =
		uint(len(targetDelimiters.newLineContextDelimiter))

	if incomingDelimiters.lenNewLineContextDelimiter !=
		targetDelimiters.lenNewLineContextDelimiter {
		incomingDelimiters.lenNewLineContextDelimiter =
			targetDelimiters.lenNewLineContextDelimiter
	}

	return nil
}
