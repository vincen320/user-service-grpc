package exception

type NotFoundError struct {
	ErrorString string
}

func NewNotFoundError(errString string) *NotFoundError {
	return &NotFoundError{
		ErrorString: errString,
	}
}

//implement error tidak menerima pointer
func (b NotFoundError) Error() string {
	return b.ErrorString
}
