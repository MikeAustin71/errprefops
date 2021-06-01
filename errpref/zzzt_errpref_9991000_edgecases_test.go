package errpref

import (
	"strings"
	"testing"
)

func TestErrPref_Mechanics_000100(t *testing.T) {

	/*	func (ePrefMech *errPrefMechanics) assembleErrPrefix(
		oldErrPrefix string,
		newErrPrefix string,
		newErrContext string,
		maxErrStringLength uint,
		delimiters ErrPrefixDelimiters) string
	*/

	ePrefMech := errPrefMechanics{}

	delimiters := ErrPrefixDelimiters{}

	delimiters.SetToDefault()

	oldErrPrefix := "OldStuff()"
	newErrPrefix := "Tx2.DoSomething()"
	newErrContext := ""
	maxErrStringLength := uint(0)

	outputStr :=
		ePrefMech.assembleErrPrefix(
			oldErrPrefix,
			newErrPrefix,
			newErrContext,
			maxErrStringLength,
			delimiters)

	if len(outputStr) == 0 {
		t.Error("ERROR: ePrefMech.assembleErrPrefix() " +
			"returned a zero length output string!\n")
		return
	}

	/*	func (ePrefMech *errPrefMechanics) formatErrPrefix(
		maxErrStringLength uint,
		delimiters ErrPrefixDelimiters,
		errPrefix string) string
	*/

	maxErrStringLength = 0

	outputStr =
		ePrefMech.formatErrPrefix(
			maxErrStringLength,
			delimiters,
			newErrPrefix)

	if len(outputStr) == 0 {
		t.Error("ERROR: ePrefMech.formatErrPrefix() " +
			"returned a zero length output string!\n")
		return
	}

	/*
		func (ePrefMech *errPrefMechanics) setErrorContext(
			oldErrPref string,
			newErrContext string,
			maxErrStringLength uint,
			delimiters ErrPrefixDelimiters) string
	*/

	maxErrStringLength = 0
	newErrContext = "X=7; Y=92"

	outputStr =
		ePrefMech.setErrorContext(
			oldErrPrefix,
			newErrContext,
			maxErrStringLength,
			delimiters)

	if len(outputStr) == 0 {
		t.Error("ERROR: ePrefMech.setErrorContext() " +
			"returned a zero length output string!\n")
	}

}

