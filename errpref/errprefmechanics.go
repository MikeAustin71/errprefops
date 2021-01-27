package errpref

import "sync"

type errPrefMechanics struct {
	lock *sync.Mutex
}

func (ePrefMech *errPrefMechanics) assembleErrPrefix(
	oldErrPref string,
	newErrPref string,
	newContext string) string {

	return ""
}
