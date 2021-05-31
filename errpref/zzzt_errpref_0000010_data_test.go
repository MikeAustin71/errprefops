package errpref

import "fmt"

func getValidEPrefixLineLenCalc(errPref string) (
	EPrefixLineLenCalc,
	error) {

	errPref += "getValidEPrefixLineLenCalc()\n"

	ePrefLineLenOne := EPrefixLineLenCalc{}

	ePrefInfoOne := ErrorPrefixInfo{}.Ptr()

	testMethod := "TxDoSomethingNow()"

	ePrefInfoOne.SetErrPrefixStr(testMethod)

	errContext := "A/B=C B=0"

	ePrefInfoOne.SetErrContextStr(errContext)

	ePrefInfoOne.SetIsFirstIndex(true)

	ePrefInfoOne.SetIsLastIndex(true)

	ePrefInfoOne.SetIsPopulated(true)

	err :=
		ePrefLineLenOne.SetErrPrefixInfo(
			ePrefInfoOne,
			errPref)

	if err != nil {

		return EPrefixLineLenCalc{},
			fmt.Errorf("%v\n"+
				"ERROR returned by ePrefLineLenOne.SetErrPrefixInfo()\n"+
				"%v\n",
				errPref,
				err.Error())
	}

	return ePrefLineLenOne, nil
}
