package main

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
)

func main() {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()"

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx4.TrySomethingNew()")

	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.FmtString(ePrefix)

	fmt.Printf(ePrefix)
}
