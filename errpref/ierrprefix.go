package errpref

type IErrorPrefix interface {
	GetIsLastLineTerminatedWithNewLine() bool

	GetEPrefCollection() []ErrorPrefixInfo

	GetEPrefCollectionLen() int

	SetCtx(newErrContext string)

	SetCtxEmpty()

	SetEPref(newErrPrefix string)

	String() string

	SetEPrefCollection(newEPrefCollection []ErrorPrefixInfo)

	SetEPrefCtx(newErrPrefix string, newErrContext string)

	SetEPrefOld(oldErrPrefix string)

	SetMaxTextLineLen(maxErrPrefixTextLineLength int)

	SetIsLastLineTermWithNewLine(isLastLineTerminatedWithNewLine bool)
}
