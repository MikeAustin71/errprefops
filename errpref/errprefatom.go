package errpref

import (
	"fmt"
	"strings"
	"sync"
)

type errPrefAtom struct {
	lock *sync.Mutex
}

// copyInErrPrefDto - Copies the data fields from an incoming
// instance of ErrPrefixDto ('inComingErrPrefixDto') to the data
// fields of the target ErrPrefixDto instance ('targetErrPrefixDto')
//
// All of the data fields in the 'targetErrPrefixDto' instance will
// be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetErrPrefixDto         *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This method
//       WILL CHANGE AND OVERWRITE all values of internal member
//       variables encapsulated in this object to achieve the
//       method's objectives.
//
//       This ErrPrefixDto instance will receive all the data
//       values contained from input parameter
//       'inComingErrPrefixDto'. When the copy operation is
//       completed, the data values in 'targetErrPrefixDto' will be
//       identical to those contained in input parameter,
//       'inComingErrPrefixDto'.
//
//
//  inComingErrPrefixDto       *ErrPrefixDto
//     - A pointer to an instance of ErrPrefixDto. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this ErrPrefixDto instance will be
//       copied to input parameter 'targetErrPrefixDto'.
//
//
//  eMsg                       string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end of
//       'eMsg'.
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
//       chain and text passed by input parameter, 'eMsg'.
//
func (ePrefAtom *errPrefAtom) copyInErrPrefDto(
	targetErrPrefixDto *ErrPrefixDto,
	inComingErrPrefixDto *ErrPrefixDto,
	eMsg string) error {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	eMsg += "errPrefAtom.copyInErrPrefDto() "

	if targetErrPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'targetErrPrefixDto' is INVALID!\n"+
			"'targetErrPrefixDto' is a nil pointer!\n",
			eMsg)

	}

	if inComingErrPrefixDto == nil {
		return fmt.Errorf("%v\n"+
			"\nInput parameter 'inComingErrPrefixDto' is INVALID!\n"+
			"'inComingErrPrefixDto' is a nil pointer!\n",
			eMsg)
	}

	targetErrPrefixDto.isLastLineTerminatedWithNewLine =
		inComingErrPrefixDto.isLastLineTerminatedWithNewLine

	if inComingErrPrefixDto.maxErrPrefixTextLineLength < 10 {

		inComingErrPrefixDto.maxErrPrefixTextLineLength =
			errPrefQuark{}.ptr().getMasterErrPrefDisplayLineLength()

	}

	targetErrPrefixDto.maxErrPrefixTextLineLength =
		inComingErrPrefixDto.maxErrPrefixTextLineLength

	if inComingErrPrefixDto.ePrefCol == nil {
		inComingErrPrefixDto.ePrefCol =
			make([]ErrorPrefixInfo, 0, 256)
	}

	lenIncomingEPrefCol :=
		len(inComingErrPrefixDto.ePrefCol)

	targetErrPrefixDto.ePrefCol =
		make(
			[]ErrorPrefixInfo,
			lenIncomingEPrefCol,
			lenIncomingEPrefCol+256)

	if lenIncomingEPrefCol == 0 {
		return nil
	}

	copy(
		targetErrPrefixDto.ePrefCol,
		inComingErrPrefixDto.ePrefCol)

	return nil
}

