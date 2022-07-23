package exception

//Conflict code is 409 for data already exists
type ConflictError struct {
	ErrorString string
}

func NewConflictError(errString string) *ConflictError {
	return &ConflictError{
		ErrorString: errString,
	}
}

//implement error tidak menerima pointer
func (b ConflictError) Error() string {
	return b.ErrorString
}
