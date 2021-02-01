package errpref

import (
	"strings"
	"sync"
)

type errorPrefixDtoQuark struct {
	lock *sync.Mutex
}