// getEPrefContextArray - Receives a string containing a series of
// error prefix strings and error context strings. This method then
// proceeds to parse the error prefix and error context elements
// returning the results in an array of ErrorPrefixInfo objects
// passed as input parameter, 'prefixContextCol' .
//
// All error prefix elements parsed from the 'errPrefix' string
// will be appended to 'prefixContextCol'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefix           string
//     - This string contains a series of error prefix and error
//       context strings. Input parameter 'errPrefix' will be
//       parsed into error prefix and error context components
//       before being appended to array of ErrorPrefixInfo objects
//       passed in input parameter, 'prefixContextCol'
//
//       This string should consist of a series of error prefix
//       strings. Error prefixes should be delimited by either a
//       new line character ('\n') or the in-line delimiter string,
//       " - ".
//
//       If this string contains associated error context strings
//       as well, they should be delimited with either a new line
//       delimiter string, "\n :  " or an in-line delimiter string,
//       " : ".
//
//
//  prefixContextCol    *[]ErrorPrefixInfo
//     - A pointer to an array of ErrorPrefixInfo objects. The
//       error prefix and error context data elements parsed from
//       input parameter, 'errPrefix', will be appended to this
//       array.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (ePrefAtom *errPrefAtom) getEPrefContextArray(
	errPrefix string,
	prefixContextCol *[]ErrorPrefixInfo) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	if *prefixContextCol == nil {
		*prefixContextCol = make([]ErrorPrefixInfo, 0, 256)
	}

	if len(errPrefix) == 0 {
		return
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	var lenCleanPrefixStr int

	errPrefix,
		lenCleanPrefixStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanPrefixStr == 0 {
		return
	}

	errPrefix = strings.ReplaceAll(
		errPrefix,
		delimiters.GetNewLineContextDelimiter(),
		delimiters.GetInLineContextDelimiter())

	errPrefix = strings.ReplaceAll(
		errPrefix,
		delimiters.GetInLinePrefixDelimiter(),
		delimiters.GetNewLinePrefixDelimiter())

	errPrefix += delimiters.GetNewLinePrefixDelimiter()

	errPrefixContextCollection := strings.Split(
		errPrefix,
		delimiters.GetNewLinePrefixDelimiter())

	lCollection := len(errPrefixContextCollection)

	if lCollection == 0 {
		return
	}

	var contextIdx int
	var idxLenInLineContextDelimiter = int(delimiters.GetLengthInLineContextDelimiter() - 1)
	var (
		errPrefixStr,
		errCtxStr string
	)

	ePrefNeutron := errPrefNeutron{}

	for i := 0; i < lCollection; i++ {

		s := errPrefixContextCollection[i]

		contextIdx = strings.Index(s,
			delimiters.GetInLineContextDelimiter())

		if contextIdx == -1 {

			element,
				err := ePrefNeutron.createNewEPrefInfo(
				s,
				"",
				"")

			if err != nil {
				continue
			}

			*prefixContextCol = append(*prefixContextCol, element)

		} else {

			errPrefixStr =
				s[0:contextIdx]

			errCtxStr = s[contextIdx+
				idxLenInLineContextDelimiter:]

			element,
				err := ePrefNeutron.createNewEPrefInfo(
				errPrefixStr,
				errCtxStr,
				"")

			if err != nil {
				continue
			}

			*prefixContextCol = append(*prefixContextCol, element)

		}
	}

	return
}

// setFlagsErrorPrefixInfoArray - Sets internal flags in an array
// of ErrorPrefixInfo objects
//
func (ePrefAtom *errPrefAtom) setFlagsErrorPrefixInfoArray(
	prefixContextCol []ErrorPrefixInfo) {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	if prefixContextCol == nil {
		prefixContextCol = make([]ErrorPrefixInfo, 0, 256)
	}

	lenCollection := len(prefixContextCol)

	if lenCollection == 0 {
		return
	}

	lastIdx := lenCollection - 1

	for i := 0; i < lenCollection; i++ {

		_ = prefixContextCol[i].GetIsPopulated()

		switch i {
		case 0:
			prefixContextCol[i].SetIsFirstIndex(true)
			prefixContextCol[i].SetIsLastIndex(false)
		case lastIdx:
			prefixContextCol[i].SetIsFirstIndex(false)
			prefixContextCol[i].SetIsLastIndex(true)
		default:
			prefixContextCol[i].SetIsFirstIndex(false)
			prefixContextCol[i].SetIsLastIndex(false)
		}
	}

	return
}

// ptr() - Returns a pointer to a new instance of
// errPrefAtom.
//
func (ePrefAtom errPrefAtom) ptr() *errPrefAtom {

	if ePrefAtom.lock == nil {
		ePrefAtom.lock = new(sync.Mutex)
	}

	ePrefAtom.lock.Lock()

	defer ePrefAtom.lock.Unlock()

	return &errPrefAtom{}
}
