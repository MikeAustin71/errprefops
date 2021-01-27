package errpref

import (
	"strings"
	"sync"
)

type errPrefNanobot struct {
	lock *sync.Mutex
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

	if len(errPrefix) == 0 {
		return ""
	}

	if maxErrStringLength == 0 {
		maxErrStringLength = 40
	}

	errPrefix = strings.TrimRight(errPrefix, " ")
	errPrefix = strings.TrimRight(errPrefix, "\n")

	if len(errPrefix) == 0 {
		return ""
	}

	errPrefix = strings.ReplaceAll(errPrefix, "\n : ", " : ")

	errPrefix = strings.ReplaceAll(errPrefix, " - ", "\n")

	errPrefix += "\n"

	lErrPrefix := len(errPrefix)

	errPrefixContextCollection := strings.Split(errPrefix, "\n")

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

	for _, s := range errPrefixContextCollection {

		s = strings.TrimLeft(strings.TrimRight(s, " "), " ")
		s = strings.TrimLeft(strings.TrimRight(s, "\n"), "\n")

		if len(s) == 0 {
			continue
		}

		contextIdx = strings.Index(s, " : ")

		if contextIdx == -1 {
			funcNames = append(funcNames, s)
			contextStrs = append(contextStrs, "0")
		} else {

			funcNames = append(funcNames, s[0:contextIdx])
			contextStrs = append(contextStrs, s[contextIdx+2:])
		}

	}

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

			if uint(len(lastStr)+len(funcName)+3) > maxErrStringLength {

				if len(lastStr) > 0 {

					b1.WriteString(lastStr + "\n")
					lastStr = ""

					if uint(len(funcName)) > maxErrStringLength {
						b1.WriteString(funcName)
					} else {
						// uint(len(funcName)) <= maxErrStringLength
						if uint(len(funcName)) > maxErrStringLength {
							b1.WriteString(funcName)
						} else {

							if len(lastStr) > 0 {
								lastStr += " - " + funcName
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
						// len(lastStr) > 0
						if uint(len(funcName)) > maxErrStringLength {
							b1.WriteString(funcName)
						} else {
							// uint(len(funcName)) <= maxErrStringLength
							// And len(lastStr) == 0
							lastStr = funcName
							continue
						}
					}
				}
			} else {
				// uint(len(lastStr) + len(funcName) + 3) <= maxErrStringLength

				if len(lastStr) > 0 {
					lastStr += " - " + funcName
				} else {
					lastStr = funcName
				}

				continue
			}
		} else {
			// funcName Len > 0 && contextStr Len > 0

			if uint(len(lastStr)+3+len(funcName)+3+len(contextStr)) > maxErrStringLength {

				if len(lastStr) > 0 {
					b1.WriteString(lastStr + "\n")
					lastStr = ""
				}

				if uint(len(funcName)+3+len(contextStr)) > maxErrStringLength {
					b1.WriteString(funcName + "\n" + " : " + contextStr)
				} else {
					// uint(len(funcName) + 3 + len(contextStr)) <= maxErrStringLength
					lastStr = funcName + " : " + contextStr
					continue
				}
			} else {
				//uint(len(lastStr) + 3 + len(funcName) + 3 + len(contextStr)) <= maxErrStringLength
				if len(lastStr) == 0 {
					lastStr = funcName + " : " + contextStr
				} else {
					lastStr += lastStr + " - " + funcName + " : " + contextStr
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
