package testmain

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/errpref"
	"strings"
)

type TestMain struct {
	testStr01 string
}

func (tMain *TestMain) TestMain003() {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	// errpref.ErrPref{}.SetMaxErrPrefTextLineLength(60)

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx8.TryAnyCombination()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx9.TryAHammer()")
	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "x->y")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx10.X()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx11.TryAnything()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx12.TryASalad()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx13.SomeFabulousAndComplexStuff()")
	ePrefix = errpref.ErrPref{}.NewContext(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = errpref.ErrPref{}.FmtString(ePrefix)

	fmt.Printf(ePrefix)

}

func (tMain *TestMain) TestMain002() {
	tStr := "Test01%Test02%Test03%Test04%Test05%Test06"

	fmt.Printf("Test String: %v\n", tStr)

	tStrArray := strings.Split(tStr, "%")

	for _, s := range tStrArray {
		fmt.Printf("%v\n", s)
	}

}

func (tMain *TestMain) TestMain001() {

	ePrefix := "Tx1.Something()\nTx2.SomethingElse()\nTx3.DoSomething()\nTx4() - Tx5()\nTx6.DoSomethingElse()\n"

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx7.TrySomethingNew()")

	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "something->newSomething")

	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx8.TryAnyCombination()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx9.TryAHammer()")
	ePrefix = errpref.ErrPref{}.AddContext(ePrefix, "x->y")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx10.X()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx11.TryAnything()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx12.TryASalad()")
	ePrefix = errpref.ErrPref{}.New(ePrefix, "Tx13.SomeFabulousAndComplexStuff()")
	ePrefix = errpref.ErrPref{}.NewContext(
		ePrefix,
		"Tx14.MoreAwesomeGoodness",
		"A=7 B=8 C=9")

	ePrefix = errpref.ErrPref{}.FmtString(ePrefix)

	fmt.Printf(ePrefix)

}
