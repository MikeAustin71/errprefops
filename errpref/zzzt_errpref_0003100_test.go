package errpref

import "testing"

func TestErrPrefixDto_Series_000100(t *testing.T) {

	var err error

	tFuncAlpha01 := testFuncAlpha01{}

	errPrefix := new(ErrPrefixDto)

	errPrefix.SetEPref("StartMethod()")

	err = tFuncAlpha01.Tx1DoSomething(
		errPrefix)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncAlpha01.Tx1DoSomething()\n"+
			"%v\n", err.Error())

		return
	}

	tFuncBravo01 := testFuncBravo01{}

	err = tFuncBravo01.Tx1DoSomething(
		errPrefix)

	if err != nil {
		t.Errorf("Unexpected error return from\n"+
			"tFuncBravo01.Tx1DoSomething()\n"+
			"%v\n", err.Error())

		return
	}

}
