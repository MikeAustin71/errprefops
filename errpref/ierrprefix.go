package errpref

type IErrorPrefix interface {
	GetIsLastLineTerminatedWithNewLine() bool

	GetEPrefCollection() []ErrorPrefixInfo

	GetEPrefCollectionLen() int

	SetCtx(newErrContext string)

	SetCtxEmpty()

	SetEPref(newErrPrefix string)

	SetEPrefCollection(newEPrefCollection []ErrorPrefixInfo)

	SetEPrefCtx(newErrPrefix string, newErrContext string)

	SetEPrefOld(oldErrPrefix string)

	SetMaxTextLineLen(maxErrPrefixTextLineLength int)

	SetIsLastLineTermWithNewLine(isLastLineTerminatedWithNewLine bool)

	String() string

	ZCtx(newErrContext string) ErrPrefixDto

	ZCtxEmpty() ErrPrefixDto

	ZEPref(newErrPrefix string) ErrPrefixDto

	ZEPrefCtx(newErrPrefix string,
		newErrContext string) ErrPrefixDto

	ZEPrefOld(oldErrPrefix string) ErrPrefixDto
}
