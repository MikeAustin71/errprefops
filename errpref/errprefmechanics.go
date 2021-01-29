package errpref

import "sync"

type errPrefMechanics struct {
	lock *sync.Mutex
}

// assembleErrPrefix - Assembles, consolidates and formats an error
// prefix string from a previous error prefix, a new error prefix
// and the new error context.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  oldErrPref                    string
//     - The existing or previous error prefix string. This text
//       string usually consists of a series of function names and
//       associated error context strings.
//
//
//  newErrPref                    string
//     - This is the new error prefix string which will be added to
//       existing error prefix string represented by input
//       parameter, 'oldErrPref'.
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
//  consolidatedOldNewErrPref     string
//     - This method will concatenate the old error prefix, new
//       error prefix and the new error context. This consolidated
//       error prefix will then be returned to the calling
//       function.
//
func (ePrefMech *errPrefMechanics) assembleErrPrefix(
	oldErrPref string,
	newErrPref string,
	newContext string,
	maxErrStringLength uint) (
	consolidatedOldNewErrPref string) {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	if maxErrStringLength == 0 {
		maxErrStringLength = 40
	}

	ePrefElectron := errPrefElectron{}

	var lenOldErrPrefCleanStr,
		lenNewErrPrefCleanStr int

	oldErrPref,
		lenOldErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(oldErrPref)

	ePrefMolecule := errPrefMolecule{}

	newErrPref,
		lenNewErrPrefCleanStr,
		_,
		_ =
		ePrefMolecule.assembleNewErrPref(
			newErrPref,
			newContext,
			maxErrStringLength)

	if lenOldErrPrefCleanStr == 0 &&
		lenNewErrPrefCleanStr == 0 {

		consolidatedOldNewErrPref = ""

	} else if lenOldErrPrefCleanStr == 0 &&

		lenNewErrPrefCleanStr > 0 {

	} else if lenOldErrPrefCleanStr > 0 &&
		lenNewErrPrefCleanStr == 0 {

		consolidatedOldNewErrPref =
			oldErrPref

	} else {
		//lenOldErrPrefCleanStr > 0 &&
		//lenNewErrPrefCleanStr > 0

		if uint(lenOldErrPrefCleanStr+
			lenNewErrPrefCleanStr+3) >
			maxErrStringLength {

			consolidatedOldNewErrPref =
				oldErrPref +
					"\n" +
					newErrPref

		} else {
			//uint(lenOldErrPrefCleanStr +
			//	lenNewErrPrefCleanStr + 3) <=
			//		maxErrPrefixTextLineLength
			consolidatedOldNewErrPref =
				oldErrPref +
					" - " +
					newErrPref
		}
	}

	return consolidatedOldNewErrPref
}
