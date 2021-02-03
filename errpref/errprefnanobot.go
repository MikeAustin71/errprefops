package errpref

import (
	"strings"
	"sync"
)

type errPrefNanobot struct {
	lock *sync.Mutex
}

// oldFormatErrPrefix - Returns a string of formatted error prefix information
//func (ePrefNanobot *errPrefNanobot) oldFormatErrPrefix(
//	maxErrStringLength uint,
//	errPrefix string) string {
//
//	if ePrefNanobot.lock == nil {
//		ePrefNanobot.lock = new(sync.Mutex)
//	}
//
//	ePrefNanobot.lock.Lock()
//
//	defer ePrefNanobot.lock.Unlock()
//
//	if len(errPrefix) == 0 {
//		return ""
//	}
//
//	ePrefQuark := errPrefQuark{}
//
//	if maxErrStringLength == 0 {
//		maxErrStringLength = ePrefQuark.getErrPrefDisplayLineLength()
//	}
//
//	delimiters := ePrefQuark.getDelimiters()
//
//	ePrefElectron := errPrefElectron{}
//	var lenCleanStr int
//
//	errPrefix,
//		lenCleanStr = ePrefElectron.cleanErrorPrefixStr(errPrefix)
//
//	if lenCleanStr == 0 {
//		return ""
//	}
//
//	errPrefix = strings.ReplaceAll(
//		errPrefix,
//		delimiters.GetNewLineContextDelimiter(),
//		delimiters.GetInLineContextDelimiter())
//
//	errPrefix = strings.ReplaceAll(
//		errPrefix,
//		delimiters.GetInLinePrefixDelimiter(),
//		delimiters.GetNewLinePrefixDelimiter())
//
//	errPrefix += delimiters.GetNewLinePrefixDelimiter()
//
//	lErrPrefix := len(errPrefix)
//
//	errPrefixContextCollection := strings.Split(
//		errPrefix,
//		delimiters.GetNewLinePrefixDelimiter())
//
//	lCollection := len(errPrefixContextCollection)
//
//	if lCollection == 0 {
//		return ""
//	}
//
//	lenInLineContextDelimiter :=
//		int(delimiters.GetLengthInLineContextDelimiter())
//
//	var b1 strings.Builder
//	lErrPrefix += 512
//	b1.Grow(lErrPrefix)
//	var funcNames = make([]string, 0, 100)
//	var contextStrs = make([]string, 0, 100)
//	var contextIdx int
//
//	for _, s := range errPrefixContextCollection {
//
//		s,
//		lenCleanStr = ePrefElectron.cleanErrorPrefixStr(s)
//
//		if lenCleanStr == 0 {
//			continue
//		}
//
//		contextIdx = strings.Index(s,
//			delimiters.GetInLineContextDelimiter())
//
//		if contextIdx == -1 {
//			funcNames = append(funcNames, s)
//			contextStrs = append(contextStrs, "0")
//		} else {
//
//			funcNames = append(funcNames, s[0:contextIdx])
//			contextStrs = append(contextStrs, s[
//			contextIdx+lenInLineContextDelimiter-1:])
//		}
//
//	}
//
//	lenInLinePrefixDelimiter :=
//		int(delimiters.GetLengthInLinePrefixDelimiter())
//
//	lCollection = len(funcNames)
//
//	if lCollection == 0 {
//		return ""
//	}
//
//	var contextStr, funcName string
//	lastStr := ""
//	lastIdx := lCollection - 1
//
//	for i := 0; i < lCollection; i++ {
//		funcName = funcNames[i]
//		contextStr = contextStrs[i]
//
//		if contextStr == "0" {
//			contextStr = ""
//		}
//
//		if len(funcName) == 0 {
//			continue
//		}
//
//		if len(contextStr) == 0 {
//			// len(funcName) > 0 len(contextStr) == 0
//
//			if uint(len(lastStr)+len(funcName)) >
//				maxErrStringLength {
//
//				if len(lastStr) > 0 {
//
//					b1.WriteString(lastStr)
//					lastStr = ""
//
//					if uint(len(funcName)) > maxErrStringLength {
//
//						b1.WriteString(funcName)
//
//					} else {
//						// uint(len(funcName)) <= maxErrPrefixTextLineLength
//						if uint(len(funcName) + lenInLinePrefixDelimiter) > maxErrStringLength {
//
//							b1.WriteString(funcName)
//
//						} else {
//
//							if len(lastStr) > 0 {
//
//								lastStr +=
//								delimiters.GetInLinePrefixDelimiter() +
//									funcName
//
//							} else {
//
//								lastStr = funcName
//
//							}
//
//							continue
//						}
//					}
//				} else {
//					// len(lastStr) == 0
//
//					if uint(len(funcName)) > maxErrStringLength {
//
//						b1.WriteString(funcName)
//
//					} else {
//						// uint(len(funcName)) <= maxErrStringLength
//
//						if uint(len(funcName)) > maxErrStringLength {
//
//							b1.WriteString(funcName)
//
//						} else {
//							// uint(len(funcName)) <= maxErrPrefixTextLineLength
//							// And len(lastStr) == 0
//
//							lastStr = funcName
//
//							continue
//						}
//					}
//				}
//			} else {
//				// uint(len(lastStr) + len(funcName) + 3) <= maxErrPrefixTextLineLength
//
//				if len(lastStr) > 0 {
//
//					lastStr +=
//						delimiters.GetInLinePrefixDelimiter() +
//							funcName
//
//				} else {
//
//					lastStr = funcName
//
//				}
//
//				continue
//			}
//		} else {
//			// funcName Len > 0 && contextStr Len > 0
//
//			if uint(
//				len(lastStr)+
//					lenInLinePrefixDelimiter+
//					len(funcName)+
//					lenInLineContextDelimiter+
//					len(contextStr)) > maxErrStringLength {
//
//				if len(lastStr) > 0 {
//					b1.WriteString(lastStr + "\n")
//					lastStr = ""
//				}
//
//				if
//					uint(len(funcName)+
//						lenInLineContextDelimiter+
//						len(contextStr)) > maxErrStringLength {
//
//					b1.WriteString(funcName +
//						delimiters.GetNewLineContextDelimiter() +
//						contextStr)
//
//				} else {
//					// uint(len(funcName) + 3 + len(contextStr)) <= maxErrPrefixTextLineLength
//					lastStr = funcName +
//						delimiters.GetInLineContextDelimiter() +
//						contextStr
//
//					continue
//				}
//			} else {
//				//uint(len(lastStr) + 3 + len(funcName) + 3 + len(contextStr)) <= maxErrPrefixTextLineLength
//				if len(lastStr) == 0 {
//
//					lastStr =
//						funcName +
//							delimiters.GetInLineContextDelimiter() +
//							contextStr
//
//				} else {
//					lastStr +=
//						delimiters.GetInLinePrefixDelimiter() +
//							funcName +
//							delimiters.GetInLineContextDelimiter() +
//							contextStr
//				}
//				continue
//			}
//		}
//
//		if i != lastIdx {
//			b1.WriteString("\n")
//		}
//
//		if b1.Len() > lErrPrefix/2 {
//			b1.Grow(lErrPrefix)
//		}
//
//	}
//
//	if len(lastStr) > 0 {
//		b1.WriteString(lastStr)
//	}
//
//	return b1.String()
//}

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
	localErrPrefix := "errPrefNanobot.formatErrPrefix() "

	ePrefNeutron := errPrefNeutron{}

	prefixContextCol :=
		ePrefNeutron.getEPrefContextArray(
			errPrefix)

	lenPrefixContextCol := len(prefixContextCol)

	if lenPrefixContextCol == 0 {
		return localErrPrefix +
			"len(prefixContextCol)==0\n"
	}

	ePrefElectron := errPrefElectron{}

	delimiters := ePrefElectron.getDelimiters()

	lineLenCalculator := EPrefixLineLenCalc{}.Ptr()

	err :=
		lineLenCalculator.SetEPrefDelimiters(
			delimiters,
			localErrPrefix)

	if err != nil {
		return err.Error()
	}

	lineLenCalculator.SetMaxErrStringLength(
		maxErrStringLength)

	errPrefixStrLen := len(errPrefix)

	var b1 strings.Builder

	b1.Grow(errPrefixStrLen + 512)

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
				lineLenCalculator)

			if err != nil {
				return err.Error()
			}

			continue
		}

	}

	//if len(lastLineStr) > 0 {
	//	b1.WriteString(lastLineStr)
	//}

	return b1.String()
}
