package errpref

type IBasicErrorPrefix interface {
	GetEPrefStrings() [][2]string
}

type IErrorPrefix interface {
	GetEPrefStrings() [][2]string

	GetEPrefCollectionLen() int

	GetIsLastLineTerminatedWithNewLine() bool

	SetCtx(newErrContext string)

	SetCtxEmpty()

	SetEPref(newErrPrefix string)

	SetEPrefCtx(newErrPrefix string, newErrContext string)

	SetEPrefOld(oldErrPrefix string)

	SetMaxTextLineLen(maxErrPrefixTextLineLength int)

	SetIsLastLineTermWithNewLine(isLastLineTerminatedWithNewLine bool)

	String() string
}
