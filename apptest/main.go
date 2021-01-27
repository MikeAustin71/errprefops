package main

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
	"strings"
)

func main() {
	TestMain01()
}

func TestMain02() {
	tStr := "Test01%Test02%Test03%Test04%Test05%Test06"

	fmt.Printf("Test String: %v\n", tStr)

	tStrArray := strings.Split(tStr, "%")

	for _, s := range tStrArray {
		fmt.Printf("%v\n", s)
	}

}

func TestMain01() {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx8.TryCombination()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx9.TryHammer()")
	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "x->y")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx10.X()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx11.TryAnything()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx12.TryToGetAJob()")

	ePrefix = errpref.ErrPref{}.FmtString(ePrefix)

	fmt.Printf(ePrefix)
}
