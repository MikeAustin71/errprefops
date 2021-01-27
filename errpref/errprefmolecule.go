package errpref

import "sync"

type errPrefMolecule struct {
	lock *sync.Mutex
}

func (ePrefMolecule *errPrefMolecule) assembleNewErrPref(
	newErrPref string,
	newContext string) string {

	return ""
}
