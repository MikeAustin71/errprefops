package errpref

import (
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
