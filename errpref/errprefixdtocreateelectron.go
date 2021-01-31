package errpref

import (
	"sync"
)

type createErrPrefixDtoElectron struct {
	lock *sync.Mutex
}
