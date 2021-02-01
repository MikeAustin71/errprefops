package errpref

import (
	"strings"
	"sync"
)

type createErrPrefixDtoAtom struct {
	lock *sync.Mutex
}

// writeNewEPrefWithContext
// Designed for Error Prefixes that DO have an associated error
// context string.
//
func (createEPrefDtoAtom *createErrPrefixDtoAtom) writeNewEPrefWithContext(
	strBuilder *strings.Builder,
	crEPrefDto *createErrPrefixDto,
	delimiters *EPrefixDelimiters,
	lastStr string,
	remainingLineLen uint) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if createEPrefDtoAtom.lock == nil {
		createEPrefDtoAtom.lock = new(sync.Mutex)
	}

	createEPrefDtoAtom.lock.Lock()

	defer createEPrefDtoAtom.lock.Unlock()

	newRemainingLineLen = remainingLineLen
	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))

	if strBuilder == nil ||
		crEPrefDto == nil ||
		delimiters == nil ||
		crEPrefDto.errPrefIsEmpty {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	lenEPrefWithContext :=
		delimiters.GetLengthInLinePrefixDelimiter() +
			crEPrefDto.lenNewErrPrefStr +
			delimiters.GetLengthInLineContextDelimiter() +
			crEPrefDto.lenNewErrContextStr

	createEPrefDtoQuark := createErrPrefixDtoQuark{}

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
				crEPrefDto.newErrPrefStr)

			strBuilder.WriteString(
				delimiters.GetNewLineContextDelimiter())

			strBuilder.WriteString(
				crEPrefDto.newErrContextStr)

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
	newLastStr += crEPrefDto.newErrPrefStr
	newLastStr += delimiters.GetInLineContextDelimiter()
	newLastStr += crEPrefDto.newErrContextStr
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
func (createEPrefDtoAtom *createErrPrefixDtoAtom) writeNewEPrefWithOutContext(
	strBuilder *strings.Builder,
	crEPrefDto *createErrPrefixDto,
	delimiters *EPrefixDelimiters,
	lastStr string,
	remainingLineLen uint) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if createEPrefDtoAtom.lock == nil {
		createEPrefDtoAtom.lock = new(sync.Mutex)
	}

	createEPrefDtoAtom.lock.Lock()

	defer createEPrefDtoAtom.lock.Unlock()

	newRemainingLineLen = remainingLineLen
	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))

	if strBuilder == nil ||
		crEPrefDto == nil ||
		crEPrefDto.errPrefIsEmpty {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	lenEPrefWithoutContext :=
		delimiters.GetLengthInLinePrefixDelimiter() +
			crEPrefDto.lenNewErrPrefStr +
			delimiters.GetLengthInLinePrefixDelimiter()

	createEPrefDtoQuark := createErrPrefixDtoQuark{}

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
				crEPrefDto.newErrPrefStr)

			if lenEPrefWithoutContext >
				remainingLineLen {

				strBuilder.WriteString(
					delimiters.GetNewLineContextDelimiter())

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

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
					crEPrefDto.newErrContextStr)

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

			newLastStr += crEPrefDto.newErrPrefStr

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
				crEPrefDto.newErrPrefStr)

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
			newLastStr += crEPrefDto.newErrPrefStr
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
		newLastStr += crEPrefDto.newErrPrefStr
		newLastStr += delimiters.GetInLineContextDelimiter()
		newLastStr += crEPrefDto.newErrContextStr
		newLenLastStr = uint(len(newLastStr))
		newRemainingLineLen =
			delimiters.GetMaxErrStringLength() -
				newLenLastStr
	}

	return newLastStr, newLenLastStr, newRemainingLineLen
}
