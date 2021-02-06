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
//  oldErrPrefix                  string
//     - The existing or previous error prefix string. This text
//       string usually consists of a series of function names and
//       associated error context strings.
//
//
//  newErrPrefix                  string
//     - This is the new error prefix string which will be added to
//       existing error prefix string represented by input
//       parameter, 'oldErrPref'.
//
//
//  newErrContext                 string
//     - An optional error context description. This is the error
//       context information associated with the new error prefix
//       ('newErrPref'). Typically context descriptions might
//       include variable names or input values. The text
//       description is expected to help identify and explain any
//       errors triggered in the immediate vicinity of the function
//       documented by error prefix 'newErrPrefix'.
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
	oldErrPrefix string,
	newErrPrefix string,
	newErrContext string,
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
		lenOldErrPrefCleanStr int
	)

	var errPrefixInfo ErrorPrefixInfo
	var err error

	eMsgPrefix := "errPrefMechanics.assembleErrPrefix() "

	errPrefixInfo,
		err = errPrefNeutron{}.ptr().createNewEPrefInfo(
		newErrPrefix,
		newErrContext,
		eMsgPrefix)

	if err != nil {
		return ""
	}

	ePrefElectron := errPrefElectron{}

	oldErrPrefix,
		lenOldErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(oldErrPrefix)

	var prefixContextCol []ErrorPrefixInfo

	var lenPrefixContextCol int

	prefixContextCol = make([]ErrorPrefixInfo, 0, 256)

	lenPrefixContextCol = 0

	if lenOldErrPrefCleanStr > 0 {

		prefixContextCol =
			errPrefAtom{}.ptr().getEPrefContextArray(
				oldErrPrefix)

		lenPrefixContextCol = len(prefixContextCol)

		if lenPrefixContextCol > 0 {

			prefixContextCol[lenPrefixContextCol-1].
				SetIsLastIndex(false)

		}

	}

	errPrefixInfo.SetIsLastIndex(true)

	prefixContextCol = append(prefixContextCol, errPrefixInfo)

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

	if maxErrStringLength == 0 {
		maxErrStringLength = errPrefQuark{}.ptr().getErrPrefDisplayLineLength()
	}
	localErrPrefix := "errPrefMechanics.formatErrPrefix() "

	prefixContextCol :=
		errPrefAtom{}.ptr().getEPrefContextArray(
			errPrefix)

	lenPrefixContextCol := len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return localErrPrefix +
			"len(prefixContextCol)==0\n"
	}

	return errPrefNanobot{}.ptr().
		formatErrPrefixComponents(
			maxErrStringLength,
			prefixContextCol)
}

// ptr() - Returns a pointer to a new instance of
// errPrefMechanics.
//
func (ePrefMech errPrefMechanics) ptr() *errPrefMechanics {

	if ePrefMech.lock == nil {
		ePrefMech.lock = new(sync.Mutex)
	}

	ePrefMech.lock.Lock()

	defer ePrefMech.lock.Unlock()

	return &errPrefMechanics{}
}

func (ePrefMech *errPrefMechanics) setErrorContext(
	oldErrPref string,
	newErrContext string,
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
		lenPrefixContextCol int
	)

	ePrefElectron := errPrefElectron{}

	oldErrPref,
		lenOldErrPrefCleanStr =
		ePrefElectron.cleanErrorPrefixStr(oldErrPref)

	if lenOldErrPrefCleanStr == 0 {
		return ""
	}

	newErrContext,
		lenNewErrContextCleanStr =
		ePrefElectron.cleanErrorContextStr(newErrContext)

	if lenNewErrContextCleanStr == 0 {
		return oldErrPref
	}

	var prefixContextCol []ErrorPrefixInfo

	prefixContextCol =
		errPrefAtom{}.ptr().getEPrefContextArray(
			oldErrPref)

	lenPrefixContextCol = len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return ""
	}

	lenPrefixContextCol--

	prefixContextCol[lenPrefixContextCol].
		SetErrPrefixHasContext(true)

	prefixContextCol[lenPrefixContextCol].
		SetErrContextStr(newErrContext)

	return errPrefNanobot{}.ptr().formatErrPrefixComponents(
		maxErrStringLength,
		prefixContextCol)
}
