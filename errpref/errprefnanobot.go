package errpref

import (
	"strings"
	"sync"
)

type errPrefNanobot struct {
	lock *sync.Mutex
}

// extractLastErrPrefInStringSeries - Receives a string containing error
// prefixes and proceeds to extract and return the last error
// prefix in the series along with the last error context and the
// modified error prefix series.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errPref             string
//     - A string containing error prefixes in series. This method
//       will extract the last error prefix and error context from
//       this string.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  oldErrPref          string
//     - This string holds the modified series of error prefixes.
//       It is identical to input parameter, 'errPref', except that
//       the last error prefix and error context have been removed.
//
//
//  newErrPref          string
//     - A string containing the last error prefix extracted from
//       input parameter, 'errPref'.
//
//
//  newErrContext       string
//     - A string containing the last error context description
//       extracted from input parameter, 'errPref'.
//
func (ePrefNanobot *errPrefNanobot) extractLastErrPrefInStringSeries(
	errPref string) (
	oldErrPref string,
	newErrPref string,
	newErrContext string) {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	ePrefElectron := errPrefElectron{}
	var lenCleanStr int

	errPref,
		lenCleanStr =
		ePrefElectron.cleanErrorPrefixStr(errPref)

	if lenCleanStr == 0 {
		return oldErrPref, newErrPref, newErrContext
	}

	delimiters := ePrefElectron.getDelimiters()

	var lastIdxDashPref, lastIdxNewLinePref,
		lastPrefixIdx int

	lastIdxDashPref =
		strings.LastIndex(
			errPref,
			delimiters.GetInLinePrefixDelimiter())

	lastIdxNewLinePref =
		strings.LastIndex(
			errPref,
			delimiters.GetNewLinePrefixDelimiter())

	if lastIdxDashPref > lastIdxNewLinePref {

		lastPrefixIdx = lastIdxDashPref

	} else if lastIdxNewLinePref > lastIdxDashPref {

		lastPrefixIdx = lastIdxNewLinePref

	} else {
		// lastIdxNewLinePref == lastIdxDashPref
		// This signals that there is only one
		// error prefix in the series.

		lastPrefixIdx = 0
	}

	if lastPrefixIdx > 0 {
		oldErrPref = errPref[0:lastPrefixIdx]

		oldErrPref,
			lenCleanStr =
			ePrefElectron.cleanErrorPrefixStr(oldErrPref)

	}

	tempNewErrPref := errPref[lastPrefixIdx:]

	tempNewErrPref,
		lenCleanStr =
		ePrefElectron.cleanErrorPrefixStr(tempNewErrPref)

	if lenCleanStr == 0 {
		return oldErrPref, newErrPref, newErrContext
	}

	var (
		lastIdxColonContext, lastIdxNewLineContext,
		lastContextIdx int
	)

	lastIdxNewLineContext =
		strings.LastIndex(
			tempNewErrPref,
			delimiters.GetNewLineContextDelimiter())

	lastIdxColonContext =
		strings.LastIndex(
			tempNewErrPref,
			delimiters.GetInLineContextDelimiter())

	if lastIdxNewLineContext > lastIdxColonContext {

		lastContextIdx = lastIdxNewLineContext

	} else if lastIdxNewLineContext < lastIdxColonContext {

		lastContextIdx = lastIdxColonContext

	} else {
		//  lastIdxNewLineContext == lastIdxColonContext
		// This signals that there is no context
		// present in tempNewErrPref
		lastContextIdx = -1
	}

	if lastContextIdx > -1 {
		newErrPref = tempNewErrPref[0:lastContextIdx]
		newErrContext = tempNewErrPref[lastContextIdx:]

		newErrPref,
			_ =
			ePrefElectron.cleanErrorPrefixStr(newErrPref)

		newErrContext,
			_ =
			ePrefElectron.cleanErrorContextStr(newErrContext)

	} else {
		// lastContextIdx = -1
		// This signals that there is no context
		// present in tempNewErrPref
		newErrPref = tempNewErrPref
	}

	return oldErrPref, newErrPref, newErrContext
}

// formatErrPrefixComponents - Returns a string of formatted error
// prefix information based on an array of error prefix elements.
//
func (ePrefNanobot *errPrefNanobot) formatErrPrefixComponents(
	maxErrStringLength uint,
	prefixContextCol []ErrorPrefixInfo) string {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	localErrPrefix := "errPrefNanobot.formatErrPrefixComponents() "

	lenPrefixContextCol := len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return localErrPrefix +
			"len(prefixContextCol)==0\n"
	}

	if maxErrStringLength == 0 {
		maxErrStringLength =
			errPrefQuark{}.ptr().
				getErrPrefDisplayLineLength()
	}

	delimiters := errPrefElectron{}.
		ptr().getDelimiters()

	lineLenCalculator := EPrefixLineLenCalc{}

	err :=
		lineLenCalculator.SetEPrefDelimiters(
			delimiters,
			localErrPrefix)

	if err != nil {
		return err.Error()
	}

	lineLenCalculator.SetMaxErrStringLength(
		maxErrStringLength)

	var b1 strings.Builder

	b1.Grow(1024)

	lineLenCalculator.SetCurrentLineStr("")

	lastIdx := lenPrefixContextCol - 1

	ePrefMolecule := errPrefMolecule{}

	for i := 0; i < lenPrefixContextCol; i++ {

		if i == lastIdx {
			prefixContextCol[i].SetIsLastIndex(true)
		}

		err =
			lineLenCalculator.SetErrPrefixInfo(
				&prefixContextCol[i],
				localErrPrefix)

		if err != nil {
			return err.Error()
		}

		err =
			lineLenCalculator.IsValidInstanceError(
				localErrPrefix)

		if err != nil {
			return err.Error()
		}

		if lineLenCalculator.ErrPrefixHasContext() {

			err = ePrefMolecule.writeNewEPrefWithContext(
				&b1,
				&lineLenCalculator)

			if err != nil {
				return err.Error()
			}

			continue
		} else {

			err = ePrefMolecule.writeNewEPrefWithOutContext(
				&b1,
				&lineLenCalculator)

			if err != nil {
				return err.Error()
			}
		}

	}

	if lineLenCalculator.GetCurrLineStrLength() > 0 {
		b1.WriteString(lineLenCalculator.GetDelimiterNewLineErrPrefix())
		errPrefNeutron{}.ptr().writeCurrentLineStr(
			&b1,
			&lineLenCalculator)
	}

	return b1.String()
}

// ptr() - Returns a pointer to a new instance of
// errPrefNanobot.
//
func (ePrefNanobot errPrefNanobot) ptr() *errPrefNanobot {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	return &errPrefNanobot{}
}
