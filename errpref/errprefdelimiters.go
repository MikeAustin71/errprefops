package errpref

import (
	"sync"
)

type ErrPrefixDelimiters struct {
	inLinePrefixDelimiter      string
	lenInLinePrefixDelimiter   uint
	newLinePrefixDelimiter     string
	lenNewLinePrefixDelimiter  uint
	inLineContextDelimiter     string
	lenInLineContextDelimiter  uint
	newLineContextDelimiter    string
	lenNewLineContextDelimiter uint
	lock                       *sync.Mutex
}

// CopyIn - Receives an instance of type ErrPrefixDelimiters and
// proceeds to copy the internal member data variable values to the
// current ErrPrefixDelimiters instance.
//
func (ePrefDelims *ErrPrefixDelimiters) CopyIn(
	incomingDelimiters *ErrPrefixDelimiters,
	ePrefix string) error {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += "ErrPrefixDelimiters.CopyIn() "

	ePrefDelimsElectron := errPrefixDelimitersElectron{}

	return ePrefDelimsElectron.copyIn(
		ePrefDelims,
		incomingDelimiters,
		ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// ErrPrefixDelimiters. After completion of this operation, the
// returned copy and the current ErrPrefixDelimiters instance are
// identical in all respects.
//
func (ePrefDelims *ErrPrefixDelimiters) CopyOut() ErrPrefixDelimiters {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	newErrPrefDelims := ErrPrefixDelimiters{}

	newErrPrefDelims.inLinePrefixDelimiter =
		ePrefDelims.inLinePrefixDelimiter

	newErrPrefDelims.lenInLinePrefixDelimiter =
		ePrefDelims.lenInLinePrefixDelimiter

	newErrPrefDelims.newLinePrefixDelimiter =
		ePrefDelims.newLinePrefixDelimiter

	newErrPrefDelims.lenNewLinePrefixDelimiter =
		ePrefDelims.lenNewLinePrefixDelimiter

	newErrPrefDelims.inLineContextDelimiter =
		ePrefDelims.inLineContextDelimiter

	newErrPrefDelims.lenNewLinePrefixDelimiter =
		ePrefDelims.lenNewLinePrefixDelimiter

	newErrPrefDelims.newLineContextDelimiter =
		ePrefDelims.newLineContextDelimiter

	newErrPrefDelims.lenNewLineContextDelimiter =
		ePrefDelims.lenNewLineContextDelimiter

	return newErrPrefDelims
}

// GetInLineContextDelimiter - Returns ePrefDelims.inLineContextDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetInLineContextDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.inLineContextDelimiter
}

// GetInLinePrefixDelimiter - Returns ePrefDelims.inLinePrefixDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetInLinePrefixDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.inLinePrefixDelimiter
}

// GetLengthInLineContextDelimiter - Returns the number of
// characters in the 'In Line Context Delimiter' string as an unsigned
// integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthInLineContextDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenInLineContextDelimiter == 0 {

		ePrefDelims.lenInLineContextDelimiter =
			uint(len(ePrefDelims.inLineContextDelimiter))

	}

	return ePrefDelims.lenInLineContextDelimiter
}

// GetLengthInLinePrefixDelimiter - Returns the number of
// characters in the 'In Line Prefix Delimiter' string as an unsigned
// integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthInLinePrefixDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenInLinePrefixDelimiter == 0 {

		ePrefDelims.lenInLinePrefixDelimiter =
			uint(len(ePrefDelims.inLinePrefixDelimiter))

	}

	return ePrefDelims.lenInLinePrefixDelimiter
}

// GetLengthNewLineContextDelimiter - Returns the number of
// characters in the 'New Line Context Delimiter' string as an
// unsigned integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthNewLineContextDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenNewLineContextDelimiter == 0 {

		ePrefDelims.lenNewLineContextDelimiter =
			uint(len(ePrefDelims.newLineContextDelimiter))

	}

	return ePrefDelims.lenNewLineContextDelimiter
}

// GetLengthNewLinePrefixDelimiter - Returns the number of
// characters in the 'New Line Prefix Delimiter' string as an unsigned
// integer.
//
func (ePrefDelims *ErrPrefixDelimiters) GetLengthNewLinePrefixDelimiter() uint {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	if ePrefDelims.lenNewLinePrefixDelimiter == 0 {

		ePrefDelims.lenNewLinePrefixDelimiter =
			uint(len(ePrefDelims.newLinePrefixDelimiter))

	}

	return ePrefDelims.lenNewLinePrefixDelimiter
}

// GetNewLineContextDelimiter - Returns ePrefDelims.newLineContextDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetNewLineContextDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.newLineContextDelimiter
}

