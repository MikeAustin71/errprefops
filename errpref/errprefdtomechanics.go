package errpref

import (
	"fmt"
	"sync"
)

type errPrefixDtoMechanics struct {
	lock *sync.Mutex
}

// Ptr - Returns a pointer to a new instance of errPrefixDtoMechanics.
//
func (ePrefDtoMech errPrefixDtoMechanics) ptr() *errPrefixDtoMechanics {

	if ePrefDtoMech.lock == nil {
		ePrefDtoMech.lock = new(sync.Mutex)
	}

	ePrefDtoMech.lock.Lock()

	defer ePrefDtoMech.lock.Unlock()

	return &errPrefixDtoMechanics{
		lock: new(sync.Mutex),
	}
}

// setFromIBasicErrorPrefix - Receives an ErrPrefixDto object and
// an empty interface. If that empty interface is convertible to
// one of the 7-valid types listed below, this method then proceeds
// to overwrite and reset the ErrPrefixDto object's ErrorPrefixInfo
// collection containing error prefix and error context
// information. New error information will be extracted from input
// parameter, 'iEPref'. 'iEPref' is an empty interface which must
// be convertible to one of the following valid types:
//
//
//   1. nil - A nil value is valid and generates an empty
//            collection of error prefix and error context
//            information.
//
//   2. string - A string containing error prefix information.
//
//   3. []string A one-dimensional slice of strings containing
//               error prefix information
//
//   4. [][2]string A two-dimensional slice of strings containing
//                  error prefix and error context information.
//
//   5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                     ErrorPrefixInfo from this object will be
//                     copied to 'errPrefDto'.
//
//   6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                      ErrorPrefixInfo from this object will be
//                     copied to 'errPrefDto'.
//
//   7. IBasicErrorPrefix - An interface to a method generating
//                          a two-dimensional slice of strings
//                          containing error prefix and error
//                          context information.
//
// Any types not listed above will be considered invalid and
// trigger the return of an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. Data values in
//       this ErrPrefixDto object will be overwritten an set to new
//       values extracted from input parameter, 'iEPref'.
//
//
//  iEPref              interface{}
//     - An empty interface containing one of the 7-valid type
//       listed above. Any one of these valid types will generate
//       valid error prefix and error context information.
//
//       If 'iEPref' is NOT convertible to one of the 7-valid types
//       listed above, this method will return an error.
//
//
//  errorPrefStr        string
//     - This parameter contains an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to an empty string: "".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefStr'. The
//       'errorPrefStr' text will be attached to the beginning of
//       the error message.
//
func (ePrefDtoMech *errPrefixDtoMechanics) setFromEmptyInterface(
	errPrefDto *ErrPrefixDto,
	iEPref interface{},
	errorPrefStr string) error {

	if ePrefDtoMech.lock == nil {
		ePrefDtoMech.lock = new(sync.Mutex)
	}

	ePrefDtoMech.lock.Lock()

	defer ePrefDtoMech.lock.Unlock()

	methodNames := errorPrefStr + "\n" +
		"errPrefixDtoMechanics.setFromEmptyInterface()"

	if errPrefDto == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'errPrefDto' is invalid!\n"+
			"'errPrefDto' is a 'nil' pointer.\n",
			methodNames)
	}

	errNanobot := errPrefixDtoNanobot{}

	if iEPref == nil {

		errPrefDto.ePrefCol = nil

		return nil
	}

	var ok bool

	var str string

	str,
		ok = iEPref.(string)

	if ok {
		return errNanobot.setFromString(
			errPrefDto,
			str,
			methodNames)
	}

	var strs []string

	strs,
		ok = iEPref.([]string)

	if ok {
		return errNanobot.setFromStringArray(
			errPrefDto,
			strs,
			methodNames)
	}

	var twoDStrArray [][2]string

	twoDStrArray,
		ok = iEPref.([][2]string)

	if ok {
		return errNanobot.setFromTwoDStrArray(
			errPrefDto,
			twoDStrArray,
			methodNames)
	}

	ePrfAtom := errPrefAtom{}

	var dto ErrPrefixDto

	dto,
		ok = iEPref.(ErrPrefixDto)

	if ok {
		return ePrfAtom.
			copyInErrPrefDto(
				errPrefDto,
				&dto,
				methodNames)
	}

	var dtoPtr *ErrPrefixDto

	dtoPtr,
		ok = iEPref.(*ErrPrefixDto)

	if ok {
		return ePrfAtom.
			copyInErrPrefDto(
				errPrefDto,
				dtoPtr,
				methodNames)
	}

	var iBasicEPref IBasicErrorPrefix

	iBasicEPref,
		ok = iEPref.(IBasicErrorPrefix)

	if ok {
		return errNanobot.setFromIBasicErrorPrefix(
			errPrefDto,
			iBasicEPref,
			methodNames)
	}

	return fmt.Errorf("%v\n"+
		"Error: Could NOT extract error prefix information\n"+
		"from input parameter 'iEPref'.\n"+
		"'iEPref' IS NOT convertible to one of the seven\n"+
		"supported types listed as follows:\n"+
		"    nil\n"+
		"    string\n"+
		"    []string\n"+
		"    [][2]string\n"+
		"    ErrPrefixDto\n"+
		"    *ErrPrefixDto\n"+
		"    IBasicErrorPrefix\n",
		methodNames)
}
