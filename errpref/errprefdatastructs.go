package errpref

import "sync"

type EPrefixDelimiters struct {
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

// GetInLineContextDelimiter - Returns ePrefDelims.inLineContextDelimiter
func (ePrefDelims *EPrefixDelimiters) GetInLineContextDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.inLineContextDelimiter
}

// GetInLinePrefixDelimiter - Returns ePrefDelims.inLinePrefixDelimiter
func (ePrefDelims *EPrefixDelimiters) GetInLinePrefixDelimiter() string {

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
func (ePrefDelims *EPrefixDelimiters) GetLengthInLineContextDelimiter() uint {

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
func (ePrefDelims *EPrefixDelimiters) GetLengthInLinePrefixDelimiter() uint {

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
func (ePrefDelims *EPrefixDelimiters) GetLengthNewLineContextDelimiter() uint {

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
func (ePrefDelims *EPrefixDelimiters) GetLengthNewLinePrefixDelimiter() uint {

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
func (ePrefDelims *EPrefixDelimiters) GetNewLineContextDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.newLineContextDelimiter
}

// GetNewLinePrefixDelimiter - Returns ePrefDelims.newLinePrefixDelimiter
func (ePrefDelims *EPrefixDelimiters) GetNewLinePrefixDelimiter() string {

	if ePrefDelims.lock == nil {
		ePrefDelims.lock = new(sync.Mutex)
	}

	ePrefDelims.lock.Lock()

	defer ePrefDelims.lock.Unlock()

	return ePrefDelims.newLinePrefixDelimiter
}

// SetInLineContextDelimiter - Sets ePrefDelims.inLineContextDelimiter
//
// This method also sets the line length value for this parameter
// and stores it in an internal member variable. This value may
// be accessed through a 'Getter' method.
//
func (ePrefDelims *EPrefixDelimiters) SetInLineContextDelimiter(
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
func (ePrefDelims *EPrefixDelimiters) SetInLinePrefixDelimiter(
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
func (ePrefDelims *EPrefixDelimiters) SetNewLineContextDelimiter(
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
func (ePrefDelims *EPrefixDelimiters) SetNewLinePrefixDelimiter(
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
func (ePrefDelims *EPrefixDelimiters) SetLineLengthValues() {

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