// GetNewLinePrefixDelimiter - Returns ePrefDelims.newLinePrefixDelimiter
func (ePrefDelims *ErrPrefixDelimiters) GetNewLinePrefixDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.newLinePrefixDelimiter
}

// IsValidInstance - Returns a boolean flag signalling whether the
// current ErrPrefixDelimiters instance is valid, or not.
//
// If this method returns a boolean value of 'false', it signals
// that the current ErrPrefixDelimiters instance is invalid.
//
// If this method returns a boolean value of 'true', it signals
// that the current ErrPrefixDelimiters instance is valid in all
// respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - This boolean flag signals whether the current
//       ErrPrefixDelimiters instance is valid.
//
//       If this method returns a value of 'false', it signals that
//       the current ErrPrefixDelimiters instance is invalid.
//
//       If this method returns a value of 'true', it signals that
//       the current ErrPrefixDelimiters instance is valid in all
//       respects.
//
func (ePrefDelims *ErrPrefixDelimiters) IsValidInstance(
	ePrefix string) bool {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += "ErrPrefixDelimiters.IsValidInstance() "

	ePrefDelimsQuark := errPrefixDelimitersQuark{}

	isValid,
		_ := ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		ePrefDelims,
		ePrefix)

	return isValid
}

// IsValidInstanceError - Returns an error type signalling whether
// the current ErrPrefixDelimiters instance is valid, or not.
//
// If this method returns an error value NOT equal to 'nil', it
// signals that the current ErrPrefixDelimiters instance is
// invalid.
//
// If this method returns an error value which IS equal to 'nil',
// it signals that the current ErrPrefixDelimiters instance is
// valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this returned error type is set equal to 'nil', it
//       signals that the current ErrPrefixDelimiters is valid in
//       all respects.
//
//       If this returned error type is NOT equal to 'nil', it
//       signals that the current ErrPrefixDelimiters is invalid.
//
func (ePrefDelims *ErrPrefixDelimiters) IsValidInstanceError(
	ePrefix string) error {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefix += "ErrPrefixDelimiters.IsValidInstanceError() "

	ePrefDelimsQuark := errPrefixDelimitersQuark{}

	_,
		err := ePrefDelimsQuark.testValidityOfErrPrefixDelimiters(
		ePrefDelims,
		ePrefix)

	return err
}

// SetInLineContextDelimiter - Sets ePrefDelims.inLineContextDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetInLineContextDelimiter(
	inLineContextDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.inLineContextDelimiter = inLineContextDelimiter

	ePrefDelims.lenInLineContextDelimiter =
		uint(len(ePrefDelims.inLineContextDelimiter))

	return
}

// SetInLinePrefixDelimiter - Sets ePrefDelims.inLinePrefixDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetInLinePrefixDelimiter(
	inLinePrefixDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.inLinePrefixDelimiter = inLinePrefixDelimiter

	ePrefDelims.lenInLinePrefixDelimiter =
		uint(len(ePrefDelims.inLinePrefixDelimiter))

	return
}

// SetNewLineContextDelimiter - Sets ePrefDelims.newLineContextDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetNewLineContextDelimiter(
	newLineContextDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.newLineContextDelimiter = newLineContextDelimiter

	ePrefDelims.lenNewLineContextDelimiter =
		uint(len(ePrefDelims.newLineContextDelimiter))

	return
}

// SetNewLinePrefixDelimiter - Sets ePrefDelims.newLinePrefixDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *ErrPrefixDelimiters) SetNewLinePrefixDelimiter(
	newLinePrefixDelimiter string) {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.newLinePrefixDelimiter = newLinePrefixDelimiter

	ePrefDelims.lenNewLinePrefixDelimiter =
		uint(len(ePrefDelims.newLinePrefixDelimiter))

	return
}

// SetLineLengthValues - Automatically calculates, sets and stores
// the line lengths for all delimiters. These values are stored in
// internal member variables and may be accessed using 'Getter'
// methods.
//
func (ePrefDelims *ErrPrefixDelimiters) SetLineLengthValues() {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	ePrefDelims.lenInLineContextDelimiter =
		uint(len(ePrefDelims.inLineContextDelimiter))

	ePrefDelims.lenInLinePrefixDelimiter =
		uint(len(ePrefDelims.inLinePrefixDelimiter))

	ePrefDelims.lenNewLinePrefixDelimiter =
		uint(len(ePrefDelims.newLinePrefixDelimiter))

	ePrefDelims.lenNewLineContextDelimiter =
		uint(len(ePrefDelims.newLineContextDelimiter))

	return
}
