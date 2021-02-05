package errpref

import (
	"sync"
)

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
//  string
//     - This method will concatenate the old error prefix, new
//       error prefix and the new error context. This consolidated
//       error prefix will then be returned to the calling
//       function.
//
func (ePrefMech *errPrefMechanics) assembleErrPrefix(
	oldErrPref string,
	newErrPref string,
	newContext string,
	maxErrStringLength uint) string {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	if maxErrStringLength == 0 {
		maxErrStringLength =
			errPrefQuark{}.ptr().getErrPrefDisplayLineLength()
	}

	var (
		lenOldErrPrefCleanStr,
		lenNewErrContextCleanStr,
		lenNewErrPrefCleanStr int
	)

	ePrefElectron := errPrefElectron{}

	oldErrPref,
		lenOldErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(oldErrPref)

	newErrPref,
		lenNewErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(newErrPref)

	if lenOldErrPrefCleanStr+
		lenNewErrPrefCleanStr == 0 {

		return "Error: Cleaned Old Error Prefix and" +
			" Cleaned New Error Prefix\n" +
			"strings have zero string length!\n"
	}

	newContext,
		lenNewErrContextCleanStr =
		ePrefElectron.cleanErrorContextStr(newContext)

	var prefixContextCol []ErrorPrefixInfo

	var lenPrefixContextCol int

	prefixContextCol = make([]ErrorPrefixInfo, 0, 256)

	lenPrefixContextCol = 0

	if lenOldErrPrefCleanStr > 0 {

		prefixContextCol =
			errPrefNeutron{}.ptr().getEPrefContextArray(
				oldErrPref)

		lenPrefixContextCol = len(prefixContextCol)

		if lenPrefixContextCol > 0 {

			prefixContextCol[lenPrefixContextCol-1].
				SetIsLastIndex(false)

		}

	}

	newErrPrefInfo := ErrorPrefixInfo{}

	newErrPrefInfo.SetIsLastIndex(true)
	newErrPrefInfo.SetErrPrefixStr(newErrPref)

	if lenNewErrContextCleanStr > 0 {

		newErrPrefInfo.SetErrPrefixHasContext(true)
		newErrPrefInfo.SetErrContextStr(newContext)

	} else {

		newErrPrefInfo.SetErrPrefixHasContext(false)

	}

	prefixContextCol = append(prefixContextCol, newErrPrefInfo)

	return errPrefNanobot{}.ptr().formatErrPrefixComponents(
		maxErrStringLength,
		prefixContextCol)
}

// formatErrPrefix - Returns a string of formatted error prefix information
func (ePrefMech *errPrefMechanics) formatErrPrefix(
	maxErrStringLength uint,
	errPrefix string) string {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	if maxErrStringLength == 0 {
		maxErrStringLength = ePrefQuark.getErrPrefDisplayLineLength()
	}
	localErrPrefix := "errPrefMechanics.formatErrPrefix() "

	prefixContextCol :=
		errPrefNeutron{}.ptr().getEPrefContextArray(
			errPrefix)

	lenPrefixContextCol := len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return localErrPrefix +
			"len(prefixContextCol)==0\n"
	}

	return errPrefNanobot{}.ptr().formatErrPrefixComponents(
		maxErrStringLength,
		prefixContextCol)
}
