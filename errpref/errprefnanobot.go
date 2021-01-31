package errpref

import (
	"strings"
	"sync"
)

type errPrefNanobot struct {
	lock *sync.Mutex
}

// formatErrPrefix - Returns a string of formatted error prefix information
func (ePrefNanobot *errPrefNanobot) oldFormatErrPrefix(
	maxErrStringLength uint,
	errPrefix string) string {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	if len(errPrefix) == 0 {
		return ""
	}

	ePrefQuark := errPrefQuark{}

	if maxErrStringLength == 0 {
		maxErrStringLength = ePrefQuark.getErrPrefDisplayLineLength()
	}

	inLinePrefixDelimiter,
	newLinePrefixDelimiter,
	inLineContextDelimiter,
	newLineContextDelimiter := ePrefQuark.getDelimiters()


	ePrefElectron := errPrefElectron{}
	var lenCleanStr int

	errPrefix,
		lenCleanStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)

	if lenCleanStr == 0 {
		return ""
	}

	errPrefix = strings.ReplaceAll(
		errPrefix,
		newLineContextDelimiter,
		inLineContextDelimiter)

	errPrefix = strings.ReplaceAll(
		errPrefix,
		inLinePrefixDelimiter, newLinePrefixDelimiter)

	errPrefix += newLinePrefixDelimiter

	lErrPrefix := len(errPrefix)

	errPrefixContextCollection := strings.Split(
		errPrefix,
		newLinePrefixDelimiter)

	lCollection := len(errPrefixContextCollection)

	if lCollection == 0 {
		return ""
	}

	var b1 strings.Builder
	lErrPrefix += 512
	b1.Grow(lErrPrefix)
	var funcNames = make([]string, 0, 100)
	var contextStrs = make([]string, 0, 100)
	var contextIdx int

	lenInLineContextDelimiter := len(inLineContextDelimiter)

	for _, s := range errPrefixContextCollection {

		s,
		lenCleanStr = ePrefElectron.cleanErrorPrefixStr(s)

		if lenCleanStr == 0 {
			continue
		}

		contextIdx = strings.Index(s, inLineContextDelimiter)

		if contextIdx == -1 {
			funcNames = append(funcNames, s)
			contextStrs = append(contextStrs, "0")
		} else {

			funcNames = append(funcNames, s[0:contextIdx])
			contextStrs = append(contextStrs, s[
			contextIdx+lenInLineContextDelimiter-1:])
		}

	}

	lenInLinePrefixDelimiter := len(inLinePrefixDelimiter)

	lCollection = len(funcNames)

	if lCollection == 0 {
		return ""
	}

	var contextStr, funcName string
	lastStr := ""
	lastIdx := lCollection - 1

	for i := 0; i < lCollection; i++ {
		funcName = funcNames[i]
		contextStr = contextStrs[i]

		if contextStr == "0" {
			contextStr = ""
		}

		if len(funcName) == 0 {
			continue
		}

		if len(contextStr) == 0 {
			// len(funcName) > 0 len(contextStr) == 0

			if uint(len(lastStr)+len(funcName)) >
				maxErrStringLength {

				if len(lastStr) > 0 {

					b1.WriteString(lastStr)
					lastStr = ""

					if uint(len(funcName)) > maxErrStringLength {

						b1.WriteString(funcName)

					} else {
						// uint(len(funcName)) <= maxErrPrefixTextLineLength
						if uint(len(funcName) + lenInLinePrefixDelimiter) > maxErrStringLength {

							b1.WriteString(funcName)

						} else {

							if len(lastStr) > 0 {

								lastStr +=
									inLinePrefixDelimiter + funcName

							} else {

								lastStr = funcName

							}

							continue
						}
					}
				} else {
					// len(lastStr) == 0

					if uint(len(funcName)) > maxErrStringLength {

						b1.WriteString(funcName)

					} else {
						// uint(len(funcName)) <= maxErrStringLength

						if uint(len(funcName)) > maxErrStringLength {

							b1.WriteString(funcName)

						} else {
							// uint(len(funcName)) <= maxErrPrefixTextLineLength
							// And len(lastStr) == 0

							lastStr = funcName

							continue
						}
					}
				}
			} else {
				// uint(len(lastStr) + len(funcName) + 3) <= maxErrPrefixTextLineLength

				if len(lastStr) > 0 {
					lastStr += inLinePrefixDelimiter + funcName
				} else {
					lastStr = funcName
				}

				continue
			}
		} else {
			// funcName Len > 0 && contextStr Len > 0

			if uint(
				len(lastStr)+
					lenInLinePrefixDelimiter+
					len(funcName)+
					lenInLineContextDelimiter+
					len(contextStr)) > maxErrStringLength {

				if len(lastStr) > 0 {
					b1.WriteString(lastStr + "\n")
					lastStr = ""
				}

				if
					uint(len(funcName)+
						lenInLineContextDelimiter+
						len(contextStr)) > maxErrStringLength {

					b1.WriteString(funcName +
						newLineContextDelimiter +
						contextStr)

				} else {
					// uint(len(funcName) + 3 + len(contextStr)) <= maxErrPrefixTextLineLength
					lastStr = funcName +
						inLineContextDelimiter +
						contextStr

					continue
				}
			} else {
				//uint(len(lastStr) + 3 + len(funcName) + 3 + len(contextStr)) <= maxErrPrefixTextLineLength
				if len(lastStr) == 0 {

					lastStr =
						funcName +
							inLineContextDelimiter +
							contextStr

				} else {
					lastStr +=
						inLinePrefixDelimiter +
							funcName +
							inLineContextDelimiter +
							contextStr
				}
				continue
			}
		}

		if i != lastIdx {
			b1.WriteString("\n")
		}

		if b1.Len() > lErrPrefix/2 {
			b1.Grow(lErrPrefix)
		}

	}

	if len(lastStr) > 0 {
		b1.WriteString(lastStr)
	}

	return b1.String()
}

