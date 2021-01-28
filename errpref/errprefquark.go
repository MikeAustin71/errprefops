package errpref

import (
	"sync"
)

type errPrefQuark struct {
	lock *sync.Mutex
}

// isEmptyOrWhiteSpace - Analyzes the incoming string and returns
// 'true' if the strings is empty or consists of all white space.
//
func (ePrefQuark *errPrefQuark) isEmptyOrWhiteSpace(
	targetStr string) bool {

	if ePrefQuark.lock == nil {
		ePrefQuark.lock = new(sync.Mutex)
	}

	ePrefQuark.lock.Lock()

	defer ePrefQuark.lock.Unlock()

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}
