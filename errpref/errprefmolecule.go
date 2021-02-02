package errpref

import (
	"strings"
	"sync"
)

type errPrefMolecule struct {
	lock *sync.Mutex
}

// assembleNewErrPref - Assembles, consolidates and formats a new
// error prefix string from two subsidiary components: the new
// error prefix string and the associated new error context
// description.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newErrPref                    string
//     - This is the new error prefix string which will be combined
//       with the new error context description to create a
//       consolidated new error prefix text description. Typically,
//       a new error prefix string identifies the name of the
//       method which begin executing next.
//
//
//  newContext                    string
//     - An optional error context description. This is the error
//       context information associated with the new error prefix
//       ('newErrPref'). Typically context descriptions might
//       include variable names or input values. The text
//       description is expected to help identify and explain any
//       errors triggered in the immediate vicinity of this function.
//
//
//  maxErrPrefixTextLineLength            uint
//      - Specifies the maximum number of text characters which can
//        be on a single line for a new line character ('\n') is
//        inserted. If this value is zero, it will be set to the
//        default value of 40.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  consolidatedNewEPrefContext   string
//     - The text string which combines and format both the new
//       error prefix and the associated new error context
//       description.
//
//
//  lenConsolidatedNewEPrefContext int
//     - This parameter returns the length of the
//       'consolidatedNewEPrefContext' text string.
//
//  lenNewErrPrefCleanStr         int
//     - This parameter returns the length of the error prefix
//       string included in 'consolidatedNewEPrefContext'.
//
//  lenNewErrContextCleanStr      int
//     - This parameter returns the length of the error context
//       string included in 'consolidatedNewEPrefContext'.
//
func (ePrefMolecule *errPrefMolecule) assembleNewErrPref(
	newErrPref string,
	newContext string,
	maxErrStringLength uint) (
	consolidatedNewEPrefContext string,
	lenConsolidatedNewEPrefContext int,
	lenNewErrPrefCleanStr int,
	lenNewErrContextCleanStr int) {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	if maxErrStringLength == 0 {
		maxErrStringLength =
			ePrefQuark.getErrPrefDisplayLineLength()
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	newErrPref,
		lenNewErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(newErrPref)

	newContext,
		lenNewErrContextCleanStr =
		ePrefElectron.cleanErrorContextStr(newContext)

	if lenNewErrPrefCleanStr == 0 &&
		lenNewErrContextCleanStr == 0 {

		consolidatedNewEPrefContext = ""

	} else if lenNewErrPrefCleanStr > 0 &&
		lenNewErrContextCleanStr == 0 {

		consolidatedNewEPrefContext =
			newErrPref

	} else if lenNewErrPrefCleanStr == 0 &&
		lenNewErrContextCleanStr > 0 {

		consolidatedNewEPrefContext =
			newContext
	} else {
		//lenNewErrPrefCleanStr > 0 &&
		//lenNewErrContextCleanStr > 0
		if uint(lenNewErrPrefCleanStr+
			lenNewErrContextCleanStr+
			3) > maxErrStringLength {

			consolidatedNewEPrefContext =
				newErrPref +
					delimiters.GetNewLineContextDelimiter() +
					newContext

		} else {
			//uint(lenNewErrPrefCleanStr +
			//	lenNewErrContextCleanStr +
			//	3) <= maxErrPrefixTextLineLength
			consolidatedNewEPrefContext =
				newErrPref +
					delimiters.GetInLineContextDelimiter() +
					newContext
		}
	}

	lenConsolidatedNewEPrefContext =
		len(consolidatedNewEPrefContext)

	return consolidatedNewEPrefContext,
		lenConsolidatedNewEPrefContext,
		lenNewErrPrefCleanStr,
		lenNewErrContextCleanStr
}

// writeNewEPrefWithContext
// Designed for Error Prefixes that DO have an associated error
// context string.
//
func (ePrefMolecule *errPrefMolecule) writeNewEPrefWithContext(
	strBuilder *strings.Builder,
	crEPrefDto *ErrorPrefixDto,
	delimiters *ErrPrefixDelimiters,
	lastStr string,
	remainingLineLen uint) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	newRemainingLineLen = remainingLineLen
	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))

	if strBuilder == nil ||
		crEPrefDto == nil ||
		delimiters == nil ||
		crEPrefDto.isValid {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	lenEPrefWithContext :=
		delimiters.GetLengthInLinePrefixDelimiter() +
			crEPrefDto.lenErrorPrefixStr +
			delimiters.GetLengthInLineContextDelimiter() +
			crEPrefDto.lenErrorContextStr

	createEPrefDtoQuark := errorPrefixDtoQuark{}

	if newLenLastStr > remainingLineLen {
		// The lastStr is already longer than
		// remaining line length

		newLastStr,
			newLenLastStr,
			remainingLineLen =
			createEPrefDtoQuark.writeLastStr(
				strBuilder,
				lastStr,
				remainingLineLen,
				crEPrefDto,
				delimiters)
	}

	if newLenLastStr+
		lenEPrefWithContext > remainingLineLen {

		if newLenLastStr > 0 {

			newLastStr,
				newLenLastStr,
				remainingLineLen =
				createEPrefDtoQuark.writeLastStr(
					strBuilder,
					lastStr,
					remainingLineLen,
					crEPrefDto,
					delimiters)

		}

		if lenEPrefWithContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.errorPrefixStr)

			strBuilder.WriteString(
				delimiters.GetNewLineContextDelimiter())

			strBuilder.WriteString(
				crEPrefDto.errorContextStr)

			if !crEPrefDto.isLastIdx {
				strBuilder.WriteString(
					delimiters.GetNewLinePrefixDelimiter())
			}

			newRemainingLineLen =
				delimiters.GetMaxErrStringLength()

			return newLastStr, newLenLastStr, newRemainingLineLen
		}
		// End Of
		//newLenLastStr +
		//	lenEPrefWithContext > remainingLineLen
	}

	//newLenLastStr +
	//	lenEPrefWithContext <= remainingLineLen
	// The line length of the next write block
	// will fit on the end of the 'lastStr'

	newLastStr += delimiters.GetInLinePrefixDelimiter()
	newLastStr += crEPrefDto.errorPrefixStr
	newLastStr += delimiters.GetInLineContextDelimiter()
	newLastStr += crEPrefDto.errorContextStr
	newLenLastStr = uint(len(newLastStr))
	newRemainingLineLen =
		delimiters.GetMaxErrStringLength() -
			newLenLastStr

	return newLastStr, newLenLastStr, newRemainingLineLen
}

