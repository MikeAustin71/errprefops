package errpref

import (
	"strings"
	"sync"
)

type errPrefQuark struct {
	lock *sync.Mutex
}

// formatErrPrefix - Returns a string of formatted error prefix information
func (ePrefQuark *errPrefQuark) formatErrPrefix(
	maxErrStringLength uint,
	errPrefix string) string {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

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

	errPrefixContextCollection := strings.Split(errPrefix, "\n")

	lCollection := len(errPrefixContextCollection)
	lastIdx := lCollection - 1

	var b1 strings.Builder
	var contextIdx int
	var s string

	for i := 0; i < lCollection; i++ {

		s = errPrefixContextCollection[i]

		contextIdx = strings.Index(s, " : ")

		if contextIdx == -1 {
			s = strings.TrimLeft(strings.TrimRight(s, " "), " ")

			b1.WriteString(s + "\n")

		} else {

			ePrefix := s[0:contextIdx]
			eContext := s[contextIdx+2:]

			if uint(contextIdx+len(eContext)+3) > maxErrStringLength {
				b1.WriteString(ePrefix + "\n : " + eContext)
			} else {

				b1.WriteString(ePrefix + " : " + eContext)
			}
		}

		if i != lastIdx {
			b1.WriteString("\n")
		}
	}

	return b1.String()
}
