package main

import (
	"fmt"
	"github.com/MikeAustin71/errprefops/apptest/testmain"
)

func main() {
	tm := testmain.TestMain{}
	actualResult := tm.TestMain003()

	expectedResult :=
		"Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\\nTx4() - Tx5() - Tx6.DoSomethingElse()\\nTx7.TrySomethingNew() :  something->newSomething\\nTx8.TryAnyCombination() - Tx9.TryAHammer() :  x->y\\nTx10.X() - Tx11.TryAnything() - Tx12.TryASalad()\\nTx13.SomeFabulousAndComplexStuff()\\nTx14.MoreAwesomeGoodness :  A=7 B=8 C=9"

	fmt.Printf("\n\nmain()-Analysis\n\n")

	if actualResult == expectedResult {
		fmt.Printf("SUCCESS! - Actual MATCHES Expected!\n")
	} else {
		fmt.Printf("FAILURE! - Actual DOES NOT MATCH Expected!\n")
	}
	fmt.Println("----------------------------------------")
	fmt.Println()
}

/*
Tx1.Something() - Tx2.SomethingElse() - Tx3.DoSomething()\nTx4() - Tx5() - Tx6.DoSomethingElse()\nTx7.TrySomethingNew() :  something->newSomething\nTx8.TryAnyCombination() - Tx9.TryAHammer() :  x->y\nTx10.X() - Tx11.TryAnything() - Tx12.TryASalad()\nTx13.SomeFabulousAndComplexStuff()\nTx14.MoreAwesomeGoodness :  A=7 B=8 C=9
*/
