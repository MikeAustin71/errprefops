package errpref

import "testing"

func TestErrPrefixDelimiters_CopyIn_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyIn()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne,
		err := ErrPrefixDelimiters{}.New(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ErrPrefixDelimiters{}.New()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefDelimsTwo := ErrPrefixDelimiters{}

	err = ePrefDelimsTwo.CopyIn(
		&ePrefDelimsOne,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsTwo.CopyIn()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefDelimsTwo.IsValidInstance(funcName) {
		t.Error("ERROR: " +
			"Expected ePrefDelimsTwo.IsValidInstance()=='true'\n" +
			"Instead, ePrefDelimsTwo.IsValidInstance()=='false'\n")

		return
	}

	if ePrefDelimsTwo.GetNewLinePrefixDelimiter() !=
		newLinePrefixDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetNewLinePrefixDelimiter()=="+
			"newLinePrefixDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetNewLinePrefixDelimiter()='%v'\n"+
			"newLinePrefixDelimiters='%v'\n",
			ePrefDelimsTwo.GetNewLinePrefixDelimiter(),
			newLinePrefixDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter() !=
		uint(len(newLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter()=="+
			"len(newLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter()='%v'\n"+
			"len(newLinePrefixDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthNewLinePrefixDelimiter(),
			uint(len(newLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsTwo.GetInLinePrefixDelimiter() !=
		inLinePrefixDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetInLinePrefixDelimiter()=="+
			"inLinePrefixDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetInLinePrefixDelimiter()='%v'\n"+
			"inLinePrefixDelimiters='%v'\n",
			ePrefDelimsTwo.GetInLinePrefixDelimiter(),
			inLinePrefixDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthInLinePrefixDelimiter() !=
		uint(len(inLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthInLinePrefixDelimiter()=="+
			"len(inLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthInLinePrefixDelimiter()='%v'\n"+
			"len(inLinePrefixDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthInLinePrefixDelimiter(),
			uint(len(inLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsTwo.GetNewLineContextDelimiter() !=
		newLineContextDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetNewLineContextDelimiter()=="+
			"newLineContextDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetNewLineContextDelimiter()='%v'\n"+
			"newLineContextDelimiters='%v'\n",
			ePrefDelimsTwo.GetNewLineContextDelimiter(),
			newLineContextDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthNewLineContextDelimiter() !=
		uint(len(newLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthNewLineContextDelimiter()=="+
			"len(newLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthNewLineContextDelimiter()='%v'\n"+
			"len(newLineContextDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthNewLineContextDelimiter(),
			uint(len(newLineContextDelimiters)))

		return
	}

	if ePrefDelimsTwo.GetInLineContextDelimiter() !=
		inLineContextDelimiters {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetInLineContextDelimiter()=="+
			"inLineContextDelimiters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetInLineContextDelimiter()='%v'\n"+
			"inLineContextDelimiters='%v'\n",
			ePrefDelimsTwo.GetInLineContextDelimiter(),
			inLineContextDelimiters)

		return
	}

	if ePrefDelimsTwo.GetLengthInLineContextDelimiter() !=
		uint(len(inLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsTwo.GetLengthInLineContextDelimiter()=="+
			"len(inLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsTwo.GetLengthInLineContextDelimiter()='%v'\n"+
			"len(inLineContextDelimiters)='%v'\n",
			ePrefDelimsTwo.GetLengthInLineContextDelimiter(),
			uint(len(inLineContextDelimiters)))

		return
	}

	if !ePrefDelimsOne.Equal(&ePrefDelimsTwo) {
		t.Error("ERROR: " +
			"Expected ePrefDelimsOne==ePrefDelimsTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!!")
	}

}

func TestErrPrefixDelimiters_CopyOut_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyOut()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne := ErrPrefixDelimiters{}

	err := ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.SetDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	var ePrefDelimsTwo ErrPrefixDelimiters

	ePrefDelimsTwo,
		err = ePrefDelimsOne.CopyOut(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.CopyOut()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefDelimsOne.Equal(&ePrefDelimsTwo) {
		t.Error("ERROR: " +
			"Expected ePrefDelimsOne==ePrefDelimsTwo\n" +
			"HOWEVER, THEY ARE NOT EQUAL!!!")
	}
}

func TestErrPrefixDelimiters_Empty_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_CopyIn()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne,
		err := ErrPrefixDelimiters{}.New(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ErrPrefixDelimiters{}.New()\n"+
			"%v\n",
			err.Error())

		return
	}

	if !ePrefDelimsOne.IsValidInstance(funcName) {
		t.Error("ERROR: " +
			"Expected ePrefDelimsOne.IsValidInstance()=='true'\n" +
			"Instead, ePrefDelimsOne.IsValidInstance()=='false'\n")

		return
	}

	ePrefDelimsOne.Empty()

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err == nil {
		t.Error("ERROR: " +
			"Expected error return from ePrefDelimsOne.IsValidInstanceError()\n" +
			"after call to ePrefDelimsOne.Empty()\n" +
			"HOWEVER NO ERROR WAS RETURNED!\n")
	}

}

func TestErrPrefixDelimiters_SetLineLengthValues_000100(t *testing.T) {

	funcName := "TestErrPrefixDelimiters_SetLineLengthValues()"

	newLinePrefixDelimiters := "\n"
	inLinePrefixDelimiters := " - "
	newLineContextDelimiters := "\n : "
	inLineContextDelimiters := " : "

	ePrefDelimsOne := ErrPrefixDelimiters{}

	err := ePrefDelimsOne.SetDelimiters(
		newLinePrefixDelimiters,
		inLinePrefixDelimiters,
		newLineContextDelimiters,
		inLineContextDelimiters,
		funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.SetDelimiters()\n"+
			"%v\n",
			err.Error())

		return
	}

	err = ePrefDelimsOne.IsValidInstanceError(funcName)

	if err != nil {
		t.Errorf("ERROR returned by ePrefDelimsOne.IsValidInstanceError()\n"+
			"%v\n",
			err.Error())

		return
	}

	ePrefDelimsOne.SetLineLengthValues()

	if ePrefDelimsOne.GetLengthNewLinePrefixDelimiter() !=
		uint(len(newLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthNewLinePrefixDelimiter()=="+
			"len(newLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthNewLinePrefixDelimiter()='%v'\n"+
			"len(newLinePrefixDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthNewLinePrefixDelimiter(),
			uint(len(newLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsOne.GetLengthInLinePrefixDelimiter() !=
		uint(len(inLinePrefixDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthInLinePrefixDelimiter()=="+
			"len(inLinePrefixDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthInLinePrefixDelimiter()='%v'\n"+
			"len(inLinePrefixDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthInLinePrefixDelimiter(),
			uint(len(inLinePrefixDelimiters)))

		return
	}

	if ePrefDelimsOne.GetLengthNewLineContextDelimiter() !=
		uint(len(newLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthNewLineContextDelimiter()=="+
			"len(newLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthNewLineContextDelimiter()='%v'\n"+
			"len(newLineContextDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthNewLineContextDelimiter(),
			uint(len(newLineContextDelimiters)))

		return
	}

	if ePrefDelimsOne.GetLengthInLineContextDelimiter() !=
		uint(len(inLineContextDelimiters)) {
		t.Errorf("ERROR: "+
			"Expected ePrefDelimsOne.GetLengthInLineContextDelimiter()=="+
			"len(inLineContextDelimiters)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"ePrefDelimsOne.GetLengthInLineContextDelimiter()='%v'\n"+
			"len(inLineContextDelimiters)='%v'\n",
			ePrefDelimsOne.GetLengthInLineContextDelimiter(),
			uint(len(inLineContextDelimiters)))
	}

}
