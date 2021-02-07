package main

import (
	"github.com/MikeAustin71/errprefops/apptest/testmain"
)

func main() {
	tm := testmain.TestMain{}
	tm.TestMain007()
}

/*
Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\nTx4() - Tx5() - Tx6.DoSomethingElse()\nTx7.TrySomethingNew() :  something->newSomething\nTx8.TryAnyCombination() - Tx9.TryAHammer() :  x->y\nTx10.X() - Tx11.TryAnything() - Tx12.TryASalad()\nTx13.SomeFabulousAndComplexStuff()\nTx14.MoreAwesomeGoodness :  A=7 B=8 C=9



*/
