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
func (createEPrefDtoElectron *createErrPrefixDtoElectron) writeNewEPrefWithContext(
	strBuilder *strings.Builder,
	crEPrefDto *createErrPrefixDto,
	lastStr string,
	remainingLineLen uint) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if createEPrefDtoElectron.lock == nil {
		createEPrefDtoElectron.lock = new(sync.Mutex)
	}

	createEPrefDtoElectron.lock.Lock()

	defer createEPrefDtoElectron.lock.Unlock()

	newRemainingLineLen = remainingLineLen
	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))

	if strBuilder == nil ||
		crEPrefDto == nil ||
		crEPrefDto.errPrefIsEmpty {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	lenEPrefWithContext :=
		crEPrefDto.lenInLinePrefixDelimiter +
			crEPrefDto.lenNewErrPrefStr +
			crEPrefDto.lenInLineContextDelimiter +
			crEPrefDto.lenNewErrContextStr

	lenEPrefWithoutContext :=
		crEPrefDto.lenInLinePrefixDelimiter +
			crEPrefDto.lenNewErrPrefStr +
			crEPrefDto.lenNewLinePrefixDelimiter

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
				crEPrefDto)

		if lenEPrefWithContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.newErrPrefStr)

			if lenEPrefWithoutContext >
				remainingLineLen {

				strBuilder.WriteString(
					crEPrefDto.newLineContextDelimiter)

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						crEPrefDto.newLinePrefixDelimiter)
				}

				newRemainingLineLen =
					crEPrefDto.maxErrStringLength

				return newLastStr, newLenLastStr, newRemainingLineLen
				// End of lenEPrefWithoutContext >
				//				remainingLineLen
			} else {
				// lenEPrefWithContext > remainingLineLen

				strBuilder.WriteString(
					crEPrefDto.inLinePrefixDelimiter)

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						crEPrefDto.newLinePrefixDelimiter)
				}

				newRemainingLineLen =
					crEPrefDto.maxErrStringLength
				newLenLastStr = 0
				newLastStr = ""
			}
			// End of if lenEPrefWithContext > remainingLineLen
		} else {
			// lenEPrefWithContext <= remainingLineLen

			newLastStr += crEPrefDto.inLinePrefixDelimiter
			newLastStr += crEPrefDto.newErrPrefStr
			newLastStr += crEPrefDto.inLineContextDelimiter
			newLastStr += crEPrefDto.newErrContextStr
			newLenLastStr = uint(len(newLastStr))
			newRemainingLineLen =
				crEPrefDto.maxErrStringLength -
					newLenLastStr
		}

		return newLastStr, newLenLastStr, newRemainingLineLen
		// End Of newLenLastStr > remainingLineLen
	}

	//newLenLastStr <= remainingLineLen

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
					crEPrefDto)

		}

		if lenEPrefWithContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.newErrPrefStr)

			if lenEPrefWithoutContext >
				remainingLineLen {

				strBuilder.WriteString(
					crEPrefDto.newLineContextDelimiter)

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						crEPrefDto.newLinePrefixDelimiter)
				}

				newRemainingLineLen =
					crEPrefDto.maxErrStringLength

				return newLastStr, newLenLastStr, newRemainingLineLen
				// End of lenEPrefWithoutContext >
				//				remainingLineLen
			} else {
				// lenEPrefWithContext > remainingLineLen

				strBuilder.WriteString(
					crEPrefDto.inLinePrefixDelimiter)

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						crEPrefDto.newLinePrefixDelimiter)
				}

				newRemainingLineLen =
					crEPrefDto.maxErrStringLength
				newLenLastStr = 0
				newLastStr = ""
			}
			// End of if lenEPrefWithContext > remainingLineLen
		} else {
			// lenEPrefWithContext <= remainingLineLen

			newLastStr += crEPrefDto.inLinePrefixDelimiter
			newLastStr += crEPrefDto.newErrPrefStr
			newLastStr += crEPrefDto.inLineContextDelimiter
			newLastStr += crEPrefDto.newErrContextStr
			newLenLastStr = uint(len(newLastStr))
			newRemainingLineLen =
				crEPrefDto.maxErrStringLength -
					newLenLastStr
		}

		return newLastStr, newLenLastStr, newRemainingLineLen
		// End Of
		//newLenLastStr +
		//	lenEPrefWithContext > remainingLineLen
	}

	//newLenLastStr +
	//	lenEPrefWithContext <= remainingLineLen
	// The line length of the next write block
	// will fit on the end of the 'lastStr'

	newLastStr += crEPrefDto.inLinePrefixDelimiter
	newLastStr += crEPrefDto.newErrPrefStr
	newLastStr += crEPrefDto.inLineContextDelimiter
	newLastStr += crEPrefDto.newErrContextStr
	newLenLastStr = uint(len(newLastStr))
	newRemainingLineLen =
		crEPrefDto.maxErrStringLength -
			newLenLastStr

	return newLastStr, newLenLastStr, newRemainingLineLen
}