func TestErrPref_Molecule_000100(t *testing.T) {

	ePrefMolecule := errPrefMolecule{}

	/*
		func (ePrefMolecule *errPrefMolecule) writeNewEPrefWithContext(
			strBuilder *strings.Builder,
		lineLenCalc *EPrefixLineLenCalc) error
	*/

	var strBuilder *strings.Builder

	var lineLenCalc *EPrefixLineLenCalc

	xLineLenCalc := EPrefixLineLenCalc{}.New()

	lineLenCalc = &xLineLenCalc

	err := ePrefMolecule.writeNewEPrefWithContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithContext()\n" +
			"because strBuilder is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err = ePrefMolecule.writeNewEPrefWithOutContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithOutContext()\n" +
			"because strBuilder is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	strBuilder = &strings.Builder{}

	lineLenCalc = nil

	err = ePrefMolecule.writeNewEPrefWithContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithContext()\n" +
			"because lineLenCalc is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err = ePrefMolecule.writeNewEPrefWithOutContext(
		strBuilder,
		lineLenCalc)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefMolecule.writeNewEPrefWithOutContext()\n" +
			"because lineLenCalc is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPrefixDelimiters_Quark_000100(t *testing.T) {
	/*
		func (ePrefDelimsQuark *errPrefixDelimitersQuark) empty(
			delimiters *ErrPrefixDelimiters,
			ePrefix string) (err error)
	*/

	ePrefDelimsQuark := errPrefixDelimitersQuark{}

	var delimiters *ErrPrefixDelimiters
	ePrefix := "TestCallingMethod()"

	err := ePrefDelimsQuark.empty(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.empty()\n" +
			"because 'delimiters' is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters' is nil.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters = &ErrPrefixDelimiters{}

	delimiters.SetToDefault()

	delimiters.inLinePrefixDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.inLinePrefixDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters.SetToDefault()

	delimiters.newLinePrefixDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.newLinePrefixDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters.SetToDefault()

	delimiters.inLineContextDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.inLineContextDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	delimiters.SetToDefault()

	delimiters.newLineContextDelimiter = ""

	_,
		err = ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		delimiters,
		ePrefix)

	if err == nil {
		t.Error("ERROR: Expected an error return from " +
			"ePrefDelimsQuark.testValidityOfErrPrefixDelimiters()\n" +
			"because 'delimiters.newLineContextDelimiter' is an empty string.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestEPrefixLineLenCalc_Electron_000100(t *testing.T) {

	funcName := "TestEPrefixLineLenCalc_Electron_000100() "

	ePrefLineLenCalcElectron := ePrefixLineLenCalcElectron{}

	var targetLineLenCalc, incomingLineLenCalc *EPrefixLineLenCalc

	incomingLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	ePrefix := "TestCallingMethod"

	err :=
		ePrefLineLenCalcElectron.copyIn(
			targetLineLenCalc,
			incomingLineLenCalc,
			ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyIn()\n" +
			"because targetLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	incomingLineLenCalc = nil

	err =
		ePrefLineLenCalcElectron.copyIn(
			targetLineLenCalc,
			incomingLineLenCalc,
			ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyIn()\n" +
			"because incomingLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	incomingLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	incomingLineLenCalc.ePrefDelimiters.newLinePrefixDelimiter = ""

	err =
		ePrefLineLenCalcElectron.copyIn(
			targetLineLenCalc,
			incomingLineLenCalc,
			ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyIn()\n" +
			"because incomingLineLenCalc is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetLineLenCalc = nil

	_,
		err = ePrefLineLenCalcElectron.copyOut(
		targetLineLenCalc,
		ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyOut()\n" +
			"because targetLineLenCalc is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetLineLenCalc = EPrefixLineLenCalc{}.Ptr()

	targetLineLenCalc.ePrefDelimiters.SetToDefault()

	targetLineLenCalc.ePrefDelimiters.newLinePrefixDelimiter = ""

	_,
		err = ePrefLineLenCalcElectron.copyOut(
		targetLineLenCalc,
		ePrefix)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.copyOut()\n" +
			"because targetLineLenCalc is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	/*
		func (ePrefLineLenCalcElectron *ePrefixLineLenCalcElectron) equal(
			lineLenCalcOne *EPrefixLineLenCalc,
			lineLenCalcTwo *EPrefixLineLenCalc,
			ePrefix string) (
			areEqual bool,
			err error)
	*/

	var lineLenCalcOne, lineLenCalcTwo EPrefixLineLenCalc

	lineLenCalcOne,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	_,
		err = ePrefLineLenCalcElectron.equal(
		nil,
		&lineLenCalcTwo,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.equal()\n" +
			"because lineLenCalcOne is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefLineLenCalcElectron.equal()\n" +
			"because lineLenCalcTwo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	var areEqual bool

	lineLenCalcTwo.ePrefDelimiters.newLinePrefixDelimiter = "@@@@@@@@@"

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because ePrefDelimiters.newLinePrefixDelimiter\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo.errorPrefixInfo.errorPrefixStr = ""

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because errorPrefixInfo.errorPrefixStr\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo.currentLineStr = "999"

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because lineLenCalcTwo.currentLineStr\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

	lineLenCalcTwo,
		err = getValidEPrefixLineLenCalc(funcName)

	if err != nil {
		t.Error(err.Error())
		return
	}

	lineLenCalcTwo.maxErrStringLength = 15

	areEqual,
		err = ePrefLineLenCalcElectron.equal(
		&lineLenCalcOne,
		&lineLenCalcTwo,
		funcName)

	if err != nil {
		t.Error("Equal test # 1" + err.Error())
		return
	}

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected 'areEqual' == false because lineLenCalcTwo.maxErrStringLength\n" +
			"values are not equal.\n" +
			"HOWEVER, the method showed the two EPrefixLineLenCalc\n" +
			"objects WERE EQUAL!\n")
		return
	}

}

func TestErrorPrefixInfo_Electron_000100(t *testing.T) {

	funcName := "TestErrorPrefixInfo_Electron_000100() "

	ePrefInfoElectron := errorPrefixInfoElectron{}

	targetErrPrefixInfo := ErrorPrefixInfo{}

	inComingErrPrefixInfo := getValidErrorPrefixInfo()

	err :=
		ePrefInfoElectron.copyIn(
			nil,
			&inComingErrPrefixInfo,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyIn()\n" +
			"because targetErrPrefixInfo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefInfoElectron.copyIn(
			&targetErrPrefixInfo,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyIn()\n" +
			"because inComingErrPrefixInfo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo.errorPrefixStr = ""

	err =
		ePrefInfoElectron.copyIn(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyIn()\n" +
			"because inComingErrPrefixInfo is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err =
		ePrefInfoElectron.copyOut(
			&inComingErrPrefixInfo,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyOut()\n" +
			"because inComingErrPrefixInfo is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	_,
		err =
		ePrefInfoElectron.copyOut(
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefInfoElectron.copyOut()\n" +
			"because inComingErrPrefixInfo is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	targetErrPrefixInfo = getValidErrorPrefixInfo()
	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	areEqual :=
		ePrefInfoElectron.equal(
			nil,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because ePrefixInfo01 is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			nil)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because ePrefixInfo02 is a nil pointer.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo.isFirstIdx = false

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because isFirstIdx is 'false'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.isLastIdx = false

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because isLastIdx is 'false'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.isPopulated = false

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because isPopulated is 'false'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.errorPrefixStr = "Nothing"

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the errorPrefixStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.lenErrorPrefixStr = 9999999

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the lenErrorPrefixStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.errorContextStr = "XXX-NOTHING-XXX"

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the errorContextStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	inComingErrPrefixInfo = getValidErrorPrefixInfo()

	inComingErrPrefixInfo.lenErrorContextStr = 55998823

	areEqual =
		ePrefInfoElectron.equal(
			&targetErrPrefixInfo,
			&inComingErrPrefixInfo)

	if areEqual {
		t.Error("ERROR:\n" +
			"Expected areEqual=='false' from ePrefInfoElectron.equal()\n" +
			"because the lenErrorContextStr values are not equal.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPrefixDto_Mechanics_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_Mechanics_000100() "

	/*
		func (ePrefDtoMech *errPrefixDtoMechanics) get2dEPrefStrings(
			errPrefDto *ErrPrefixDto,
		errorPrefStr string) (
			twoDStrArray [][2]string,
			err error)
	*/

	ePrefDtoMech := errPrefixDtoMechanics{}

	var ePDto ErrPrefixDto

	_,
		err := ePrefDtoMech.get2dEPrefStrings(
		nil,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDtoMech.get2dEPrefStrings()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto = getValidErrorPrefixDto()

	ePDto.ePrefCol = nil

	_,
		err = ePrefDtoMech.get2dEPrefStrings(
		&ePDto,
		funcName)

	if err != nil {
		t.Errorf("ERROR:\n"+
			"ePrefDtoMech.get2dEPrefStrings() SHOULD NOT HAVE RETURNED AND Error.\n"+
			"'ePDto.ePrefCol' parameter is 'nil'.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	/*
		func (ePrefDtoMech *errPrefixDtoMechanics) setFromEmptyInterface(
			errPrefDto *ErrPrefixDto,
		iEPref interface{},
		errorPrefStr string) error
	*/

	ePDto = getValidErrorPrefixDto()

	var twoDStrings [][2]string

	twoDStrings = ePDto.GetEPrefStrings()

	ePrefDtoMech = errPrefixDtoMechanics{}

	err = ePrefDtoMech.setFromEmptyInterface(
		nil,
		twoDStrings,
		funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefDtoMech.setFromEmptyInterface()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestErrPrefixDto_Nanobot_000100(t *testing.T) {

	funcName := "TestErrPrefixDto_Nanobot_000100() "

	ePrefixDtoNanobot := errPrefixDtoNanobot{}

	_,
		err :=
		ePrefixDtoNanobot.copyOutErrPrefDtoPtr(
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.copyOutErrPrefDtoPtr()\n" +
			"because the 'ePrefixDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	ePrefixDtoNanobot.deleteLastErrContext(nil)

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	ePDto2 := getValidErrorPrefixDto()

	err =
		ePrefixDtoNanobot.setFromIBuilder(
			nil,
			&ePDto2,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBuilder()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefixDtoNanobot.setFromIBuilder(
			&ePDto2,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBuilder()\n" +
			"because the 'iEPref' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	err =
		ePrefixDtoNanobot.setFromIBasicErrorPrefix(
			nil,
			&ePDto2,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBasicErrorPrefix()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefixDtoNanobot.setFromIBasicErrorPrefix(
			&ePDto2,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromIBasicErrorPrefix()\n" +
			"because the 'iEPref' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	err =
		ePrefixDtoNanobot.setFromString(
			nil,
			"Hello",
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromString()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	err =
		ePrefixDtoNanobot.setFromString(
			&ePDto2,
			"",
			funcName)

	if ePDto2.ePrefCol != nil {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoNanobot.setFromString() would return\n" +
			"ePDto2.ePrefCol==nil because the 'iEPref' parameter is" +
			"an empty string.\n" +
			"HOWEVER,  ePDto2.ePrefCol != nil!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	strBuilder := strings.Builder{}

	err =
		ePrefixDtoNanobot.setFromStringBuilder(
			nil,
			&strBuilder,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromStringBuilder()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto2 = getValidErrorPrefixDto()

	err =
		ePrefixDtoNanobot.setFromStringBuilder(
			&ePDto2,
			nil,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromStringBuilder()\n" +
			"because the 'iEPref' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	strArray := []string{"Hello", "Goodbye", "Come to dinner"}

	err =
		ePrefixDtoNanobot.setFromStringArray(
			nil,
			strArray,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromStringArray()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto2 = getValidErrorPrefixDto()
	strArray = nil

	_ =
		ePrefixDtoNanobot.setFromStringArray(
			&ePDto2,
			strArray,
			funcName)

	if ePDto2.ePrefCol != nil {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoNanobot.setFromStringArray() would return\n" +
			"ePDto2.ePrefCol==nil because the 'iEPref' parameter is" +
			"nil.\n" +
			"HOWEVER,  ePDto2.ePrefCol != nil!\n")
		return
	}

	ePrefixDtoNanobot = errPrefixDtoNanobot{}

	str2DArray := make([][2]string, 3)

	str2DArray[0][0] = "Hello()"
	str2DArray[1][0] = "GoodBye()"
	str2DArray[2][0] = "ComeAgain()"

	err =
		ePrefixDtoNanobot.setFromTwoDStrArray(
			nil,
			str2DArray,
			funcName)

	if err == nil {
		t.Error("ERROR:\n" +
			"Expected an error return from ePrefixDtoNanobot.setFromTwoDStrArray()\n" +
			"because the 'errPrefDto' parameter is 'nil'.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

	ePDto2 = getValidErrorPrefixDto()

	_ =
		ePrefixDtoNanobot.setFromTwoDStrArray(
			&ePDto2,
			nil,
			funcName)

	if ePDto2.ePrefCol != nil {
		t.Error("ERROR:\n" +
			"Expected ePrefixDtoNanobot.setFromTwoDStrArray() would return\n" +
			"ePDto2.ePrefCol==nil because the 'iEPref' parameter is" +
			"nil.\n" +
			"HOWEVER,  ePDto2.ePrefCol != nil!\n")
		return
	}

}