// formatErrPrefix - Returns a string of formatted error prefix information
func (ePrefNanobot *errPrefNanobot) formatErrPrefix(
	maxErrStringLength uint,
	errPrefix string) string {

	if ePrefNanobot.lock == nil {
		ePrefNanobot.lock = new(sync.Mutex)
	}

	ePrefNanobot.lock.Lock()

	defer ePrefNanobot.lock.Unlock()

	ePrefQuark := errPrefQuark{}

	if maxErrStringLength == 0 {
		maxErrStringLength = ePrefQuark.getErrPrefDisplayLineLength()
	}


	ePrefNeutron := errPrefNeutron{}

	prefixContextCol :=
		ePrefNeutron.getEPrefContextArray(
				errPrefix,
			maxErrStringLength)

	if lenPrefixContextCol == 0 {
		return ""
	}

	inLinePrefixDelimiter,
	newLinePrefixDelimiter,
	inLineContextDelimiter,
	newLineContextDelimiter := ePrefQuark.getDelimiters()

	lastIdx := lenPrefixContextCol - 1

	var b1 strings.Builder

	b1.Grow(errPrefixStrLen+512)

	lastLineStr := ""

	funcNameType := "1"
	contextType := "2"
	lenInLinePrefixDelimiter := len(inLinePrefixDelimiter)
	lenInLineContextDelimiter := len(inLineContextDelimiter)

	for i := 0; i < lenPrefixContextCol; i++ {

		remainingLineLen := int(maxErrStringLength)

		remainingLineLen -= len(lastLineStr)

		if remainingLineLen < 0 {
			b1.WriteString(lastLineStr)
			b1.WriteString(newLinePrefixDelimiter)
			lastLineStr = ""
			remainingLineLen = int(maxErrStringLength)
		}

		remainingLineLen -= len(prefixContextCol[i][0])
		remainingLineLen -= lenInLinePrefixDelimiter

		lenContextStr := len(prefixContextCol[i][1])

		if remainingLineLen < 0 {
			
		}

		if lenContextStr > 0 {
			remainingLineLen -=
		}


		nextWriteBlock := make([][]string, 0, 10)

		lenFuncName := uint(len(prefixContextCol[i][0]))

		if lenFuncName == 0 {
			continue
		}

		nextWriteBlock = append(nextWriteBlock,
			[]string{
				prefixContextCol[i][0],
				funcNameType})

		if prefixContextCol[i][1] != "0" {
			nextWriteBlock = append(nextWriteBlock,
				[]string{
					prefixContextCol[i][1],
					contextType})
		}


		if i != lastIdx {
			b1.WriteString("\n")
		}

		if b1.Len() > errPrefixStrLen  {
			b1.Grow(errPrefixStrLen + 512 )
		}

	}

	if len(lastLineStr) > 0 {
		b1.WriteString(lastLineStr)
	}

	return b1.String()
}