// writeNewEPrefWithOutContext
// Designed for Error Prefixes that do NOT have an associated error
// context string.
//
func (ePrefMolecule *errPrefMolecule) writeNewEPrefWithOutContext(
	strBuilder *strings.Builder,
	crEPrefDto *ErrorPrefixDto,
	delimiters *ErrPrefixDelimiters,
	lastStr string,
	remainingLineLen uint) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if ePrefMolecule.lock == nil {
		ePrefMolecule.lock = new(sync.Mutex)
	}

	ePrefMolecule.lock.Lock()

	defer ePrefMolecule.lock.Unlock()

	newRemainingLineLen = remainingLineLen
	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))

	if strBuilder == nil ||
		crEPrefDto == nil ||
		crEPrefDto.isValid {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	lenEPrefWithoutContext :=
		delimiters.GetLengthInLinePrefixDelimiter() +
			crEPrefDto.lenErrorPrefixStr +
			delimiters.GetLengthInLinePrefixDelimiter()

	createEPrefDtoQuark := errorPrefixDtoQuark{}

	if newLenLastStr > remainingLineLen {
		// The lastStr is already longer than
		// remaining line length

		newLastStr,
			newLenLastStr,
			remainingLineLen =
			createEPrefDtoQuark.writeLastStr(
				strBuilder,
				lastStr,
				remainingLineLen,
				crEPrefDto,
				delimiters)

		if lenEPrefWithoutContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.errorPrefixStr)

			if lenEPrefWithoutContext >
				remainingLineLen {

				strBuilder.WriteString(
					delimiters.GetNewLineContextDelimiter())

				strBuilder.WriteString(
					crEPrefDto.errorContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						delimiters.GetNewLinePrefixDelimiter())
				}

				newRemainingLineLen =
					delimiters.GetMaxErrStringLength()

				return newLastStr, newLenLastStr, newRemainingLineLen
				// End of lenEPrefWithoutContext >
				//				remainingLineLen
			} else {
				// lenEPrefWithoutContext <= remainingLineLen

				strBuilder.WriteString(
					delimiters.GetInLinePrefixDelimiter())

				strBuilder.WriteString(
					crEPrefDto.errorContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						delimiters.GetNewLinePrefixDelimiter())
				}

				newRemainingLineLen =
					delimiters.GetMaxErrStringLength()

				newLenLastStr = 0

				newLastStr = ""

			}
			// End of if lenEPrefWithoutContext > remainingLineLen

		} else {
			// lenEPrefWithoutContext <= remainingLineLen
			// Add the next write block to 'lastStr'
			// Add to 'lastStr'.

			newLastStr += delimiters.GetInLinePrefixDelimiter()

			newLastStr += crEPrefDto.errorPrefixStr

			newLenLastStr = uint(len(newLastStr))

			newRemainingLineLen =
				delimiters.GetMaxErrStringLength() -
					newLenLastStr
		}

		return newLastStr, newLenLastStr, newRemainingLineLen
		// End Of
		// newLenLastStr > remainingLineLen
	}

	if newLenLastStr+
		lenEPrefWithoutContext > remainingLineLen {

		if newLenLastStr > 0 {
			newLastStr,
				newLenLastStr,
				remainingLineLen =
				createEPrefDtoQuark.writeLastStr(
					strBuilder,
					lastStr,
					remainingLineLen,
					crEPrefDto,
					delimiters)

		}

		if lenEPrefWithoutContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.errorPrefixStr)

			if !crEPrefDto.isLastIdx {
				strBuilder.WriteString(
					delimiters.GetNewLinePrefixDelimiter())
			}

			newRemainingLineLen =
				delimiters.GetMaxErrStringLength()

			newLenLastStr = 0

			newLastStr = ""

			// End of if lenEPrefWithoutContext > remainingLineLen
		} else {
			// lenEPrefWithoutContext <= remainingLineLen
			// Add to 'lastStr'

			newLastStr += delimiters.GetInLinePrefixDelimiter()
			newLastStr += crEPrefDto.errorPrefixStr
			newLenLastStr = uint(len(newLastStr))
			newRemainingLineLen =
				delimiters.GetMaxErrStringLength() -
					newLenLastStr
		}

		// End Of
		//newLenLastStr +
		//	lenEPrefWithoutContext > remainingLineLen
	} else {
		//newLenLastStr +
		//	lenEPrefWithoutContext <= remainingLineLen
		newLastStr += delimiters.GetInLinePrefixDelimiter()
		newLastStr += crEPrefDto.errorPrefixStr
		newLastStr += delimiters.GetInLineContextDelimiter()
		newLastStr += crEPrefDto.errorContextStr
		newLenLastStr = uint(len(newLastStr))
		newRemainingLineLen =
			delimiters.GetMaxErrStringLength() -
				newLenLastStr
	}

	return newLastStr, newLenLastStr, newRemainingLineLen
}
