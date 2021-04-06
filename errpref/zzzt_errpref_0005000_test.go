package errpref

type testFuncAlpha01 struct {
	input string
}

func (tFuncAlpha01 *testFuncAlpha01) Tx1DoSomething(
	errorPrefix interface{}) error {

	var ePrefix *ErrPrefixDto
	var err error

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"testFuncAlpha01."+
			"Tx1DoSomething()",
		"X->Y")

	if err != nil {
		return err
	}

	tFuncAlpha2 := testFuncAlpha02{}

	err = tFuncAlpha2.Tx2DoSomethingElse(
		ePrefix)

	return err
}

type testFuncAlpha02 struct {
	input string
}

func (tFuncAlpha02 *testFuncAlpha02) Tx2DoSomethingElse(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha02." +
			"Tx2DoSomethingElse()")

	tFuncAlpha03 := testFuncAlpha03{}

	err := tFuncAlpha03.Tx3DoAnything(
		ePrefix)

	return err
}

type testFuncAlpha03 struct {
	input string
}

func (tFuncAlpha03 *testFuncAlpha03) Tx3DoAnything(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha03." +
			"Tx3DoAnything()")

	tFuncAlpha04 := testFuncAlpha04{}

	err := tFuncAlpha04.Tx4DoNothing(
		ePrefix)

	return err
}

type testFuncAlpha04 struct {
	input string
}

func (tFuncAlpha04 *testFuncAlpha04) Tx4DoNothing(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha04." +
			"Tx4DoNothing()")

	tFuncAlpha05 := testFuncAlpha05{}

	err := tFuncAlpha05.Tx5DoSomethingBig(
		ePrefix.XEPrefCtx(
			"Tx4DoNothing()",
			"A/B==4"))

	return err
}

type testFuncAlpha05 struct {
	input string
}

func (tFuncAlpha05 *testFuncAlpha05) Tx5DoSomethingBig(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha05." +
			"Tx5DoSomethingBig()")

	tFuncAlpha06 := testFuncAlpha06{}

	err := tFuncAlpha06.Tx6GiveUp(
		ePrefix.XCtx(
			"A->B"))

	return err
}

type testFuncAlpha06 struct {
	input string
}

func (tFuncAlpha06 *testFuncAlpha06) Tx6GiveUp(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncAlpha06." +
			"Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	return nil
}

type testFuncBravo01 struct {
	input string
}

func (tFuncBravo01 *testFuncBravo01) Tx1DoSomething(
	errorPrefix interface{}) error {

	var ePrefix *ErrPrefixDto
	var err error

	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"testFuncBravo01."+
			"Tx1DoSomething()",
		"")

	if err != nil {
		return err
	}

	tFuncBravo02 := testFuncBravo02{}

	err = tFuncBravo02.Tx2DoSomethingElse(ePrefix)

	return err
}

type testFuncBravo02 struct {
	input string
}

func (tFuncBravo02 *testFuncBravo02) Tx2DoSomethingElse(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo02." +
			"Tx2DoSomethingElse()")

	tFuncBravo03 := testFuncBravo03{}

	err := tFuncBravo03.Tx3DoAnything(ePrefix)

	return err
}

type testFuncBravo03 struct {
	input string
}

func (tFuncBravo03 *testFuncBravo03) Tx3DoAnything(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo03." +
			"Tx3DoAnything()")

	tFuncBravo04 := testFuncBravo04{}

	err := tFuncBravo04.Tx4DoNothing(ePrefix)

	return err
}

type testFuncBravo04 struct {
	input string
}

func (Bravo04 *testFuncBravo04) Tx4DoNothing(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo04." +
			"Tx4DoNothing()")

	tFuncBravo05 := testFuncBravo05{}

	ePrefix.SetCtx("A/B==4")

	err := tFuncBravo05.Tx5DoSomethingBig(ePrefix)

	return err
}

type testFuncBravo05 struct {
	input string
}

func (tFuncBravo05 *testFuncBravo05) Tx5DoSomethingBig(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"testFuncBravo05." +
			"Tx5DoSomethingBig()")

	tFuncBravo06 := testFuncBravo06{}

	ePrefix.SetCtx("A->B")

	err := tFuncBravo06.Tx6GiveUp(ePrefix)

	return err
}

type testFuncBravo06 struct {
	input string
}

func (tFuncY6 *testFuncBravo06) Tx6GiveUp(
	ePrefix *ErrPrefixDto) error {

	if ePrefix == nil {
		ePrefix = ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("Tx6GiveUp()")

	ePrefix.SetCtx("A/B = C B==0")

	return nil
}