// writeNewEPrefWithOutContext
// Designed for Error Prefixes that do NOT have an associated error
// context string.
//
func (createEPrefDtoElectron *createErrPrefixDtoElectron) writeNewEPrefWithOutContext(
	strBuilder *strings.Builder,
	crEPrefDto *createErrPrefixDto,
	lastStr string,
	remainingLineLen uint) (
	newLastStr string,
	newLenLastStr uint,
	newRemainingLineLen uint) {

	if createEPrefDtoElectron.lock == nil {
		createEPrefDtoElectron.lock = new(sync.Mutex)
	}

	createEPrefDtoElectron.lock.Lock()

	defer createEPrefDtoElectron.lock.Unlock()

	newRemainingLineLen = remainingLineLen
	newLastStr = lastStr
	newLenLastStr = uint(len(lastStr))

	if strBuilder == nil ||
		crEPrefDto == nil ||
		crEPrefDto.errPrefIsEmpty {
		return newLastStr, newLenLastStr, newRemainingLineLen
	}

	lenEPrefWithoutContext :=
		crEPrefDto.lenInLinePrefixDelimiter +
			crEPrefDto.lenNewErrPrefStr +
			crEPrefDto.lenInLinePrefixDelimiter

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
				crEPrefDto)

		if lenEPrefWithoutContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.newErrPrefStr)

			if lenEPrefWithoutContext >
				remainingLineLen {

				strBuilder.WriteString(
					crEPrefDto.newLineContextDelimiter)

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						crEPrefDto.newLinePrefixDelimiter)
				}

				newRemainingLineLen =
					crEPrefDto.maxErrStringLength

				return newLastStr, newLenLastStr, newRemainingLineLen
				// End of lenEPrefWithoutContext >
				//				remainingLineLen
			} else {
				// lenEPrefWithoutContext <= remainingLineLen

				strBuilder.WriteString(
					crEPrefDto.inLinePrefixDelimiter)

				strBuilder.WriteString(
					crEPrefDto.newErrContextStr)

				if !crEPrefDto.isLastIdx {
					strBuilder.WriteString(
						crEPrefDto.newLinePrefixDelimiter)
				}

				newRemainingLineLen =
					crEPrefDto.maxErrStringLength
				newLenLastStr = 0
				newLastStr = ""
			}
			// End of if lenEPrefWithoutContext > remainingLineLen

		} else {
			// lenEPrefWithoutContext <= remainingLineLen
			// Add the next write block to 'lastStr'
			// Add to 'lastStr'.

			newLastStr += crEPrefDto.inLinePrefixDelimiter
			newLastStr += crEPrefDto.newErrPrefStr
			newLenLastStr = uint(len(newLastStr))
			newRemainingLineLen =
				crEPrefDto.maxErrStringLength -
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
					crEPrefDto)

		}

		if lenEPrefWithoutContext > remainingLineLen {

			strBuilder.WriteString(
				crEPrefDto.newErrPrefStr)

			if !crEPrefDto.isLastIdx {
				strBuilder.WriteString(
					crEPrefDto.newLinePrefixDelimiter)
			}

			newRemainingLineLen =
				crEPrefDto.maxErrStringLength
			newLenLastStr = 0
			newLastStr = ""
			// End of if lenEPrefWithoutContext > remainingLineLen
		} else {
			// lenEPrefWithoutContext <= remainingLineLen
			// Add to 'lastStr'

			newLastStr += crEPrefDto.inLinePrefixDelimiter
			newLastStr += crEPrefDto.newErrPrefStr
			newLenLastStr = uint(len(newLastStr))
			newRemainingLineLen =
				crEPrefDto.maxErrStringLength -
					newLenLastStr
		}

		// End Of
		//newLenLastStr +
		//	lenEPrefWithoutContext > remainingLineLen
	} else {
		//newLenLastStr +
		//	lenEPrefWithoutContext <= remainingLineLen
		newLastStr += crEPrefDto.inLinePrefixDelimiter
		newLastStr += crEPrefDto.newErrPrefStr
		newLastStr += crEPrefDto.inLineContextDelimiter
		newLastStr += crEPrefDto.newErrContextStr
		newLenLastStr = uint(len(newLastStr))
		newRemainingLineLen =
			crEPrefDto.maxErrStringLength -
				newLenLastStr
	}

	return newLastStr, newLenLastStr, newRemainingLineLen
}
